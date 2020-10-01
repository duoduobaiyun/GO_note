package reqhandle

import "net/http"

/**
 * 处理默认的首页
 */
func Index(writer http.ResponseWriter,request *http.Request){
	writer.Write([]byte("欢迎进入豆瓣信息采集系统"))
}
