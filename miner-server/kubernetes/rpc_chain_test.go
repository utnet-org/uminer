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
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
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

func TestGenerateMinerKeys(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &chainApi.GetMinerKeysRequest{}
	client := chainApi.NewChainServiceClient(conn)

	// Call the RPC method
	var response *chainApi.GetMinerKeysReply
	response, err := client.GetMinerKeys(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)

}

func TestClaimChipComputation(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &chainApi.ClaimChipComputationRequest{
		AccountId: "jackronwong",
		ChipPubK:  "",
		ChipP2K:   "",
		Signature: "CwAAAGphY2tyb253b25nAFk9JSDbTNKCWY4cSFOXJUZJLZ4g/jpKTu5cQMKcmWXegb92CAAAAAALAAAAamFja3JvbndvbmfWmg9s0aOPUaM2HO12t8KVKw+70b6Rm1Fzqpkt7EAWogEAAAAKAniNdNgaiVsdRf7GrQd7CULR6Tk8xe9c7Z5e2whlmQjBbgiMt+s+jBCcGKTFGuLew6zq6I8CSd6iT0EfFlRD3wgngjNKz7ZwgiaF9rFaVuqf6lj01FpMcNaBWmvW5KB8wHLAET7nXrDXLavbx6yK+XDIxf7OuCJo2JJk6/hdM75phBzQTMqqd0UvRo0eHf9zciazwp+tuwJ59kI2Xz6+BcwSpQ70bWmNyVT7bwURob69LIafUzATAldRJRJo+cEm6CJQd9bIRuvgxMlnI+4IIyIbSjDyeriFILllUqMnAIHxGnBYq05TZbe3x1A0kPLfYtCEkqliE0UGBR/uAe24RW5g694uwpfpgPLyzHldXDGxklaB/qj4bppnTjIeQ1LplQIQmLOmLABZPSUg20zSglmOHEhTlyVGSS2eIP46Sk7uXEDCnJll3gAAAAAAJfwHe6dA3utlog6a8UmGqKyEa6QvMhC+4T+OJGQpF5B9h4ZRTKXQsnOumb0JNVyJZd7UT7c+DtiRwBEGx2+4AA==",
	}
	client := chainApi.NewChainServiceClient(conn)

	// Call the RPC method
	var response *chainApi.ClaimChipComputationReply
	response, err := client.ClaimChipComputation(context.Background(), request, grpc.WaitForReady(true)) //txHash:"ASvFKz16BBW6J7pqF1q6kzRkS1fTfWFxfSAmo1CwDUBz"
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)

}
