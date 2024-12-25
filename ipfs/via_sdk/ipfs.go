package main

import (
	"context"
	"fmt"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
)

func main() {
	// 创建一个与本地IPFS节点的RPC接口的连接。如果连接失败，会打印错误信息并退出程序。
	node, err := rpc.NewLocalApi()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	// 将一个CID字符串解码为CID对象。如果解码失败，会打印错误信息并退出程序。
	c, err := cid.Decode("QmTyEVbNyWQeA4ZLov246YTwnzsNraLmhnrF8ewhwGH4Hu")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 使用path.FromCid函数将CID转换为路径对象。
	p := path.FromCid(c)
	// 使用节点的Pin API将文件固定到本地节点。如果固定操作失败，会打印错误信息并退出程序。
	err = node.Pin().Add(ctx, p)
	if err != nil {
		fmt.Println(err)
		return
	}
}
