<script lang="ts">
	import { goto } from '$app/navigation';
	import StatusBadge from '$lib/components/status-badge.svelte';
	import { orderStatusNames } from '$lib/types/order'
	import { Order } from '../../gen/oms/v1/types_pb'

	export let data;

	$: ({ orders } = { orders: data.orders.map(o => new Order().fromJson(o) )});
</script>

<svelte:head>
	<title>OMS</title>
	<meta name="description" content="OMS App" />
</svelte:head>

<section>
	<nav>
		<h1>Orders</h1>
		<button on:click={() => goto('/orders/new')}>New Order</button>
	</nav>
	<table>
		<thead>
			<tr>
				<th>ID</th>
				<th style="text-align: center;">Date & Time</th>
				<th style="text-align: center;">Status</th>
			</tr>
		</thead>
		<tbody>
			{#each orders as order}
				<tr>
					<td style="width: 100%;"><a href={`/orders/${order.id}`}>{order.id}</a></td>
					<td style="text-align: center;">
						<div style="width: 210px;">
							{#if order.receivedAt != undefined}
								{order.receivedAt.toDate().toLocaleDateString()}
								{order.receivedAt.toDate().toLocaleTimeString()}
							{/if}
						</div></td>
					<td style="text-align: center;"><StatusBadge status={orderStatusNames.get(order.status) ?? 'unspecified'} /></td>
				</tr>
			{:else}
				<tr>
					<td>No Active Orders</td>
					<td />
					<td />
				</tr>
			{/each}
		</tbody>
	</table>
</section>
