package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var filePath string // for serving different json files

// creates httptest server to send json files to
func serverForTest() *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(fileServe))
}

// what is passed to HandleFunc in serveForTest()
func fileServe(w http.ResponseWriter, r *http.Request) {
	fp := "testing_responses/" + filePath
	http.ServeFile(w, r, fp)
}

func TestNewClient(t *testing.T) {
	c := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")
	if c.authHeader != "Basic NDMyNWEzOGQyM2Y5Mjg5MGNmZTEyYjVmYTdlY2M0ZDY3MjljOTAyODpjY2U0MGE1ZTVhOTI5NWZkNmJkZjA2ZGRmMzRlZTNiOGMyMmQ1ZTQw" {
		t.Errorf("Incorrect token generation")
	}
}

// Base Case for Get request to Listings
func TestGetListingsBaseCase(t *testing.T) {
	ts := serverForTest()
	defer ts.Close()
	client := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")

	filePath = "firstResponse.json"

	data, _ := client.GetListings(ts.URL, "")
	fmt.Println("\nStarting Listings test... ")
	for i := 0; i < len(data.Results); i++ {
		fmt.Print(data.Results[i].Type + " ")
		fmt.Print(data.Results[i].Price.Amount)
		fmt.Print(" ")
		fmt.Println(data.Results[i].MlNum)
	}
}

// Tests if no links are in json response
func TestGetListingsNoLinks(t *testing.T) {
	client := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")
	ts := serverForTest()
	defer ts.Close()
	filePath = "noLinks.json"
	data1, _ := client.GetListings(ts.URL, "")

	for i := 0; i < len(data1.Results); i++ {
		fmt.Print(data1.Results[i].Type + " ")
		fmt.Print(data1.Results[i].Price.Amount)
		fmt.Print(" ")
		fmt.Println(data1.Results[i].MlNum)
	}
}

// Base Case for GetBrokers
func TestClient_GetBrokersBaseCase(t *testing.T) {
	client := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")
	ts := serverForTest()
	defer ts.Close()
	filePath = "firstBrokers.json"

	data, _ := client.GetBrokers(ts.URL, "")

	fmt.Println("\nStarting Brokers test... ")
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i].Name + " ")
		fmt.Print(data[i].BrokerName + " ")
		fmt.Println(data[i].BrokerId)
	}
}

//Base Case for GetAgents
func TestClient_GetAgents(t *testing.T) {
	c := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")
	ts := serverForTest()
	defer ts.Close()
	filePath = "firstAgents.json"
	data, _ := c.GetAgents(ts.URL, "")

	fmt.Println("\nStarting Agents test... ")
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i].Phone.Mobile + " ")
		fmt.Println(data[i].AgentId)
	}
}
