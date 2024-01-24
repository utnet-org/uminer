package types

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"uminer/common/log"
	"uminer/miner-server/api/containerApi"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
)

type Task struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}
type Notebook struct {
	CreatedAt         int64   `json:"createdAt"`
	UpdatedAt         int64   `json:"updatedAt"`
	ID                string  `json:"id"`
	UserID            string  `json:"userId"`
	WorkspaceID       string  `json:"workspaceId"`
	Name              string  `json:"name"`
	Desc              string  `json:"desc"`
	ImageID           string  `json:"imageId"`
	ImageName         string  `json:"imageName"`
	AlgorithmID       string  `json:"algorithmId"`
	AlgorithmVersion  string  `json:"algorithmVersion"`
	AlgorithmName     string  `json:"algorithmName"`
	ResourceSpecID    string  `json:"resourceSpecId"`
	ResourceSpecName  string  `json:"resourceSpecName"`
	Status            string  `json:"status"`
	DatasetID         string  `json:"datasetId"`
	DatasetVersion    string  `json:"datasetVersion"`
	DatasetName       string  `json:"datasetName"`
	ResourceSpecPrice float64 `json:"resourceSpecPrice"`
	NotebookJobID     string  `json:"notebookJobId"`
	ImageVersion      string  `json:"imageVersion"`
	Tasks             []Task  `json:"tasks"`
	ImageURL          string  `json:"imageUrl"`
}

type NoteBookService struct {
	containerApi.UnimplementedNotebookServiceServer
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewRentalService(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) containerApi.NotebookServiceServer {
	return &NoteBookService{
		conf: conf,
		log:  log.NewHelper("RentalService", logger),
		data: data,
	}
}

// CreateNotebook create notebook by miner after receiving request from renter
func (s *NoteBookService) CreateNotebook(ctx context.Context, req *containerApi.CreateNoteBookRequest) (*containerApi.CreateNoteBookReply, error) {

	reply := &containerApi.CreateNoteBookReply{
		Id:     "",
		Status: false,
	}

	// http addr request
	requestUrl := mainURL + "/v1/developmanage/notebook"
	jsonData := map[string]interface{}{
		"algorithmId":      req.AlgorithmId,
		"algorithmVersion": "V1",
		"datasetId":        "",
		"datasetVersion":   "",
		"desc":             req.Description,
		"imageId":          req.ImageId,
		"imageVersion":     "v1",
		"name":             req.Name,
		"resourcePool":     "common-pool",
		"resourceSpecId":   req.ResourceSpecId,
		"taskNumber":       1,
	}
	resp := HTTPRequest("POST", requestUrl, jsonData, "application/json", req.Token)

	var response struct {
		Success bool `json:"success"`
		Payload struct {
			ID string `json:"id"`
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

	reply.Id = response.Payload.ID
	reply.Status = true
	return reply, nil
}

// DeleteNotebook create notebook by miner after receiving request from renter
func (s *NoteBookService) DeleteNotebook(ctx context.Context, req *containerApi.DeleteNotebookRequest) (*containerApi.DeleteNotebookReply, error) {

	reply := &containerApi.DeleteNotebookReply{
		Id:     "",
		Status: false,
	}

	// http addr request
	requestUrl := mainURL + "/v1/developmanage/notebook/" + req.Id
	jsonData := map[string]interface{}{
		"id": req.Id,
	}
	resp := HTTPRequest("DELETE", requestUrl, jsonData, "application/json", req.Token)

	var response struct {
		Success bool `json:"success"`
		Payload struct {
			ID string `json:"id"`
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

	reply.Id = response.Payload.ID
	reply.Status = true
	return reply, nil
}

// QueryNotebookByCondition query notebook by miner
func (s *NoteBookService) QueryNotebookByCondition(ctx context.Context, req *containerApi.QueryNotebookByConditionRequest) (*containerApi.QueryNotebookByConditionReply, error) {
	list := make([]*containerApi.NotebookList, 0)
	// http addr request
	requestUrl := mainURL + "/v1/developmanage/notebook/" + req.Id
	if req.Id == "" {
		requestUrl = mainURL + "/v1/developmanage/notebook"
		jsonData := map[string]interface{}{
			"pageSize":  req.PageSize,
			"pageIndex": req.PageIndex,
		}
		resp := HTTPRequest("GET", requestUrl, jsonData, "application/json", req.Token)

		type Payload struct {
			TotalSize int64      `json:"totalSize"`
			Notebooks []Notebook `json:"notebook"`
		}
		var response struct {
			Success bool        `json:"success"`
			Payload Payload     `json:"payload"`
			Error   interface{} `json:"error"`
		}
		err := json.Unmarshal(resp, &response)
		if err != nil {
			reply := &containerApi.QueryNotebookByConditionReply{NoteBookList: list}
			return reply, err
		}
		fmt.Println(string(resp))
		for _, item := range response.Payload.Notebooks {
			list = append(list, &containerApi.NotebookList{
				ContainerId:   item.ID,
				UserId:        item.UserID,
				NotebookJobId: item.NotebookJobID,
				NotebookUrl:   item.Tasks[0].URL,
				Status:        item.Status,
				CreatedAt:     item.CreatedAt,
			})
		}
		reply := &containerApi.QueryNotebookByConditionReply{NoteBookList: list}

		return reply, nil

	}

	jsonData := map[string]interface{}{
		"id": req.Id,
	}
	resp := HTTPRequest("GET", requestUrl, jsonData, "application/json", req.Token)

	type Payload struct {
		Notebooks Notebook `json:"notebook"`
	}
	var response struct {
		Success bool        `json:"success"`
		Payload Payload     `json:"payload"`
		Error   interface{} `json:"error"`
	}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		reply := &containerApi.QueryNotebookByConditionReply{NoteBookList: list}
		return reply, err
	}

	list = append(list, &containerApi.NotebookList{
		ContainerId:   response.Payload.Notebooks.ID,
		UserId:        response.Payload.Notebooks.UserID,
		NotebookJobId: response.Payload.Notebooks.NotebookJobID,
		NotebookUrl:   response.Payload.Notebooks.Tasks[0].URL,
		Status:        response.Payload.Notebooks.Status,
		CreatedAt:     response.Payload.Notebooks.CreatedAt,
	})
	reply := &containerApi.QueryNotebookByConditionReply{NoteBookList: list}

	return reply, nil
}
