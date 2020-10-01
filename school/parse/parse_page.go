package parse

import (
	"fmt"
	"regexp"
	"school/entity"
	"strings"
)

func ParsePage(page string) ([]entity.Movie) {
	fmt.Println("=====解析开始=====")
	idReg := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(.*?)/" `)
	idSlice := idReg.FindAllStringSubmatch(page,-1)

	nameReg := regexp.MustCompile(`title="(.*?)"`)
	nameSlice := nameReg.FindAllStringSubmatch(page ,-1)
	for index,value:=range nameSlice {
		fmt.Println(index ,value[1])
	}
	descReg := regexp.MustCompile(`<p class="pl">(.*?)</p>`)
	descSlice := descReg.FindAllStringSubmatch(page ,-1)

	rateReg := regexp.MustCompile(`<span class="rating_nums">(.*?)</span>`)
	ratSlice := rateReg.FindAllStringSubmatch(page,-1)

	voteReg := regexp.MustCompile(`<span class="pl">(.*?)</span>`)
	voteSlice := voteReg.FindAllStringSubmatch(page,-1)
	for i := 0; i < len(voteSlice); i++ {
		voteSlice[i][1]	= strings.ReplaceAll(voteSlice[i][1],"(","")
		voteSlice[i][1] = strings.ReplaceAll(voteSlice[i][1],"人评价)","")
	}

	coverImgReg := regexp.MustCompile(`<img src="(.*?)" width="75" `)
	coverImgSlice := coverImgReg.FindAllStringSubmatch(page,-1)
	movies :=make([]entity.Movie,0)
	for i := 0 ;i < len(nameSlice) ; i++ {
		movie := entity.Movie{
			Id:       idSlice[i][1],
			Name:     nameSlice[i][1],
			RateNum:  ratSlice[i][1],
			VoteNum:  voteSlice[i][1],
			CoverUrl: coverImgSlice[i][1],
			Desc:     descSlice[i][1],
			}
			movies =append(movies,movie)
	}
	fmt.Println("=====解析结束=====",len(movies))
	return movies
}
