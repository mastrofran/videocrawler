package videocrawler

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func GetVideo(url string) []string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	err := chromedp.Run(ctx,
		// visit the target page
		chromedp.Navigate(url),

		// wait for the page to load
		chromedp.Sleep(3000*time.Millisecond),

		// extract the raw HTML from the page
		chromedp.ActionFunc(func(ctx context.Context) error {
			// select the root node on the page
			rootNode, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			html, err = dom.GetOuterHTML().WithNodeID(rootNode.NodeID).Do(ctx)
			return err
		}),
	)
	if err != nil {
		log.Fatal("Error while performing the automation logic:", err)
	}

	re := regexp.MustCompile(`https?://[^\s,]*?mp4[^\s]*?"`)

	found := re.FindAllString(html, -1)

	var links []string
	replacer := strings.NewReplacer("amp;", "", `"`, "")
	for _, element := range found {
		links = append(links, replacer.Replace(element))
	}
	fmt.Println("crawled element: ", url, "found: ", links)
	return links
}
