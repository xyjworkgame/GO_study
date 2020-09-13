package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
@Time : 2020/9/12 23:32
@Author : Firewine
@File : picture
@Software: GoLand
@Description:  爬取bing 壁纸
*/

/*
   分析：
		1. 手机端
https://cn.bing.com/th?id=OHR.MedievalRocamadour_ZH-CN7063423495_480x800.jpg&rf=LaDigue_1920x1080.jpg&pid=hp
https://cn.bing.com/th?id=OHR.SangreCristoDunes_ZH-CN7193190503_480x800.jpg&rf=LaDigue_1920x1080.jpg&pid=hp
		2. pc
https://cn.bing.com/th?id=OHR.MedievalRocamadour_ZH-CN7063423495_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp
		12 ： 相同参数
		rf ，pid(可忽略)
		3. 获取后面链接
https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&nc=1599925134238&pid=hp
		需要参数 ： format ，ids， n  ， nc(时间戳)  pid

*/
const (
	referer         = "https://cn.bing.com/"
	foundUrl        = "https://cn.bing.com/HPImageArchive.aspx"
	userAgentPc     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0"
	userAgentMobile = "Mozilla/5.0 (Linux; Android 7.0; SM-G892A Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/67.0.3396.87 Mobile Safari/537.36"
	sizePC = "1920x1080"
	sizeMobile = "480x800"
	)


var InputPc = flag.String("pcSize",sizePC , "input size of pc eg: 1920x1080 or 1920*1080")
var InputMobile = flag.String("mobileSize", sizeMobile, "input size of mobile eg: 480x800 or 480*800")


// 爬壁纸，两种手机，pc
func GetPicture(mobile string,pc string) {
	// client 客户端
	client := http.Client{}
	fullUrl := getBase(client)
	getPc(client, fullUrl,pc)
	getMobile(client, fullUrl,mobile)
	fmt.Println("获取完成")
}
func getBase(client http.Client) (fullUrl string) {
	// 添加参数
	params := url.Values{}
	params.Add("format", "js")
	params.Add("ids", "0")
	params.Add("n", "1")
	params.Add("nc", string(time.Now().UnixNano()))
	params.Add("pid", "pd")

	// 生成 url
	Url, err := url.Parse(foundUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	Url.RawQuery = params.Encode()

	request, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("Referer", "https://cn.bing.com/")
	request.Header.Add("Host", "cn.bing.com")
	request.Header.Add("User-Agent", userAgentPc)

	resp, _ := client.Do(request)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf(string(body))
	maps := map[string]interface{}{}
	_ = json.Unmarshal(body, &maps)
	// 获取image meta
	image := maps["images"].([]interface{})[0].(map[string]interface{})
	fullUrl = image["url"].(string)
	//urlBase = image["urlbase"].(string)

	return fullUrl
}
func getPc(client http.Client, fullUrl string, pc string) {
	uurl := referer + fullUrl
	// 生成 url
	Url, err := url.Parse(uurl)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//req.Response.Header.Add("content-type","image/jpeg")
	//fmt.Println(Url.String())
	resp, _ := client.Do(req)
	resp.Header.Add("content-type", "image/jpeg")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(time.Now().Format("20060102_PC")+".jpeg", body, 644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func getMobile(client http.Client, fullUrl string, mobile string) {

	replace := strings.Replace(fullUrl, "1920x1080", mobile, 1)
	uurl := referer + replace
	// 生成 url
	Url, err := url.Parse(uurl)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", userAgentPc)
	//req.Response.Header.Add("content-type","image/jpeg")
	//fmt.Println(Url.String())
	resp, _ := client.Do(req)
	resp.Header.Add("content-type", "image/jpeg")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(time.Now().Format("20060102_mobile")+".jpeg", body, 644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	flag.Parse()
	mobileSize := ""
	pcSize := ""
	if *InputMobile != ""{
		if strings.Contains(*InputMobile,"*"){
			mobileSize = strings.Replace(*InputMobile, "*", "x", 1)
		}
		mobileSize = *InputMobile

	}else {
		mobileSize = sizeMobile
	}
	if *InputPc != ""{
		if strings.Contains(*InputPc,"*"){
			pcSize = strings.Replace(*InputPc, "*", "x", 1)
		}
		pcSize = *InputPc
	}else {
		pcSize = sizePC
	}
	GetPicture(mobileSize,pcSize)
}
