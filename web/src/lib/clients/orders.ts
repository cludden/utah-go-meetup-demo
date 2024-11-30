import { env } from '$env/dynamic/private';
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-node";
import { Api } from '../../gen/oms/order/v1/orders_connect';

const transport = createConnectTransport({
  baseUrl: env.ORDER_API_URL, 
  httpVersion: '2',
});

export const client = createClient(Api, transport)
