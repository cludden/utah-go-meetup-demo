import { writable } from 'svelte/store';
import { 
	FulfillmentStatus,
	Item,
	OrderStatus,
	PaymentStatus,
	ShipmentStatus,
} from '../../gen/oms/v1/types_pb';
import { CreateOrderInput } from '../../gen/oms/order/v1/orders_pb';

export type orderStatus =
	| 'pending'
	| 'processing'
	| 'customerActionRequired'
	| 'completed'
	| 'failed'
	| 'cancelled';

export interface Order {
	id: string;
	customerId: string;
	items: Item[];
	fulfillments?: Fulfillment[];
	status?: orderStatus;
}

export type fulfillmentStatus =
	| 'unavailable'
	| 'pending'
	| 'processing'
	| 'completed'
	| 'cancelled'
	| 'failed';

export interface Fulfillment {
	id: string;
	shipment?: Shipment;
	items: Item[];
	payment?: Payment;
	location: string;
	status?: fulfillmentStatus;
}

export interface Shipment {
	id: string;
	status: string;
	items: Item[];
	updatedAt: string;
}

export type paymentStatus = 'pending' | 'success' | 'failed';

export interface Payment {
	shipping: number;
	tax: number;
	subTotal: number;
	total: number;
	status: paymentStatus;
}

export let orderStatusNames: Map<OrderStatus, string> = new Map();
orderStatusNames.set(OrderStatus.CANCELLED, 'cancelled');
orderStatusNames.set(OrderStatus.COMPLETED, 'completed');
orderStatusNames.set(OrderStatus.CUSTOMER_ACTION_REQUIRED, 'customerActionRequired');
orderStatusNames.set(OrderStatus.FAILED, 'failed');
orderStatusNames.set(OrderStatus.PENDING, 'pending');
orderStatusNames.set(OrderStatus.PROCESSING, 'processing');
orderStatusNames.set(OrderStatus.TIMED_OUT, 'timedOut');

export let orderStatusValues: Map<string, OrderStatus> = new Map();
orderStatusValues.set('cancelled', OrderStatus.CANCELLED);
orderStatusValues.set('completed', OrderStatus.COMPLETED);
orderStatusValues.set('customerActionRequired', OrderStatus.CUSTOMER_ACTION_REQUIRED);
orderStatusValues.set('failed', OrderStatus.FAILED);
orderStatusValues.set('pending', OrderStatus.PENDING);
orderStatusValues.set('processing', OrderStatus.PROCESSING);
orderStatusValues.set('timedOut', OrderStatus.TIMED_OUT);

export let fulfillmentStatusNames: Map<FulfillmentStatus, string> = new Map()
fulfillmentStatusNames.set(FulfillmentStatus.CANCELLED, 'cancelled');
fulfillmentStatusNames.set(FulfillmentStatus.COMPLETED, 'completed');
fulfillmentStatusNames.set(FulfillmentStatus.FAILED, 'failed');
fulfillmentStatusNames.set(FulfillmentStatus.PENDING, 'pending');
fulfillmentStatusNames.set(FulfillmentStatus.PROCESSING, 'processing');
fulfillmentStatusNames.set(FulfillmentStatus.UNAVAILABLE, 'unavailable');

export let fulfillmentStatusValues: Map<string, FulfillmentStatus> = new Map()
fulfillmentStatusValues.set('cancelled', FulfillmentStatus.CANCELLED);
fulfillmentStatusValues.set('completed', FulfillmentStatus.COMPLETED);
fulfillmentStatusValues.set('failed', FulfillmentStatus.FAILED);
fulfillmentStatusValues.set('pending', FulfillmentStatus.PENDING);
fulfillmentStatusValues.set('processing', FulfillmentStatus.PROCESSING);
fulfillmentStatusValues.set('unavailable', FulfillmentStatus.UNAVAILABLE);


export type p = 'pending' | 'success' | 'failed';

export let paymentStatusNames: Map<PaymentStatus, string> = new Map();
paymentStatusNames.set(PaymentStatus.FAILED, 'failed');
paymentStatusNames.set(PaymentStatus.PENDING, 'pending');
paymentStatusNames.set(PaymentStatus.SUCCESS, 'success');

export let paymentStatusValues: Map<string, PaymentStatus> = new Map();
paymentStatusValues.set('failed', PaymentStatus.FAILED);
paymentStatusValues.set('pending', PaymentStatus.PENDING);
paymentStatusValues.set('success', PaymentStatus.SUCCESS);

export let shipmentStatusNames: Map<ShipmentStatus, string> = new Map();
shipmentStatusNames.set(ShipmentStatus.BOOKED, 'booked');
shipmentStatusNames.set(ShipmentStatus.DELIVERED, 'delivered');
shipmentStatusNames.set(ShipmentStatus.DISPATCHED, 'dispatched');
shipmentStatusNames.set(ShipmentStatus.PENDING, 'pending');

