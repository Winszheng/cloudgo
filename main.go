package main

import (
	"github.com/Winszheng/cloudgo/handler"
	logger "github.com/Winszheng/cloudgo/logger"
	"github.com/Winszheng/cloudgo/service"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(
	handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		// http库自带的panic处理内容太多，没有必要
		// 自行修改panic的处理
		defer func() {
			if r := recover(); r != nil {
				// 日志打印到控制台
				logger.SugarLogger.Panicf("panic:%v\n", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// 调用处理函数
		err := handler(writer, request)
		if err != nil {
			// 日志信息输出到控制台
			logger.SugarLogger.Errorf("error handling request:%s\n", err.Error())

			// 判断是系统错误，还是用户可见
			if userErr, ok := err.(userError); ok {
				// 用户可见
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK // 默认200， 正常
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound // 404 not found
			case os.IsPermission(err):
				code = http.StatusForbidden // 403 没权限
			default:
				code = http.StatusInternalServerError // else
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error            // 系统看的
	Message() string // for user
}

func main() {
	logger.InitLogger()
	defer logger.SugarLogger.Sync()

	// 法2：使用mux库
	//
	// 创建路由器
	//router := mux.NewRouter()
	//// 因为mux库严格按照正则匹配，当需要处理前缀变长的路由时需要额外设置
	//router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	//// 限制路由器只处理指定的请
	//router.HandleFunc("/api/test", service.ApiTestHandler(formatter))
	//router.HandleFunc("/api/table", handler.HandleTable(formatter))
	//router.HandleFunc("/unknow", handler.UnknownHandler(formatter))
	//err := http.ListenAndServe(":8080", router)
	//if err != nil {
	//	logger.SugarLogger.Error(err)
	//} else {
	//	logger.SugarLogger.Debug("Normal start-up")
	//}
	r := mux.NewRouter()

	// 法1：用http库直接实现
	http.HandleFunc("/", errWrapper(handler.HandleFileList))
	formatter := render.New(render.Options{
		//设置
		Directory:  "template",
		Extensions: []string{".html"},
		IndentJSON: true,
	})
	http.HandleFunc("/api/test", service.ApiTestHandler(formatter))
	http.HandleFunc("/api/table", handler.HandleTable(formatter))
	http.ListenAndServe(":8080", nil)
}
