package scrapper

import (
	"context"
	"github.com/code7unner/vk-mini-app-backend/scrapper/helper"
	"sync"
	"time"
)

type Scrapper struct {
	wg        *sync.WaitGroup
	scrapTime time.Duration
	helper    *helper.HTMLHelper
}

func New(time time.Duration) *Scrapper {
	return &Scrapper{
		wg:        &sync.WaitGroup{},
		scrapTime: time,
		helper:    helper.New(),
	}
}

func (s Scrapper) Start(ctx context.Context) {
	go s.GetMatches(ctx, s.wg)
}
