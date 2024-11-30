<script lang="ts">
	import ItemDetails from './item-details.svelte';
	import ShipmentStatus from './shipment-status.svelte';
	import { shipmentStatusNames } from '$lib/types/order'
	import { Shipment } from '../../gen/oms/v1/types_pb'

	export let shipment: Shipment;
	export let status = shipmentStatusNames.get(shipment.status)
</script>

<div class="details">
	<ShipmentStatus { status } />
	<ItemDetails items={shipment.items} />
</div>

<style>
	.details {
		display: flex;
		flex-direction: column;
		gap: 2rem;
		background-color: white;
		align-items: center;
		width: 100%;
	}
</style>
