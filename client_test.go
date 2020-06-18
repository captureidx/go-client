package client

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"testing"
)

var filePath string // for serving different json files

//serves a json file tp allow our client to call - for testing purposes
func util(port string) {

	var httpDir string
	var httpPort string

	httpDir = "./testing_responses/" + filePath
	httpPort = port
	flag.Parse()

	log.Println("Starting responses server with root dir: ", httpDir)
	http.HandleFunc("/", foo)
	log.Println("Listening on: ", httpPort)
	http.ListenAndServe(":"+httpPort, nil)
}

// what is passed to HandleFunc
func foo(w http.ResponseWriter, r *http.Request) {
	fp := "testing_responses/" + filePath
	http.ServeFile(w, r, fp)
}

func TestNewClient(t *testing.T) {
	c := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")
	if c.authHeader != "Basic NDMyNWEzOGQyM2Y5Mjg5MGNmZTEyYjVmYTdlY2M0ZDY3MjljOTAyODpjY2U0MGE1ZTVhOTI5NWZkNmJkZjA2ZGRmMzRlZTNiOGMyMmQ1ZTQw" {
		t.Errorf("Incorrect token generation")
	}
}

func TestGetListingsBaseCase(t *testing.T) {

	client := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")
	// base case test
	filePath = "firstResponse.json"

	go util("8080")

	data, _ := client.GetListings("http://localhost:8080", "")

	for i := 0; i < len(data.Results); i++ {
		fmt.Print(data.Results[i].Type + " ")
		fmt.Print(data.Results[i].Price.Amount)
		fmt.Print(" ")
		fmt.Println(data.Results[i].MlNum)
	}

}

func TestGetListingsNoLinks(t *testing.T) {
	client := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")

	filePath = "noLinks.json"
	go util("8080")
	data1, _ := client.GetListings("http://localhost:8080", "")

	for i := 0; i < len(data1.Results); i++ {
		fmt.Print(data1.Results[i].Type + " ")
		fmt.Print(data1.Results[i].Price.Amount)
		fmt.Print(" ")
		fmt.Println(data1.Results[i].MlNum)
	}
}

func TestClient_GetBrokers(t *testing.T) {
	//c := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")

	//var data []Brokers

	//data, _ = c.GetBrokers("","")

	//fmt.Println(data[0].Name)
}

func TestClient_GetAgents(t *testing.T) {
	//c := NewClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")

	//var data []Agents

	//data, _ = c.GetAgents("")

	//fmt.Println(data[0].Phone.Mobile)

	//for i:= 0; i < len(data.Results); i++{
	//	fmt.Println(data.Results[i].Phone.Mobile)
	//}
}
