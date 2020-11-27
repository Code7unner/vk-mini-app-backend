package scrapper

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func (s *Scrapper) GetMatches(ctx context.Context, wg *sync.WaitGroup) {
	const url = "https://liquipedia.net/dota2/api.php?action=query&prop=revisions&titles=Liquipedia:Upcoming_and_ongoing_matches&rvprop=content&rvparse&format=json"

	ticker := time.NewTicker(time.Second * s.scrapTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			matches, err := s.helper.GetMatches(url)
			if err != nil {
				logrus.Error(err)
				return
			}

			logrus.Info(matches[0])
		case <-ctx.Done():
			return
		}
	}
}
