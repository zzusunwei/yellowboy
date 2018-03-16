package main


import (
	"fmt"
	"log"
	"github.com/sclevine/agouti"
)

var (
	PORT = 8080
    baseUrl = fmt.Sprintf("http://localhost:%v/admin", PORT)
    driver  *agouti.WebDriver
    page    *agouti.Page
)
func open(url string) {
    var err error
	driver = agouti.ChromeDriver() // 设置 driver
    driver.Start()
	page, err = driver.NewPage() // 初始化页面对象
	if err != nil{
		log.Println("get page error", err)
	}
	page.Navigate(baseUrl)
	page.Reset()
}

func close(){
	driver.Stop() // 关闭 driver
}