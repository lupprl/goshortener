package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var UrlInputForm = `
<html>
    <head>
    <title></title>
    </head>
    <body>
        <form method="post">
            <label>Enter URL to shorten: </label><input type="text" name="url">
            <input type="submit" value="OK">
        </form>
    </body>
</html>
`
var UrlStore []string

func Shortener(w http.ResponseWriter, r *http.Request) {
	UrlStore = make([]string, 0)
	switch r.Method {
	case http.MethodGet:
		if len(r.URL.Query()) == 0 {
			fmt.Fprint(w, UrlInputForm)
			return
		}
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Url is missing", http.StatusBadRequest)
			return
		}
		id64, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			http.Error(w, "URL be integer", http.StatusBadRequest)
			return
		}
		if id64 >= uint64(len(UrlStore)) {
			http.Error(w, "URL not", http.StatusBadRequest)
			return
		}
		w.Header().Set("Location", UrlStore[id64])
		w.WriteHeader(http.StatusTemporaryRedirect)
	case http.MethodPost:
		url, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		UrlStore = append(UrlStore, string(url))
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "http://localhost:8080/?id=", len(UrlStore)-1)
	}
}
