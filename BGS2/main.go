
package main

import (
"BGS2/entity"
"fmt"
"html/template"
"io/ioutil"
"log"
"net/http"
"regexp"
"strconv"
"strings"
)
var imgAll entity.IMG

func main() {

	fmt.Println("开始咯。。。")
	client:=http.Client{}
	for i:=0;i <10 ;i++  {
		var request  *http.Request
		var err error
		if i==0 {
			request,err=http.NewRequest("GET","https://www.gamersky.com/ent/202005/1288085.shtml",nil)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
		}else {
			b:=strconv.Itoa(i+1)
			request,err=http.NewRequest("GET","https://www.gamersky.com/ent/202005/1288085_"+b+".shtml",nil)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
		}

		//request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		//request.Header.Add("Accept-Encoding","gzip, deflate, br")
		//request.Header.Add("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		//request.Header.Add("Connection","keep-alive")
		//request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36 Edg/83.0.478.56")
		//request.Header.Add("Cache-Control","max-age=0")

		response,err1:=client.Do(request)
		if err1 != nil {
			log.Fatal(err1.Error())
			return
		}
		htmlBytes,err2:=ioutil.ReadAll(response.Body)
		if err2 != nil {
			log.Fatal(err2.Error())
			return
		}
		html:=string(htmlBytes)
		imgAddressReg:=regexp.MustCompile(`<p align="center"><a target="_blank" href="(.*?)" `)
		imgAddressSlice:=imgAddressReg.FindAllStringSubmatch(html,-1)

		for j:=0;j<len(imgAddressSlice) ;j++  {
			imgUrl1 := strings.Split(imgAddressSlice[j][1],"?")
			imgAddressSlice[j][1] = imgUrl1[1]
			imgUrl2:=strings.Split(imgAddressSlice[j][1],">")
			imgAddressSlice[j][1]=imgUrl2[0]
			/*imgUrl3:=strings.Split(imgAddressSlice[j][1],"'")
			imgAddressSlice[j][1]=imgUrl3[0]*/

			url := imgAddressSlice[j][1]
			if strings.HasSuffix(url,`"`){
				imgAddressSlice[j][1] = strings.ReplaceAll(url,`"`,"")
			}
			fmt.Println(j,imgAddressSlice[j][1])
			imgAll.Address = append(imgAll.Address,imgAddressSlice[j][1])
		}
		/*database,err3:=sql.Open("mysql","root:MYSQL2020@tcp(127.0.0.1:3306)/staff?charset=utf8")
		defer database.Close()
		if err3 != nil {
			log.Fatal(err3.Error())
			return
		}
		for n:=0;n<len(imgAddressSlice) ;n++  {
			_,err4:=database.Exec("insert into "+"beautiful(img)"+"value(?)",imgAddressSlice[n][1])
			if err4 != nil {
				log.Fatal(err4.Error())
				break
			}
		}*/

	}
	http.HandleFunc("/beautiful",index)
	err5:=http.ListenAndServe(":1003",nil)
	if err5 != nil {
		fmt.Println(err5.Error())
	}
	//fmt.Println("结束啦！")

}
func index(w http.ResponseWriter,req *http.Request){
	t,err:=template.ParseFiles("./views/show.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Execute(w,imgAll)
}
/*func queryCount()(int,error){
	row:=*sql.DB.QueryRow("select COUNT(beautiful.img) count from beautiful")
	var count int
	err:=row.Scan(&count)
	if err != nil {
		return 0,err
	}
	return count,nil
}*/
