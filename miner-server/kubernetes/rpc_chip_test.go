package kubernetes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	chipApi "uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/cmd"
)

func TestQueryAllChipsGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer()
	defer conn.Close()

	// Prepare the request
	request := &chipApi.ChipsRequest{
		SerialNum: "",
		BusId:     "",
	}
	client := chipApi.NewChipServiceClient(conn)

	// Call the RPC method
	var response *chipApi.ListChipsReply
	response, err := client.ListAllChips(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}
