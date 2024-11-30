import { json, type RequestHandler } from '@sveltejs/kit';
import { client } from '$lib/clients/orders';
import { CustomerAction } from '../../../gen/oms/v1/types_pb';

export const POST: RequestHandler = async ({ request }) => {
	const { id, action } = await request.json();
	let a = CustomerAction.UNSPECIFIED
	if (action == 'amend') {
		a = CustomerAction.AMEND
	} else if (action == 'cancel') {
		a = CustomerAction.CANCEL
	}
	try {
		console.log('sending action: ' + action + ' (' + a.toString() + ')')
		const response = await client.customerAction({id: id, action: a })
		return json({ status: 'ok', body: response });
	} catch (e) {
		return json({ status: 'error' });
	}
};
