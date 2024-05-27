package domain

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type IpConfConext struct {
	Ctx       *context.Context
	AppCtx    *app.RequestContext
	ClientCtx *ClientConext
}

type ClientConext struct {
	IP string `json:"ip"`
}

func BuildConfContext(c *context.Context, ctx *app.RequestContext) *IpConfConext {
	IpConfConext := &IpConfConext{
		Ctx:       c,
		AppCtx:    ctx,
		ClientCtx: &ClientConext{},
	}
	return IpConfConext
}
