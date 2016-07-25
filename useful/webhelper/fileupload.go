package webhelper

import (
	"io"
	"net/http"
	"os"
)

type HttpFile struct {
	FilePath     string
	FormFileName string
}

//文件上传功能
func (h *HttpFile) FileUpload(r *http.Request) (s string) {
	s = ""

	//获取文件内容 要这样获取
	file, head, err := r.FormFile(h.FormFileName)
	if err != nil {
		s = "获取文件内容失败，失败信息:" + err.Error()
		return
	}
	defer file.Close()

	//创建文件
	err = os.MkdirAll(h.FilePath, os.ModePerm)
	if err != nil {
		s = "创建文件夹失败，失败信息:" + err.Error()
		return
	}

	fW, err := os.Create(h.FilePath + head.Filename)
	if err != nil {
		s = "创建文件失败，失败信息:" + err.Error()
		return
	}
	defer fW.Close()

	_, err = io.Copy(fW, file)
	if err != nil {
		s = "文件保存失败，失败信息:" + err.Error()
		return
	}

	return
}
