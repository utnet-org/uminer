package kubernetes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	chainApi "uminer/miner-server/api/chainApi/rpc"
	"uminer/miner-server/cmd"
)

func TestUpateChainStatusGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer()
	defer conn.Close()

	// Prepare the request
	request := &chainApi.ReportChainsStatusRequest{}
	client := chainApi.NewChainServiceClient(conn)

	// Call the RPC method
	var response *chainApi.ReportChainsStatusReply
	response, err := client.UpdateChainsStatus(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}
