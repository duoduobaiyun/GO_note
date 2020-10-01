package bean

type IndexData struct {
	//是否需要采集数据
	IsNeedSpiderData bool
	//英雄信息
	Heros []Hero
	//一共有多少页数据
	PageNum      int64
	IsPageButton bool //是否有翻页器
	CurrentNum   int
}
