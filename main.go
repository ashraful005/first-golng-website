package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main()  {
	
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets/"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request)  {
	//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil{
		fmt.Println(err.Error())
	}
	//fmt.Fprintf(w, "Welcome")

	ptmp.Execute(w, nil)
}

func features(w http.ResponseWriter, r *http.Request)  {
	//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil{
		fmt.Println(err.Error())
	}
	//fmt.Fprintf(w, "Welcome")

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil{
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request)  {
	//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil{
		fmt.Println(err.Error())
	}
	//fmt.Fprintf(w, "Welcome")

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil{
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}