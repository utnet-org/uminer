package types

import (
	"context"
	"encoding/json"
	"errors"
	"uminer/common/log"
	"uminer/miner-server/api/containerApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
)

type ImageService struct {
	containerApi.UnimplementedImageServiceServer
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewImageService(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) containerApi.ImageServiceServer {
	return &ImageService{
		conf: conf,
		log:  log.NewHelper("ImageService", logger),
		data: data,
	}
}

// CreateImage create image by miner after receiving request from renter
func (s *ImageService) CreateImage(ctx context.Context, req *containerApi.CreateImageRequest) (*containerApi.CreateImageReply, error) {

	reply := &containerApi.CreateImageReply{
		CreatedAt: 0,
		ImageId:   "",
		Status:    false,
	}

	// http addr request
	requestUrl := mainURL + "/v1/imagemanage/image"
	jsonData := map[string]interface{}{
		"imageAddr":    req.ImageAddr,
		"imageDesc":    req.ImageDesc,
		"imageName":    req.ImageName,
		"imageVersion": req.ImageVersion,
		"sourceType":   req.SourceType,
	}
	resp := HTTPRequest("POST", requestUrl, jsonData, "application/json", req.Token)

	var response struct {
		Success bool `json:"success"`
		Payload struct {
			ImageId   string `json:"imageId"`
			CreatedAt int64  `json:"createdAt"`
		} `json:"payload"`
		Error interface{} `json:"error"`
	}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		return reply, err
	}

	switch errObj := response.Error.(type) {
	case map[string]interface{}:
		// 转换为 map 类型成功，可以提取目标字段的值
		message, ok := errObj["message"].(string)
		if !ok {
			return reply, errors.New("error message not found")
		} else {
			return reply, errors.New(message)
		}
	default:

	}

	reply.ImageId = response.Payload.ImageId
	reply.CreatedAt = response.Payload.CreatedAt
	reply.Status = true
	return reply, nil
}

// DeleteImage create Image by miner after receiving request from renter
func (s *ImageService) DeleteImage(ctx context.Context, req *containerApi.DeleteImageRequest) (*containerApi.DeleteImageReply, error) {

	reply := &containerApi.DeleteImageReply{
		DeletedAt: 0,
		Status:    false,
	}

	// http addr request
	requestUrl := mainURL + "/v1/imagemanage/image/" + req.ImageId
	jsonData := map[string]interface{}{}
	resp := HTTPRequest("DELETE", requestUrl, jsonData, "application/json", req.Token)

	var response struct {
		Success bool `json:"success"`
		Payload struct {
			DeletedAt int64 `json:"deletedAt"`
		} `json:"payload"`
		Error interface{} `json:"error"`
	}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		return reply, err
	}

	switch errObj := response.Error.(type) {
	case map[string]interface{}:
		// 转换为 map 类型成功，可以提取目标字段的值
		message, ok := errObj["message"].(string)
		if !ok {
			return reply, errors.New("error message not found")
		} else {
			return reply, errors.New(message)
		}
	default:

	}

	reply.DeletedAt = response.Payload.DeletedAt
	reply.Status = true
	return reply, nil
}
