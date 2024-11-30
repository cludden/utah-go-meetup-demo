import { json, type RequestHandler } from '@sveltejs/kit';
import { client } from '$lib/clients/orders';

export const POST: RequestHandler = async ({ request }) => {
	const order = await request.json();
	try {
		const response = await client.createOrder(order)
		return json({ status: 'ok', body: response });
	} catch (e) {
		return json({ status: 'error' });
	}
};
