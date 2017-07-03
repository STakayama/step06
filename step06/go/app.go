package app

import (
	"fmt"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	s_a_0:=strings.Split(r.Form["a"][0], "")
//	s_a_1:=strings.Split(r.Form["a"][1], "")	
	s_b_0:=strings.Split(r.Form["b"][0], "")
//	s_b_1:=strings.Split(r.Form["b"][1], "")
	
	var ans1 string	
	for i:=0;i<len(s_a_0);i++{
	    ans1+=ans1+s_a_0[i]+s_b_0[i]
}
fmt.Fprintf(w,"Hello,world")
	//fmt.Fprintf(w,ans1)
}
