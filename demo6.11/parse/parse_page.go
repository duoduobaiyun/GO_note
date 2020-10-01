package parse

import (
	"DoubanNewMovie/entity"
	"fmt"
	"regexp"
	"strings"
)

/***
  功能：将特定网页数据作为数据源，从数据源中解析除符合自己需求的数据
 */
func ParsePage(page string)([]entity.Movie){
	fmt.Println("网页解析函数")
	//电影编号、电影名、简介、评分、评价人数、封面图

	//编号
	idReg := regexp.MustCompile(`<a class="nbg" href="https://movie.douban.com/subject/(.*?)/"`)
	idSlice := idReg.FindAllStringSubmatch(page,-1)

	nameReg := regexp.MustCompile(`title="(.*?)"`)
	//返回数据格式：[][]string
	nameSlice := nameReg.FindAllStringSubmatch(page,-1)//-1表示在整个数据源内进行采集

	//description：描述，简介
	descReg := regexp.MustCompile(`<p class="pl">(.*?)</p>`)
	descSlice := descReg.FindAllStringSubmatch(page,-1)

	//rate:打分,分数
	rateReg := regexp.MustCompile(`<span class="rating_nums">(.*?)</span>`)
	rateSlice := rateReg.FindAllStringSubmatch(page,-1)

	//评价人数 vote:投票
	voteReg := regexp.MustCompile(`<span class="pl">(.*?)</span>`)
	voteSlice := voteReg.FindAllStringSubmatch(page,-1)
	for i :=0;i<len(voteSlice);i++{
		//原始格式：(31052人评价)
		voteSlice[i][1] = strings.ReplaceAll(voteSlice[i][1],"(","")// 31052人评价)
		voteSlice[i][1] = strings.ReplaceAll(voteSlice[i][1],"人评价)","")//31052
	}

	//封面图： cover覆盖
	coverImgReg := regexp.MustCompile(`<img src="(.*?)" width="75" alt`)
	coverImgSlice := coverImgReg.FindAllStringSubmatch(page,-1)

	movies := make([]entity.Movie,0)
	for i :=0; i <len(nameSlice); i++{
		movie := entity.Movie{
			Id:       idSlice[i][1],
			Name:     nameSlice[i][1],
			RateNum:  rateSlice[i][1],
			VoteNum:  voteSlice[i][1],
			CoverUrl: coverImgSlice[i][1],
			Desc:     descSlice[i][1],
		}
		//追加
		movies = append(movies,movie)
	}
	return movies
}
