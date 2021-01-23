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

func OpenChrome(inEmail string, inPassword string) {
	defer func() {

	}()
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
		chromedp.WindowSize(1200, 800),
		chromedp.Flag("headless", false),
		// Like in Puppeteer.
		chromedp.Flag("hide-scrollbars", false),
		//chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer func() {
		fmt.Println("close browser")
		cancel()
	}()

	taskLogin, mapValue := taskLogin(inEmail, inPassword)
	email, password := *mapValue["email"], *mapValue["password"]
	fmt.Println("email:", email)
	fmt.Println("password:", password)

	taskOpenMenuCsv := taskOpenMenuCsv()
	taskEntryCsv := taskEntryCsv()

	wdPath, _ := os.Getwd()
	csvPath := wdPath + "/gl_csv_import/test-import.csv"
	fmt.Println("Csv Path: ", csvPath)

	taskGlImport := taskGlImport(csvPath)

	err = chromedp.Run(taskCtx,
		chromedp.Navigate("http://vm.globaloutlet-backend.com:8011/login"),
		taskLogin,
		taskOpenMenuCsv,
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

func taskLogin(inEmail string, inPassword string) (tasks chromedp.Tasks, mapValue map[string](*string)) {

	var email, password string
	var mapValueLogin = map[string]*string{
		"email":    &email,
		"password": &password,
	}

	mapValue = mapValueLogin

	tasks = chromedp.Tasks{
		chromedp.WaitVisible(`#email`, chromedp.ByID),
		chromedp.WaitVisible(`#password`, chromedp.ByID),
		chromedp.SendKeys("#email", inEmail, chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 600),
		chromedp.SendKeys("#password", inPassword, chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 600),
		chromedp.Value("#email", mapValue["email"], chromedp.ByID),
		chromedp.Value("#password", mapValue["password"], chromedp.ByID),
		chromedp.WaitVisible("#login-form .btn"),
		chromedp.Submit("#login-form .btn"),
		//chromedp.WaitVisible("body .login-description", chromedp.ByQuery),
	}

	return tasks, mapValue
}

func taskOpenMenuCsv() chromedp.Tasks {
	return chromedp.Tasks{
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
}

func taskEntryCsv() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"actionChoose\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 新增产品 IS VISIBLE")),
		chromedp.Click("//*[@id=\"actionChoose\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 新增产品 IS Click")),
		chromedp.Sleep(time.Millisecond * 600),

		chromedp.WaitVisible("//*[@id=\"import\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 点击Gl IS VISIBLE")),
		chromedp.Click("//*[@id=\"import\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 点击Gl IS Click")),
	}
}

func taskGlImport(csvPath string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"csv\"]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl file IS VISIBLE")),
		chromedp.SendKeys("//*[@id=\"csv\"]", csvPath),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl file IS Ready")),
		chromedp.Sleep(time.Millisecond * 600),

		chromedp.WaitVisible("//*[@id=\"currency\"]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl currency IS VISIBLE")),
		chromedp.SendKeys("//*[@id=\"currency\"]", kb.ArrowDown),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl currency IS Ready")),
		chromedp.Sleep(time.Millisecond * 600),

		chromedp.WaitVisible("//*[@id=\"uploadForm\"]/div[3]/input"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Validation IS VISIBLE")),
		chromedp.Click("//*[@id=\"uploadForm\"]/div[3]/input"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Validation IS Click")),
	}
}
