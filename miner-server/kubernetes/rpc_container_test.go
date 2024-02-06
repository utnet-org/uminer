package kubernetes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	"uminer/miner-server/api/containerApi"
	"uminer/miner-server/cmd"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJlM2IxYzE0ZDBiM2M0NGZkYWU2NzEyZGRjYjE3NjU0MyIsImNyZWF0ZWRBdCI6MTcwNzIxMjMyMCwiZXhwIjoxNzA3Mjk4NzIwLCJpYXQiOjE3MDcyMTIzMjB9.gDRURqM5jgFDxFePkOw_YET5hDWzOi1FWzg-3yH6Mt8"

/* exemplaire */
/* image */

func TestCreateImageGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.CreateImageRequest{
		Token:        token,
		ImageAddr:    "example",
		ImageDesc:    "miner test",
		ImageName:    "image1",
		ImageVersion: "V1",
		SourceType:   2,
	}
	client := containerApi.NewImageServiceClient(conn)

	// Call the RPC method
	var response *containerApi.CreateImageReply
	response, err := client.CreateImage(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

func TestDeleteImageGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.DeleteImageRequest{
		Token:   token,
		ImageId: "0707c1f2c09843768681d869e013f669",
	}
	client := containerApi.NewImageServiceClient(conn)

	// Call the RPC method
	var response *containerApi.DeleteImageReply
	response, err := client.DeleteImage(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

func TestQueryImageGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.QueryImageByConditionRequest{
		Token:     token,
		Id:        "e6805f27a79a43f793d93757ba70d03a",
		PageSize:  10,
		PageIndex: 1,
	}
	client := containerApi.NewImageServiceClient(conn)

	// Call the RPC method
	var response *containerApi.QueryImageByConditionReply
	response, err := client.QueryImageByCondition(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

/* notebook */

func TestCreateNoteBookGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.CreateNoteBookRequest{
		Token:          token,
		Name:           "notebook-20240123-uy01",
		Description:    "miner rpc test",
		AlgorithmId:    "6c61e286272a4c5f92a1efe9b45cd714",
		ImageId:        "1e67cad5c01042a6b64eff4437bc21f8",
		ResourceSpecId: "639ecfb98cdf4e54860f6fb2b0a7b65f",
	}
	client := containerApi.NewNotebookServiceClient(conn)

	// Call the RPC method
	var response *containerApi.CreateNoteBookReply
	response, err := client.CreateNotebook(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

func TestDeleteNoteBookGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.DeleteNotebookRequest{
		Token: token,
		Id:    "s24876b7bf6e4788b2412c7f28c72cc2",
	}
	client := containerApi.NewNotebookServiceClient(conn)

	// Call the RPC method
	var response *containerApi.DeleteNotebookReply
	response, err := client.DeleteNotebook(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

func TestStartStopNoteBookGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.StartStopNotebookRequest{
		Token: token,
		Id:    "s24876b7bf6e4788b2412c7f28c72cc2",
	}
	client := containerApi.NewNotebookServiceClient(conn)

	// Call the RPC method
	var response *containerApi.StartStopNotebookReply
	response, err := client.StopNotebook(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

func TestQueryNoteBookGRPC(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.QueryNotebookByConditionRequest{
		Token:     token,
		Id:        "", //s24876b7bf6e4788b2412c7f28c72cc2",
		PageSize:  10,
		PageIndex: 1,
	}
	client := containerApi.NewNotebookServiceClient(conn)

	// Call the RPC method
	var response *containerApi.QueryNotebookByConditionReply
	response, err := client.QueryNotebookByCondition(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

func TestQueryNotebookEventRecord(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.QueryNotebookEventRecordRequest{
		Token:      token,
		NotebookId: "ge7b4a7225464d81a240eed8d38a8b53",
		PageSize:   10,
		PageIndex:  1,
	}
	client := containerApi.NewNotebookServiceClient(conn)

	// Call the RPC method
	var response *containerApi.QueryNotebookEventRecordReply
	response, err := client.QueryNotebookEventRecord(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}

func TestObtainNotebookEvents(t *testing.T) {
	// Connect to the RPC server
	conn := cmd.ConnectRPCServer(cmd.MinerServerIP, "9001")
	defer conn.Close()

	// Prepare the request
	request := &containerApi.ObtainNotebookEventRequest{
		Token:         token,
		NotebookJobId: "fb15ec98802d4858ad3a1a197b3bb262",
		PageSize:      10,
		PageIndex:     1,
		TaskIndex:     1,
		ReplicaIndex:  1,
	}
	client := containerApi.NewNotebookServiceClient(conn)

	// Call the RPC method
	var response *containerApi.ObtainNotebookEventReply
	response, err := client.ObtainNotebookEvent(context.Background(), request, grpc.WaitForReady(true))
	if err != nil {
		fmt.Println("调用 gRPC 方法失败:", err)
		return
	}

	// 处理响应
	fmt.Println("gRPC Response: ", response)
}
