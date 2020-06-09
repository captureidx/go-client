package client

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token := createToken("4325a38d23f92890cfe12b5fa7ecc4d6729c9028", "cce40a5e5a9295fd6bdf06ddf34ee3b8c22d5e40\n")
	fmt.Println(token)
}

func TestGetListings(t *testing.T) {
	//var data Listings
	c := Client{
		authHeader: "Basic " + authToken,
		//authHeader: createToken(),
		//accept:  "application/vnd.southernweb.idx.v" + version + "+json", // this is from php file
		accept:         "*/*", // this is what works on postman
		acceptEncoding: "gzip, deflate",

		timeout: defaultTimeout,
	}
	data, _ := c.GetListings("")

	for i := 0; i < len(data.Results); i++ {
		fmt.Println(data.Results[i].Type)
		//fmt.Println(data.Results[i].Price)
		fmt.Println(data.Results[i].MlNum)
	}

}
