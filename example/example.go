package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/horechek/сityads"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	client := cityads.NewClient(
		"https://cityads.com/api/rest/webmaster/json",
		"xxx",
	)

	type Response struct {
		Status int    `json:"status"`
		Error  string `json:"error"`
		Data   struct {
			Total int `json:"total"`
			Items map[string]struct {
				Id    int    `json:"id"`
				Name  string `json:"name"`
				Items []struct {
					Title     string `json:"title"`
					IsDefault bool   `json:"is_default"`
					DeepLink  string `json:"deep_link"`
				} `json:"items"`
			} `json:"items"`
		} `json:"data"`
	}

	r := Response{}
	params := url.Values{}
	params.Add("user_has_offer", "true")
	err := client.Call("offers/web", "GET", params, &r)
	if err != nil {
		log.Fatal(err)
	}

	for _, offer := range r.Data.Items {
		fmt.Println("==============")
		fmt.Println("name:", offer.Name)
		for _, item := range offer.Items {
			fmt.Println("---")
			fmt.Println("title:", item.Title)
			fmt.Println("is_default:", item.IsDefault)
			fmt.Println("deep_link:", item.DeepLink)
		}
	}
}
