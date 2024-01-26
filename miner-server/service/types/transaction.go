package types

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/tidwall/gjson"
	"net/http"
	"time"
	"uminer/miner-server/util"
)

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

type RentalOrder struct {
	RentalAddress string
	UserAddress   string
	ContainerId   string
	Computation   int64
	Duration      int64
	Fee           float64
	CreateTime    string
}

func sendTransactionAsync(ctx context.Context, signature string) (string, error) {

	jsonData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  "broadcast_tx_async",
		"params":  []interface{}{signature},
	}

	jsonStr, _ := json.Marshal(jsonData)

	// POST request
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, nodeURL, bytes.NewReader(jsonStr))
	if err != nil {
		return "", err
	}
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("accept-encoding", "gzip,deflate")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	gzipBytes := util.GzipApi(resp)

	res := gjson.Get(string(gzipBytes), "result").String()
	return res, nil
}
