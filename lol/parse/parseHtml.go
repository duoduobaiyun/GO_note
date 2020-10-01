package parse

import (
	"project1/LoL/object"
)

func Pagehtml(page []object.Hero) []object.Loldatabase {
	heros:=make([]object.Loldatabase,0)
	for i:=0;i<=len(page)-1 ; i++ {
		hero:=object.Loldatabase{
			Id:      page[i].HeroId,
			Name:     page[i].Name,
			Alias:    page[i].Alias,
			Title:    page[i].Title,
			VoiceOne: page[i].SelectAudio,
			VoiceTwo: page[i].BanAudio,
			Label:    page[i].ChangeLabel,
			ImgUrl:   "//game.gtimg.cn/images/lol/act/img/champion/"+page[i].Alias+".png",
		}
		heros =append(heros,hero)

	}
	return  heros
}