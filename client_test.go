package client

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	// token := createToken("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")
	//fmt.Println(token)
}

func TestGetListings(t *testing.T) {
	//var data Listings
	c := newClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")
	data, _ := c.GetListings("?limit=2")

	for i := 0; i < len(data.Results); i++ {
		fmt.Print(data.Results[i].Type + " ")
		fmt.Print(data.Results[i].Price.Amount)
		fmt.Print(" ")
		fmt.Println(data.Results[i].MlNum)
	}

	//fmt.Println(data.Results[0].Price.Amount)
	//fmt.Println(data.Results[0].Description)
	//fmt.Println(data.Results[0].Location.Coordinates.GeoJson.Type)

}

func TestClient_GetBrokers(t *testing.T) {
	c := newClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")

	var data []Brokers

	data, _ = c.GetBrokers("")

	fmt.Println(data[0].Name)
}

func TestClient_GetAgents(t *testing.T) {
	c := newClient("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40")

	var data []Agents

	data, _ = c.GetAgents("")

	fmt.Println(data[0].Phone.Mobile)

	//for i:= 0; i < len(data.Results); i++{
	//	fmt.Println(data.Results[i].Phone.Mobile)
	//}
}
