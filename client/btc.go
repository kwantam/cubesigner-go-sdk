package client

import (
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

func (cli *Client) SignTaproot(pubkey string, request *v0.TaprootSignRequest) (*v0.TaprootSignResponse, error) {
	encoded, err := encodeJSONRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	// replace path variables
	endpoint := strings.Replace("/v0/org/:org_id/btc/taproot/sign/:pubkey", ":pubkey", url.PathEscape(pubkey), -1)

	response, err := cli.post(endpoint, encoded, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request SignTaproot")
	}
	decoded, err := decodeJSONResponse[v0.TaprootSignResponse](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}