package impl

import (
	"context"
	"encoding/xml"
	"fmt"
	"home.hourong.me/HenryHou/drone-baidu-sitemap/pkg/controller"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type SiteMapControllerImpl struct {
}

type SiteMap struct {
	XMLName xml.Name `xml:"urlset"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLS    []struct {
		Text       string `xml:",chardata"`
		Loc        string `xml:"loc"`
		Lastmod    string `xml:"lastmod"`
		Priority   string `xml:"priority"`
		Changefreq string `xml:"changefreq"`
	} `xml:"url"`
}

var _ controller.SiteMapController = (*SiteMapControllerImpl)(nil)

var DefaultSiteMapController *SiteMapControllerImpl

func init() {
	DefaultSiteMapController = NewSiteMapController()
}

func NewSiteMapController() *SiteMapControllerImpl {
	return &SiteMapControllerImpl{}
}

func (c *SiteMapControllerImpl) GetSiteMapUrls(ctx context.Context, path string) (urls []string) {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var siteMap SiteMap
	xml.Unmarshal(byteValue, &siteMap)

	for _, url := range siteMap.URLS {
		urls = append(urls, url.Loc)
	}
	return
}

func (c *SiteMapControllerImpl) SendData(ctx context.Context, site, token, data string) {
	body := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://data.zz.baidu.com/urls?site=%s&token=%s", site, token), body)
	req.Header.Add("Content-Type", "text/plain")
	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failure : ", err)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("百度 URL 数据更新成功，具体信息：", string(respBody))
}
