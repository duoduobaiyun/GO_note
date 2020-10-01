package network_request

import (
	"LeagueOfLegends/object"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetHerolist (url string) (*object.HeroList,error) {
	//1.1获取网络数据
	client := http.Client{}
	//
	request, err := http.NewRequest("GET",url,nil)
	if err != nil {
		return nil,err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil,err
	}

	heroData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil,err
	}
	//fmt.Println(string(heroData))
	var herolist object.HeroList

	//1.2解析网络数据
	err = json.Unmarshal(heroData,&herolist)
	if err != nil {
		return nil,err
	}
	return &herolist,nil
}
