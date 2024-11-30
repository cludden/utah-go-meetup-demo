<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { invalidate } from '$app/navigation';
	import { orderStatusNames } from '$lib/types/order'
	import { Order, OrderStatus } from '../../../../gen/oms/v1/types_pb'
	import FulfillmentDetails from '$lib/components/fulfillment-details.svelte';
	import OrderActions from '$lib/components/order-actions.svelte';
	import StatusBadge from '$lib/components/status-badge.svelte';

	let order: Order;

	$: order = new Order().fromJson($page.data.order)

	onMount(() => {
		const finalStatuses = [OrderStatus.COMPLETED, OrderStatus.FAILED,  OrderStatus.CANCELLED];

		const interval = setInterval(() => {
			const isFinal = finalStatuses.includes(order.status);
			if (!isFinal) {
				invalidate('data:order');
			} else {
				clearInterval(interval);
			}
		}, 500);

		return () => {
			clearInterval(interval);
		};
	});
</script>

<section>
	<div class="container">
		<div class="order">
			<h1>{$page.params.id}</h1>
			<div class="status">
				<StatusBadge status={orderStatusNames.get(order.status) ?? 'unspecified'}></StatusBadge>
			</div>
		</div>
		
		
		<FulfillmentDetails {order} />
		{#if order.status === OrderStatus.CUSTOMER_ACTION_REQUIRED}
			<OrderActions id={$page.params.id} />
		{/if}
	</div>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.order {
		display: flex;
		align-items: start;
		justify-content: space-between;
		width: 100%;
	}

	.status {
		display: flex;
		align-items: end;
		justify-content: space-between;
	}

	.container {
		display: flex;
		flex-direction: column;
		align-items: start;
		gap: 2rem;
		background-color: white;
		padding: 2rem;
		border-radius: 0.5rem;
	}
</style>
