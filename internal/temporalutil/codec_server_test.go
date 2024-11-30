package temporalutil

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cludden/protoc-gen-go-temporal/pkg/codec"
	"github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	"github.com/stretchr/testify/require"
	billingv1 "github.com/temporalio/reference-app-orders-go/gen/oms/billing/v1"
	orderv1 "github.com/temporalio/reference-app-orders-go/gen/oms/order/v1"
	shipmentv1 "github.com/temporalio/reference-app-orders-go/gen/oms/shipment/v1"
	"go.temporal.io/sdk/converter"
)

type decoded struct {
	Payloads []struct {
		Metadata map[string]any `json:"metadata"`
		Data     string         `json:"data"`
	} `json:"payloads"`
}

func TestCodecServer(t *testing.T) {
	srv := httptest.NewServer(converter.NewPayloadCodecHTTPHandler(
		&Codec{},
		codec.NewJSONCodec(scheme.New(
			billingv1.WithWorkerSchemeTypes(),
			orderv1.WithWorkerSchemeTypes(),
			shipmentv1.WithWorkerSchemeTypes(),
		)),
	))
	t.Cleanup(srv.Close)

	resp, err := http.Post(srv.URL+"/decode", "application/json", strings.NewReader(`{
	"payloads": [
		{
			"metadata": {
				"encoding": "YmluYXJ5L2VuY3J5cHRlZA==",
				"encryption-key-id": "dXRhaGdv"
			},
			"data": "0jwCD0P7kXQxfSUDlpz3N8BMXaipcOSDy6AGF8l8GmKmQ+Bp80hCDGUuOCothG84CvsDFlijGroGpbcSbKbJXS2DmfWqT5lrBbewqqz//qbzyjhDNaxWzO7O4GIsfI6aM7ziKe30uIA+rDXX3nrMHTrngHCpSS/CD91fvSrMpRRixiXGvkGQxAcKsHzkb8rs/ZldTFlQbdzSRwBK7knQvY6dzMvaZGhosmuuyUDTU43WuV8E7sfqVR0SUCDL26Se+woOo5ZtGw43MiSX5AQqnvucjT9h"
		}
	]
}`))
	require.NoError(t, err)
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	resp.Body.Close()

	var d decoded
	require.NoError(t, json.Unmarshal(b, &d))
}
