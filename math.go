package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Mathval struct {
	Mathval int
}

func _gcd(xi, yi int) Mathval {
	val := 0
	if xi == 0 && yi == 0 {
		return Mathval{Mathval: 0}
	}

	if xi == 0 {
		return Mathval{Mathval: yi}
	} else if yi == 0 {
		return Mathval{Mathval: xi}
	}

	for yi > 0 {
		val = xi % yi

		xi = yi
		yi = val
	}

	return Mathval{Mathval: xi}
}

func gcd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/math/gcd/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"template/math.gcd.tmpl",
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

		x := r.Form["numberX"]
		y := r.Form["numberY"]

		xi, _ := strconv.Atoi(x[0])
		yi, _ := strconv.Atoi(y[0])

		data := Mathval{Mathval: 0}

		if xi > yi {
			data = _gcd(xi, yi)
		} else {
			data = _gcd(yi, xi)
		}

		err = ts.Execute(w, data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", 500)
		}
	}

}
