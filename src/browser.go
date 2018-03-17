package main


import (
	"fmt"
	"log"
	"github.com/sclevine/agouti"
)

var (
    baseUrl = fmt.Sprintf("http://localhost%v/index", PORT)
    driver  *agouti.WebDriver
    page    *agouti.Page
)
func openBrowser(url string) {
	log.Println("in method openBrowser",baseUrl)
    var err error
	driver = agouti.ChromeDriver() // 设置 driver
    driver.Start()
	page, err = driver.NewPage() // 初始化页面对象
	if err != nil{
		log.Println("get page error", err)
	}
	page.Navigate("www.baidu.com")
	page.Reset()
}

func closeBrowser(){
	driver.Stop() // 关闭 driver
}

func refeshBrowser(){
	driver.Stop() // 关闭 driver
}