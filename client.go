package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//baseURL = "http://api.idx.io:5000/"
	baseURL = "http://localhost:5000"
	// for local
	authToken = "NDMyNWEzOGQyM2Y5Mjg5MGNmZTEyYjVmYTdlY2M0ZDY3MjljOTAyODpjY2U0MGE1ZTVhOTI5NWZkNmJkZjA2ZGRmMzRlZTNiOGMyMmQ1ZTQw"
	// for actual api
	//authToken      = "MzY0ZDE5MTc0MzZlNTM0ZTE5OWEyMDM5M2FmYmFkNmViMDRhZDEwODphNDczZDBjNjI5MTg0MmMwYjVjMjFiNzc1OTIxZWM3YTFlMGQ4Nzkz"
	defaultTimeout = 30
	version        = "1"
)

type Client struct {
	authHeader     string
	accept         string
	acceptEncoding string
	timeout        int
}

func startUp() *Client {

	return &Client{
		authHeader: "Basic " + authToken,
		//authHeader: createToken(),
		//accept:  "application/vnd.southernweb.idx.v" + version + "+json", // this is from php file
		accept:         "*/*", // this is what works on postman
		acceptEncoding: "gzip, deflate",

		timeout: defaultTimeout,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	client := &http.Client{}

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
	//checks for status code 200
	//fmt.Println(resp.Request)    // print out req for debug
	//fmt.Println(resp.StatusCode) // print out status for debug
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (c *Client) GetListings(query string) (*Listings, error) {
	//form URL for request
	path := baseURL + "/listings/" + query

	// create request and check for error
	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		fmt.Print("Error in request")
		return nil, err
	}

	req.Header.Add("Authorization", c.authHeader)
	req.Header.Add("Accept", c.accept)
	req.Header.Add("Accept-Encoding", c.acceptEncoding)

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

	//fmt.Println("Path: " + path) // making sure have correct path

	err = json.Unmarshal(bytesReturned, &data)

	//this error is executing, invalid character error - maybe not in JSON
	if err != nil {
		fmt.Println("Error in json Unmarshal")
		log.Fatal(err)
		return nil, err
	}

	// return data, nil
	return &data, nil
}

func (c *Client) AddListing(listing *Listings, query string) error {
	url := baseURL + "/" + query
	j, err := json.Marshal(listing)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	return err
}

//func (c *Client) GetBrokers(query string) (*[]Brokers, error){

//}

func createToken(key string, secret string) string {
	unencodedToken := key + ":" + secret
	return base64.StdEncoding.EncodeToString([]byte(unencodedToken))
}
