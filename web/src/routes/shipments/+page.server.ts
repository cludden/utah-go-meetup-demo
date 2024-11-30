import { client} from '$lib/clients/shipment'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ depends }) => {
	depends('data:shipments')
	const { shipments } = await client.listShipments({})
	return {
		shipments: shipments.map((s) => s.toJson())
	}
}
