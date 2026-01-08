package emailengine

import (
	"context"

	"github.com/stax-ai/emailengine/api"
)

func (c *Client) ListAccounts(ctx context.Context, page, pageSize int32, opts ...func(*api.ApiGetV1AccountsRequest) api.ApiGetV1AccountsRequest) (*api.AccountsFilterResponse, error) {
	req := c.api.AccountAPI.GetV1Accounts(ctx).Page(page).PageSize(pageSize)
	for _, o := range opts {
		req = o(&req)
	}
	res, httpRes, err := c.api.AccountAPI.GetV1AccountsExecute(req)
	if err != nil {
		if httpRes != nil && httpRes.StatusCode >= 400 {
			return nil, decodeError("ListAccounts", httpRes)
		}
		return nil, err
	}
	return res, nil
}
