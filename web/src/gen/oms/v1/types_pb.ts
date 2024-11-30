// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file oms/v1/types.proto (package oms.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * CustomerAction enumerates the set of supported customer actions
 *
 * @generated from enum oms.v1.CustomerAction
 */
export enum CustomerAction {
  /**
   * @generated from enum value: CUSTOMER_ACTION_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: CUSTOMER_ACTION_CANCEL = 1;
   */
  CANCEL = 1,

  /**
   * @generated from enum value: CUSTOMER_ACTION_AMEND = 2;
   */
  AMEND = 2,

  /**
   * @generated from enum value: CUSTOMER_ACTION_TIMED_OUT = 3;
   */
  TIMED_OUT = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(CustomerAction)
proto3.util.setEnumType(CustomerAction, "oms.v1.CustomerAction", [
  { no: 0, name: "CUSTOMER_ACTION_UNSPECIFIED" },
  { no: 1, name: "CUSTOMER_ACTION_CANCEL" },
  { no: 2, name: "CUSTOMER_ACTION_AMEND" },
  { no: 3, name: "CUSTOMER_ACTION_TIMED_OUT" },
]);

/**
 * FulfillmentStatus enumerates the set of supported fulfillment statuses
 *
 * @generated from enum oms.v1.FulfillmentStatus
 */
export enum FulfillmentStatus {
  /**
   * @generated from enum value: FULFILLMENT_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: FULFILLMENT_STATUS_UNAVAILABLE = 1;
   */
  UNAVAILABLE = 1,

  /**
   * @generated from enum value: FULFILLMENT_STATUS_PENDING = 2;
   */
  PENDING = 2,

  /**
   * @generated from enum value: FULFILLMENT_STATUS_PROCESSING = 3;
   */
  PROCESSING = 3,

  /**
   * @generated from enum value: FULFILLMENT_STATUS_COMPLETED = 4;
   */
  COMPLETED = 4,

  /**
   * @generated from enum value: FULFILLMENT_STATUS_CANCELLED = 5;
   */
  CANCELLED = 5,

  /**
   * @generated from enum value: FULFILLMENT_STATUS_FAILED = 6;
   */
  FAILED = 6,
}
// Retrieve enum metadata with: proto3.getEnumType(FulfillmentStatus)
proto3.util.setEnumType(FulfillmentStatus, "oms.v1.FulfillmentStatus", [
  { no: 0, name: "FULFILLMENT_STATUS_UNSPECIFIED" },
  { no: 1, name: "FULFILLMENT_STATUS_UNAVAILABLE" },
  { no: 2, name: "FULFILLMENT_STATUS_PENDING" },
  { no: 3, name: "FULFILLMENT_STATUS_PROCESSING" },
  { no: 4, name: "FULFILLMENT_STATUS_COMPLETED" },
  { no: 5, name: "FULFILLMENT_STATUS_CANCELLED" },
  { no: 6, name: "FULFILLMENT_STATUS_FAILED" },
]);

/**
 * OrderStatus enumerates the set of supported order statuses
 *
 * @generated from enum oms.v1.OrderStatus
 */
export enum OrderStatus {
  /**
   * @generated from enum value: ORDER_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ORDER_STATUS_PENDING = 1;
   */
  PENDING = 1,

  /**
   * @generated from enum value: ORDER_STATUS_PROCESSING = 2;
   */
  PROCESSING = 2,

  /**
   * @generated from enum value: ORDER_STATUS_CUSTOMER_ACTION_REQUIRED = 3;
   */
  CUSTOMER_ACTION_REQUIRED = 3,

  /**
   * @generated from enum value: ORDER_STATUS_COMPLETED = 4;
   */
  COMPLETED = 4,

  /**
   * @generated from enum value: ORDER_STATUS_FAILED = 5;
   */
  FAILED = 5,

  /**
   * @generated from enum value: ORDER_STATUS_CANCELLED = 6;
   */
  CANCELLED = 6,

