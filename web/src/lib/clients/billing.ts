import { env } from '$env/dynamic/private';
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-node";
import { Api } from '../../gen/oms/billing/v1/billing_connect';

const transport = createConnectTransport({
  baseUrl: env.BILLING_API_URL, 
  httpVersion: '2',
});

export const client = createClient(Api, transport)
