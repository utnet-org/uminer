package HTTP

type MapWorkersAddressRequest struct {
	MinerAddr string `json:"minerAddr"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}
type MapWorkersAddressReply struct {
	MinerAddr  string   `json:"minerAddr"`
	AuthToken  string   `json:"authToken"`
	WorkerAddr []string `json:"workerAddr"`
}
type GetMinerIdReply struct {
	MinerId string `json:"minerId"`
	UserId  string `json:"userId"`
}

type ReportNodesStatusReply struct {
	Computation       string `json:"computation"`
	NumberOfMiners    string `json:"numberOfMiners"`
	Rewards           string `json:"rewards"`
	LatestBlockHeight string `json:"latestBlockHeight"`
	GasFee            string `json:"gasFee"`
	//LatestBlockHash   string `son:"latestBlockHash"`
	//LatestBlockTime   string `json:"latestBlockTime"`
	//MyComputation string `json:"myComputation"`
	//MyRewards     string `json:"myRewards"`
	//MyBlocks      string `json:"myBlocks"`
	//MyWorkerNum   string `json:"myWorkerNum"`
}

type ViewAccountReply struct {
	Total   string `json:"total"`
	Pledge  string `json:"pledge"`
	Rewards string `json:"rewards"`
	Slashed string `json:"slashed"`
	Power   string `json:"power"`
}

type GetRentalOrderListReply struct {
	RentalOrders []RentalOrderDetails `json:"rentalOrders"`
}
type RentalOrderDetails struct {
	ID         string `json:"id"`
	HASH       string `json:"hash"`
	MinerAddr  string `json:"minerAddr"`
	RentalAddr string `json:"rentalAddr"`
	Resource   string `json:"resource"`
	Power      string `json:"power"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}
