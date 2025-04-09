package internal

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func GetAllGTFOBins() ([]string, error) {
	url := "https://gtfobins.github.io/"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GTFOBins homepage: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GTFOBins returned HTTP %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse GTFOBins HTML: %v", err)
	}

	var bins []string
	doc.Find("table tbody tr td a.bin-name").Each(func(i int, s *goquery.Selection) {
		bin := s.Text()
		bins = append(bins, bin)
	})

	return bins, nil
}
