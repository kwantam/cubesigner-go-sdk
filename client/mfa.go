package client

import (
	"fmt"
	"net/url"

	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/pkg/errors"
)

// Retrieves and returns all pending MFA requests that are accessible to the
// current user, i.e., those in which the current user is listed as an approver.
// No pagination.
func (cli *Client) ListMfaRequests() (*v0.ListMfaResponse, error) {
	response, err := cli.get("/v0/org/:org_id/mfa", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request ListMfaRequests")
	}
	decoded, err := decodeJSONResponse[v0.ListMfaResponse](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}

// Retrieves and returns a pending MFA request by its id.
func (cli *Client) GetMfaRequest(mfaId string) (*v0.MfaRequestInfo, error) {
	response, err := cli.get(fmt.Sprintf("/v0/org/:org_id/mfa/%s", url.PathEscape(mfaId)), nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request GetMfaRequest")
	}
	decoded, err := decodeJSONResponse[v0.MfaRequestInfo](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, nil
}

// Approve or Reject MFA Request after logging in with CubeSigner.
// If approving, adds the currently-logged user as an approver of a pending MFA
// request of the [Status::RequiredApprovers] kind. If the required number of
// approvers is reached, the MFA request is approved; the confirmation receipt
// can be used to resume the original HTTP request.
// If rejecting, immediately deletes the pending MFA request.
func (cli *Client) ApproveOrRejectMfaRequest(mfaId string, mfaVote v0.MfaVote) (*v0.MfaRequestInfo, error) {
	var voteQuery string
	switch mfaVote {
	case v0.Approve:
		voteQuery = "approve"
	case v0.Reject:
		voteQuery = "reject"
	default:
		return nil, errors.New("invalid MfaVote value")
	}

	response, _, err := cli.patch(fmt.Sprintf("/v0/org/:org_id/mfa/%s?mfa_vote=%s", url.PathEscape(mfaId), voteQuery), nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request ApproveOrRejectMfaRequest")
	}

	decoded, err := decodeJSONResponse[v0.MfaRequestInfo](response)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return &decoded, err
}
