package generator

import (
	"context"
	"github.com/evleria-trading/price-service/internal/chanPool"
	"github.com/evleria-trading/price-service/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"time"
)

type Price interface {
	GeneratePrices(ctx context.Context) error
}

type price struct {
	pool chanPool.ChanPool
	rate time.Duration
}

func NewPricesGenerator(pool chanPool.ChanPool, rate time.Duration) Price {
	return &price{
		pool: pool,
		rate: rate,
	}
}

var symbols = []string{
	"AAPL",
	"TSLA",
	"GOOG",
	"MSFT",
	"AMZN",
	"FB",
	"NVDA",
}

func (g *price) GeneratePrices(ctx context.Context) error {
	m := map[string]model.Price{}
	log.Info("Starting prices generation at rate: ", g.rate)
	for {
		select {
		case <-time.After(g.rate):
			s := symbols[rand.Intn(len(symbols))]
			var p model.Price
			if prev, ok := m[s]; ok {
				c := getRandomCoefficient()
				p = newPrice(s, prev.Ask*c, prev.Bid*c)
			} else {
				base := 200 + rand.Float64()*500
				p = newPrice(s, base, base*0.9)
			}
			m[s] = p

			log.WithFields(log.Fields{"symbol": p.Symbol, "ask": p.Ask, "bid": p.Bid}).Info("Generated price")
			g.pool.WriteToEach(p)
		case <-ctx.Done():
			return nil
		}
	}
}

func newPrice(symbol string, ask, bid float64) model.Price {
	return model.Price{
		Id:     uuid.New().String(),
		Symbol: symbol,
		Ask:    roundToCents(ask),
		Bid:    roundToCents(bid),
	}
}

func roundToCents(f float64) float64 {
	return math.Round(f*100) / 100
}

func getRandomCoefficient() float64 {
	return rand.Float64()*0.0201010101 + 0.99
}
