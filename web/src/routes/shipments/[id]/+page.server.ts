import { client } from '$lib/clients/shipment';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (request) => {
	const { id } = request.params;

	const { shipment } = await client.getShipment({ id })

	return { shipment: shipment?.toJson() ?? null };
};
