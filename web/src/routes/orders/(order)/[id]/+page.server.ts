import { client } from '$lib/clients/orders';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, depends }) => {
	depends('data:order')
	const { id } = params;

	const order = await client.getOrder({id});

	return { order: order.order?.toJson() };
};
