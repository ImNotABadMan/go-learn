package gl_csv_import

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Path struct {
	PathStr string `json:"path"`
}

func GetCsvPath(pathStr string) Path {
	var path Path
	wdPath, _ := os.Getwd()
	configFullPath := wdPath + pathStr

	b, err := ioutil.ReadFile(configFullPath)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bytes.Buffer{}
	buffer.Write(b)

	decoder := json.NewDecoder(&buffer)
	if err := decoder.Decode(&path); err != nil {
		log.Fatal(err)
	}
	return path
}

func logAction(logStr string) func(context.Context) error {
	return func(context.Context) error {
		log.Printf(logStr)
		return nil
	}
}

func ReStartCsvQueue() {
	// linux
	fmt.Println("/home/ubuntu/jenkins/gl/kill.sh")
	killCmd := exec.Command("/home/ubuntu/jenkins/gl/kill.sh")
	if err := killCmd.Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(killCmd.Stdout)

	fmt.Println("/home/ubuntu/jenkins/gl/test.sh")
	cmd := exec.Command("/home/ubuntu/jenkins/gl/test.sh")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(cmd.Stdout)
}

func OpenChrome(inEmail string, inPassword string, configPath string) {
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
		//chromedp.WindowSize(1200, 800),
		chromedp.WindowSize(1500, 900),
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
	email, password := *mapValue["username"], *mapValue["password"]
	fmt.Println("username:", email)
	fmt.Println("password:", password)

	//taskOpenMenuCsv, taskEntryCsv, taskImport, taskClickImport := taskImportGlCsv(configPath)
	taskOpenMenuCsv, taskEntryCsv, taskImport, taskClickImport := taskImportShopifyCsv(configPath)

	wdPath, _ := os.Getwd()
	//csvFullPath := wdPath + "/gl_csv_import/test-import.csv"
	//configFullPath := wdPath + configPath
	pathStruct := GetCsvPath(configPath)
	csvPath := wdPath + pathStruct.PathStr

	fmt.Println("Csv Path: ", csvPath)

	err = chromedp.Run(taskCtx,
		//chromedp.Navigate("http://v2.globaloutlet-backend.com:8011/login"),
		chromedp.Navigate("http://192.168.10.113:8011/login"),
		taskLogin,
		taskOpenMenuCsv,
		taskEntryCsv,
		taskImport,
		taskClickImport,
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

	var username, password string
	var mapValueLogin = map[string]*string{
		"username": &username,
		"password": &password,
	}

	mapValue = mapValueLogin

	tasks = chromedp.Tasks{
		chromedp.WaitVisible(`#username`, chromedp.ByID),
		chromedp.WaitVisible(`#password`, chromedp.ByID),
		chromedp.SendKeys("#username", inEmail, chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 1000),
		chromedp.SendKeys("#password", inPassword, chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 2000),
		chromedp.Value("#username", mapValue["username"], chromedp.ByID),
		chromedp.Value("#password", mapValue["password"], chromedp.ByID),
		chromedp.WaitVisible("#login-form .btn"),
		chromedp.Submit("#login-form .btn"),
		//chromedp.WaitVisible("body .login-description", chromedp.ByQuery),
	}

	return tasks, mapValue
}

func taskImportGlCsv(configPath string) (inTaskOpenMenuCsv chromedp.Tasks, inTaskEntryCsv chromedp.Tasks,
	inTaskGlImport chromedp.Tasks, inTaskClickGlImport chromedp.Tasks) {

	taskOpenMenuCsv := taskOpenMenuCsv()
	taskEntryCsv := taskEntryGlCsv()

	wdPath, _ := os.Getwd()
	//csvFullPath := wdPath + "/gl_csv_import/test-import.csv"
	//configFullPath := wdPath + configPath
	pathStruct := GetCsvPath(configPath)
	csvPath := wdPath + pathStruct.PathStr

	fmt.Println("Csv Path: ", csvPath)

	taskGlImport := taskGlImport(csvPath)
	taskClickGlImport := taskClickGlImport()

	return taskOpenMenuCsv, taskEntryCsv, taskGlImport, taskClickGlImport
}

func taskImportShopifyCsv(configPath string) (inTaskOpenMenuCsv chromedp.Tasks, inTaskEntryCsv chromedp.Tasks,
	inTaskShopifyImport chromedp.Tasks, inTaskClickShopifyImport chromedp.Tasks) {

	taskOpenMenuCsv := taskOpenMenuCsv()
	taskEntryCsv := taskEntryShopifyCsv()

	wdPath, _ := os.Getwd()
	//csvFullPath := wdPath + "/gl_csv_import/test-import.csv"
	//configFullPath := wdPath + configPath
	pathStruct := GetCsvPath(configPath)
	csvPath := wdPath + pathStruct.PathStr

	fmt.Println("Csv Path: ", csvPath)

	taskShopifyImport := taskShopifyImport(csvPath)
	taskClickShopifyImport := taskClickShopifyImport()

	return taskOpenMenuCsv, taskEntryCsv, taskShopifyImport, taskClickShopifyImport
}

func taskOpenMenuCsv() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible("/html/body/div/aside[1]/div/div[4]/div/div/nav/ul/li[1]"),
		chromedp.ActionFunc(logAction((">>>>>>>>>>>>>>>>>>>> Product IS VISIBLE"))),
		chromedp.Sleep(time.Millisecond * 1500),
		chromedp.Click("/html/body/div/aside/div/div[4]/div/div/nav/ul/li[1]/a"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Product IS Click")),
		chromedp.Sleep(time.Millisecond * 2500),

		chromedp.WaitVisible("/html/body/div/aside[1]/div/div[4]/div/div/nav/ul/li[1]/ul/li[3]/a/p"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Import Product IS VISIBLE")),
		chromedp.Click("/html/body/div/aside[1]/div/div[4]/div/div/nav/ul/li[1]/ul/li[3]/a"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Import Product IS Click")),
	}
}

func taskEntryGlCsv() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"actionChoose\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 新增产品 IS VISIBLE")),
		chromedp.Click("//*[@id=\"actionChoose\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 新增产品 IS Click")),
		chromedp.Sleep(time.Millisecond * 3000),

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
		chromedp.Sleep(time.Millisecond * 2000),

		chromedp.WaitVisible("//*[@id=\"currency\"]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl currency IS VISIBLE")),
		chromedp.SendKeys("//*[@id=\"currency\"]", kb.ArrowDown),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> gl currency IS Ready")),
		chromedp.Sleep(time.Millisecond * 2000),

		chromedp.WaitVisible("//*[@id=\"uploadForm\"]/div[4]/input[2]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Validation IS VISIBLE")),
		chromedp.Click("//*[@id=\"uploadForm\"]/div[4]/input[2]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Validation IS Click")),
	}
}

