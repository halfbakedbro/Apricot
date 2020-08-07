package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"
)

type Hashval struct {
	HashVal string
}

func mdHash(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hash/md5/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"template/hash.md5.tmpl",
		"template/base.layout.tmpl",
		"template/sidenave.page.tmpl",
		"template/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if r.Method == "GET" {
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", 500)
		}
	} else {

		r.ParseForm()

		input := r.FormValue("text")
		bytes := []byte(input)
		val := md5.Sum(bytes)
		data := Hashval{
			HashVal: hex.EncodeToString(val[:]),
		}

		err = ts.Execute(w, data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", 500)
		}
	}
}

func shas1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hash/sha1/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"template/hash.sha1.tmpl",
		"template/base.layout.tmpl",
		"template/sidenave.page.tmpl",
		"template/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if r.Method == "GET" {
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", 500)
		}
	} else {

		r.ParseForm()

		input := r.FormValue("text")
		bytes := []byte(input)
		val := sha1.Sum(bytes)
		data := Hashval{
			HashVal: hex.EncodeToString(val[:]),
		}

		err = ts.Execute(w, data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", 500)
		}
	}
}

func shas256(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hash/sha256/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"template/hash.sha256.tmpl",
		"template/base.layout.tmpl",
		"template/sidenave.page.tmpl",
		"template/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	if r.Method == "GET" {
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", 500)
		}
	} else {

		r.ParseForm()

		input := r.FormValue("text")
		bytes := []byte(input)
		val := sha256.Sum256(bytes)
		data := Hashval{
			HashVal: hex.EncodeToString(val[:]),
		}

		err = ts.Execute(w, data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", 500)
		}
	}
}
