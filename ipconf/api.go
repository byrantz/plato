package ipconf

import (
	"context"

	"github.com/byrantz/plato/ipconf/domain"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// GetIpInfoList API 适配应用层
func GetIpInfoList(c context.Context, ctx *app.RequestContext) {
	// 如果在 GetIpInfoList 函数中发生了 panic，我们可以通过 recover 来捕获这个 panic，并返回一个错误信息
	// 这么做的好处是即使在 GetIpInfoList 函数中发生了 panic，我们也可以保证程序不会崩溃
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(consts.StatusBadRequest, utils.H{"err": err})
		}
	}()
	// Step0 : 构建客户请求信息
	// ipConfCtx 是一个结构体，它包含了客户端请求的所有信息，可以在协程之间共享
	ipConfCtx := domain.BuildConfContext(&c, ctx)
	// Step1 : 进行 ip 调度
	eds := domain.Dispatch(ipConfCtx)
	// step2: 根据得分取top5返回
	ipConfCtx.AppCtx.JSON(consts.StatusOK, packRes(top5Endports(eds)))
}