func taskClickGlImport() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"importConfirmModal\"]/div/div/div[4]/input"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Import IS VISIBLE")),
		chromedp.Sleep(time.Second * 5),
		chromedp.Click("//*[@id=\"importConfirmModal\"]/div/div/div[4]/input"),
		//chromedp.Click("//*[@id=\"importConfirmModal\"]/div/div/div[4]/button[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Import IS Click")),
	}
}

func taskEntryShopifyCsv() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"actionChoose\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 新增产品 IS VISIBLE")),
		chromedp.Click("//*[@id=\"actionChoose\"]/div[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 新增产品 IS Click")),
		chromedp.Sleep(time.Millisecond * 3000),

		chromedp.WaitVisible("//*[@id=\"import\"]/div[3]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 点击Shopify //*[@id=\"import\"]/div[3] IS VISIBLE")),
		chromedp.Click("//*[@id=\"import\"]/div[3]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> 点击Shopify //*[@id=\"import\"]/div[3] IS Click")),
	}
}

func taskShopifyImport(csvPath string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"csv\"]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> shopify file IS VISIBLE")),
		chromedp.SendKeys("//*[@id=\"csv\"]", csvPath),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> shopify file IS Ready")),
		chromedp.Sleep(time.Millisecond * 2000),

		chromedp.WaitVisible("//*[@id=\"currency\"]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> shopify currency IS VISIBLE")),
		chromedp.SendKeys("//*[@id=\"currency\"]", kb.ArrowDown+kb.ArrowDown),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> shopify currency IS Ready")),
		chromedp.Sleep(time.Millisecond * 2000),

		chromedp.WaitVisible("//*[@id=\"uploadForm\"]/div[4]/input[2]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Validation IS VISIBLE")),
		chromedp.Click("//*[@id=\"uploadForm\"]/div[4]/input[2]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Validation IS Click")),
	}
}

func taskClickShopifyImport() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible("//*[@id=\"importConfirmModal\"]/div/div/div[4]/input"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Import IS VISIBLE")),
		chromedp.Sleep(time.Second * 5),
		chromedp.Click("//*[@id=\"importConfirmModal\"]/div/div/div[4]/input"),
		//chromedp.Click("//*[@id=\"importConfirmModal\"]/div/div/div[4]/button[1]"),
		chromedp.ActionFunc(logAction(">>>>>>>>>>>>>>>>>>>> Button Import IS Click")),
	}
}

//#### 1
// docker run --name=pms_mysql \
//  --network=ubuntu_static_22 \
//  --entrypoint /bin/bash \
//  --ip=172.22.0.2 \
//  -p 3301:3306 -u root \
//  -e MYSQL_ROOT_PASSWORD=123 \
//  -e MYSQL_ROOT_HOST=172.22.0.2 \
//  -e MYSQL_DATABASE=globaloutletcom \
//  -e MYSQL_USER=root \
//  -e MYSQL_PASSWOR=123 \
//  -d mysql:5.7 \
//  -c 'exec echo sql_mode="STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION" | tee -a /etc/mysql/mysql.conf.d/mysqld.cnf ; cat /etc/mysql/mysql.conf.d/mysqld.cnf;cat entrypoint.sh; docker-entrypoint.sh mysqld;'

//#### 2
//docker exec -i pms_mysql sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD" --database=globaloutletcom' < /home/ubuntu/dockers/volumes/pms_mysql/gl_innodb_bak.sql

//#### 3
//docker exec -i pms_mysql sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD" --database=globaloutletcom -e "GRANT ALL PRIVILEGES ON *.* TO "root"@"%" IDENTIFIED BY "123";FLUSH PRIVILEGES;"'
