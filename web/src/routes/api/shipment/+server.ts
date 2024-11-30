import { json, type RequestHandler } from '@sveltejs/kit';
import { client } from '$lib/clients/shipment';
import { ShipmentStatus } from '../../../gen/oms/v1/types_pb';

export const POST: RequestHandler = async ({ request }) => {
	const { signal, shipment } = await request.json();

	try {
		let s = ShipmentStatus.UNSPECIFIED;
		if (signal.status == 'dispatched') {
			s = ShipmentStatus.DISPATCHED;
		} else if (signal.status == 'delivered') {
			s = ShipmentStatus.DELIVERED;
		}
		const response = await client.updateShipmentStatus({shipmentId: shipment.id, status: s})
		return json({ status: 'ok', body: response });
	} catch (e) {
		return json({ status: 'error' });
	}
};
