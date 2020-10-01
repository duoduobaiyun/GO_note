package object

type Hero struct {
	HeroId string //英雄ID
	Name string //英雄中文名
	Alias string //英语名
	Title string //昵称
	Roles []string //角色类型
	Attack string //攻击力
	Defense string //防守力
	Magic string //魔力
	Difficulty string //操作难度
	SelectAudio string //选择语音
	BanAudio string //禁选语音
	MainImg  string//主皮肤
	IconImg []string//副皮肤
}
