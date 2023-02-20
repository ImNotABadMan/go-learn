package code

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func Login() {
	//ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	//defer cancel()
	//
	//chromedp.Run(ctx,
	//	openPage("https://marketplace.catch.com.au/login",
	//		"/html/body/div/div/div[1]/div[3]/form/div[3]/div[1]/button"),
	//)

	chromedp.Run(context.Background())
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-background-networking", false),
		chromedp.Flag("disable-renderer-backgrounding", false),
		chromedp.Flag("disable-popup-blocking", false),
		chromedp.Flag("disable-ipc-flooding-protection", false),
		chromedp.Flag("disable-client-side-phishing-detection", false),
		chromedp.Flag("disable-background-timer-throttling", false),
		//chromedp.WindowSize(1200, 800),
		chromedp.WindowSize(1500, 900),
		chromedp.Flag("headless", false),
		// Like in Puppeteer.
		chromedp.Flag("hide-scrollbars", false),
		//chromedp.DisableGPU,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer func() {
		fmt.Println("close browser")
		cancel()
	}()

	chromedp.Run(taskCtx,
		//chromedp.Navigate("http://v2.globaloutlet-backend.com:8011/login"),
		chromedp.Navigate("https://marketplace.catch.com.au/login"),
		chromedp.WaitVisible("/html/body/div/div/div[1]/div[3]/form/div[3]/div[1]/button"),
		login(),
		choseShippingImport(),
	)

}

func login() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(`#username`, chromedp.ByID),
		chromedp.SendKeys("#username", "", chromedp.ByID), // TODO::
		chromedp.Sleep(time.Millisecond * 1000),
		chromedp.WaitVisible("#submitButton", chromedp.ByID),
		chromedp.Click("#submitButton", chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 1000),
		chromedp.WaitVisible(`#password`, chromedp.ByID),
		chromedp.SendKeys("#password", "", chromedp.ByID), // TODO::
		chromedp.Sleep(time.Millisecond * 2000),
		chromedp.WaitVisible(`#btn-login`, chromedp.ByID),
		chromedp.Click("#btn-login", chromedp.ByID),
	}
}

func choseShippingImport() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(`body > div.mui-gateway-menu > nav > div.mui-top-menu-container > ul.mui-top-menu-right > li.mui-top-menu-shop`, chromedp.ByQuery),
		chromedp.Click("body > div.mui-gateway-menu > nav > div.mui-top-menu-container > ul.mui-top-menu-right > li.mui-top-menu-shop", chromedp.ByQuery),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> body > div.mui-gateway-menu > nav > div.mui-top-menu-container > ul.mui-top-menu-right > li.mui-top-menu-shop.mui-top-menu-item-open IS Click")),
		chromedp.Sleep(time.Millisecond * 1000),
		chromedp.WaitVisible(`#changeToShop9875-small`, chromedp.ByID),
		chromedp.Click(`#changeToShop9875-small`, chromedp.ByID),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> #changeToShop9875-small IS Click")),
		chromedp.Sleep(time.Millisecond * 1000),

		chromedp.WaitVisible(`body > div.mui-gateway-menu > nav > div.mui-top-menu-container > ul.mui-top-menu-left > li:nth-child(4)`, chromedp.ByQuery),
		chromedp.Click("body > div.mui-gateway-menu > nav > div.mui-top-menu-container > ul.mui-top-menu-left > li:nth-child(4)", chromedp.ByQuery),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> body > div.mui-gateway-menu > nav > div.mui-top-menu-container > ul.mui-top-menu-left > li:nth-child(4) IS Click")),
		chromedp.Sleep(time.Millisecond * 1000),

		chromedp.WaitVisible(`#shopParameter`, chromedp.ByID),
		chromedp.Click(`#shopParameter`, chromedp.ByID),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> #shopParameter IS Click")),
		chromedp.Sleep(time.Millisecond * 1000),

		chromedp.WaitVisible(`#shipping`, chromedp.ByID),
		chromedp.Click(`#shipping`, chromedp.ByID),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> #shipping IS Click")),
		chromedp.Sleep(time.Millisecond * 1000),

		chromedp.WaitVisible(`body > div.mui-app > div > div.mui-page > div.mui-page-body.mui-page-body-centered > div > div > div > div > div > div.panel-body > div > div > div > div:nth-child(10) > div > div > div > div > div:nth-child(3) > a`, chromedp.ByQuery),
		chromedp.Click(`body > div.mui-app > div > div.mui-page > div.mui-page-body.mui-page-body-centered > div > div > div > div > div > div.panel-body > div > div > div > div:nth-child(10) > div > div > div > div > div:nth-child(3) > a`, chromedp.ByQuery),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> body > div.mui-app > div > div.mui-page > div.mui-page-body.mui-page-body-centered > div > div > div > div > div > div.panel-body > div > div > div > div:nth-child(8) > div > div > div > div > div:nth-child(3) > a IS Click")),

		chromedp.Sleep(time.Millisecond * 1000),

		chromedp.WaitVisible(".rate-enable", chromedp.ByQueryAll),
		//chromedp.SetValue(".rate-enable", kb.ArrowUp, chromedp.ByQueryAll),
		chromedp.SetValue(".rate-enable", "ENABLED_OVERRIDED", chromedp.ByQueryAll),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> #shopShippingFragment .rate-enable > a IS Click")),

		chromedp.Sleep(time.Millisecond * 1000),

		chromedp.WaitVisible(".input-number", chromedp.ByQueryAll),
		chromedp.SetValue(".input-number", "9.95", chromedp.ByQueryAll),

		//chromedp.WaitVisible("#ratesForm > form > div > div.panel-footer.datatable-toolbar > button", chromedp.ByQuery),
		//chromedp.Click("#ratesForm > form > div > div.panel-footer.datatable-toolbar > button", chromedp.ByQuery),

		chromedp.Sleep(10000),
	}
}

func logAction(logStr string) func(context.Context) error {
	return func(context.Context) error {
		log.Printf(logStr)
		return nil
	}
}
