package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func RequestHtmlPage(page_url string) (string,error) {
	fmt.Println("=====开始爬取=====")
	client := http.Client{}
	request, err := http.NewRequest("GET",page_url,nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	request.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
	request.Header.Add("Cache-Control","max-age=0")
	request.Header.Add("Connection","keep-alive")
	request.Header.Add("Host","movie.douban.com")
	request.Header.Add("Sec-Fetch-Mode","navigate")
	request.Header.Add("Sec-Fetch-Site","none")
	request.Header.Add("Sec-Fetch-User","?1")
	request.Header.Add("Upgrade-Insecure-Requests","1")
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.87 Safari/537.36 SLBrowser/6.0.1.4221")
	request.Header.Add("Cookie","bid=53_pKkPE_S8; _pk_id.100001.4cf6=9654b327ad873bf8.1591360810.1.1591360810.1591360810.; _pk_ses.100001.4cf6=*; ap_v=0,6.0; __utma=30149280.1534143555.1591360810.1591360810.1591360810.1; __utmb=30149280.0.10.1591360810; __utmc=30149280; __utmz=30149280.1591360810.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utma=223695111.1657223101.1591360810.1591360810.1591360810.1; __utmb=223695111.0.10.1591360810; __utmc=223695111; __utmz=223695111.1591360810.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __yadk_uid=2osgLwFtl4cfrl36BVl9IjkZIkoOdG4V")

	resp,err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	pageByes,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}

	html := string(pageByes)
	//获取页面信息

	//电影名称、
	nameRG := regexp.MustCompile(`alt="(.*?)"`)
	nameSlice := nameRG.FindAllStringSubmatch(html, -1)

	//电影评分、
	gradeRG := regexp.MustCompile(`<span class="rating_nums">(.*?)</span>`)
	gradeSlice := gradeRG.FindAllStringSubmatch(html,-1)

	//评价人数、
	voteRG := regexp.MustCompile(`<span class="pl">(.*?)</span>`)
	voteSlice := voteRG.FindAllStringSubmatch(html,-1)
	for i := 0; i < len(nameSlice)-1; i++ {
		fmt.Printf("电影名是：%s   评分：%s分   %s\n",
			nameSlice[i][1],//名字
			gradeSlice[i][1],//评分
			voteSlice[i][1],//评价人数
		)
	}
	fmt.Println("=====爬取结束=====")
	return string(pageByes),nil
}



