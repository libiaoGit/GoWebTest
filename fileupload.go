package main

import (
	"fmt"
	"gitlabdemo/gowebtest/useful/webhelper"
	"io"
	"net/http"
)

const (
	upload_path string = "F:/workspace/Go/src/gitlabdemo/gowebtest/image/"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}

//上传
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<html><head><title>我的第一个页面</title></head><body><form action='' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file'  /><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		file := &webhelper.HttpFile{
			FilePath:     upload_path,
			FormFileName: "file",
		}

		str := file.FileUpload(r)
		if str == "" {
			io.WriteString(w, "图片上传成功！")
		} else {
			io.WriteString(w, str)
		}
	}
	/*//从请求当中判断方法
	if r.Method == "GET" {
		io.WriteString(w, "<html><head><title>我的第一个页面</title></head><body><form action='' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file'  /><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		//获取文件内容 要这样获取
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println(head.Filename)
		//创建文件
		fW, err := os.Create(upload_path + head.Filename)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		//io.WriteString(w, head.Filename+" 保存成功")
		http.Redirect(w, r, "/hello", http.StatusFound)
		//io.WriteString(w, head.Filename)
	}*/
}

func main() {
	//启动一个http 服务器
	http.HandleFunc("/hello", helloHandle)
	//上传
	http.HandleFunc("/image", uploadHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
		return
	}
	fmt.Println("服务器启动成功")
}
