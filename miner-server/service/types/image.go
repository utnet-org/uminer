package types

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"uminer/common/log"
	"uminer/miner-server/api/containerApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service/connect"
)

type Image struct {
	CreatedAt     int64  `json:"createdAt"`
	ID            string `json:"id"`
	ImageAddr     string `json:"imageAddr"`
	ImageDesc     string `json:"imageDesc"`
	ImageFullAddr string `json:"imageFullAddr"`
	ImageName     string `json:"imageName"`
	ImageStatus   int    `json:"imageStatus"`
	ImageVersion  string `json:"imageVersion"`
	SourceType    int    `json:"sourceType"`
	SpaceID       string `json:"spaceId"`
	UpdatedAt     int64  `json:"updatedAt"`
	UserID        string `json:"userId"`
	Username      string `json:"username"`
}

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
	requestUrl := connect.MainURL + "/v1/imagemanage/image"
	jsonData := map[string]interface{}{
		"imageAddr":    req.ImageAddr,
		"imageDesc":    req.ImageDesc,
		"imageName":    req.ImageName,
		"imageVersion": req.ImageVersion,
		"sourceType":   req.SourceType,
	}
	resp := connect.HTTPRequest("POST", requestUrl, jsonData, "application/json", req.Token)

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
	requestUrl := connect.MainURL + "/v1/imagemanage/image/" + req.ImageId
	jsonData := map[string]interface{}{}
	resp := connect.HTTPRequest("DELETE", requestUrl, jsonData, "application/json", req.Token)

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

// QueryImageByCondition query image by miner
func (s *ImageService) QueryImageByCondition(ctx context.Context, req *containerApi.QueryImageByConditionRequest) (*containerApi.QueryImageByConditionReply, error) {
	list := make([]*containerApi.ImageList, 0)
	// http addr request
	requestUrl := connect.MainURL + "/v1/imagemanage/image/" + req.Id
	if req.Id == "" {
		requestUrl = connect.MainURL + "/v1/imagemanage/userimage" + "?pageSize=" + strconv.FormatInt(req.PageSize, 10) + "&pageIndex=" + strconv.FormatInt(req.PageIndex, 10)
		jsonData := map[string]interface{}{
			"pageSize":  req.PageSize,
			"pageIndex": req.PageIndex,
		}
		resp := connect.HTTPRequest("GET", requestUrl, jsonData, "application/json", req.Token)

		type imageStru struct {
			Image    Image `json:"image"`
			IsShared bool  `json:"isShared"`
		}
		type Payload struct {
			TotalSize int64       `json:"totalSize"`
			Images    []imageStru `json:"images"`
		}
		var response struct {
			Success bool        `json:"success"`
			Payload Payload     `json:"payload"`
			Error   interface{} `json:"error"`
		}
		err := json.Unmarshal(resp, &response)
		if err != nil {
			reply := &containerApi.QueryImageByConditionReply{ImageList: list}
			return reply, err
		}
		//fmt.Println(string(resp))

		switch errObj := response.Error.(type) {
		case map[string]interface{}:
			// 转换为 map 类型成功，可以提取目标字段的值
			message, ok := errObj["message"].(string)
			reply := &containerApi.QueryImageByConditionReply{ImageList: list}
			if !ok {
				return reply, errors.New("error message not found")
			} else {
				return reply, errors.New(message)
			}
		default:

		}

		for _, item := range response.Payload.Images {
			list = append(list, &containerApi.ImageList{
				UserName:     item.Image.Username,
				UserId:       item.Image.UserID,
				ImageId:      item.Image.ID,
				ImageAddr:    item.Image.ImageAddr,
				ImageDesc:    item.Image.ImageDesc,
				ImageName:    item.Image.ImageName,
				ImageVersion: item.Image.ImageVersion,
				ImageStatus:  int64(item.Image.ImageStatus),
				SourceType:   int64(item.Image.SourceType),
				SpaceId:      item.Image.SpaceID,
				CreatedAt:    item.Image.CreatedAt,
				UpdatedAt:    item.Image.UpdatedAt,
			})
		}
		reply := &containerApi.QueryImageByConditionReply{ImageList: list}

		return reply, nil

	}

	jsonData := map[string]interface{}{
		"id": req.Id,
	}
	resp := connect.HTTPRequest("GET", requestUrl, jsonData, "application/json", req.Token)

	type Payload struct {
		Image Image `json:"image"`
	}
	var response struct {
		Success bool        `json:"success"`
		Payload Payload     `json:"payload"`
		Error   interface{} `json:"error"`
	}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		reply := &containerApi.QueryImageByConditionReply{ImageList: list}
		return reply, err
	}

	list = append(list, &containerApi.ImageList{
		UserName:     response.Payload.Image.Username,
		UserId:       response.Payload.Image.UserID,
		ImageId:      response.Payload.Image.ID,
		ImageAddr:    response.Payload.Image.ImageAddr,
		ImageDesc:    response.Payload.Image.ImageDesc,
		ImageName:    response.Payload.Image.ImageName,
		ImageVersion: response.Payload.Image.ImageVersion,
		ImageStatus:  int64(response.Payload.Image.ImageStatus),
		SourceType:   int64(response.Payload.Image.SourceType),
		SpaceId:      response.Payload.Image.SpaceID,
		CreatedAt:    response.Payload.Image.CreatedAt,
		UpdatedAt:    response.Payload.Image.UpdatedAt,
	})
	reply := &containerApi.QueryImageByConditionReply{ImageList: list}

	return reply, nil
}
