package parse

import (
	_struct "demo6.12/struct"
	"fmt"
	"regexp"
)

func Parse(html string)([]_struct.Movie)  {
	fmt.Println("网页解析函数")
    //编号，片名，评分，评分人数，描述，图片

    //编号
    IdMovie:=regexp.MustCompile(`<a href="https://movie.douban.com/subject/(.*?)/">`)
    IdMovieSlice:=IdMovie.FindAllStringSubmatch(html,-1)

    //片名
    NameMovie:=regexp.MustCompile(` <img width="100" alt="(.*?)"`)
    NameMovieSlice:=NameMovie.FindAllStringSubmatch(html,-1)

    //评分
	RatRegMovie:=regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	RatRegMovieSlice:=RatRegMovie.FindAllStringSubmatch(html,-1)

    //评分人数
    VoteMovie:=regexp.MustCompile(` <span>(.*?)人评价</span>`)
    VoteMovieSlice:=VoteMovie.FindAllStringSubmatch(html,-1)

    //描述
    DescMovie:=regexp.MustCompile(` <span class="inq">(.*?)</span>`)
    DescMovieSlice:=DescMovie.FindAllStringSubmatch(html,-1)

    //图片
    ImageMovie:=regexp.MustCompile(`src="(.*?)" class="">`)
    ImageMovieSlice:=ImageMovie.FindAllStringSubmatch(html,-1)

    movies:=make([]_struct.Movie,0)//创建了一个容器一样的东西，用来存放数据
	for i:=0;i<len(NameMovieSlice) ;i++  {
		movie:=_struct.Movie{//这里面就是爬取到的数据
			Id:     IdMovieSlice[i][1],
			Name:   NameMovieSlice[i][1],
			RatReg: RatRegMovieSlice[i][1],
			Vote:   VoteMovieSlice[i][1],
			Desc:   DescMovieSlice[i][1],
			Image:  ImageMovieSlice[i][1],
		}
		//追加
		movies=append(movies,movie)
	}
	return movies
}
