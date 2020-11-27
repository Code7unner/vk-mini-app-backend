package helper

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/code7unner/vk-mini-app-backend/scrapper/helper/models"
	"github.com/code7unner/vk-mini-app-backend/scrapper/pools"
	"io/ioutil"
	"regexp"
	"strings"
)

type HTMLHelper struct {
	httpClientPool *pools.HttpClientPool
	collyPool      *pools.CollyPool
}

func New() *HTMLHelper {
	return &HTMLHelper{
		httpClientPool: pools.NewHttpClientPool(),
		collyPool:      pools.NewCollyPool(),
	}
}

func (h *HTMLHelper) findMatches(html string) ([]models.MatchData, error) {
	spaceRegExp := regexp.MustCompile(`\r?\n`)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	matchesData := make([]models.MatchData, 0)

	s := doc.Find(`.wikitable.wikitable-striped.infobox_matches_content`)
	s = s.Each(func(i int, s1 *goquery.Selection) {
		var matchData models.MatchData
		//leftTeam
		_ = s1.Find(".team-left").Each(func(i int, s2 *goquery.Selection) {
			matchData.LeftTeamName = spaceRegExp.ReplaceAllString(s2.Text(), " ")
			matchData.ImageSrc, _ = s2.Find("img").Attr("src")
		})

		//versus
		versus := s1.Find(".versus").Text()
		matchData.Score = spaceRegExp.ReplaceAllString(versus, " ")

		//teamRight
		_ = s1.Find(".team-right").Each(func(i int, s2 *goquery.Selection) {
			matchData.RightTeamName = spaceRegExp.ReplaceAllString(s2.Text(), " ")
			matchData.ImageSrc, _ = s2.Find("img").Attr("src")
		})

		const dateDivider = 5
		//matchInfo
		_ = s1.Find(".match-filler").Each(func(i int, s2 *goquery.Selection) {
			dateAndTournament := spaceRegExp.ReplaceAllString(s2.Text(), "")
			dateAndTournamentSlice := strings.Fields(dateAndTournament)
			matchData.DateTime = strings.Join(dateAndTournamentSlice[:dateDivider], " ")
			matchData.TournamentName = strings.Join(dateAndTournamentSlice[dateDivider:], " ")
		})

		matchesData = append(matchesData, matchData)
	})

	return matchesData, nil
}

func (h *HTMLHelper) getStringFromHTML(url string) ([]byte, error) {
	client := h.httpClientPool.Get()
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	h.httpClientPool.Put(client)

	return result, nil
}

func (h *HTMLHelper) GetMatches(url string) ([]models.MatchData, error) {
	result, err := h.getStringFromHTML(url)
	if err != nil {
		return nil, err
	}

	var data models.Data
	if err := json.Unmarshal(result, &data); err != nil {
		return nil, err
	}

	html := data.Query.Pages.Page
	if html.Revisions == nil {
		return nil, errors.New("matches string html is empty")
	}

	matchesData, err := h.findMatches(html.Revisions[0].Html)
	if err != nil {
		return nil, err
	}

	return matchesData, nil
}
