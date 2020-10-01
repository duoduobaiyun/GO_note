package parse

import (
	"lo(1)/object"
)

func PageHtml(page []object.Hero)[]object.Name {
	heroSlcie := make([]object.Name, 0)
	for i := 0; i < len(page); i++ {
		hero := object.Name{
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
	heroSlcie=append(heroSlcie,hero)
	}
	return heroSlcie
}
