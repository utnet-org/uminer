package types

import (
	"context"
	"errors"
	"strconv"
	"uminer/common/log"
	"uminer/miner-server/api/chipApi"
	"uminer/miner-server/api/chipApi/HTTP"
	"uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/data"
	"uminer/miner-server/serverConf"
)

type ChipServiceHTTP struct {
	conf *serverConf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewChipServiceHTTP(conf *serverConf.Bootstrap, logger log.Logger, data *data.Data) *ChipServiceHTTP {
	return &ChipServiceHTTP{
		conf: conf,
		log:  log.NewHelper("ChipService", logger),
		data: data,
	}
}

// show details of all chips
func (s *ChipServiceHTTP) ListAllChipsHTTP(ctx context.Context, req *HTTP.ChipsRequest) (*HTTP.ListChipsReply, error) {

	cards := make([]*HTTP.CardItem, 0)

	//cardLists := chipApi.BMChipsInfos("../../api/chipApi/bm_smi.txt")
	cardLists := chipApi.RemoteGetChipInfo(req.Url)
	listLen := 0

	for _, card := range cardLists {
		tpus := make([]*HTTP.ChipItem, 0)
		if req.SerialNum != "" && card.SerialNum != req.SerialNum {
			continue
		}
		// tpu chips
		for _, chip := range card.Chips {
			if req.BusId != "" && chip.BusId != req.BusId {
				continue
			}
			tpus = append(tpus, &HTTP.ChipItem{
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
				Status:  chip.Status,
			})
			listLen += 1
		}
		// all card infos
		// get claim status
		claimStatus := "unclaimed"
		// claimStatus := getStatusOfCard(card.SerialNum)
		cards = append(cards, &HTTP.CardItem{
			CardID:      card.CardID,
			Name:        card.Name,
			Mode:        card.Mode,
			SerialNum:   card.SerialNum,
			Atx:         card.ATX,
			MaxP:        card.MaxP,
			BoardP:      card.BoardP,
			BoardT:      card.BoardT,
			Minclk:      card.Minclk,
			Maxclk:      card.Maxclk,
			Chips:       tpus,
			ClaimStatus: claimStatus,
		})
	}

	return &HTTP.ListChipsReply{
		TotalSize: int64(listLen),
		Cards:     cards,
	}, nil

}

// start chip CPU
func (s *ChipServiceHTTP) StartChipCPUHTTP(ctx context.Context, req *HTTP.ChipsRequest) (*HTTP.ChipStatusReply, error) {

	fipBin := "../../../bm_chip/src/fip.bin"
	rambootRootfs := "../../../bm_chip/src/ramboot_rootfs.itb"
	busId, _ := strconv.ParseInt(req.BusId, 10, 64)

	res := chipApi.StartChips(int(busId), fipBin, rambootRootfs)
	if res {
		return &HTTP.ChipStatusReply{
			Status: true,
		}, nil
	}
	return &HTTP.ChipStatusReply{
		Status: false,
	}, errors.New("unable to start chip cpu")
}

// burn chip at efuse
func (s *ChipServiceHTTP) BurnChipEfuseHTTP(ctx context.Context, req *HTTP.ChipsRequest) (*HTTP.ChipStatusReply, error) {

	busId, _ := strconv.ParseInt(req.BusId, 10, 64)
	res := chipApi.BurnChips(req.SerialNum, req.BusId, int(busId))
	if res {
		return &HTTP.ChipStatusReply{
			Status: true,
		}, nil
	}
	return &HTTP.ChipStatusReply{
		Status: false,
	}, errors.New("unable to burn aes key at efuse")
}

// generate p2 + pubkey at chip and store them into files
func (s *ChipServiceHTTP) GenerateChipKeyPairsHTTP(ctx context.Context, req *HTTP.ChipsRequest) (*HTTP.ChipStatusReply, error) {

	busId, _ := strconv.ParseInt(req.BusId, 10, 64)
	res := chipApi.GenChipsKeyPairs(req.SerialNum, req.BusId, int(busId))
	if res {
		return &HTTP.ChipStatusReply{
			Status: true,
		}, nil
	}
	return &HTTP.ChipStatusReply{
		Status: false,
	}, errors.New("unable to generate key pairs")
}

// read stored files to get p2 + pubkey at chip
func (s *ChipServiceHTTP) ObtainChipKeyPairsHTTP(ctx context.Context, req *rpc.ChipsRequest) (*HTTP.ReadChipReply, error) {

	busId, _ := strconv.ParseInt(req.BusId, 10, 64)
	keyPairs := chipApi.ReadChipKeyPairs(req.SerialNum, req.BusId, int(busId))
	if keyPairs.P2 == "" {
		return &HTTP.ReadChipReply{
			SerialNumber: req.SerialNum,
			BusId:        req.BusId,
			P2:           "",
			PublicKey:    "",
		}, errors.New("unable to read key pairs")
	}

	// packed as transaction, upload to the chain

	return &HTTP.ReadChipReply{
		SerialNumber: req.SerialNum,
		BusId:        req.BusId,
		P2:           keyPairs.P2,
		PublicKey:    keyPairs.PubKey,
	}, nil
}

// sign a chip by p2 to get signature
func (s *ChipServiceHTTP) SignChipHTTP(ctx context.Context, req *HTTP.SignChipsRequest) (*HTTP.SignChipsReply, error) {

	if req.SerialNum == "" && req.BusId == "" {
		return &HTTP.SignChipsReply{
			Signature: "",
			Status:    false,
		}, errors.New("no chip is selected")
	}

	// obtain the devId of the chip
	devId := -1
	cardLists := chipApi.BMChipsInfos("../../api/chipApi/bm_smi.txt")
	for _, card := range cardLists {
		if card.SerialNum == req.SerialNum {
			for _, chip := range card.Chips {
				if chip.BusId == req.BusId {
					id, _ := strconv.ParseInt(chip.DevId, 10, 64)
					devId = int(id)
				}
			}
		}

	}
	if devId == -1 {
		return &HTTP.SignChipsReply{
			Signature: "",
			Status:    false,
		}, errors.New("no chip is selected")
	}

	sign := chipApi.SignMinerChips(devId, req.P2, req.PublicKey, int(req.P2Size), int(req.PublicKeySize), req.Msg)
	if sign.Signature == "" {
		return &HTTP.SignChipsReply{
			Signature: sign.Signature,
			Status:    false,
		}, errors.New("unable to sign the chip")
	}

	// packed as transaction, upload to the chain

	return &HTTP.SignChipsReply{
		Signature: sign.Signature,
		Status:    true,
	}, nil

}
