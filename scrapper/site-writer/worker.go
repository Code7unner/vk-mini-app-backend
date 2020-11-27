package site_writer

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type SiteReceiver struct {
	HttpClientPool *HttpClientPool
	Logger         *logrus.Logger
}

func NewSiteReceiver(logger *logrus.Logger) *SiteReceiver {
	r := &SiteReceiver{
		HttpClientPool: NewHttpClientPool(),
		Logger:         logger,
	}

	return r
}

func (r *SiteReceiver) Receive(ctx context.Context, url string) (successProcessed int) {
	var defaultTickTime = time.Duration(1)
	tickTime := viper.GetDuration("time")
	if tickTime == time.Duration(0) {
		tickTime = defaultTickTime
	}

	ticker := time.NewTicker(time.Minute * tickTime)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

			successProcessed++
		}
	}
}
