package chanPool

import (
	"github.com/evleria-trading/price-service/internal/model"
	log "github.com/sirupsen/logrus"
	"sync"
)

type ChanPool interface {
	WriteToEach(p model.Price)
	Acquire() (ch chan model.Price, id int)
	Release(id int)
}

type chanPool struct {
	lastId int
	mx     sync.Mutex
	chans  map[int]chan model.Price
}

func NewChanPool() ChanPool {
	return &chanPool{
		chans: map[int]chan model.Price{},
	}
}

func (c *chanPool) WriteToEach(p model.Price) {
	c.mx.Lock()
	defer c.mx.Unlock()

	for _, ch := range c.chans {
		ch <- p
	}
}

func (c *chanPool) Acquire() (ch chan model.Price, id int) {
	c.mx.Lock()
	defer c.mx.Unlock()

	ch = make(chan model.Price)
	c.lastId++
	id = c.lastId

	log.WithFields(log.Fields{
		"chan_id": id,
	}).Debug("Channel acquired")

	c.chans[id] = ch
	return ch, id
}

func (c *chanPool) Release(id int) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if ch, ok := c.chans[id]; ok {
		delete(c.chans, id)
		close(ch)

		log.WithFields(log.Fields{
			"chan_id":    id,
			"chans_left": len(c.chans),
		}).Debug("Channel released")
	}
}
