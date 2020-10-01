package network_request1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lol(new)/object1"
	"net/http"
)

func GetHeroList(url string) (*object1.HeroList, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	herodata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var herolist object1.HeroList
	err =json.Unmarshal(herodata,&herolist)
	if err!=nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return &herolist,nil
}
