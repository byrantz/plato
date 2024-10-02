package source

import (
	"context"

	"github.com/byrantz/plato/common/config"
	"github.com/byrantz/plato/common/discovery"
	"github.com/bytedance/gopkg/util/logger"
)

func Init() {
	eventChan = make(chan *Event)
	// 创建一个空的上下文，传递给 DataHandler
	ctx := context.Background()
	go DataHandler(&ctx)
	if config.IsDebug() {
		ctx := context.Background()
		// 代替 ip 网关，做 mock 数据
		testServiceRegister(&ctx, "7896", "node1")
		testServiceRegister(&ctx, "7897", "node2")
		testServiceRegister(&ctx, "7898", "node3")
	}
}

// 服务发现处理
// DataHandler 函数通过服务发现机制监听服务的变化，并根据变化触发相应的事件。
// 它定义了两个回调函数 setFunc 和 delFunc 来处理新增和删除节点的情况，并将事件发送到事件通道。
func DataHandler(ctx *context.Context) {
	// dis 创建服务发现监听器
	dis := discovery.NewServiceDiscovery(ctx)
	defer dis.Close()
	// 定义处理新增节点的函数
	setFunc := func(key, value string) {
		if ed, err := discovery.UnMarshal([]byte(value)); err == nil {
			if event := NewEvent(ed); ed != nil {
				event.Type = AddNodeEvent
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.setFunc.err :%s", err.Error())
		}
	}
	// 定义处理删除节点的函数
	delFunc := func(key, value string) {
		if ed, err := discovery.UnMarshal([]byte(value)); err == nil {
			if event := NewEvent(ed); ed != nil {
				event.Type = DelNodeEvent
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.delFunc.err :%s", err.Error())
		}
	}
	// 开始监听服务变化
	err := dis.WatchService(config.GetServicePathForIPConf(), setFunc, delFunc)
	if err != nil {
		panic(err)
	}
}