  /**
   * @generated from enum value: ORDER_STATUS_TIMED_OUT = 7;
   */
  TIMED_OUT = 7,
}
// Retrieve enum metadata with: proto3.getEnumType(OrderStatus)
proto3.util.setEnumType(OrderStatus, "oms.v1.OrderStatus", [
  { no: 0, name: "ORDER_STATUS_UNSPECIFIED" },
  { no: 1, name: "ORDER_STATUS_PENDING" },
  { no: 2, name: "ORDER_STATUS_PROCESSING" },
  { no: 3, name: "ORDER_STATUS_CUSTOMER_ACTION_REQUIRED" },
  { no: 4, name: "ORDER_STATUS_COMPLETED" },
  { no: 5, name: "ORDER_STATUS_FAILED" },
  { no: 6, name: "ORDER_STATUS_CANCELLED" },
  { no: 7, name: "ORDER_STATUS_TIMED_OUT" },
]);

/**
 * PaymentStatus enumerates the set of supported payment statuses
 *
 * @generated from enum oms.v1.PaymentStatus
 */
export enum PaymentStatus {
  /**
   * @generated from enum value: PAYMENT_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: PAYMENT_STATUS_PENDING = 1;
   */
  PENDING = 1,

  /**
   * @generated from enum value: PAYMENT_STATUS_SUCCESS = 2;
   */
  SUCCESS = 2,

  /**
   * @generated from enum value: PAYMENT_STATUS_FAILED = 3;
   */
  FAILED = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(PaymentStatus)
proto3.util.setEnumType(PaymentStatus, "oms.v1.PaymentStatus", [
  { no: 0, name: "PAYMENT_STATUS_UNSPECIFIED" },
  { no: 1, name: "PAYMENT_STATUS_PENDING" },
  { no: 2, name: "PAYMENT_STATUS_SUCCESS" },
  { no: 3, name: "PAYMENT_STATUS_FAILED" },
]);

/**
 * @generated from enum oms.v1.ShipmentStatus
 */
export enum ShipmentStatus {
  /**
   * @generated from enum value: SHIPMENT_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SHIPMENT_STATUS_PENDING = 1;
   */
  PENDING = 1,

  /**
   * @generated from enum value: SHIPMENT_STATUS_BOOKED = 2;
   */
  BOOKED = 2,

  /**
   * @generated from enum value: SHIPMENT_STATUS_DISPATCHED = 3;
   */
  DISPATCHED = 3,

  /**
   * @generated from enum value: SHIPMENT_STATUS_DELIVERED = 4;
   */
  DELIVERED = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(ShipmentStatus)
proto3.util.setEnumType(ShipmentStatus, "oms.v1.ShipmentStatus", [
  { no: 0, name: "SHIPMENT_STATUS_UNSPECIFIED" },
  { no: 1, name: "SHIPMENT_STATUS_PENDING" },
  { no: 2, name: "SHIPMENT_STATUS_BOOKED" },
  { no: 3, name: "SHIPMENT_STATUS_DISPATCHED" },
  { no: 4, name: "SHIPMENT_STATUS_DELIVERED" },
]);

/**
 * holds a set of items that will be delivered in one shipment (due to location and stock level).
 *
 * @generated from message oms.v1.Fulfillment
 */
export class Fulfillment extends Message<Fulfillment> {
  /**
   * @generated from field: string order_id = 1;
   */
  orderId = "";

  /**
   * @generated from field: string customer_id = 2;
   */
  customerId = "";

  /**
   * @generated from field: string id = 3;
   */
  id = "";

  /**
   * @generated from field: repeated oms.v1.Item items = 4;
   */
  items: Item[] = [];

  /**
   * @generated from field: string location = 5;
   */
  location = "";

  /**
   * @generated from field: oms.v1.FulfillmentStatus status = 6;
   */
  status = FulfillmentStatus.UNSPECIFIED;

  /**
   * @generated from field: oms.v1.Payment payment = 7;
   */
  payment?: Payment;

  /**
   * @generated from field: oms.v1.Shipment shipment = 8;
   */
  shipment?: Shipment;

