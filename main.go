package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	etihad "github.com/oceanm7n/planes/connections/etihad"
	"github.com/ssgreg/logf"
)

func get_bearer() string {
	// todo: automate this

	return "Bearer T1RLAQJDmFRVL84iCqkJBI5YF6O1XPDcd0NxbJ/skxDwKKaU9hCH33wU1nAK6JbX408C0LbpAADQJJ8M058T7j2Jbesr2AKyTTwzfTfKhHdUTozZtyC+kfdUsRd/ulPMW2zpjUgFSceXWSgOynZ+vX9z1+9RBH7F8BFM9u1uAqTVzd4Bo8QQ16Y9c9wY3tEeBT1wLKy+uCVhushi1aEMpsFpYMYX9AgOHKh416IxOkN2EYDIVWfsDIah1/L2YkukslkzFpaOO1HnADUi1yQVv1h3p/e+ow92rIEK2FqiAfT5TXJjdpuy1BFVWayYQYdBaPT3OmXksb3pQZ1LNmodfhqq9bnLlwZqkA**"
}

func set_headers(req *http.Request, logger *logf.Logger) {
	bearer := get_bearer()

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

func logger() (*logf.Logger, logf.ChannelWriterCloseFunc) {
	writer, writerClose := logf.NewChannelWriter.Default()

	return logf.NewLogger(logf.LevelInfo, writer), writerClose
}

func read_response(resp *http.Response, logger *logf.Logger) {
	logger.Info("response Status:" + resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	logger.Debug("response Body:")
	logger.Debug(strings.ReplaceAll(string(body), `"`, `'`))

	var response etihad.Response
	json.Unmarshal(body, &response)

	seats_left, price := response.UnbundledOffers[0][0].SeatsRemaining.Count, response.UnbundledOffers[0][0].Total.Alternatives[0][0].Amount
	logger.Info("Price fetched", logf.Int("seats_left", seats_left), logf.Int("price", price))
}

func main() {
	logger, writerClose := logger()
	defer writerClose()

	logger.Info("Application started")

	// Set URL
	EXECGUID := "95c1c492-b008-4981-9a12-03c9c1c2ff5d"
	JIPCC := "EYDX"
	url := "https://dc.etihad.com/v4.3/dc/products/air/search?execution=" + EXECGUID + "&jipcc=" + JIPCC

	// Create request
	req_body := etihad.Create_body("SVO", "AUH", "2022-09-01")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(req_body))
	if err != nil {
		log.Fatal(err)
	}

	set_headers(req, logger)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	read_response(resp, logger)
}
