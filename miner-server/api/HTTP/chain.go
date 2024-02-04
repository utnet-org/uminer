package HTTP

type MapWorkersAddressRequest struct {
	MinerAddr string `json:"minerAddr"`
}
type MapWorkersAddressReply struct {
	MinerAddr  string   `json:"minerAddr"`
	WorkerAddr []string `json:"workerAddr"`
}

type ReportNodesStatusReply struct {
	Computation       string `json:"computation"`
	NumberOfMiners    string `json:"numberOfMiners"`
	Rewards           string `json:"rewards"`
	LatestBlockHeight string `json:"latestBlockHeight"`
	GasFee            string `json:"gasFee"`
	//LatestBlockHash   string `son:"latestBlockHash"`
	//LatestBlockTime   string `json:"latestBlockTime"`
	MyComputation string `json:"myComputation"`
	MyRewards     string `json:"myRewards"`
	MyBlocks      string `json:"myBlocks"`
	MyWorkerNum   string `json:"myWorkerNum"`
}
