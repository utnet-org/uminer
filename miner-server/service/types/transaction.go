package types

type transaction struct {
	Address   string
	From      string
	To        string
	Amount    int64
	txData    string
	TimeStamp string
	GasFee    int64
	Hash      string
}

type MinerChip struct {
	SN            string
	BusID         string
	P2            string
	PublicKey     string
	P2Size        int64
	PublicKeySize int64
}
type reportMinerComputation struct {
	Address    string
	ServerIP   string
	BMChips    []MinerChip
	totalPower int64
}
