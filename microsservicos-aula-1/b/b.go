package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9091", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	ccNumber := r.PostFormValue("ccNumber")

	resultCoupon := makeHttpCall("http://localhost:9092", coupon)

	result := Result{Status: "declined"}

	if ccNumber == "1" {
		result.Status = "approved" 
	}

	if resultCoupon.Status != "valid" && result.Status != "approved" {
		result.Status = "invalid coupon and card"
	}
	if resultCoupon.Status == "valid" && result.Status != "approved" {
		result.Status = "Invalid card but valid coupon"
	}
	if resultCoupon.Status == "invalid" && result.Status == "approved" {
		result.Status = "Valid card but invalid coupon"
	}
	if resultCoupon.Status == "valid" && result.Status == "approved" {
		result.Status = "Valid card and coupon"
	}
	if resultCoupon.Status == "cashback" {
		result.Status = "valid! and you got cashback! Wowzers!"
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}


func makeHttpCall(urlMicroservice string, coupon string) Result {

	values := url.Values{}
	values.Add("coupon", coupon)

	res, err := http.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: "Servidor fora do ar!"}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result

}
