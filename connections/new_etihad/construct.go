package new_etihad

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	structs "github.com/oceanm7n/planes/common"
)

func Url() (string, string) {
	return "https://dxbooking.etihad.com/api/graphql", "POST"
}

func Constructor(query structs.Query) (string, map[string]string, map[string]string) {
	// вернуть:
	// - string - payload
	// - map[string]string - headers
	// - map[string]string - params

	body := get_body(query)
	headers := get_headers()
	params := get_params()

	return body, headers, params
}

func ResponseParser(res *http.Response) string {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var response responseBody
	json.Unmarshal(body, &response)

	seats_left, price := response.Data.BookingAirSearch.OriginalResponse.UnbundledOffers[0][0].SeatsRemaining.Count, response.Data.BookingAirSearch.OriginalResponse.UnbundledOffers[0][0].Total.Alternatives[0][0].Amount

	return "Seats left: " + fmt.Sprint(seats_left) + ", Price: " + fmt.Sprint(price)
}
