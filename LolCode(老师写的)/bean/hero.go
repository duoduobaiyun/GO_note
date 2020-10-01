package bean

type Hero struct {
	HeroId      string `json:"heroId"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Avatar      string `json:"avatar"`
	Title       string `json:"title"`
	SelectAudio string `json:"selectAudio"`
	BanAudio    string `json:"banAudio"`
}
