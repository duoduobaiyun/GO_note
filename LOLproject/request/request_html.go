package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Request_html(page string)(string , error)  {
	client := http.Client{}
	request, err := http.NewRequest("GET", page, nil)
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
	return string(herohtml),err
}