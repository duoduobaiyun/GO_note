package parse

import (
	"LOLproject/object"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Hero_html(pages string)[]object.Hero  {
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js", nil)
	if err != nil {
		fmt.Println(request)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(request)
	}
	herohtml, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var hero object.HeroList
	err = json.Unmarshal(herohtml, &hero)
	if err != nil {
		fmt.Println(err.Error())
	}
	herolols := make([]object.Hero, 0)
	for i := 0; i < 148; i++ {
		herolol := object.Hero{
			HeroId:              hero.Hero[i].HeroId,
			Name:                hero.Hero[i].Name,
			Alias:               hero.Hero[i].Alias,
			Title:               hero.Hero[i].Title,
			IsWeekFree:          hero.Hero[i].IsWeekFree,
			SelectAudio:         hero.Hero[i].SelectAudio,
			BanAudio:            hero.Hero[i].BanAudio,
		}
		herolols = append(herolols, herolol)
	}
	return herolols
}
