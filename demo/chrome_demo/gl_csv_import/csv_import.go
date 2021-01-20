package gl_csv_import

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func logAction(logStr string) func(context.Context) error {
	return func(context.Context) error {
		log.Printf(logStr)
		return nil
	}
}

func OpenChrome() {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir)

	chromedp.Run(context.Background())
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-background-networking", false),
		chromedp.Flag("disable-renderer-backgrounding", false),
		chromedp.Flag("disable-popup-blocking", false),
		chromedp.Flag("disable-ipc-flooding-protection", false),
		chromedp.Flag("disable-client-side-phishing-detection", false),
		chromedp.Flag("disable-background-timer-throttling", false),
		chromedp.WindowSize(1600, 900),
		chromedp.Flag("headless", false),
		// Like in Puppeteer.
		chromedp.Flag("hide-scrollbars", false),
		//chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var (
		email    string
		password string
	)

	taskLogin := chromedp.Tasks{
		chromedp.WaitVisible(`#email`, chromedp.ByID),
		chromedp.WaitVisible(`#password`, chromedp.ByID),
		chromedp.SendKeys("#email", "binz", chromedp.ByID),
		chromedp.SendKeys("#password", "binz123", chromedp.ByID),
		chromedp.Value("#email", &email, chromedp.ByID),
		chromedp.Value("#password", &password, chromedp.ByID),
		chromedp.WaitVisible("#login-form .btn"),
		chromedp.Submit("#login-form .btn"),
		//chromedp.WaitVisible("body .login-description", chromedp.ByQuery),
	}

	taskOpenCsv := chromedp.Tasks{
		chromedp.WaitVisible("/html/body/div/aside[1]/div/div[4]/div/div/nav/ul/li[1]"),
		chromedp.ActionFunc(logAction((">>>>>>>>>>>>>>>>>>>> Product IS VISIBLE"))),
		chromedp.Sleep(time.Millisecond * 600),
		chromedp.Click("/html/body/div/aside/div/div[4]/div/div/nav/ul/li[1]/a"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Product IS Click")),
		chromedp.Sleep(time.Millisecond * 600),

		chromedp.WaitVisible("/html/body/div/aside[1]/div/div[4]/div/div/nav/ul/li[1]/ul/li[2]/a/p"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Import Product IS VISIBLE")),
		chromedp.Click("/html/body/div/aside[1]/div/div[4]/div/div/nav/ul/li[1]/ul/li[2]/a"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Import Product IS Click")),
	}

	taskEntryCsv := chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"actionChoose\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 新增产品 IS VISIBLE")),
		chromedp.Click("//*[@id=\"actionChoose\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 新增产品 IS Click")),
		chromedp.WaitVisible("//*[@id=\"import\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 点击Gl IS VISIBLE")),
		chromedp.Click("//*[@id=\"import\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 点击Gl IS Click")),
	}

	wdPath, _ := os.Getwd()
	csvPath := wdPath + "/gl_csv_import/test-import.csv"
	fmt.Println("Csv Path: ", csvPath)

	taskGlImport := chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"csv\"]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl file IS VISIBLE")),
		chromedp.SendKeys("//*[@id=\"csv\"]", csvPath),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl file IS Ready")),

		chromedp.WaitVisible("//*[@id=\"currency\"]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl currency IS VISIBLE")),
		chromedp.SendKeys("//*[@id=\"currency\"]", kb.ArrowDown),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl currency IS Ready")),

		chromedp.WaitVisible("//*[@id=\"uploadForm\"]/div[3]/input"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Validation IS VISIBLE")),
		chromedp.Click("//*[@id=\"uploadForm\"]/div[3]/input"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Validation IS Click")),
	}

	err = chromedp.Run(taskCtx,
		chromedp.Navigate("http://v2.globaloutlet-backend.com:8011/login"),
		taskLogin,
		taskOpenCsv,
		taskEntryCsv,
		taskGlImport,
		chromedp.WaitVisible("body"),
	)

	if err != nil {
		panic(err)
	}

	path := filepath.Join(dir, "DevToolsActivePort")

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(bs, []byte("\n"))

	fmt.Printf("DevToolsActivePort has %d lines \n", len(lines))

}
