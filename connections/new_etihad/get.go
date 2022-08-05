package new_etihad

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	structs "github.com/oceanm7n/planes/common"
)

func get_body(query structs.Query) string {
	f, err := os.Open("./connections/new_etihad/body.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	var struct_body requestBody

	json.Unmarshal(byteValue, &struct_body)

	// преобразования
	struct_body.Variables.AirSearchInput.ItineraryParts[0].From.Code = query.From
	struct_body.Variables.AirSearchInput.ItineraryParts[0].To.Code = query.To
	struct_body.Variables.AirSearchInput.ItineraryParts[0].When.Date = query.Date

	b, err := json.Marshal(struct_body)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func get_headers() map[string]string {
	return map[string]string{
		"authority":          "dc.etihad.com",
		"accept-language":    "en-US,en;q=0.9,ru-RU;q=0.8,ru;q=0.7,he;q=0.6,de;q=0.5",
		"authorization":      "Bearer T1RLAQLeJ6kT9SsmxWx6wxoJKhU405XWZGEYh6CuCqT9WCBraBCooZeDJyVa9vRqvHyB7yFwAADQHvPqBiB/TGQQ8bOCq35WpGkRdJXArg4IoibY6yI2Qy++wdQo2Pxuf4wvoK73zdz1A3oZbw5FRff9T5dR3qvuorJqkPYppPu8rZAznmk1kf7c+SyhlkMlLg7o5hIK9ms0ryaejxqa05dFWKVxfALdwstnrDlpblSs53JlS50dTzXsLkYKRfMPhOIepI7qTHGlk6BcVI4MkshraCCVasHzfwyD+P0OyOyavVth6bzQZJ2OOOuFB8NHawTA36DkVQvKD6kimjTVPKwiPdFd8afVWQ**",
		"content-type":       "application/json",
		"conversation-id":    "cl4eb78dwjow0szsqk0k1h6q5",
		"origin":             "https://dxbooking.etihad.com",
		"referer":            "https://dxbooking.etihad.com/",
		"x-sabre-storefront": "EYDX",
		"Cookie":             "incap_ses_1456_2042947=XqK9C0gPazo9kQaoqMA0FJtJ7GIAAAAAesVMpRpL2FEBXlfvySJsdg==; incap_ses_1456_2042951=eaRrHwov41L0dwGoqMA0FIo97GIAAAAAdVLffqKgLFpW9qavnTR6uQ==; nlbi_2042947_2612571=ZhnWSR9ityI8MHpl6k45QQAAAACKat5JivoI8/HRYBp0Nueu; nlbi_2042951=JOcAeQwtlxoOo+R0MzxorgAAAAAE9a4DX3/RH7UHt43B/vF6; visid_incap_2042947=Gv7B+WKuQx69FdA1t3VsDbM/7GIAAAAAQUIPAAAAAACIgpifBRepBgp2gpuBsLeL; visid_incap_2042951=ao9GGsLpQ2y6ewBJSeuuDYg97GIAAAAAQUIPAAAAAABG9HWfzp6OeCHc+AxjdmE5",
	}
}

func get_params() map[string]string {
	return map[string]string{}
}
