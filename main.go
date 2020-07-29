package main

import (
	"encoding/json"
	"fmt"
	"github.com/domainr/whois"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func whoisQuery(data string) (string, error) {

	response, err := whois.Fetch(data)

	if err != nil {
		return "", err
	}

	return string(response.Body), nil
}

func jsonResponse(w http.ResponseWriter, x interface{}) {

	bytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "applicaton/json")
	w.Write(bytes)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"template/home.page.tmpl",
		"template/base.layout.tmpl",
		"template/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server Error", 500)
	}
}

func whoist(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		files := []string{
			"template/util.whois.tmpl",
			"template/base.layout.tmpl",
			"template/footer.partial.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", 500)
		}
		return
	}

	data := r.PostFormValue("data")

	result, err := whoisQuery(data)

	if err != nil {
		jsonResponse(w, Response{Error: err.Error()})
		return
	}

	jsonResponse(w, Response{Result: result})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	mux := http.NewServeMux()
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", home)

	mux.HandleFunc("/hash/md5/", mdHash)
	mux.HandleFunc("/hash/sha1/", shas1)
	mux.HandleFunc("/hash/sha256/", shas256)
	mux.HandleFunc("/whois/", whoist)

	fmt.Println("Starting server on :8081")
	s := strconv.Itoa(port)
	str := ":" + s
	err := http.ListenAndServe(str, mux)
	log.Fatal(err)
}
