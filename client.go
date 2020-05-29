

package client

import "fmt"

type Client struct {
	// var apiURL string = "https://api:5000/"
}

func (c *Client) Hello() {
	fmt.Println("Hello")
}
