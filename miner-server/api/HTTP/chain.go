package HTTP

type ReportNodesStatusReply struct {
	Computation       string `json:"computation"`
	NumberOfMiners    string `json:"NumberOfMiners"`
	Rewards           string `json:"rewards"`
	LatestBlockHeight string `json:"latestBlockHeight"`
	//LatestBlockHash   string `son:"latestBlockHash"`
	//LatestBlockTime   string `json:"latestBlockTime"`
	MyComputation string `json:"myComputation"`
	MyRewards     string `json:"myRewards"`
	MyBlocks      string `json:"myBlocks"`
	MyWorkerNum   string `json:"myWorkerNum"`
}
