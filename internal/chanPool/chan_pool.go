package chanPool

import (
	"github.com/evleria/price-service/internal/model"
	log "github.com/sirupsen/logrus"
	"sync"
)

type ChanPool interface {
	WriteToEach(p model.Price)
	Acquire() chan model.Price
	Release(ch chan model.Price)
}

type chanPool struct {
	mx    sync.Mutex
	chans []chan model.Price
}

func NewChanPool() ChanPool {
	return &chanPool{}
}

func (c *chanPool) WriteToEach(p model.Price) {
	c.mx.Lock()
	defer c.mx.Unlock()

	for _, ch := range c.chans {
		ch <- p
	}
}

func (c *chanPool) Acquire() chan model.Price {
	c.mx.Lock()
	defer c.mx.Unlock()

	log.Debug("Acquire")

	ch := make(chan model.Price)
	c.chans = append(c.chans, ch)
	return ch
}

func (c *chanPool) Release(ch chan model.Price) {
	c.mx.Lock()
	defer c.mx.Unlock()

	log.Debug("Release")

	close(ch)
	for i, x := range c.chans {
		if x == ch {
			c.chans[i], c.chans[len(c.chans)-1] = c.chans[len(c.chans)-1], c.chans[i]
			c.chans = c.chans[:len(c.chans)-1]
			break
		}
	}
}
