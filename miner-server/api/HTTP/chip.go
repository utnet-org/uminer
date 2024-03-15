package HTTP

//type ChipServiceServer interface {
//	StartChipCPUHTTP(context.Context, *ChipsRequest) (*ChipStatusReply, error)
//	ListAllChipsHTTP(context.Context, *ChipsRequest) (*ListChipsReply, error)
//	BurnChipEfuseHTTP(context.Context, *ChipsRequest) (*ChipStatusReply, error)
//	GenerateChipKeyPairsHTTP(context.Context, *ChipsRequest) (*ChipStatusReply, error)
//	ObtainChipKeyPairsHTTP(context.Context, *ChipsRequest) (*ReadChipReply, error)
//	SignChipHTTP(context.Context, *SignChipsRequest) (*SignChipsReply, error)
//}

type ChipItem struct {
	DevId       string `json:"devId,omitempty"`
	BusId       string `json:"busId,omitempty"`
	Memory      string `json:"memory,omitempty"`
	Tpuuti      string `json:"tpuuti,omitempty"`
	BoardT      string `json:"boardT,omitempty"`
	ChipT       string `json:"chipT,omitempty"`
	TpuP        string `json:"tpuP,omitempty"`
	TpuV        string `json:"tpuV,omitempty"`
	TpuC        string `json:"tpuC,omitempty"`
	Currclk     string `json:"currclk,omitempty"`
	ClaimStatus string `json:"claimStatus,omitempty"`
}

type CardItem struct {
	CardID    string      `json:"cardID"`
	Name      string      `json:"name"`
	Mode      string      `json:"mode"`
	SerialNum string      `json:"serialNum"`
	Atx       string      `json:"atx"`
	MaxP      string      `json:"maxP"`
	BoardP    string      `json:"boardP"`
	BoardT    string      `json:"boardT"`
	Minclk    string      `json:"minclk"`
	Maxclk    string      `json:"maxclk"`
	Chips     []*ChipItem `json:"chips"`
}

type ListChipsRequest struct {
	Addr      []string `json:"addr"`
	SerialNum string   `json:"serialNum"`
	BusId     string   `json:"busId"`
	Account   string   `json:"account"`
}
type StartChipsRequest struct {
	Addr  string `json:"addr"`
	DevId string `json:"busId"`
}

// start/burn/gen chip
type ChipStatusReply struct {
	Status bool `json:"status"`
}

// details information
type ListWorkersReply struct {
	NumOfWorkers int64       `json:"numOfWorkers"`
	Workers      []ListCards `json:"workers"`
}

type ListCards struct {
	TotalSize int64       `json:"totalSize"`
	Addr      string      `json:"address"`
	Cards     []*CardItem `json:"cards"`
	Status    string      `json:"status"` // Connected/Disconnected
}

// read keys
type ReadChipReply struct {
	SerialNumber string `json:"SerialNumber"`
	BusId        string `json:"BusId"`
	P2           string `json:"p2"`
	PublicKey    string `json:"publicKey"`
}

// sign
type SignChipsRequest struct {
	SerialNum     string `json:"serialNum,omitempty"`
	BusId         string `json:"busId,omitempty"`
	P2            string `json:"p2,omitempty"`
	PublicKey     string `json:"publicKey,omitempty"`
	P2Size        int64  `json:"p2Size,omitempty"`
	PublicKeySize int64  `json:"publicKeySize,omitempty"`
	Msg           string `json:"msg,omitempty"`
}

type SignChipsReply struct {
	Signature string `json:"signature"`
	Status    bool   `json:"status"`
}
