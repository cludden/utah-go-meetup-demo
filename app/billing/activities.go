package billing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/temporalio/reference-app-orders-go/app/fraud"
	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	omsv1 "github.com/temporalio/reference-app-orders-go/gen/oms/v1"
	"go.temporal.io/sdk/activity"
)

// Activities implements the billing package's Activities.
// Any state shared by the worker among the activities is stored here.
type Activities struct {
	FraudCheckURL string
}

// GenerateInvoice activity creates an invoice for a fulfillment.
func (a *Activities) GenerateInvoice(ctx context.Context, input *billingv1.GenerateInvoiceInput) (*billingv1.GenerateInvoiceResult, error) {
	var result billingv1.GenerateInvoiceResult

	if input.GetCustomerId() == "" {
		return nil, fmt.Errorf("CustomerID is required")
	}
	if input.GetReference() == "" {
		return nil, fmt.Errorf("OrderReference is required")
	}
	if len(input.GetItems()) == 0 {
		return nil, fmt.Errorf("invoice must have items")
	}

	result.InvoiceReference = input.GetReference()

	for _, item := range input.Items {
		cost, tax := calculateCosts(item)
		result.SubTotal += cost
		result.Tax += tax
		result.Shipping += calculateShippingCost(item)
		result.Total += result.SubTotal + result.Tax + result.Shipping
	}

	activity.GetLogger(ctx).Info(
		"Invoice",
		"Customer", input.GetCustomerId(),
		"Total", result.GetTotal(),
		"Reference", result.GetInvoiceReference(),
	)

	return &result, nil
}

// CheckFraud uses as external vendor to determine whether a given charge is fraudulent
func (a *Activities) CheckFraud(ctx context.Context, input *billingv1.CheckFraudInput) (*billingv1.CheckFraudResult, error) {
	if a.FraudCheckURL == "" {
		return &billingv1.CheckFraudResult{Declined: false}, nil
	}

	checkInput := fraud.FraudCheckInput{
		CustomerID: input.GetCustomerId(),
		Charge:     input.GetCharge(),
	}
	jsonInput, err := json.Marshal(checkInput)
	if err != nil {
		return nil, fmt.Errorf("failed to encode input: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.FraudCheckURL+"/check", bytes.NewReader(jsonInput))
	if err != nil {
		return nil, fmt.Errorf("failed to build fraud check request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("fraud check request failed: %s: %s", http.StatusText(res.StatusCode), body)
	}

	var checkResult fraud.FraudCheckResult

	err = json.NewDecoder(res.Body).Decode(&checkResult)
	return &billingv1.CheckFraudResult{Declined: checkResult.Declined}, err
}

// ProcessPayment activity charges a customer for a fulfillment.
func (a *Activities) ProcessPayment(ctx context.Context, input *billingv1.ProcessPaymentInput) (*billingv1.ProcessPaymentResult, error) {
	var result billingv1.ProcessPaymentResult

	result.AuthCode = "1234"
	result.Success = true

	activity.GetLogger(ctx).Info(
		"Charge",
		"Customer", input.GetCustomerId(),
		"Amount", input.GetCharge(),
		"Reference", input.GetReference(),
		"Success", result.GetSuccess(),
	)
	return &result, nil
}

// =============================================================================

// calculateCosts calculates the cost and tax for an item.
func calculateCosts(item *omsv1.Item) (cost int32, tax int32) {
	// This is just a simulation, so make up a cost
	// Normally this would be looked up on the SKU
	costPerUnit := 3500 + rand.Int31n(8500)
	// Return tax at 20%
	return costPerUnit * int32(item.Quantity), costPerUnit * int32(item.Quantity) / 5
}

// calculateShippingCost calculates the shipping cost for an item.
func calculateShippingCost(item *omsv1.Item) int32 {
	// This is just a simulation, so make up a cost
	// Normally this would be looked up on the SKU
	costPerUnit := 500 + rand.Int31n(500)
	return costPerUnit * int32(item.Quantity)
}
