package types

import (
	"context"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
	"strings"
	"uminer/common/log"
	"uminer/miner-server/api/chipApi"
	"uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
)

type ChipService struct {
	rpc.UnimplementedChipServiceServer
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewChipService(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) rpc.ChipServiceServer {
	return &ChipService{
		conf: conf,
		log:  log.NewHelper("ChipService", logger),
		data: data,
	}
}

// ListAllChips show details of all available bm-chips on a worker
func (s *ChipService) ListAllChips(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ListChipsReply, error) {

	cards := make([]*rpc.CardItem, 0)

	//cardLists := chipApi.BMChipsInfos("../../api/chipApi/bm_smi.txt")
	cardLists := chipApi.RemoteGetChipInfo(req.Url)
	listLen := 0

	// all cards
	for _, card := range cardLists {
		tpus := make([]*rpc.ChipItem, 0)
		if req.SerialNum != "" && card.SerialNum != req.SerialNum {
			continue
		}
		// all tpu chips in one card
		for _, chip := range card.Chips {
			if req.BusId != "" && chip.BusId != req.BusId {
				continue
			}
			tpus = append(tpus, &rpc.ChipItem{
				DevId:  chip.DevId,
				BusId:  chip.BusId,
				Memory: chip.UsedMemory + "/" + chip.TotalMemory,
				Tpuuti: chip.TPUUti,
				//BoardT:  chip.BoardT,
				ChipT:   chip.ChipT,
				TpuP:    chip.TPUP,
				TpuV:    chip.TPUV,
				TpuC:    chip.TPUC,
				Currclk: chip.Currclk,
			})
			listLen += 1
		}
		// all cards infos
		cards = append(cards, &rpc.CardItem{
			CardID:    card.CardID,
			Name:      card.Name,
			Mode:      card.Mode,
			SerialNum: card.SerialNum,
			Atx:       card.ATX,
			MaxP:      card.MaxP,
			BoardP:    card.BoardP,
			BoardT:    card.BoardT,
			Minclk:    card.Minclk,
			Maxclk:    card.Maxclk,
			Chips:     tpus,
		})
	}

	return &rpc.ListChipsReply{
		TotalSize: int64(listLen),
		Cards:     cards,
	}, nil

}

// StartChipCPU to start chip CPU before calling chips to perform their task
func (s *ChipService) StartChipCPU(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ChipStatusReply, error) {

	fipBin := "../../../bm_chip/src/fip.bin"
	rambootRootfs := "../../../bm_chip/src/ramboot_rootfs.itb"
	chipId, _ := strconv.ParseInt(req.DevId, 10, 64)

	res := chipApi.StartChips(int(chipId), fipBin, rambootRootfs)
	if res {
		return &rpc.ChipStatusReply{
			Status: "1",
			Msg:    "start chip cpu success",
		}, nil
	}
	return &rpc.ChipStatusReply{
		Status: "0",
		Msg:    "unable to start chip cpu",
	}, nil
}

// BurnChipEfuse burn the secret key of the chip at EFUSE area
func (s *ChipService) BurnChipEfuse(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ChipStatusReply, error) {

	chipId, _ := strconv.ParseInt(req.DevId, 10, 64)
	res := chipApi.BurnChips(req.SerialNum, req.BusId, int(chipId))
	if res {
		return &rpc.ChipStatusReply{
			Status: "1",
			Msg:    "burn aes key at efuse success",
		}, nil
	}
	return &rpc.ChipStatusReply{
		Status: "0",
		Msg:    "unable to burn aes key at efuse",
	}, nil
}

// GenerateChipKeyPairs generate p2 key (private key encrypted by secret key at EFUSE) + public key at chip and store them into files
func (s *ChipService) GenerateChipKeyPairs(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ChipStatusReply, error) {

	chipId, _ := strconv.ParseInt(req.DevId, 10, 64)
	res := chipApi.GenChipsKeyPairs(req.SerialNum, req.BusId, int(chipId))
	if res {
		return &rpc.ChipStatusReply{
			Status: "1",
			Msg:    "generate key pairs success and stored in files",
		}, nil
	}
	return &rpc.ChipStatusReply{
		Status: "0",
		Msg:    "unable to generate key pairs/key pairs already generated",
	}, nil
}

// ObtainChipKeyPairs read stored files to get p2 + pubkey of a chip
func (s *ChipService) ObtainChipKeyPairs(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ReadChipReply, error) {

	chipId, _ := strconv.ParseInt(req.DevId, 10, 64)
	keyPairs := chipApi.ReadChipKeyPairs(req.SerialNum, req.BusId, int(chipId))
	if keyPairs.P2 == "" {
		return &rpc.ReadChipReply{
			SerialNumber: req.SerialNum,
			BusId:        req.BusId,
			P2:           "",
			PublicKey:    "",
		}, errors.New("unable to read key pairs")
	}

	// encode both keys with base58 encoding
	block, _ := pem.Decode([]byte(keyPairs.PubKey))
	pubKeyBase58 := base58.Encode(block.Bytes)
	paddingLength := 402 - len(pubKeyBase58)
	paddedPubKeyBase58 := pubKeyBase58 + strings.Repeat("u", paddingLength)

	P2Bytes, _ := hex.DecodeString(keyPairs.P2)
	p2KeyBase58 := base58.Encode(P2Bytes)

	/* obtain all chip information and fill it into the chip.json at the jsonfile directory, and get prepare to upload to the chain  */

	return &rpc.ReadChipReply{
		SerialNumber:  req.SerialNum,
		BusId:         req.BusId,
		DevId:         req.DevId,
		P2:            p2KeyBase58,
		PublicKey:     paddedPubKeyBase58,
		P2Size:        int64(keyPairs.P2Size),
		PublicKeySize: int64(keyPairs.PubKeySize),
	}, nil
}

// SignChip sign at the chip by p2 key to get the unique signature for verification
func (s *ChipService) SignChip(ctx context.Context, req *rpc.SignChipsRequest) (*rpc.SignChipsReply, error) {

	devId, _ := strconv.ParseInt(req.DevId, 10, 64)

	// recover the p2 and pubKey by base58 decode
	originalPubKeyBase58 := strings.TrimSuffix(req.PublicKey, strings.Repeat("u", 33))
	pubKeyBytes := base58.Decode(originalPubKeyBase58)
	// use x509.ParsePKCS1PublicKey parse the bytes, and recovery the pubkey hex
	pubKey, err := x509.ParsePKCS1PublicKey(pubKeyBytes)
	if err != nil {
		return &rpc.SignChipsReply{
			Signature: "",
			Status:    false,
		}, err
	}
	serializedPubKey := x509.MarshalPKCS1PublicKey(pubKey)
	pemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: serializedPubKey,
	}
	PubKeyBytes := pem.EncodeToMemory(pemBlock)
	fmt.Println("recovered pubKey is")
	fmt.Println(string(PubKeyBytes))
	// use hex.EncodeToString encode the bytes, and recovery the p2 key hex
	P2Bytes := base58.Decode(req.P2)
	P2Key := hex.EncodeToString(P2Bytes)
	fmt.Println("recovered p2Key is")
	fmt.Println(P2Key)

	sign := chipApi.SignMinerChips(int(devId), P2Key, string(PubKeyBytes), int(req.P2Size), int(req.PublicKeySize), req.Msg)
	if sign.Signature == "" {
		return &rpc.SignChipsReply{
			Signature: sign.Signature,
			Status:    false,
		}, errors.New("unable to sign the chip")
	}

	return &rpc.SignChipsReply{
		Signature: sign.Signature,
		Status:    true,
	}, nil

}
