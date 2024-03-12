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
	conn := cmd.ConnectRPCServer(cmd.WorkerServerIP, "7001")
	defer conn.Close()

	// Prepare the request
	request := &chipApi.ChipsRequest{
		Url:       "http://119.120.92.239:30345",
		SerialNum: "",
		BusId:     "",
		DevId:     "",
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

func TestReadChipKeyPairsGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.WorkerServerIP, "7001")
	defer conn.Close()

	// Prepare the request
	request := &chipApi.ChipsRequest{
		Url:       "http://119.120.92.239:30345",
		SerialNum: "",
		BusId:     "",
		DevId:     "0",
	}
	client := chipApi.NewChipServiceClient(conn)

	// Call the RPC method
	var response *chipApi.ReadChipReply
	response, err := client.ObtainChipKeyPairs(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

func TestSignChipsGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.WorkerServerIP, "7001")
	defer conn.Close()

	// Prepare the request
	request := &chipApi.SignChipsRequest{
		//Url:           "http://119.120.92.239:30345",
		//SerialNum:     "HQDZKC5BAAABJ0309",
		//BusId:         "000:D9:00.1",
		DevId:         "0",
		P2:            "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",
		PublicKey:     "4e1BUTgGBfqVXqrJHJi28YVj6qv6ma9G3FhoVoGqk5V3ZjTtzCYN2ahWGBbZHcWRGhCtyVUW1SV6VJQ5ckXB4fBc5ajDE4xERbqE2w4Cnok1RwrzeWfhQc35ha7z472vyoMds2uD1G6aTqk692rYQaCrr7Buc8M23sffqHZsy2V9SE1ggnAdgwnZXbwf92icSprUBGckaNtrftLcwSUYARBjARkd32bFLGt5HbK4J3CyyMUXBTbJbtK3rbnXmAAyR4MvD9dEKM1a1hiYf3mXES3mXozxzZqJsb9vdoY4vgBcrpHV1EX5rWdABGEWVqAPCFC95BcoJrkWKxfUCvunMeJZEGYpyP1WkTaf8dJC6xjrmAzUQuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu",
		P2Size:        425,
		PublicKeySize: 1680,
		Msg:           "test",
	}
	client := chipApi.NewChipServiceClient(conn)

	// Call the RPC method
	var response *chipApi.SignChipsReply
	response, err := client.SignChip(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}
