package emailengine

import (
	"context"

	"github.com/stax-ai/emailengine/api"
)

func (c *Client) SendMessage(ctx context.Context, account string, msg api.SubmitMessage, opts ...func(*api.ApiPostV1AccountAccountSubmitRequest) api.ApiPostV1AccountAccountSubmitRequest) (*api.SubmitMessageResponse, error) {
	req := c.api.SubmitAPI.PostV1AccountAccountSubmit(ctx, account).SubmitMessage(msg)
	for _, o := range opts {
		req = o(&req)
	}
	res, httpRes, err := c.api.SubmitAPI.PostV1AccountAccountSubmitExecute(req)
	if err != nil {
		if httpRes != nil && httpRes.StatusCode >= 400 {
			return nil, decodeError("SendMessage", httpRes)
		}
		return nil, err
	}
	return res, nil
}
