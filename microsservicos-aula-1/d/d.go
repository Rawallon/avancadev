package main

import (
    "math/rand"
	"time"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
 
type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}


func flipCoin() int {
	rand.Seed(time.Now().UnixNano())
    min := 1
    max := 2
	return rand.Intn(max - min + 1) + min
}

func home(w http.ResponseWriter, r *http.Request) {
	chance := flipCoin()

	result := Result{Status: "unlucky"}
	
	if chance == 1 {
		result.Status = "lucky"
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}