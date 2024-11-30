package billing

import (
	"strings"

	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
)

type Workflows struct{}

type ChargeWorkflow struct {
	*billingv1.ChargeWorkflowInput
	logger log.Logger
}

func (w *Workflows) Charge(ctx workflow.Context, input *billingv1.ChargeWorkflowInput) (billingv1.ChargeWorkflow, error) {
	return &ChargeWorkflow{input, workflow.GetLogger(ctx)}, nil
}

// Charge Workflow invoices and processes payment for a fulfillment.
func (wf *ChargeWorkflow) Execute(ctx workflow.Context) (*billingv1.ChargeResult, error) {
	idempotencyKey := wf.Req.GetIdempotencyKey()
	if idempotencyKey == "" {
		idempotencyKey = strings.TrimPrefix(workflow.GetInfo(ctx).WorkflowExecution.ID, "Charge:")
	}

	invoice, err := billingv1.GenerateInvoice(ctx, &billingv1.GenerateInvoiceInput{
		CustomerId: wf.Req.GetCustomerId(),
		Reference:  idempotencyKey,
		Items:      wf.Req.GetItems(),
	})
	if err != nil {
		return nil, err
	}

	result, err := billingv1.CheckFraud(ctx, &billingv1.CheckFraudInput{
		CustomerId: wf.Req.GetCustomerId(),
		Charge:     invoice.GetTotal(),
	})
	if err != nil {
		return nil, err
	} else if result.GetDeclined() {
		wf.logger.Warn("Fraudulent charge detected", "customer_id", wf.Req.GetCustomerId(), "error", err)
		return &billingv1.ChargeResult{Success: false}, nil
	}

	charge, err := billingv1.ProcessPayment(ctx, &billingv1.ProcessPaymentInput{
		CustomerId: wf.Req.GetCustomerId(),
		Reference:  invoice.GetInvoiceReference(),
		Charge:     invoice.GetTotal(),
	})
	if err != nil {
		wf.logger.Warn("Charge failed", "customer_id", wf.Req.GetCustomerId(), "error", err)
		charge.Success = false
	}

	return &billingv1.ChargeResult{
		InvoiceReference: invoice.InvoiceReference,
		SubTotal:         invoice.SubTotal,
		Tax:              invoice.Tax,
		Shipping:         invoice.Shipping,
		Total:            invoice.Total,

		Success:  charge.Success,
		AuthCode: charge.AuthCode,
	}, nil
}
