package handler

import (
	"github.com/evleria-trading/price-service/internal/chanPool"
	"github.com/evleria-trading/price-service/protocol/pb"
	"github.com/golang/protobuf/ptypes/empty"
)

type PriceService struct {
	pb.UnimplementedPriceServiceServer
	pool chanPool.ChanPool
}

func NewPriceService(pool chanPool.ChanPool) pb.PriceServiceServer {
	return &PriceService{
		pool: pool,
	}
}

func (p *PriceService) GetPrices(_ *empty.Empty, stream pb.PriceService_GetPricesServer) error {
	ch, id := p.pool.Acquire()
	defer p.pool.Release(id)

	for {
		select {
		case pr := <-ch:
			err := stream.Send(&pb.Price{
				Id:     pr.Id,
				Symbol: pr.Symbol,
				Ask:    pr.Ask,
				Bid:    pr.Bid,
			})
			if err != nil {
				return err
			}

		case <-stream.Context().Done():
			return nil
		}
	}
}
