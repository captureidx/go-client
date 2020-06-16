package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	//baseURL = "http://api.idx.io:5000/"
	baseURL = "http://localhost:5000"
	// for local
	//authToken = "NDMyNWEzOGQyM2Y5Mjg5MGNmZTEyYjVmYTdlY2M0ZDY3MjljOTAyODpjY2U0MGE1ZTVhOTI5NWZkNmJkZjA2ZGRmMzRlZTNiOGMyMmQ1ZTQw"
	// for actual api
	//authToken      = "MzY0ZDE5MTc0MzZlNTM0ZTE5OWEyMDM5M2FmYmFkNmViMDRhZDEwODphNDczZDBjNjI5MTg0MmMwYjVjMjFiNzc1OTIxZWM3YTFlMGQ4Nzkz"
	defaultTimeout = 30 * time.Second
	version        = "1"
)

type Client struct {
	authHeader     string
	accept         string
	acceptEncoding string
}

// user will pass in key and secret when instantiating the Client
func newClient(key string, secret string) *Client {
	unencodedToken := key + ":" + secret
	auth := base64.StdEncoding.EncodeToString([]byte(unencodedToken))
	return &Client{
		authHeader:     "Basic " + auth,
		accept:         "*/*",
		acceptEncoding: "gzip, deflate",
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.Header.Add("Authorization", c.authHeader)
	req.Header.Add("Accept", c.accept)
	req.Header.Add("Accept-Encoding", c.acceptEncoding)

	//adds 30 second timeout
	client := &http.Client{
		Timeout: defaultTimeout,
	}

	//client.Do sends req, req made in function that calls doRequest
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// defer - executed last
	defer resp.Body.Close()

	//ReadAll could be risky, look into other functions
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// Look into simplifying all the Get functions into one as they are very similar
// When passing in query strings for Get functions format as such Ex: ?query[key=value]&limit=value
func (c *Client) GetListings(query string) (*Listings, error) {
	//form URL for request
	path := baseURL + "/listings/" + query

	// create request and check for error
	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		fmt.Print("Error in request")
		return nil, err
	}
	// pass request to doRequest function which sends the request and error check
	bytesReturned, err := c.doRequest(req)

	if err != nil {
		fmt.Println("Error sending request: ")
		//fmt.Println(baseURL)
		log.Fatal(err)
		return nil, err
	}
	// create struct to save data in, struct in separate go file
	var data Listings

	err = json.Unmarshal(bytesReturned, &data)

	if err != nil {
		fmt.Println("Error in json Unmarshal")
		log.Fatal(err)
		return nil, err
	}
	// return data, nil
	return &data, nil
}

func (c *Client) GetBrokers(query string) ([]Brokers, error) {
	path := baseURL + "/brokers/" + query

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Print("Error in request")
		return nil, err
	}

	bytesReturned, err := c.doRequest(req)
	if err != nil {
		fmt.Println("Error sending request: ")
		log.Fatal(err)
		return nil, err
	}

	var data []Brokers

	err = json.Unmarshal(bytesReturned, &data)
	if err != nil {
		fmt.Println("Error in json Unmarshal")
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}

func (c *Client) GetAgents(query string) ([]Agents, error) {
	path := baseURL + "/agents/" + query

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Print("Error in request")
		return nil, err
	}

	bytesReturned, err := c.doRequest(req)
	if err != nil {
		fmt.Println("Error sending request: ")
		log.Fatal(err)
		return nil, err
	}

	var data []Agents

	err = json.Unmarshal(bytesReturned, &data)
	if err != nil {
		fmt.Println("Error in json Unmarshal")
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}
