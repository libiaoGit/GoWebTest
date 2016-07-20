package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GoodsComment struct {
	OrderCode     string
	GoodsCode     int64
	GoodsName     string
	GoodsPicUrl   string
	GoodsPrice    string
	ModelCode     int64
	ModelName     string
	ShopOwnerCode string
	CustomerCode  string
	BuyerPhotoUrl string
	IsAnonymous   bool
	Content       string
	CreateTime    string
	CreateTimeStr time.Time
	PicUrls       []string
	Score         int
	CommentLevel  int
	Reply         string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//解析参数，默认是不会解析的
	fmt.Println("Form表单数据：", r.Form)
	fmt.Println("Form表单中Data数据", r.Form["data"])

	var data GoodsComment
	json.Unmarshal([]byte(r.Form["data"][0]), &data)
	data.CreateTimeStr = time.ParseInLocation("2006-01-02 15:04:05", data.CreateTime, time.Local)

	fmt.Println("订单号", data.OrderCode)
	fmt.Println("商品编号", data.GoodsCode)
	fmt.Println("商品名称", data.GoodsName)
	fmt.Println("图片地址", data.GoodsPicUrl)
	fmt.Println("商品价格", data.GoodsPrice)
	fmt.Println("型号编号", data.ModelCode)
	fmt.Println("型号名称", data.ModelName)
	fmt.Println("卖家编号", data.ShopOwnerCode)
	fmt.Println("买家编号", data.CustomerCode)
	fmt.Println("买家头像地址", data.BuyerPhotoUrl)
	fmt.Println("是否匿名", data.IsAnonymous)
	fmt.Println("评价内容", data.Content)
	fmt.Println("评价图片", data.PicUrls)
	fmt.Println("评价时间(string)", data.CreateTime)
	fmt.Println("评价时间(time.Time)", data.CreateTimeStr)
	fmt.Println("评分", data.Score)
	fmt.Println("评价等级", data.CommentLevel)
	fmt.Println("回复", data.Reply)
}

func main() {
	http.HandleFunc("/test.html", sayhelloName)
	//设置访问的路由
	err := http.ListenAndServe(":9090", nil)
	//设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
