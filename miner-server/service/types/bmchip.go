package types

import (
	"context"
	"errors"
	"strconv"
	"uminer/common/log"
	"uminer/miner-server/chipApi"
	"uminer/miner-server/chipApi/rpc"
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

// show details of all chips
func (s *ChipService) ListAllChips(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ListChipsReply, error) {

	cards := make([]*rpc.CardItem, 0)

	cardLists := chipApi.BMChipsInfos()
	listLen := len(cardLists)

	for _, card := range cardLists {
		tpus := make([]*rpc.ChipItem, 0)
		for _, chip := range card.Chips {
			tpus = append(tpus, &rpc.ChipItem{
				DevId:   chip.DevId,
				BusId:   chip.BusId,
				Memory:  chip.Memory,
				Tpuuti:  chip.TPUUti,
				BoardT:  chip.BoardT,
				ChipT:   chip.ChipT,
				TpuP:    chip.TPUP,
				TpuV:    chip.TPUV,
				TpuC:    chip.TPUC,
				Currclk: chip.Currclk,
				Status:  chip.Status,
			})
		}
		cards = append(cards, &rpc.CardItem{
			CardID:    card.CardID,
			Name:      card.Name,
			Mode:      card.Mode,
			SerialNum: card.SerialNum,
			Atx:       card.ATX,
			MaxP:      card.MaxP,
			BoardP:    card.BoardP,
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

// start chip CPU
func (s *ChipService) StartChipCPU(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ChipStatusReply, error) {

	fipBin := "../../bm_chip/src/fip.bin"
	rambootRootfs := "../../bm_chip/src/ramboot_rootfs.itb"
	busId, _ := strconv.ParseInt(req.BusId, 10, 64)

	res := chipApi.StartChips(int(busId), fipBin, rambootRootfs)
	if res {
		return &rpc.ChipStatusReply{
			Status: true,
		}, nil
	}
	return &rpc.ChipStatusReply{
		Status: false,
	}, errors.New("unable to start chip cpu")
}

// burn chip at efuse
func (s *ChipService) BurnChipEfuse(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ChipStatusReply, error) {

	busId, _ := strconv.ParseInt(req.BusId, 10, 64)
	res := chipApi.BurnChips(req.SerialNum, req.BusId, int(busId))
	if res {
		return &rpc.ChipStatusReply{
			Status: true,
		}, nil
	}
	return &rpc.ChipStatusReply{
		Status: false,
	}, errors.New("unable to burn aes key at efuse")
}

// generate p2 + pubkey at chip and store them into files
func (s *ChipService) GenerateChipKeyPairs(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ChipStatusReply, error) {

	busId, _ := strconv.ParseInt(req.BusId, 10, 64)
	res := chipApi.GenChipsKeyPairs(req.SerialNum, req.BusId, int(busId))
	if res {
		return &rpc.ChipStatusReply{
			Status: true,
		}, nil
	}
	return &rpc.ChipStatusReply{
		Status: false,
	}, errors.New("unable to generate key pairs")
}

// read stored files to get p2 + pubkey at chip
func (s *ChipService) ObtainChipKeyPairs(ctx context.Context, req *rpc.ChipsRequest) (*rpc.ReadChipReply, error) {

	busId, _ := strconv.ParseInt(req.BusId, 10, 64)
	keyPairs := chipApi.ReadChipKeyPairs(req.SerialNum, req.BusId, int(busId))
	if keyPairs.P2 == "" {
		return &rpc.ReadChipReply{
			SerialNumber: req.SerialNum,
			BusId:        req.BusId,
			P2:           "",
			PublicKey:    "",
		}, errors.New("unable to read key pairs")
	}

	// packed as transaction, upload to the chain

	return &rpc.ReadChipReply{
		SerialNumber: req.SerialNum,
		BusId:        req.BusId,
		P2:           keyPairs.P2,
		PublicKey:    keyPairs.PubKey,
	}, nil
}

// sign a chip by p2 to get signature
func (s *ChipService) SignChip(ctx context.Context, req *rpc.SignChipsRequest) (*rpc.SignChipsReply, error) {

	busId, _ := strconv.ParseInt(req.BusId, 10, 64)
	sign := chipApi.SignMinerChips(req.SerialNum, req.BusId, int(busId), req.P2, req.PublicKey, int(req.P2Size), int(req.PublicKeySize), req.Msg)
	if sign.Signature == "" {
		return &rpc.SignChipsReply{
			Signature: sign.Signature,
			Status:    false,
		}, errors.New("unable to sign the chip")
	}

	// packed as transaction, upload to the chain

	return &rpc.SignChipsReply{
		Signature: sign.Signature,
		Status:    true,
	}, nil

}
