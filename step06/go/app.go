package app
//https://www.slideshare.net/takuyaueda967/gaegoweb

import (
	"fmt"
//	"html"
	"net/http"
	"strings"
//	"strconv"
)

func init() {
	http.HandleFunc("/", handlePata)// /というパターンのパスできたリクエストをハンドリングする関数を登録　GAEでmainは使わない
}

func handlePata(w http.ResponseWriter, r *http.Request) { //レスポンスを書き込むwriter　リクエスト
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
form := `

<html>
  <head>
    <title>パタトクカシーー</title>
  </head>
  <body>
    
    <hr>
    <form action="/post" method="post">
      <input type=text name=a><br>
      <input type=text name=b><br>
      <input type=submit>
    </form>
  </body>
</html>
 `
	fmt.Fprint(w,form)	

	s_a_0:=strings.Split(r.FormValue("a"), "")
	s_b_0:=strings.Split(r.FormValue("b"), "")
	
	ans1:=""//var ans1 string	
//	fmt.Fprintf(w,strconv.Itoa(len(s_a_0)))

	longer:=[]string{}
	shorter:=[]string{}

	if len(s_a_0)>len(s_b_0){
	longer=s_a_0
	shorter=s_b_0
}else{
	longer=s_b_0
	shorter=s_a_0
} 

	for i:=0;i<len(longer);i++{
	if i<len(shorter){
	    ans1+=s_a_0[i]+s_b_0[i]
}else{
	ans1+=longer[i]
}
}
	fmt.Fprintf(w,ans1)
}
