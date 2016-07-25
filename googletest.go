package main

import (
	"container/list"
	"fmt"
	"net/http"
	"strconv"
)

var i int
var alli *list.List = list.New()

func myfunc(w http.ResponseWriter, r *http.Request) {
	i++
	alli.PushBack(strconv.Itoa(i))
	for e := alli.Front(); e != nil; e = e.Next() {
		s := "第" + e.Value.(string) + "次访问网页。\n"
		fmt.Fprintf(w, s)
	}

}

func main() {
	http.HandleFunc("/", myfunc)
	http.ListenAndServe(":8080", nil)
}
