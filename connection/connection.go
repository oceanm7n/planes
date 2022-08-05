package connection

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	structs "github.com/oceanm7n/planes/common"
	new_etihad "github.com/oceanm7n/planes/connections/new_etihad"
	"moul.io/http2curl"
)

type Connection struct {
	AirlineName string
	AirlineCode string
	Url         string
	Method      string
	Headers     map[string]string
	Params      map[string]string
	Body        string

	Query structs.Query
}

func NewConnection(AirlineName string, AirlineCode string, query structs.Query) *Connection {
	c := new(Connection)
	c.AirlineName = AirlineName
	c.AirlineCode = AirlineCode
	c.Query = query

	url, req_type := Url()[c.AirlineCode]()
	c.Url = url
	c.Method = req_type

	return c
}

func Url() map[string]func() (string, string) {
	return map[string]func() (string, string){
		"ETH": new_etihad.Url,
	}
}

func RequestConnections() map[string]func(structs.Query) (string, map[string]string, map[string]string) {
	return map[string]func(structs.Query) (string, map[string]string, map[string]string){
		"ETH": new_etihad.Constructor,
	}
}

func ResponseConnections() map[string]func(*http.Response) string {
	return map[string]func(*http.Response) string{
		"ETH": new_etihad.ResponseParser,
	}
}

func (c Connection) SendRequest() *http.Response {
	req := c.createRequest()

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	return res
}

func (c Connection) ParseRequest(res *http.Response) string {
	return construct_response(c.AirlineCode, res)
}

func (c Connection) createRequest() *http.Request {

	body, headers, params := construct_request(c.AirlineCode, c.Query)
	c.Body = body
	c.Headers = headers
	c.Params = params

	payload := strings.NewReader(c.Body)
	req, err := http.NewRequest(c.Method, c.Url, payload)

	if err != nil {
		fmt.Println(err)
		os.Exit(12)
	}

	for k, v := range c.Headers {
		req.Header.Add(k, v)
	}

	command, _ := http2curl.GetCurlCommand(req)
	fmt.Println(command)
	return req
}

func construct_request(AirlineCode string, query structs.Query) (string, map[string]string, map[string]string) {
	c := RequestConnections()[AirlineCode]
	return c(query)
}

func construct_response(AirlineCode string, res *http.Response) string {
	c := ResponseConnections()[AirlineCode]
	return c(res)
}
