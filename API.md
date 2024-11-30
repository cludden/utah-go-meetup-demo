# Table of Contents

- [oms.billing.v1](#oms-billing-v1)
  - Services
    - [oms.billing.v1.Worker](#oms-billing-v1-worker)
      - [Workflows](#oms-billing-v1-worker-workflows)
        - [billing.v1.Charge](#billing-v1-charge-workflow)
      - [Activities](#oms-billing-v1-worker-activities)
        - [oms.billing.v1.Worker.CheckFraud](#oms-billing-v1-worker-checkfraud-activity)
        - [oms.billing.v1.Worker.GenerateInvoice](#oms-billing-v1-worker-generateinvoice-activity)
        - [oms.billing.v1.Worker.ProcessPayment](#oms-billing-v1-worker-processpayment-activity)
  - Messages
    - [oms.billing.v1.ChargeInput](#oms-billing-v1-chargeinput)
    - [oms.billing.v1.ChargeResult](#oms-billing-v1-chargeresult)
    - [oms.billing.v1.CheckFraudInput](#oms-billing-v1-checkfraudinput)
    - [oms.billing.v1.CheckFraudResult](#oms-billing-v1-checkfraudresult)
    - [oms.billing.v1.GenerateInvoiceInput](#oms-billing-v1-generateinvoiceinput)
    - [oms.billing.v1.GenerateInvoiceResult](#oms-billing-v1-generateinvoiceresult)
    - [oms.billing.v1.ProcessPaymentInput](#oms-billing-v1-processpaymentinput)
    - [oms.billing.v1.ProcessPaymentResult](#oms-billing-v1-processpaymentresult)
- [oms.order.v1](#oms-order-v1)
  - Services
    - [oms.order.v1.Worker](#oms-order-v1-worker)
      - [Workflows](#oms-order-v1-worker-workflows)
        - [order.v1.Order](#order-v1-order-workflow)
      - [Queries](#oms-order-v1-worker-queries)
        - [order.v1.GetStatus](#order-v1-getstatus-query)
      - [Signals](#oms-order-v1-worker-signals)
        - [order.v1.CustomerAction](#order-v1-customeraction-signal)
      - [Updates](#oms-order-v1-worker-updates)
        - [order.v1.UpdateOrderStatus](#order-v1-updateorderstatus-update)
      - [Activities](#oms-order-v1-worker-activities)
        - [order.v1.CreateOrder](#order-v1-createorder-activity)
        - [order.v1.ReserveItems](#order-v1-reserveitems-activity)
        - [order.v1.UpdateOrderStatus](#order-v1-updateorderstatus-activity)
  - Messages
    - [oms.order.v1.CreateOrderInput](#oms-order-v1-createorderinput)
    - [oms.order.v1.CreateOrderResult](#oms-order-v1-createorderresult)
    - [oms.order.v1.CustomerActionInput](#oms-order-v1-customeractioninput)
    - [oms.order.v1.GetOrderInput](#oms-order-v1-getorderinput)
    - [oms.order.v1.GetOrderResult](#oms-order-v1-getorderresult)
    - [oms.order.v1.ListOrdersResult](#oms-order-v1-listordersresult)
    - [oms.order.v1.ReserveItemsInput](#oms-order-v1-reserveitemsinput)
    - [oms.order.v1.ReserveItemsResult](#oms-order-v1-reserveitemsresult)
    - [oms.order.v1.UpdateOrderStatusInput](#oms-order-v1-updateorderstatusinput)
- [oms.shipment.v1](#oms-shipment-v1)
  - Services
    - [oms.shipment.v1.Worker](#oms-shipment-v1-worker)
      - [Workflows](#oms-shipment-v1-worker-workflows)
        - [shipment.v1.Shipment](#shipment-v1-shipment-workflow)
      - [Queries](#oms-shipment-v1-worker-queries)
        - [shipment.v1.Status](#shipment-v1-status-query)
      - [Signals](#oms-shipment-v1-worker-signals)
        - [shipment.v1.ShipmentStatusUpdated](#shipment-v1-shipmentstatusupdated-signal)
      - [Updates](#oms-shipment-v1-worker-updates)
        - [shipment.v1.UpdateShipmentStatus](#shipment-v1-updateshipmentstatus-update)
      - [Activities](#oms-shipment-v1-worker-activities)
        - [oms.shipment.v1.Worker.CreateShipment](#oms-shipment-v1-worker-createshipment-activity)
        - [oms.shipment.v1.Worker.UpdateShipmentStatus](#oms-shipment-v1-worker-updateshipmentstatus-activity)
  - Messages
    - [oms.shipment.v1.CreateShipmentInput](#oms-shipment-v1-createshipmentinput)
    - [oms.shipment.v1.CreateShipmentResult](#oms-shipment-v1-createshipmentresult)
    - [oms.shipment.v1.GetShipmentInput](#oms-shipment-v1-getshipmentinput)
    - [oms.shipment.v1.GetShipmentResult](#oms-shipment-v1-getshipmentresult)
    - [oms.shipment.v1.ListShipmentsInput](#oms-shipment-v1-listshipmentsinput)
    - [oms.shipment.v1.ListShipmentsResult](#oms-shipment-v1-listshipmentsresult)
    - [oms.shipment.v1.UpdateShipmentStatusInput](#oms-shipment-v1-updateshipmentstatusinput)
- [google.protobuf](#google-protobuf)
  - Messages
    - [google.protobuf.Timestamp](#google-protobuf-timestamp)
- [oms.v1](#oms-v1)
  - Messages
    - [oms.v1.CustomerAction](#oms-v1-customeraction)
    - [oms.v1.Fulfillment](#oms-v1-fulfillment)
    - [oms.v1.FulfillmentStatus](#oms-v1-fulfillmentstatus)
    - [oms.v1.Item](#oms-v1-item)
    - [oms.v1.Order](#oms-v1-order)
    - [oms.v1.OrderStatus](#oms-v1-orderstatus)
    - [oms.v1.Payment](#oms-v1-payment)
    - [oms.v1.PaymentStatus](#oms-v1-paymentstatus)
    - [oms.v1.Shipment](#oms-v1-shipment)
    - [oms.v1.ShipmentStatus](#oms-v1-shipmentstatus)

<a name="oms-billing-v1"></a>
# oms.billing.v1

<a name="oms-billing-v1-services"></a>
## Services

<a name="oms-billing-v1-worker"></a>
## oms.billing.v1.Worker

<a name="oms-billing-v1-worker-workflows"></a>
### Workflows

---
<a name="billing-v1-charge-workflow"></a>
### billing.v1.Charge

<pre>
durably process an order from a customer
</pre>

**Input:** [oms.billing.v1.ChargeInput](#oms-billing-v1-chargeinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>fulfillment_id</td>
<td>string</td>
<td><pre>
json_name: fulfillmentId
go_name: FulfillmentId</pre></td>
</tr><tr>
<td>idempotency_key</td>
<td>string</td>
<td><pre>
json_name: idempotencyKey
go_name: IdempotencyKey</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>order_id</td>
<td>string</td>
<td><pre>
json_name: orderId
go_name: OrderId</pre></td>
</tr><tr>
<td>reference</td>
<td>string</td>
<td><pre>
json_name: reference
go_name: Reference</pre></td>
</tr>
</table>

**Output:** [oms.billing.v1.ChargeResult](#oms-billing-v1-chargeresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>auth_code</td>
<td>string</td>
<td><pre>
json_name: authCode
go_name: AuthCode</pre></td>
</tr><tr>
<td>invoice_reference</td>
<td>string</td>
<td><pre>
json_name: invoiceReference
go_name: InvoiceReference</pre></td>
</tr><tr>
<td>shipping</td>
<td>int32</td>
<td><pre>
json_name: shipping
go_name: Shipping</pre></td>
</tr><tr>
<td>sub_total</td>
<td>int32</td>
<td><pre>
json_name: subTotal
go_name: SubTotal</pre></td>
</tr><tr>
<td>success</td>
<td>bool</td>
<td><pre>
json_name: success
go_name: Success</pre></td>
</tr><tr>
<td>tax</td>
<td>int32</td>
<td><pre>
json_name: tax
go_name: Tax</pre></td>
</tr><tr>
<td>total</td>
<td>int32</td>
<td><pre>
json_name: total
go_name: Total</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>Charge:${! idempotencyKey.or(uuid_v4()) }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
<tr><td>search_attributes</td><td><pre><code>CustomerId = customerId 
FulfillmentId = fulfillmentId 
OrderId = orderId</code></pre></td></tr>
</table>    

<a name="oms-billing-v1-worker-activities"></a>
### Activities

---
<a name="oms-billing-v1-worker-checkfraud-activity"></a>
### oms.billing.v1.Worker.CheckFraud

<pre>
determines whether the charge is fraudulent
</pre>

**Input:** [oms.billing.v1.CheckFraudInput](#oms-billing-v1-checkfraudinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>charge</td>
<td>int32</td>
<td><pre>
json_name: charge
go_name: Charge</pre></td>
</tr><tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr>
</table>

**Output:** [oms.billing.v1.CheckFraudResult](#oms-billing-v1-checkfraudresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>declined</td>
<td>bool</td>
<td><pre>
json_name: declined
go_name: Declined</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.initial_interval</td><td>5 seconds</td></tr>
<tr><td>retry_policy.max_interval</td><td>5 seconds</td></tr>
<tr><td>start_to_close_timeout</td><td>5 minutes</td></tr>
</table> 

---
<a name="oms-billing-v1-worker-generateinvoice-activity"></a>
### oms.billing.v1.Worker.GenerateInvoice

<pre>
generates an invoice file
</pre>

**Input:** [oms.billing.v1.GenerateInvoiceInput](#oms-billing-v1-generateinvoiceinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>reference</td>
<td>string</td>
<td><pre>
json_name: reference
go_name: Reference</pre></td>
</tr>
</table>

**Output:** [oms.billing.v1.GenerateInvoiceResult](#oms-billing-v1-generateinvoiceresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>invoice_reference</td>
<td>string</td>
<td><pre>
json_name: invoiceReference
go_name: InvoiceReference</pre></td>
</tr><tr>
<td>shipping</td>
<td>int32</td>
<td><pre>
json_name: shipping
go_name: Shipping</pre></td>
</tr><tr>
<td>sub_total</td>
<td>int32</td>
<td><pre>
json_name: subTotal
go_name: SubTotal</pre></td>
</tr><tr>
<td>tax</td>
<td>int32</td>
<td><pre>
json_name: tax
go_name: Tax</pre></td>
</tr><tr>
<td>total</td>
<td>int32</td>
<td><pre>
json_name: total
go_name: Total</pre></td>
</tr>
</table> 

---
<a name="oms-billing-v1-worker-processpayment-activity"></a>
### oms.billing.v1.Worker.ProcessPayment

<pre>
processes the customer payment
</pre>

**Input:** [oms.billing.v1.ProcessPaymentInput](#oms-billing-v1-processpaymentinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>charge</td>
<td>int32</td>
<td><pre>
json_name: charge
go_name: Charge</pre></td>
</tr><tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>reference</td>
<td>string</td>
<td><pre>
json_name: reference
go_name: Reference</pre></td>
</tr>
</table>

**Output:** [oms.billing.v1.ProcessPaymentResult](#oms-billing-v1-processpaymentresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>auth_code</td>
<td>string</td>
<td><pre>
json_name: authCode
go_name: AuthCode</pre></td>
</tr><tr>
<td>success</td>
<td>bool</td>
<td><pre>
json_name: success
go_name: Success</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>2 minutes</td></tr>
</table>   

<a name="oms-billing-v1-messages"></a>
## Messages

<a name="oms-billing-v1-chargeinput"></a>
### oms.billing.v1.ChargeInput

<pre>
ChargeInput is the input for the Charge workflow.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>fulfillment_id</td>
<td>string</td>
<td><pre>
json_name: fulfillmentId
go_name: FulfillmentId</pre></td>
</tr><tr>
<td>idempotency_key</td>
<td>string</td>
<td><pre>
json_name: idempotencyKey
go_name: IdempotencyKey</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>order_id</td>
<td>string</td>
<td><pre>
json_name: orderId
go_name: OrderId</pre></td>
</tr><tr>
<td>reference</td>
<td>string</td>
<td><pre>
json_name: reference
go_name: Reference</pre></td>
</tr>
</table>



<a name="oms-billing-v1-chargeresult"></a>
### oms.billing.v1.ChargeResult

<pre>
ChargeResult is the result for the Charge workflow.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>auth_code</td>
<td>string</td>
<td><pre>
json_name: authCode
go_name: AuthCode</pre></td>
</tr><tr>
<td>invoice_reference</td>
<td>string</td>
<td><pre>
json_name: invoiceReference
go_name: InvoiceReference</pre></td>
</tr><tr>
<td>shipping</td>
<td>int32</td>
<td><pre>
json_name: shipping
go_name: Shipping</pre></td>
</tr><tr>
<td>sub_total</td>
<td>int32</td>
<td><pre>
json_name: subTotal
go_name: SubTotal</pre></td>
</tr><tr>
<td>success</td>
<td>bool</td>
<td><pre>
json_name: success
go_name: Success</pre></td>
</tr><tr>
<td>tax</td>
<td>int32</td>
<td><pre>
json_name: tax
go_name: Tax</pre></td>
</tr><tr>
<td>total</td>
<td>int32</td>
<td><pre>
json_name: total
go_name: Total</pre></td>
</tr>
</table>



<a name="oms-billing-v1-checkfraudinput"></a>
### oms.billing.v1.CheckFraudInput

<pre>
CheckFraudInput describes the input to a CheckFraud activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>charge</td>
<td>int32</td>
<td><pre>
json_name: charge
go_name: Charge</pre></td>
</tr><tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr>
</table>



<a name="oms-billing-v1-checkfraudresult"></a>
### oms.billing.v1.CheckFraudResult

<pre>
CheckFraudResult describes the output from a CheckFraud activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>declined</td>
<td>bool</td>
<td><pre>
json_name: declined
go_name: Declined</pre></td>
</tr>
</table>



<a name="oms-billing-v1-generateinvoiceinput"></a>
### oms.billing.v1.GenerateInvoiceInput

<pre>
GenerateInvoiceInput describes the input to a GenerateInvoice activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>reference</td>
<td>string</td>
<td><pre>
json_name: reference
go_name: Reference</pre></td>
</tr>
</table>



<a name="oms-billing-v1-generateinvoiceresult"></a>
### oms.billing.v1.GenerateInvoiceResult

<pre>
GenerateInvoiceResult describes the output from a GenerateInvoice activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>invoice_reference</td>
<td>string</td>
<td><pre>
json_name: invoiceReference
go_name: InvoiceReference</pre></td>
</tr><tr>
<td>shipping</td>
<td>int32</td>
<td><pre>
json_name: shipping
go_name: Shipping</pre></td>
</tr><tr>
<td>sub_total</td>
<td>int32</td>
<td><pre>
json_name: subTotal
go_name: SubTotal</pre></td>
</tr><tr>
<td>tax</td>
<td>int32</td>
<td><pre>
json_name: tax
go_name: Tax</pre></td>
</tr><tr>
<td>total</td>
<td>int32</td>
<td><pre>
json_name: total
go_name: Total</pre></td>
</tr>
</table>



<a name="oms-billing-v1-processpaymentinput"></a>
### oms.billing.v1.ProcessPaymentInput

<pre>
ProcessPaymentInput describes the input to a ProcessPayment activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>charge</td>
<td>int32</td>
<td><pre>
json_name: charge
go_name: Charge</pre></td>
</tr><tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>reference</td>
<td>string</td>
<td><pre>
json_name: reference
go_name: Reference</pre></td>
</tr>
</table>



<a name="oms-billing-v1-processpaymentresult"></a>
### oms.billing.v1.ProcessPaymentResult

<pre>
ProcessPaymentResult describes the output from a ProcessPayment activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>auth_code</td>
<td>string</td>
<td><pre>
json_name: authCode
go_name: AuthCode</pre></td>
</tr><tr>
<td>success</td>
<td>bool</td>
<td><pre>
json_name: success
go_name: Success</pre></td>
</tr>
</table>



<a name="oms-order-v1"></a>
# oms.order.v1

<a name="oms-order-v1-services"></a>
## Services

<a name="oms-order-v1-worker"></a>
## oms.order.v1.Worker

<a name="oms-order-v1-worker-workflows"></a>
### Workflows

---
<a name="order-v1-order-workflow"></a>
### order.v1.Order

<pre>
manage the lifecycle of an order
</pre>

**Input:** [oms.order.v1.CreateOrderInput](#oms-order-v1-createorderinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr>
</table>

**Output:** [oms.order.v1.CreateOrderResult](#oms-order-v1-createorderresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>order</td>
<td><a href="#oms-v1-order">oms.v1.Order</a></td>
<td><pre>
json_name: order
go_name: Order</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>Order:${! id.or(throw("id required")) }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE</code></pre></td></tr>
<tr><td>search_attributes</td><td><pre><code>CustomerId = customerId 
OrderId = id</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#oms-order-v1-worker-getstatus-query">oms.order.v1.Worker.GetStatus</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#oms-order-v1-worker-customeraction-signal">oms.order.v1.Worker.CustomerAction</a></td><td>false</td></tr>
<tr><td><a href="#oms-order-v1-worker-oms-shipment-v1-worker-shipmentstatusupdated-signal">oms.order.v1.Worker.oms.shipment.v1.Worker.ShipmentStatusUpdated</a></td><td>false</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#oms-order-v1-worker-updateorderstatus-update">oms.order.v1.Worker.UpdateOrderStatus</a></td></tr>
</table>  

<a name="oms-order-v1-worker-queries"></a>
### Queries

---
<a name="order-v1-getstatus-query"></a>
### order.v1.GetStatus

<pre>
returns information about the order
</pre>

**Output:** [oms.order.v1.GetOrderResult](#oms-order-v1-getorderresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>order</td>
<td><a href="#oms-v1-order">oms.v1.Order</a></td>
<td><pre>
json_name: order
go_name: Order</pre></td>
</tr>
</table>  

<a name="oms-order-v1-worker-signals"></a>
### Signals

---
<a name="order-v1-customeraction-signal"></a>
### order.v1.CustomerAction

<pre>
process a customer action
</pre>

**Input:** [oms.order.v1.CustomerActionInput](#oms-order-v1-customeractioninput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>action</td>
<td><a href="#oms-v1-customeraction">oms.v1.CustomerAction</a></td>
<td><pre>
json_name: action
go_name: Action</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr>
</table>  

<a name="oms-order-v1-worker-updates"></a>
### Updates

---
<a name="order-v1-updateorderstatus-update"></a>
### order.v1.UpdateOrderStatus

<pre>
updates the order status in the database
</pre>

**Input:** [oms.order.v1.UpdateOrderStatusInput](#oms-order-v1-updateorderstatusinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-orderstatus">oms.v1.OrderStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

<a name="oms-order-v1-worker-activities"></a>
### Activities

---
<a name="order-v1-createorder-activity"></a>
### order.v1.CreateOrder

<pre>
initialize a new order
</pre>

**Input:** [oms.order.v1.CreateOrderInput](#oms-order-v1-createorderinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr>
</table>

**Output:** [oms.order.v1.CreateOrderResult](#oms-order-v1-createorderresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>order</td>
<td><a href="#oms-v1-order">oms.v1.Order</a></td>
<td><pre>
json_name: order
go_name: Order</pre></td>
</tr>
</table> 

---
<a name="order-v1-reserveitems-activity"></a>
### order.v1.ReserveItems

<pre>
reserves items to satisfy an order and returns a list of reservations for the items
Any unavailable items will be returned in a Reservation with Available set to false.
In a real system this would involve an inventory database of some kind.
For our purposes we just split orders arbitrarily.
</pre>

**Input:** [oms.order.v1.ReserveItemsInput](#oms-order-v1-reserveitemsinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>order_id</td>
<td>string</td>
<td><pre>
json_name: orderId
go_name: OrderId</pre></td>
</tr>
</table>

**Output:** [oms.order.v1.ReserveItemsResult](#oms-order-v1-reserveitemsresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>fulfillments</td>
<td><a href="#oms-v1-fulfillment">oms.v1.Fulfillment</a></td>
<td><pre>
json_name: fulfillments
go_name: Fulfillments</pre></td>
</tr>
</table> 

---
<a name="order-v1-updateorderstatus-activity"></a>
### order.v1.UpdateOrderStatus

<pre>
updates the order status in the database
</pre>

**Input:** [oms.order.v1.UpdateOrderStatusInput](#oms-order-v1-updateorderstatusinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-orderstatus">oms.v1.OrderStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table>   

<a name="oms-order-v1-messages"></a>
## Messages

<a name="oms-order-v1-createorderinput"></a>
### oms.order.v1.CreateOrderInput

<pre>
CreateOrderInput describes the input to an CreateOrder workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr>
</table>



<a name="oms-order-v1-createorderresult"></a>
### oms.order.v1.CreateOrderResult

<pre>
CreateOrderResult describes the output from an CreateOrder workflow
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>order</td>
<td><a href="#oms-v1-order">oms.v1.Order</a></td>
<td><pre>
json_name: order
go_name: Order</pre></td>
</tr>
</table>



<a name="oms-order-v1-customeractioninput"></a>
### oms.order.v1.CustomerActionInput

<pre>
CustomerActionInput describes the input to a CustomerAction signal
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>action</td>
<td><a href="#oms-v1-customeraction">oms.v1.CustomerAction</a></td>
<td><pre>
json_name: action
go_name: Action</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr>
</table>



<a name="oms-order-v1-getorderinput"></a>
### oms.order.v1.GetOrderInput

<pre>
GetOrderInput describes the input to a GetOrderStatus query
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr>
</table>



<a name="oms-order-v1-getorderresult"></a>
### oms.order.v1.GetOrderResult

<pre>
GetOrderResult describes the output from a GetOrderStatus query
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>order</td>
<td><a href="#oms-v1-order">oms.v1.Order</a></td>
<td><pre>
json_name: order
go_name: Order</pre></td>
</tr>
</table>



<a name="oms-order-v1-listordersresult"></a>
### oms.order.v1.ListOrdersResult

<pre>
ListOrdersResult describes the output from a ListOrders operation
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>orders</td>
<td><a href="#oms-v1-order">oms.v1.Order</a></td>
<td><pre>
json_name: orders
go_name: Orders</pre></td>
</tr>
</table>



<a name="oms-order-v1-reserveitemsinput"></a>
### oms.order.v1.ReserveItemsInput

<pre>
ReserveItemsInput describes the input to a ReserveItems activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>order_id</td>
<td>string</td>
<td><pre>
json_name: orderId
go_name: OrderId</pre></td>
</tr>
</table>



<a name="oms-order-v1-reserveitemsresult"></a>
### oms.order.v1.ReserveItemsResult

<pre>
ReserveItemsResult describes the output from a ReserveItems activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>fulfillments</td>
<td><a href="#oms-v1-fulfillment">oms.v1.Fulfillment</a></td>
<td><pre>
json_name: fulfillments
go_name: Fulfillments</pre></td>
</tr>
</table>



<a name="oms-order-v1-updateorderstatusinput"></a>
### oms.order.v1.UpdateOrderStatusInput

<pre>
UpdateOrderStatusInput describes the input to an UpdateOrderStatus activity
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-orderstatus">oms.v1.OrderStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="oms-shipment-v1"></a>
# oms.shipment.v1

<a name="oms-shipment-v1-services"></a>
## Services

<a name="oms-shipment-v1-worker"></a>
## oms.shipment.v1.Worker

<a name="oms-shipment-v1-worker-workflows"></a>
### Workflows

---
<a name="shipment-v1-shipment-workflow"></a>
### shipment.v1.Shipment

<pre>
process a shipment
</pre>

**Input:** [oms.shipment.v1.CreateShipmentInput](#oms-shipment-v1-createshipmentinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>fulfillment_id</td>
<td>string</td>
<td><pre>
json_name: fulfillmentId
go_name: FulfillmentId</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>order_id</td>
<td>string</td>
<td><pre>
json_name: orderId
go_name: OrderId</pre></td>
</tr><tr>
<td>requestor_wid</td>
<td>string</td>
<td><pre>
json_name: requestorWid
go_name: RequestorWid</pre></td>
</tr>
</table>

**Output:** [oms.shipment.v1.CreateShipmentResult](#oms-shipment-v1-createshipmentresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>courier_reference</td>
<td>string</td>
<td><pre>
json_name: courierReference
go_name: CourierReference</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>Shipment:${! id.or(throw("id required")) }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
<tr><td>search_attributes</td><td><pre><code>CustomerId = customerId 
FulfillmentId = fulfillmentId 
OrderId = orderId</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#oms-shipment-v1-worker-getstatus-query">oms.shipment.v1.Worker.GetStatus</a></td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#oms-shipment-v1-worker-updateshipmentstatus-update">oms.shipment.v1.Worker.UpdateShipmentStatus</a></td></tr>
</table>  

<a name="oms-shipment-v1-worker-queries"></a>
### Queries

---
<a name="shipment-v1-status-query"></a>
### shipment.v1.Status

<pre>
get shipment status
</pre>

**Output:** [oms.shipment.v1.GetShipmentResult](#oms-shipment-v1-getshipmentresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>shipment</td>
<td><a href="#oms-v1-shipment">oms.v1.Shipment</a></td>
<td><pre>
json_name: shipment
go_name: Shipment</pre></td>
</tr>
</table>  

<a name="oms-shipment-v1-worker-signals"></a>
### Signals

---
<a name="shipment-v1-shipmentstatusupdated-signal"></a>
### shipment.v1.ShipmentStatusUpdated

<pre>
notify the requestor of an update to a shipment's status.
</pre>

**Input:** [oms.shipment.v1.UpdateShipmentStatusInput](#oms-shipment-v1-updateshipmentstatusinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>shipment_id</td>
<td>string</td>
<td><pre>
json_name: shipmentId
go_name: ShipmentId</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-shipmentstatus">oms.v1.ShipmentStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>updated_at</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: updatedAt
go_name: UpdatedAt</pre></td>
</tr>
</table>  

<a name="oms-shipment-v1-worker-updates"></a>
### Updates

---
<a name="shipment-v1-updateshipmentstatus-update"></a>
### shipment.v1.UpdateShipmentStatus

<pre>
process shiptment status update from carrier
</pre>

**Input:** [oms.shipment.v1.UpdateShipmentStatusInput](#oms-shipment-v1-updateshipmentstatusinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>shipment_id</td>
<td>string</td>
<td><pre>
json_name: shipmentId
go_name: ShipmentId</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-shipmentstatus">oms.v1.ShipmentStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>updated_at</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: updatedAt
go_name: UpdatedAt</pre></td>
</tr>
</table>

<a name="oms-shipment-v1-worker-activities"></a>
### Activities

---
<a name="oms-shipment-v1-worker-createshipment-activity"></a>
### oms.shipment.v1.Worker.CreateShipment

<pre>
books shipment with carrier and persists record to database
</pre>

**Input:** [oms.shipment.v1.CreateShipmentInput](#oms-shipment-v1-createshipmentinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>fulfillment_id</td>
<td>string</td>
<td><pre>
json_name: fulfillmentId
go_name: FulfillmentId</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>order_id</td>
<td>string</td>
<td><pre>
json_name: orderId
go_name: OrderId</pre></td>
</tr><tr>
<td>requestor_wid</td>
<td>string</td>
<td><pre>
json_name: requestorWid
go_name: RequestorWid</pre></td>
</tr>
</table>

**Output:** [oms.shipment.v1.CreateShipmentResult](#oms-shipment-v1-createshipmentresult)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>courier_reference</td>
<td>string</td>
<td><pre>
json_name: courierReference
go_name: CourierReference</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_interval</td><td>5 seconds</td></tr>
<tr><td>start_to_close_timeout</td><td>5 seconds</td></tr>
</table> 

---
<a name="oms-shipment-v1-worker-updateshipmentstatus-activity"></a>
### oms.shipment.v1.Worker.UpdateShipmentStatus

<pre>
process shiptment status update from carrier
</pre>

**Input:** [oms.shipment.v1.UpdateShipmentStatusInput](#oms-shipment-v1-updateshipmentstatusinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>shipment_id</td>
<td>string</td>
<td><pre>
json_name: shipmentId
go_name: ShipmentId</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-shipmentstatus">oms.v1.ShipmentStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>updated_at</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: updatedAt
go_name: UpdatedAt</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>start_to_close_timeout</td><td>5 seconds</td></tr>
</table>   

<a name="oms-shipment-v1-messages"></a>
## Messages

<a name="oms-shipment-v1-createshipmentinput"></a>
### oms.shipment.v1.CreateShipmentInput

<pre>
CreateShipmentInput describes the input to a CreateShipment operation
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>fulfillment_id</td>
<td>string</td>
<td><pre>
json_name: fulfillmentId
go_name: FulfillmentId</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>order_id</td>
<td>string</td>
<td><pre>
json_name: orderId
go_name: OrderId</pre></td>
</tr><tr>
<td>requestor_wid</td>
<td>string</td>
<td><pre>
json_name: requestorWid
go_name: RequestorWid</pre></td>
</tr>
</table>



<a name="oms-shipment-v1-createshipmentresult"></a>
### oms.shipment.v1.CreateShipmentResult

<pre>
CreateShipmentResult describes the output from a CreateShipment operation
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>courier_reference</td>
<td>string</td>
<td><pre>
json_name: courierReference
go_name: CourierReference</pre></td>
</tr>
</table>



<a name="oms-shipment-v1-getshipmentinput"></a>
### oms.shipment.v1.GetShipmentInput

<pre>
GetShipmentInput describes the input to a GetShipment operation
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr>
</table>



<a name="oms-shipment-v1-getshipmentresult"></a>
### oms.shipment.v1.GetShipmentResult

<pre>
GetShipmentResult describes the output from a GetShipment operation
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>shipment</td>
<td><a href="#oms-v1-shipment">oms.v1.Shipment</a></td>
<td><pre>
json_name: shipment
go_name: Shipment</pre></td>
</tr>
</table>



<a name="oms-shipment-v1-listshipmentsinput"></a>
### oms.shipment.v1.ListShipmentsInput

<pre>
ListShipmentsInput describes the input to a ListShipments operation
</pre>



<a name="oms-shipment-v1-listshipmentsresult"></a>
### oms.shipment.v1.ListShipmentsResult

<pre>
ListShipmentsResult describes the output from a ListShipments operation
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>shipments</td>
<td><a href="#oms-v1-shipment">oms.v1.Shipment</a></td>
<td><pre>
json_name: shipments
go_name: Shipments</pre></td>
</tr>
</table>



<a name="oms-shipment-v1-updateshipmentstatusinput"></a>
### oms.shipment.v1.UpdateShipmentStatusInput

<pre>
UpdateShipmentStatusInput describes the input to an UpdateShipmentStatus
operation
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>shipment_id</td>
<td>string</td>
<td><pre>
json_name: shipmentId
go_name: ShipmentId</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-shipmentstatus">oms.v1.ShipmentStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>updated_at</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: updatedAt
go_name: UpdatedAt</pre></td>
</tr>
</table>




<a name="google-protobuf"></a>
# google.protobuf

<a name="google-protobuf-messages"></a>
## Messages

<a name="google-protobuf-timestamp"></a>
### google.protobuf.Timestamp

<pre>
A Timestamp represents a point in time independent of any time zone or local
calendar, encoded as a count of seconds and fractions of seconds at
nanosecond resolution. The count is relative to an epoch at UTC midnight on
January 1, 1970, in the proleptic Gregorian calendar which extends the
Gregorian calendar backwards to year one.

All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
second table is needed for interpretation, using a [24-hour linear
smear](https://developers.google.com/time/smear).

The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
restricting to that range, we ensure that we can convert to and from [RFC
3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.

# Examples

Example 1: Compute Timestamp from POSIX `time()`.

    Timestamp timestamp;
    timestamp.set_seconds(time(NULL));
    timestamp.set_nanos(0);

Example 2: Compute Timestamp from POSIX `gettimeofday()`.

    struct timeval tv;
    gettimeofday(&tv, NULL);

    Timestamp timestamp;
    timestamp.set_seconds(tv.tv_sec);
    timestamp.set_nanos(tv.tv_usec * 1000);

Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

    FILETIME ft;
    GetSystemTimeAsFileTime(&ft);
    UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

    // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
    // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
    Timestamp timestamp;
    timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
    timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

    long millis = System.currentTimeMillis();

    Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
        .setNanos((int) ((millis % 1000) * 1000000)).build();

Example 5: Compute Timestamp from Java `Instant.now()`.

    Instant now = Instant.now();

    Timestamp timestamp =
        Timestamp.newBuilder().setSeconds(now.getEpochSecond())
            .setNanos(now.getNano()).build();

Example 6: Compute Timestamp from current time in Python.

    timestamp = Timestamp()
    timestamp.GetCurrentTime()

# JSON Mapping

In JSON format, the Timestamp type is encoded as a string in the
[RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
where {year} is always expressed using four digits while {month}, {day},
{hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
is required. A proto3 JSON serializer should always use UTC (as indicated by
"Z") when printing the Timestamp type and a proto3 JSON parser should be
able to accept both UTC and other timezones (as indicated by an offset).

For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
01:30 UTC on January 15, 2017.

In JavaScript, one can convert a Date object to this format using the
standard
[toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
method. In Python, a standard `datetime.datetime` object can be converted
to this format using
[`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
the Joda Time's [`ISODateTimeFormat.dateTime()`](
http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()
) to obtain a formatter capable of generating timestamps in this format.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>nanos</td>
<td>int32</td>
<td><pre>
Non-negative fractions of a second at nanosecond resolution. Negative
second values with fractions must still have non-negative nanos values
that count forward in time. Must be from 0 to 999,999,999
inclusive.<br>

json_name: nanos
go_name: Nanos</pre></td>
</tr><tr>
<td>seconds</td>
<td>int64</td>
<td><pre>
Represents seconds of UTC time since Unix epoch
1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
9999-12-31T23:59:59Z inclusive.<br>

json_name: seconds
go_name: Seconds</pre></td>
</tr>
</table>




<a name="oms-v1"></a>
# oms.v1

<a name="oms-v1-messages"></a>
## Messages

<a name="oms-v1-customeraction"></a>
### oms.v1.CustomerAction

<pre>
CustomerAction enumerates the set of supported customer actions
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>CUSTOMER_ACTION_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>CUSTOMER_ACTION_CANCEL</td>
<td></td>
</tr><tr>
<td>CUSTOMER_ACTION_AMEND</td>
<td></td>
</tr><tr>
<td>CUSTOMER_ACTION_TIMED_OUT</td>
<td></td>
</tr>
</table>

<a name="oms-v1-fulfillment"></a>
### oms.v1.Fulfillment

<pre>
holds a set of items that will be delivered in one shipment (due to location and stock level).
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>location</td>
<td>string</td>
<td><pre>
json_name: location
go_name: Location</pre></td>
</tr><tr>
<td>order_id</td>
<td>string</td>
<td><pre>
json_name: orderId
go_name: OrderId</pre></td>
</tr><tr>
<td>payment</td>
<td><a href="#oms-v1-payment">oms.v1.Payment</a></td>
<td><pre>
json_name: payment
go_name: Payment</pre></td>
</tr><tr>
<td>shipment</td>
<td><a href="#oms-v1-shipment">oms.v1.Shipment</a></td>
<td><pre>
json_name: shipment
go_name: Shipment</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-fulfillmentstatus">oms.v1.FulfillmentStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="oms-v1-fulfillmentstatus"></a>
### oms.v1.FulfillmentStatus

<pre>
FulfillmentStatus enumerates the set of supported fulfillment statuses
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>FULFILLMENT_STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>FULFILLMENT_STATUS_UNAVAILABLE</td>
<td></td>
</tr><tr>
<td>FULFILLMENT_STATUS_PENDING</td>
<td></td>
</tr><tr>
<td>FULFILLMENT_STATUS_PROCESSING</td>
<td></td>
</tr><tr>
<td>FULFILLMENT_STATUS_COMPLETED</td>
<td></td>
</tr><tr>
<td>FULFILLMENT_STATUS_CANCELLED</td>
<td></td>
</tr><tr>
<td>FULFILLMENT_STATUS_FAILED</td>
<td></td>
</tr>
</table>

<a name="oms-v1-item"></a>
### oms.v1.Item

<pre>
Item represents an item being ordered
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>description</td>
<td>string</td>
<td><pre>
json_name: description
go_name: Description</pre></td>
</tr><tr>
<td>quantity</td>
<td>int32</td>
<td><pre>
json_name: quantity
go_name: Quantity</pre></td>
</tr><tr>
<td>sku</td>
<td>string</td>
<td><pre>
json_name: sku
go_name: Sku</pre></td>
</tr>
</table>



<a name="oms-v1-order"></a>
### oms.v1.Order

<pre>
OrderStatus holds the status of an Order workflow.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>fulfillments</td>
<td><a href="#oms-v1-fulfillment">oms.v1.Fulfillment</a></td>
<td><pre>
json_name: fulfillments
go_name: Fulfillments</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>received_at</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: receivedAt
go_name: ReceivedAt</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-orderstatus">oms.v1.OrderStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="oms-v1-orderstatus"></a>
### oms.v1.OrderStatus

<pre>
OrderStatus enumerates the set of supported order statuses
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>ORDER_STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>ORDER_STATUS_PENDING</td>
<td></td>
</tr><tr>
<td>ORDER_STATUS_PROCESSING</td>
<td></td>
</tr><tr>
<td>ORDER_STATUS_CUSTOMER_ACTION_REQUIRED</td>
<td></td>
</tr><tr>
<td>ORDER_STATUS_COMPLETED</td>
<td></td>
</tr><tr>
<td>ORDER_STATUS_FAILED</td>
<td></td>
</tr><tr>
<td>ORDER_STATUS_CANCELLED</td>
<td></td>
</tr><tr>
<td>ORDER_STATUS_TIMED_OUT</td>
<td></td>
</tr>
</table>

<a name="oms-v1-payment"></a>
### oms.v1.Payment

<pre>
Payment holds the status of a Payment.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>shipping</td>
<td>int32</td>
<td><pre>
json_name: shipping
go_name: Shipping</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-paymentstatus">oms.v1.PaymentStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>sub_total</td>
<td>int32</td>
<td><pre>
json_name: subTotal
go_name: SubTotal</pre></td>
</tr><tr>
<td>tax</td>
<td>int32</td>
<td><pre>
json_name: tax
go_name: Tax</pre></td>
</tr><tr>
<td>total</td>
<td>int32</td>
<td><pre>
json_name: total
go_name: Total</pre></td>
</tr>
</table>



<a name="oms-v1-paymentstatus"></a>
### oms.v1.PaymentStatus

<pre>
PaymentStatus enumerates the set of supported payment statuses
</pre>

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>PAYMENT_STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>PAYMENT_STATUS_PENDING</td>
<td></td>
</tr><tr>
<td>PAYMENT_STATUS_SUCCESS</td>
<td></td>
</tr><tr>
<td>PAYMENT_STATUS_FAILED</td>
<td></td>
</tr>
</table>

<a name="oms-v1-shipment"></a>
### oms.v1.Shipment

<pre>
Shipment describes the output from a Status query
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>courier_reference</td>
<td>string</td>
<td><pre>
json_name: courierReference
go_name: CourierReference</pre></td>
</tr><tr>
<td>id</td>
<td>string</td>
<td><pre>
json_name: id
go_name: Id</pre></td>
</tr><tr>
<td>items</td>
<td><a href="#oms-v1-item">oms.v1.Item</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr><tr>
<td>status</td>
<td><a href="#oms-v1-shipmentstatus">oms.v1.ShipmentStatus</a></td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr><tr>
<td>updated_at</td>
<td><a href="#google-protobuf-timestamp">google.protobuf.Timestamp</a></td>
<td><pre>
json_name: updatedAt
go_name: UpdatedAt</pre></td>
</tr>
</table>



<a name="oms-v1-shipmentstatus"></a>
### oms.v1.ShipmentStatus

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>SHIPMENT_STATUS_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>SHIPMENT_STATUS_PENDING</td>
<td></td>
</tr><tr>
<td>SHIPMENT_STATUS_BOOKED</td>
<td></td>
</tr><tr>
<td>SHIPMENT_STATUS_DISPATCHED</td>
<td></td>
</tr><tr>
<td>SHIPMENT_STATUS_DELIVERED</td>
<td></td>
</tr>
</table>