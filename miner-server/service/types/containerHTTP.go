package types

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"
	http2 "net/http"
	"strconv"
	"time"
	"uminer/common/log"
	"uminer/miner-server/api/HTTP"
	"uminer/miner-server/api/containerApi"
	"uminer/miner-server/cmd"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service/connect"
	"uminer/miner-server/util"
)

type ContainerServiceHTTP struct {
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewMinerContainerServiceHTTP(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) *ContainerServiceHTTP {
	return &ContainerServiceHTTP{
		conf: conf,
		log:  log.NewHelper(" ContainerService", logger),
		data: data,
	}
}

// GetRentalOrderListHandler get all rental order list according to account address
func (s *ContainerServiceHTTP) GetRentalOrderListHandler(w http.ResponseWriter, r *http.Request) {

	// method Get
	if r.Method != http2.MethodGet {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get params
	query := r.URL.Query()
	account := query.Get("account")

	// connect to utility node rpc methods to get the results

	// get response
	rental := make([]HTTP.RentalOrderDetails, 0)
	rentallist := make([]HTTP.RentalOrderDetails, 3)
	for i, _ := range rentallist {
		// mock data
		rental = append(rental, HTTP.RentalOrderDetails{
			ID:         strconv.FormatInt(int64(i), 10),
			HASH:       util.RandomString(32),
			MinerAddr:  account,
			RentalAddr: util.RandomString(11),
			Resource:   "1684x",
			Power:      strconv.FormatInt(int64((i+2)*10), 10),
			StartTime:  time.Now().Add(time.Duration(-i) * 24 * time.Hour).Format("2006-01-02 15:04"),
			EndTime:    time.Now().Add(time.Duration(180-i) * 24 * time.Hour).Format("2006-01-02 15:04"),
		})
	}
	response := HTTP.GetRentalOrderListReply{
		RentalOrders: rental,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}

// GetNotebookListHandler get all notebook list by user token and notebook id
func (s *ContainerServiceHTTP) GetNotebookListHandler(w http.ResponseWriter, r *http.Request) {

	// method Get
	if r.Method != http2.MethodGet {
		http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
		return
	}
	// get params
	query := r.URL.Query()
	token := query.Get("token")
	notebookId := query.Get("notebookId")

	// connect to worker
	clientDeadline := time.Now().Add(time.Duration(connect.Delay * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	conn, err := grpc.DialContext(ctx, cmd.MinerServerIP+":9001", grpc.WithInsecure())
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	// Prepare the request
	request := &containerApi.QueryNotebookByConditionRequest{
		Token:     token,
		Id:        notebookId,
		PageSize:  10,
		PageIndex: 1,
	}
	client := containerApi.NewNotebookServiceClient(conn)
	// Call the RPC method
	var response *containerApi.QueryNotebookByConditionReply
	response, err = client.QueryNotebookByCondition(ctx, request, grpc.WaitForReady(true))
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http2.Error(w, err.Error(), http2.StatusInternalServerError)
	}

}
