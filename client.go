package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// timeout for sending request
	defaultTimeout = 30 * time.Second
	//version        = "1"
)

//var baseURL = "http://localhost:5000"

// Stores the Headers for the request
type Client struct {
	authHeader     string
	accept         string
	acceptEncoding string
	Timeout        time.Duration
}

// User will pass in key and secret when instantiating the Client
func NewClient(key string, secret string) *Client {
	unencodedToken := key + ":" + secret
	auth := base64.StdEncoding.EncodeToString([]byte(unencodedToken))
	return &Client{
		authHeader:     "Basic " + auth,
		accept:         "*/*",
		acceptEncoding: "gzip, deflate",
		Timeout:        0, // explicit default
	}
}

// Takes a request, adds headers to it, sends request, then checks for 200 status
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	// adds headers to request
	req.Header.Add("Authorization", c.authHeader)
	req.Header.Add("Accept", c.accept)
	req.Header.Add("Accept-Encoding", c.acceptEncoding)

	clientTimeout := defaultTimeout
	if c.Timeout != 0 {
		clientTimeout = c.Timeout
	}

	//adds 30 second timeout
	client := &http.Client{
		Timeout: clientTimeout,
	}

	//client.Do sends req, req made in function that calls doRequest
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	//ReadAll could be risky depending on file size
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		fmt.Println("receieved status code: ", resp.StatusCode)
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

// When passing in query strings for Get functions format as such Ex: ?query[key=value]&limit=value
// Creates a request to Get Listings from API, sends to DoRequest(), receives the json from DoRequest and returns it
func (c *Client) GetListings(url string, query string) (*Listings, error) {

	path := url + "/listings/" + query

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		fmt.Print("Error in request")
		return nil, err
	}

	bytesReturned, err := c.doRequest(req)

	if err != nil {
		// fmt.Println("Error sending request: ")
		// //fmt.Println(baseURL)
		// log.Fatalerr)
		return nil, err
	}

	var data Listings

	err = json.Unmarshal(bytesReturned, &data)

	if err != nil {
		fmt.Println("Error in json Unmarshal")
		return nil, err
	}

	return &data, nil
}

// Creates a request to Get Brokers from API, sends to DoRequest(), receives the json from DoRequest and returns it
// API currently does not support queries for Brokers I believe
func (c *Client) GetBrokers(url string, query string) ([]Brokers, error) {
	path := url + "/brokers/" + query

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Print("Error in request")
		return nil, err
	}

	bytesReturned, err := c.doRequest(req)
	if err != nil {
		fmt.Println("Error sending request: ")
		return nil, err
	}

	var data []Brokers

	err = json.Unmarshal(bytesReturned, &data)
	if err != nil {
		fmt.Println("Error in json Unmarshal")
		return nil, err
	}

	return data, nil
}

// Creates a request to Get Agents from API, sends to DoRequest(), receives the json from DoRequest and returns it
// API currently does not support queries for Agents I believe
func (c *Client) GetAgents(url string, query string) ([]Agents, error) {

	path := url + "/agents/" + query

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Print("Error in request")
		return nil, err
	}

	bytesReturned, err := c.doRequest(req)
	if err != nil {
		fmt.Println("Error sending request: ")
		return nil, err
	}

	var data []Agents

	err = json.Unmarshal(bytesReturned, &data)
	if err != nil {
		fmt.Println("Error in json Unmarshal")
		return nil, err
	}

	return data, nil
}
