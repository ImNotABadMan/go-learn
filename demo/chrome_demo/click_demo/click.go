package click_demo

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func RunChromeClick() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://learnku.com/docs/laravel/7.x/helpers/7486#method-array-pull`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > .open-all`),
		// find and click "Expand All" link
		chromedp.Click(`.open-all`, chromedp.NodeVisible),
		// retrieve the value of the textarea
		chromedp.Value(`.open-all`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
