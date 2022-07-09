package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// url: http://127.0.0.1:8080/add?a=1&b=2
	// response: {data: 3}
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		fmt.Println(r.URL)

		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Context-Type", "application/json")
		result, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		w.Write(result)
	})
	http.ListenAndServe(":8080", nil)
}
