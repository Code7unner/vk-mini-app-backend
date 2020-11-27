package pools

import (
	"github.com/gocolly/colly"
	"sync"
)

type CollyPool struct {
	pool *sync.Pool
}

func NewCollyPool() *CollyPool {
	const domain = "liquipedia.net"
	return &CollyPool{
		pool: &sync.Pool{New: func() interface{} {
			return colly.NewCollector(colly.AllowedDomains(domain))
		}},
	}
}

func (m *CollyPool) Get() *colly.Collector {
	return m.pool.Get().(*colly.Collector)
}

func (m *CollyPool) Put(collector *colly.Collector) {
	m.pool.Put(collector)
}
