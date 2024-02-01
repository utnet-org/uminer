package HTTP

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
