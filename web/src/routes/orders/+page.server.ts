import { client } from '$lib/clients/orders';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const { orders } = await client.listOrders({})
	return {
		orders: orders.map((o) => o.toJson())
	};
};