  constructor(data?: PartialMessage<Fulfillment>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "oms.v1.Fulfillment";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "order_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "customer_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "items", kind: "message", T: Item, repeated: true },
    { no: 5, name: "location", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "status", kind: "enum", T: proto3.getEnumType(FulfillmentStatus) },
    { no: 7, name: "payment", kind: "message", T: Payment },
    { no: 8, name: "shipment", kind: "message", T: Shipment },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Fulfillment {
    return new Fulfillment().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Fulfillment {
    return new Fulfillment().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Fulfillment {
    return new Fulfillment().fromJsonString(jsonString, options);
  }

  static equals(a: Fulfillment | PlainMessage<Fulfillment> | undefined, b: Fulfillment | PlainMessage<Fulfillment> | undefined): boolean {
    return proto3.util.equals(Fulfillment, a, b);
  }
}

/**
 * Item represents an item being ordered
 *
 * @generated from message oms.v1.Item
 */
export class Item extends Message<Item> {
  /**
   * @generated from field: string sku = 1;
   */
  sku = "";

  /**
   * @generated from field: string description = 2;
   */
  description = "";

  /**
   * @generated from field: int32 quantity = 3;
   */
  quantity = 0;

  constructor(data?: PartialMessage<Item>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "oms.v1.Item";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sku", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "quantity", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Item {
    return new Item().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Item {
    return new Item().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Item {
    return new Item().fromJsonString(jsonString, options);
  }

  static equals(a: Item | PlainMessage<Item> | undefined, b: Item | PlainMessage<Item> | undefined): boolean {
    return proto3.util.equals(Item, a, b);
  }
}

/**
 * OrderStatus holds the status of an Order workflow.
 *
 * @generated from message oms.v1.Order
 */
export class Order extends Message<Order> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string customer_id = 2;
   */
  customerId = "";

  /**
   * @generated from field: google.protobuf.Timestamp received_at = 3;
   */
  receivedAt?: Timestamp;

  /**
   * @generated from field: oms.v1.OrderStatus status = 4;
   */
  status = OrderStatus.UNSPECIFIED;

  /**
   * @generated from field: repeated oms.v1.Fulfillment fulfillments = 5;
   */
  fulfillments: Fulfillment[] = [];

  constructor(data?: PartialMessage<Order>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "oms.v1.Order";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "customer_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "received_at", kind: "message", T: Timestamp },
    { no: 4, name: "status", kind: "enum", T: proto3.getEnumType(OrderStatus) },
    { no: 5, name: "fulfillments", kind: "message", T: Fulfillment, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Order {
    return new Order().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Order {
    return new Order().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Order {
    return new Order().fromJsonString(jsonString, options);
  }

  static equals(a: Order | PlainMessage<Order> | undefined, b: Order | PlainMessage<Order> | undefined): boolean {
    return proto3.util.equals(Order, a, b);
  }
}

/**
 * Payment holds the status of a Payment.
 *
 * @generated from message oms.v1.Payment
 */
export class Payment extends Message<Payment> {
  /**
   * @generated from field: int32 sub_total = 1;
   */
  subTotal = 0;

  /**
   * @generated from field: int32 tax = 2;
   */
  tax = 0;

  /**
   * @generated from field: int32 shipping = 3;
   */
  shipping = 0;

  /**
   * @generated from field: int32 total = 4;
   */
  total = 0;

  /**
   * @generated from field: oms.v1.PaymentStatus status = 5;
   */
  status = PaymentStatus.UNSPECIFIED;

  constructor(data?: PartialMessage<Payment>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "oms.v1.Payment";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sub_total", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "tax", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "shipping", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "total", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "status", kind: "enum", T: proto3.getEnumType(PaymentStatus) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Payment {
    return new Payment().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Payment {
    return new Payment().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Payment {
    return new Payment().fromJsonString(jsonString, options);
  }

  static equals(a: Payment | PlainMessage<Payment> | undefined, b: Payment | PlainMessage<Payment> | undefined): boolean {
    return proto3.util.equals(Payment, a, b);
  }
}

/**
 * Shipment describes the output from a Status query
 *
 * @generated from message oms.v1.Shipment
 */
export class Shipment extends Message<Shipment> {
  /**
   * @generated from field: string courier_reference = 1;
   */
  courierReference = "";

  /**
   * @generated from field: string id = 2;
   */
  id = "";

  /**
   * @generated from field: repeated oms.v1.Item items = 3;
   */
  items: Item[] = [];

  /**
   * @generated from field: oms.v1.ShipmentStatus status = 4;
   */
  status = ShipmentStatus.UNSPECIFIED;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 5;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Shipment>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "oms.v1.Shipment";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "courier_reference", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "items", kind: "message", T: Item, repeated: true },
    { no: 4, name: "status", kind: "enum", T: proto3.getEnumType(ShipmentStatus) },
    { no: 5, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Shipment {
    return new Shipment().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Shipment {
    return new Shipment().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Shipment {
    return new Shipment().fromJsonString(jsonString, options);
  }

  static equals(a: Shipment | PlainMessage<Shipment> | undefined, b: Shipment | PlainMessage<Shipment> | undefined): boolean {
    return proto3.util.equals(Shipment, a, b);
  }
}

