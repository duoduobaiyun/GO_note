package dao

import (
	"LolCode/util"
	"LolCode/bean"
	"strconv"
)

/**
 * 查询英雄的个数
 */
func QueryHeroCount() (int64, error) {
	row := util.LolDB.QueryRow("select count(hero_id) hero_count from hero")
	var count int64
	err := row.Scan(&count)
	return count, err
}

/**
 * 查询所有的英雄
 */
func QueryAllHeros() ([]bean.Hero, error) {
	rows, err := util.LolDB.Query("select * from hero order by hero_id asc")
	if err != nil {
		return nil, err
	}
	heros := make([]bean.Hero, 0)
	for rows.Next() {
		var hero bean.Hero
		err = rows.Scan(&hero.HeroId, &hero.Name, &hero.Alias, &hero.Title, &hero.SelectAudio, &hero.BanAudio)
		if err != nil {
			return nil, err
		}
		hero.Avatar = "https://game.gtimg.cn/images/lol/act/img/champion/" + hero.Alias + ".png"
		heros = append(heros, hero)
	}
	return heros, nil
}

func QueryHeros(start, end int) ([]bean.Hero, error) {
	rows, err := util.LolDB.Query("select * from lol_info.hero order by hero_id asc limit ?, ?", start, end)
	if err != nil {
		return nil, err
	}
	heros := make([]bean.Hero, 0)
	for rows.Next() {
		var hero bean.Hero
		err = rows.Scan(&hero.HeroId, &hero.Name, &hero.Alias, &hero.Title, &hero.SelectAudio, &hero.BanAudio)
		if err != nil {
			return nil, err
		}
		hero.Avatar = "https://game.gtimg.cn/images/lol/act/img/champion/" + hero.Alias + ".png"
		heros = append(heros, hero)
	}
	return heros, nil
}

/**
 * 保存英雄信息到数据库
 */
func SaveHero(hero bean.Hero) (int64, error) {

	//"1" "12" "a"
	id, err := strconv.Atoi(hero.HeroId)
	if err != nil {
		return 0, err
	}
	result, err := util.LolDB.Exec("insert into hero"+
		"(hero_id, name, alias, title, select_audio, ban_audio) "+
		"values"+
		"(?, ?, ?, ?, ?, ?)",
		id, hero.Name, hero.Alias, hero.Title, hero.SelectAudio, hero.BanAudio)

	if err != nil {
		return 0, err
	}
	rowsNum, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsNum, nil
}

/**
 * 删除英雄
 */
func DeleteHero(heroId int) (int64, error) {
	res, err := util.LolDB.Exec("delete from hero where hero_id = ?", heroId)
	if err != nil {
		return 0, err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return row, nil
}
