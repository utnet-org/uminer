package types

import (
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	http2 "net/http"
	"uminer/common/log"
	"uminer/miner-server/api/HTTP"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service/connect"
)

type MinerLoginServiceHTTP struct {
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewMinerLoginServiceHTTP(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) *MinerLoginServiceHTTP {
	return &MinerLoginServiceHTTP{
		conf: conf,
		log:  log.NewHelper("MinerLoginService", logger),
		data: data,
	}
}

// LoginHandler get login token and all worker address of a miner
func (s *MinerLoginServiceHTTP) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// method Get
	if r.Method != http2.MethodGet {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get params
	requestUrl := connect.MainURL + "/v1/authmanage/token"
	query := r.URL.Query()
	req := &HTTP.MapWorkersAddressRequest{
		MinerAddr: query.Get("minerAddr"),
		UserName:  query.Get("username"),
		Password:  query.Get("password"),
	}
	jsonData := map[string]interface{}{
		"username": req.UserName,
		"password": req.Password,
	}
	resp := connect.HTTPRequest("POST", requestUrl, jsonData, "application/json", "")
	var response struct {
		Success bool `json:"success"`
		Payload struct {
			Token      string `json:"token"`
			Expiration int    `json:"expiration"`
		} `json:"payload"`
		Error interface{} `json:"error"`
	}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

	switch errObj := response.Error.(type) {
	case map[string]interface{}:
		// 转换为 map 类型成功，可以提取目标字段的值
		message, ok := errObj["message"].(string)
		if !ok {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		} else {
			http2.Error(w, message, http2.StatusInternalServerError)
		}
	default:

	}

	// get mapping
	workers := make([]string, 0)
	workers = append(workers, "192.168.10.49")
	workers = append(workers, "192.168.10.50")
	workers = append(workers, "192.168.10.51")
	token := response.Payload.Token
	finalResponse := HTTP.MapWorkersAddressReply{
		MinerAddr:  req.MinerAddr,
		AuthToken:  token,
		WorkerAddr: workers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(finalResponse); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}

// GetMinerInfoHandler get minerInfo(userId) from token
func (s *MinerLoginServiceHTTP) GetMinerInfoHandler(w http.ResponseWriter, r *http.Request) {
	// method Get
	if r.Method != http2.MethodGet {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get params
	query := r.URL.Query()
	token := query.Get("token")
	ownerAddress := query.Get("ownerAddress")
	requestUrl := connect.MainURL + "/v1/usermanage/user?token=" + token
	jsonData := map[string]interface{}{
		"token": token,
	}
	// get userid
	resp := connect.HTTPRequest("GET", requestUrl, jsonData, "application/json", token)
	type User struct {
		ID            string   `json:"id"`
		CreatedAt     int64    `json:"createdAt"`
		UpdatedAt     int64    `json:"updatedAt"`
		FullName      string   `json:"fullName"`
		Email         string   `json:"email"`
		Phone         string   `json:"phone"`
		Gender        int      `json:"gender"`
		Status        int      `json:"status"`
		FTPUserName   string   `json:"ftpUserName"`
		ResourcePools []string `json:"resourcePools"`
	}
	var response struct {
		Success bool `json:"success"`
		Payload struct {
			User User `json:"user"`
		} `json:"payload"`
		Error interface{} `json:"error"`
	}
	err := json.Unmarshal(resp, &response)
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

	switch errObj := response.Error.(type) {
	case map[string]interface{}:
		// 转换为 map 类型成功，可以提取目标字段的值
		message, ok := errObj["message"].(string)
		if !ok {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		} else {
			http2.Error(w, message, http2.StatusInternalServerError)
		}
	default:
	}

	// get minerId(publicKey) by owner address
	fmt.Println("owner is: ", ownerAddress)

	finalResponse := HTTP.GetMinerIdReply{
		MinerId: "jackronwong",
		UserId:  response.Payload.User.ID,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(finalResponse); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}
