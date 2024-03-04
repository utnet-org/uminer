package kubernetes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	chainApi "uminer/miner-server/api/chainApi/rpc"
	"uminer/miner-server/cmd"
)

func TestUpdateChainStatusGRPC(t *testing.T) {
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

func TestClaimStake(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &chainApi.ClaimStakeRequest{
		AccountId: "jackronwong",
		Amount:    "100000000",
		NearPath:  "/Users/mac/sandbox/utnet/utility-cli-rs/target/debug/near",
		KeyPath:   "/Users/mac/sandbox/utnet/uminer/miner-server/cmd/miner/validator_key.json",
	}
	client := chainApi.NewChainServiceClient(conn)

	// Call the RPC method
	var response *chainApi.ClaimStakeReply
	response, err := client.ClaimStake(context.Background(), request, grpc.WaitForReady(true)) //txHash:"ASvFKz16BBW6J7pqF1q6kzRkS1fTfWFxfSAmo1CwDUBz"
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
		NearPath:  "/Users/mac/sandbox/utnet/utility-cli-rs/target/debug/near",
		KeyPath:   "/Users/mac/sandbox/utnet/uminer/miner-server/cmd/miner/miner_key.json",
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
