package source

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/byrantz/plato/common/config"
	"github.com/byrantz/plato/common/discovery"
)

//	func init() {
//		ctx := context.Background()
//		testServiceRegister(&ctx, "7896", "node1")
//		testServiceRegister(&ctx, "7897", "node2")
//		testServiceRegister(&ctx, "7898", "node3")
//	}
func testServiceRegister(ctx *context.Context, port, node string) {
	// 模拟服务发现
	go func() {
		ed := discovery.EndpointInfo{
			IP:   "127.0.0.1",
			Port: port,
			MetaData: map[string]interface{}{
				"connect_num":   float64(rand.Int63n(12312321231231131)),
				"message_bytes": float64(rand.Int63n(1231232131556)),
			},
		}
		sr, err := discovery.NewServiceRegister(ctx, fmt.Sprintf("%s%s", config.GetServicePathForIPConf(), node), &ed, time.Now().Unix())
		if err != nil {
			panic(err)
		}
		go sr.ListenLeaseRespChan()
		// 无线循环，不断更新服务信息
		for {
			// 模拟服务信息变化，更新 EndpointInfo 实例，每秒更新一次
			ed = discovery.EndpointInfo{
				IP:   "127.0.0.1",
				Port: port,
				MetaData: map[string]interface{}{
					"connect_num":   float64(rand.Int63n(12312321231231131)),
					"message_bytes": float64(rand.Int63n(1231232131556)),
				},
			}
			sr.UpdateValue(&ed)
			time.Sleep(1 * time.Second)
		}
	}()
}