export let shipmentStatusValues: Map<string, ShipmentStatus> = new Map();
shipmentStatusValues.set('booked', ShipmentStatus.BOOKED);
shipmentStatusValues.set('delivered', ShipmentStatus.DELIVERED);
shipmentStatusValues.set('dispatched', ShipmentStatus.DISPATCHED);
shipmentStatusValues.set('pending', ShipmentStatus.PENDING);


export const generateOrders = (quantity: number): CreateOrderInput[] => {
	const orders = [];
	for (let i = 0; i < quantity; i++) {
		const shuffledItems = items.sort(() => 0.5 - Math.random());
		const n = Math.floor(Math.random() * 5) + 2;
		const selected = shuffledItems.slice(0, n);
		orders.push(new CreateOrderInput({
			id: `A${i + 1}-${Date.now()}`,
			customerId: '1234',
			items: selected.map((item) => new Item({ ...item, quantity: Math.floor(Math.random() * 3) + 1 }))
		}));
	}
	// For demo purposes, we'll ensure that the first order always
	// contains a single specific item (whose SKU does not trigger
	// item unavailability). This will make it easy to demonstrate
	// the simplest possible case.
	const sortedItems = items.sort(function (a, b) {
		return a.sku < b.sku ? -1 : a.sku > b.sku ? 1 : 0;
	});
	orders[0].items = [
		new Item({
			sku: sortedItems[sortedItems.length - 1].sku,
			description: items[sortedItems.length - 1].description,
			quantity: 1
		})
	];
	return orders;
};

export const order = writable<CreateOrderInput | undefined>();

export const items: Item[] = [
	new Item({
		sku: 'Nike Air Force Ones',
		description:
			'The Nike Air Force Ones combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Adidas UltraBoost',
		description:
			'The Adidas UltraBoost combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Reebok Classic Leather White',
		description:
			'The Reebok Classic Leather combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Puma Suede Classic',
		description:
			'The Puma Suede Classic combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'New Balance 574',
		description:
			'The New Balance 574 combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Vans Old Skool',
		description:
			'The Vans Old Skool combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Converse Chuck Taylor All Star',
		description:
			'The Converse Chuck Taylor All Star combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Under Armour HOVR Sonic',
		description:
			'The Under Armour HOVR Sonic combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Jordan Air Jordan 1',
		description:
			'The Jordan Air Jordan 1 combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Asics GEL-Kayano',
		description:
			'The Asics GEL-Kayano combines timeless style with modern comfort, featuring premium materials and cutting-edge technology for unmatched performance.'
	}),
	new Item({
		sku: 'Nike Air Force Ones',
		description:
			'A second iteration of the classic, the Nike Air Force Ones Model 11 is redesigned for the modern athlete, offering enhanced cushioning and durability.'
	}),
	new Item({
		sku: 'Adidas UltraBoost',
		description:
			'Adidas UltraBoost Model 12 brings you closer to the ground for a more responsive feel, ideal for runners seeking a blend of support and speed.'
	}),
	new Item({
		sku: 'Reebok Classic Leather Black',
		description:
			"Reebok's Classic Leather Model 13 is the epitome of retro chic, offering unparalleled comfort and a sleek design for everyday wear."
	}),
	new Item({
		sku: 'Puma Suede Classic',
		description:
			'The Puma Suede Classic Model 14 updates the iconic design with advanced materials for better wearability and style.'
	}),
	new Item({
		sku: 'New Balance 574',
		description:
			'This latest version of the New Balance 574, Model 15, combines heritage styling with modern technology for an improved fit and function.'
	}),
	new Item({
		sku: 'Vans New Skool',
		description:
			'Vans New Skool Model 16 reintroduces the classic skate shoe with updated features for enhanced performance and comfort.'
	}),
	new Item({
		sku: 'Converse Chuck Taylor All Star',
		description:
			'Model 17 of the Converse Chuck Taylor All Star elevates the iconic silhouette with premium materials and an improved insole for all-day comfort.'
	}),
	new Item({
		sku: 'Under Armour HOVR Sonic',
		description:
			'The latest Under Armour HOVR Sonic, Model 18, offers an unparalleled ride, blending the perfect balance of cushioning and energy return.'
	}),
	new Item({
		sku: 'Jordan Air Jordan 2',
		description:
			'Jordan Air Jordan 2 Model 19 continues the legacy, integrating classic design lines with modern technology for a timeless look and feel.'
	}),
	new Item({
		sku: 'Asics GEL-Tres',
		description:
			'Asics GEL-Tres Model 20 is the latest in the series, offering improved stability and support for overpronators.'
	}),
];
