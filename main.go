package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var defaultName = "Joe"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", htmlHandler)
	r.HandleFunc("/monster/{name}", identiconHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func identiconHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	resp, err := http.Get(fmt.Sprintf("http://dnmonster:8080/monster/%v?size=80", name))
	if err != nil {
		log.Printf("getting image error %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "image/png")
	w.Write(body)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	name := defaultName
	if r.Method == "POST" {
		r.ParseForm()
		name = r.FormValue("name")
	}

	html := fmt.Sprintf(`<html>
				<head><title>Identidock</title></head>
				<body>
    				<form method="POST">
              			Hello <input type="text" name="name" value="">
              			<input type="submit" value="submit">
              		</form>
              		<p>You look like a:
              		<img src="/monster/%v"/>
    			</body>
    		</html>`, name)

	fmt.Fprintf(w, html)
}
