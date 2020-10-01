package parse

import "project1/lo/object"

func PageHtml(page []object.Hero)[]object.Names  {
	heroes:=make([]object.Names,0)
	for i:=0;i<len(page) ; i++ {
		hero:=object.Names{
			Id:         page[i].HeroId,
			ImgUrl:     "//game.gtimg.cn/images/lol/act/img/champion/"+page[i].Alias+".png",
			Name:       page[i].Name,
			Alias:      page[i].Alias,
			Title:      page[i].Title,
			SelectAudio:page[i].SelectAudio,
			Attack:     page[i].Attack,
			Defense:    page[i].Defense,
			Magic:      page[i].Magic,
		}
		heroes=append(heroes,hero)
	}
	return heroes
}