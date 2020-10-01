package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/upload",uploadFile)

	http.ListenAndServe("127.0.0.1:9001",nil)
}

/**
  处理文件上传
*/
func uploadFile(writer http.ResponseWriter,request *http.Request){
	// request Content-Type isn't multipart/form-data
	fmt.Println("method:",request.Method)//post
	request.ParseMultipartForm(32 << 10)
	//解析文件
	file, header, err :=request.FormFile("uploadFile")//uploadfile和前端页面input标签中的name要一致
	if err != nil {
		log.Fatal(err)
		return
	}
	//header的类型是：FileHeader
	fileName := header.Filename
	fmt.Println("文件名: ",fileName)
	fileSize := header.Size
	fmt.Println("文件大小: ",fileSize)
	//读、写：IO流  input 、output
	defer file.Close()

	//把上传的文件保存下来，保存成一个文件。
	//1、先在硬盘创建一个空文件
	newFile,err :=os.OpenFile("./"+header.Filename,os.O_CREATE|os.O_RDONLY|os.O_WRONLY,0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer newFile.Close()

	//草稿2、将上传的file文件内容拷贝到新文件当中
	io.Copy(newFile,file)
}

