package handler

import (
	"github.com/Winszheng/cloudgo/logger"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/assets/"

type userError string

func (e userError) Error() string {
	return string(e)
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	// 判断前缀是否合法
	logger.SugarLogger.Debug("path:", request.URL.Path)
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}
	// 取路径
	path := request.URL.Path[len(prefix):]
	file, err := os.Open("assets/" + path)
	if err != nil {
		return err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	writer.Write(content)
	return nil // 没有出错，返回nil
}

func HandleTable(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		logger.SugarLogger.Debug()
		formatter.HTML(writer, http.StatusOK, "table", struct {
			Name  string
			NetId string
		}{Name: request.Form["name"][0], NetId: request.Form["netid"][0]})
	}
}

func UnknownHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusNotImplemented, "501 Not Implemented")
	}
}
