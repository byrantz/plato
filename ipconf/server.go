package ipconf

import (
	"github.com/byrantz/plato/common/config"
	"github.com/byrantz/plato/ipconf/domain"
	"github.com/byrantz/plato/ipconf/source"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// RunMain 启动web容器
func RunMain(path string) {
	config.Init(path)
	source.Init() // 数据源要优先启动
	domain.Init() // 初始化调度层
	// 框架 hertz
	s := server.Default(server.WithHostPorts(":6789"))
	s.GET("/ip/list", GetIpInfoList)
	s.Spin()
}
