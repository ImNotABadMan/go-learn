package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	var (
		url = "https://reeneefashion.myshopify.com/admin/products/6639277768865"
	)
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
		//chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	//
	//options := chromedp.WithBrowserErrorf(func(s string, i ...interface{}) {
	//	fmt.Println(s)
	//})

	//broser, err := chromedp.NewBrowser(context.Background(), "https://reeneefashion.myshopify.com/admin/products/6639277768865", options)

	//if err != nil {
	//	log.Fatal(err)
	//}

	//broser.Execute(context.Background(), "GET", nil, nil)
	taskLogin, mapValue := taskLogin("binzheng@jiebeili.cn", "zrb555666")

	res := chromedp.Run(taskCtx,
		chromedp.Navigate(url),
		taskLogin,
	)

	fmt.Println(*mapValue["username"], " -- ", *mapValue["password"])

	fmt.Println(res)

}

func logAction(logStr string) func(context.Context) error {
	return func(context.Context) error {
		log.Printf(logStr)
		return nil
	}
}

func taskLogin(inEmail string, inPassword string) (tasks chromedp.Tasks, mapValue map[string](*string)) {

	var username, password string
	var mapValueLogin = map[string]*string{
		"username": &username,
		"password": &password,
	}

	mapValue = mapValueLogin

	tasks = chromedp.Tasks{
		chromedp.WaitVisible(`#account_email`, chromedp.ByID),
		chromedp.SendKeys("#account_email", inEmail, chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 1000),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> account_email IS Value")),

		chromedp.WaitVisible(".captcha__submit"),
		chromedp.Submit(".captcha__submit"),
		chromedp.Value("#account_email", mapValue["username"], chromedp.ByID),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> captcha__submit IS Click")),

		chromedp.Sleep(time.Millisecond * 2000),

		chromedp.WaitVisible(`#account_password`, chromedp.ByID),
		chromedp.SendKeys("#account_password", inPassword, chromedp.ByID),

		chromedp.Value("#account_password", mapValue["password"], chromedp.ByID),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> account_password IS Value")),

		chromedp.WaitVisible(".captcha__submit"),
		chromedp.Submit(".captcha__submit"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> captcha__submit IS Click")),

		chromedp.Sleep(time.Second * 15),

		chromedp.WaitVisible(`.Polaris-TextField__Input_30ock`),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> TextField__Input_30ock IS WaitVisible")),
		chromedp.Sleep(time.Second * 5),

		chromedp.Value(".Polaris-TextField__Input_30ock", mapValue["password"]),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> TextField__Input_30ock " + *mapValue["password"])),

		chromedp.Sleep(time.Second * 10),

		//chromedp.WaitVisible("body .login-description", chromedp.ByQuery),
	}

	return tasks, mapValue
}
