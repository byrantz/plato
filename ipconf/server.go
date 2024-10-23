package ipconf

import (
	"github.com/byrantz/plato/common/config"
	"github.com/byrantz/plato/ipconf/domain"
	"github.com/byrantz/plato/ipconf/source"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// RunMain 启动web容器
// path 是命令行传递的
func RunMain(path string) {
	config.Init(path)
	source.Init() // 数据源要优先启动, 在此处更新候选
	domain.Init() // 初始化调度层
	// 框架 hertz
	s := server.Default(server.WithHostPorts(":6789"))
	s.GET("/ip/list", GetIpInfoList)
	// 服务器将开始监听再 ：6789 端口上的请求，并处理这些请求
	s.Spin()
}
