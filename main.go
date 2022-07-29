package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func set_headers(req *http.Request) {
	bearer := "Bearer T1RLAQKLURagp1oTjbdGHeNc66cp2u2hdOyBjnrm+fjGrk+l/BDPPvpTx0rY9XmTDcoIm9LSAADQNBqqwfCK7l1AActydyb7vcka6vuqM61sKKA52qvtT7JMQRkltsKpf9/GiG6rbLzWcbdqTjwNctkbL/8OQbxJWzNhnlIkHAe3ggXrec6JsN5BHos2prdkBS1tzm7BjUgLfHwxShH6UI08kdmX7Q9Q+MnLZiKU71L0c9pkHez4DvqbVoujrdrfmePIYtzc9iqaLDKzmli3M9YMHxjhQWqtwgnFMOypIm4mpOg+aMTi/+e0B5rhJeOPGaKe5iBIER+PeBc8XRBRec+3hkurJkFJ8w**"

	headers := make(map[string]string)

	headers["authority"] = "dc.etihad.com"
	headers["accept"] = "application/json"
	headers["accept-language"] = "en-US,en;q=0.9,ru-RU;q=0.8,ru;q=0.7,he;q=0.6,de;q=0.5"
	headers["authorization"] = bearer
	headers["content-type"] = "application/json"
	headers["conversation-id"] = "cl4eb78dwjow0szsqk0k1h6q5"
	headers["origin"] = "https://dxbooking.etihad.com"
	headers["referer"] = "https://dxbooking.etihad.com/"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"

	for key, val := range headers {
		req.Header.Set(key, val)
	}
}

func main() {
	fmt.Println("Application started")

	// Set URL
	EXECGUID := "95c1c492-b008-4981-9a12-03c9c1c2ff5d"
	JIPCC := "EYDX"
	url := "https://dc.etihad.com/v4.3/dc/products/air/search?execution=" + EXECGUID + "&jipcc=" + JIPCC

	// Create request
	req_body := Generate_body("SVO", "AUH", "2022-06-01")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(req_body))
	if err != nil {
		log.Fatal(err)
	}

	set_headers(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("request Headers", req.Header)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	fmt.Println(url)
}
