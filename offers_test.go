package cityads

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OffersSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *OffersSuite) SetupSuite() {
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, offersTestData)
	}))
}

func (s *OffersSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
}

func (s *OffersSuite) TestOffers() {
	client := NewClient(
		s.server.URL+"/api/rest/webmaster/json/",
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
	s.Assert().Nil(err)

	s.Assert().Len(r.Data.Items, 2)
	s.Assert().Len(r.Data.Items["30766"].Items, 5)
	s.Assert().EqualValues("Standard", r.Data.Items["30766"].Items[0].Title)
	s.Assert().EqualValues(30766, r.Data.Items["30766"].Id)
	s.Assert().EqualValues(" IDEALICA", r.Data.Items["30766"].Name)
}

func TestOffersSuite(t *testing.T) {
	suite.Run(t, new(OffersSuite))
}

const offersTestData = `{
   "status":200,
   "error":"",
   "request_id":0,
   "data":{
      "total":1134,
      "items":{
         "30766":{
            "id":30766,
            "name":" IDEALICA",
            "user_has_offer":"1",
            "cookie_ltv":"30",
            "site":"http://ml5.de.idealica.net",
            "categories":[
               {
                  "main_id":"1906",
                  "main_title":"Health and Beauty",
                  "other_categories":[
                     {
                        "id":"1906",
                        "title":"Health and Beauty"
                     },
                     {
                        "id":"10023",
                        "title":"E-commerce"
                     }
                  ]
               }
            ],
            "traffic-types":{
               "1":{
                  "id":"1",
                  "value":true,
                  "note":"",
                  "name":"Paid Search"
               },
               "2":{
                  "id":"2",
                  "value":true,
                  "note":"",
                  "name":"Display"
               },
               "3":{
                  "id":"3",
                  "value":true,
                  "note":"",
                  "name":"AdSpot/RichMedia/Sliding"
               },
               "4":{
                  "id":"4",
                  "value":true,
                  "note":"",
                  "name":"E-mail"
               },
               "5":{
                  "id":"5",
                  "value":true,
                  "note":"",
                  "name":"Social Networks"
               },
               "6":{
                  "id":"6",
                  "value":false,
                  "note":"",
                  "name":"Price-Comparison"
               },
               "7":{
                  "id":"7",
                  "value":false,
                  "note":"",
                  "name":"Coupon/Promo Codes"
               },
               "8":{
                  "id":"8",
                  "value":false,
                  "note":"",
                  "name":"Cashback"
               },
               "9":{
                  "id":"9",
                  "value":true,
                  "note":"",
                  "name":"Teaser advertisements"
               },
               "10":{
                  "id":"10",
                  "value":true,
                  "note":"",
                  "name":"Clickunder/Popunder"
               },
               "11":{
                  "id":"11",
                  "value":true,
                  "note":"",
                  "name":"Doorways"
               },
               "12":{
                  "id":"12",
                  "value":true,
                  "note":"",
                  "name":"Content Sites"
               },
               "14":{
                  "id":"14",
                  "value":false,
                  "note":"",
                  "name":"Incentive"
               },
               "15":{
                  "id":"15",
                  "value":false,
                  "note":"",
                  "name":"Messenger/SMS"
               },
               "17":{
                  "id":"17",
                  "value":true,
                  "note":"",
                  "name":"Toolbar"
               },
               "20":{
                  "id":"20",
                  "value":true,
                  "note":"",
                  "name":"Youtube"
               },
               "32":{
                  "id":"32",
                  "value":true,
                  "note":"",
                  "name":"Adult"
               },
               "33":{
                  "id":"33",
                  "value":true,
                  "note":"",
                  "name":"Autoredirect"
               },
               "34":{
                  "id":"34",
                  "value":true,
                  "note":"",
                  "name":"Branded context ads"
               },
               "35":{
                  "id":"35",
                  "value":true,
                  "note":"",
                  "name":"Retargeting"
               },
               "39":{
                  "id":"39",
                  "value":true,
                  "note":"",
                  "name":"Pop-up"
               },
               "40":{
                  "id":"40",
                  "value":true,
                  "note":"",
                  "name":"Mobile traffic in mobile version"
               }
            },
            "additional-traffic-types":null,
            "geo":[
               {
                  "id":"25",
                  "name":"Bulgaria",
                  "country_id":"25",
                  "area_id":"0",
                  "code":"BG",
                  "iso":"",
                  "parent_id":"3",
                  "city_id":"0",
                  "timezone":"7200"
               },
               {
                  "id":"57",
                  "name":"Germany",
                  "country_id":"57",
                  "area_id":"0",
                  "code":"DE",
                  "iso":"",
                  "parent_id":"3",
                  "city_id":"0",
                  "timezone":"3600"
               },
               {
                  "id":"68",
                  "name":"Spain",
                  "country_id":"68",
                  "area_id":"0",
                  "code":"ES",
                  "iso":"",
                  "parent_id":"3",
                  "city_id":"0",
                  "timezone":"3600"
               },
               {
                  "id":"89",
                  "name":"Greece",
                  "country_id":"89",
                  "area_id":"0",
                  "code":"GR",
                  "iso":"",
                  "parent_id":"3",
                  "city_id":"0",
                  "timezone":"7200"
               },
               {
                  "id":"109",
                  "name":"Italy",
                  "country_id":"109",
                  "area_id":"0",
                  "code":"IT",
                  "iso":"",
                  "parent_id":"3",
                  "city_id":"0",
                  "timezone":"3600"
               },
               {
                  "id":"131",
                  "name":"Lithuania",
                  "country_id":"131",
                  "area_id":"0",
                  "code":"LT",
                  "iso":"",
                  "parent_id":"3",
                  "city_id":"0",
                  "timezone":"7200"
               },
               {
                  "id":"185",
                  "name":"Romania",
                  "country_id":"185",
                  "area_id":"0",
                  "code":"RO",
                  "iso":"",
                  "parent_id":"3",
                  "city_id":"0",
                  "timezone":"7200"
               }
            ],
            "items":[
               {
                  "title":"Standard",
                  "is_default":true,
                  "deep_link":"https://pwieu.com/click-EQLHL9CB-KGCQCJEL?bt=25&tl=1"
               },
               {
                  "title":"Spain",
                  "is_default":false,
                  "deep_link":"https://pwieu.com/click-EQLHL9CB-KGCQCJEK?bt=25&tl=1"
               },
               {
                  "title":"Italy",
                  "is_default":false,
                  "deep_link":"https://pwieu.com/click-EQLHL9CB-KGCQCJEM?bt=25&tl=1"
               },
               {
                  "title":"Bulgaria",
                  "is_default":false,
                  "deep_link":"https://pwieu.com/click-EQLHL9CB-KGCQCJEN?bt=25&tl=1"
               },
               {
                  "title":"GREEK",
                  "is_default":false,
                  "deep_link":"https://pwieu.com/click-EQLHL9CB-KGCQCJEO?bt=25&tl=1"
               }
            ],
            "favicon":"",
            "shop_id":"0",
            "cpl":"",
            "cpa":"14.40 - 20.80",
            "rating":"1",
            "epc7days":0,
            "epc90days":0,
            "cpo7days":0,
            "cpo90days":0,
            "screen":"http://cdn77.cityads.com/graph/n/30/766_idealica.png",
            "is_exclusive":"0"
         },
         "30767":{
            "id":30767,
            "name":" Kaspersky PL CPS",
            "user_has_offer":"0",
            "cookie_ltv":"30",
            "site":"https://www.kaspersky.pl/",
            "categories":[
               {
                  "main_id":"10141",
                  "main_title":"Soft",
                  "other_categories":[
                     {
                        "id":"10023",
                        "title":"E-commerce"
                     },
                     {
                        "id":"10141",
                        "title":"Soft"
                     }
                  ]
               }
            ],
            "traffic-types":{
               "1":{
                  "id":"1",
                  "value":false,
                  "note":"",
                  "name":"Paid Search"
               },
               "2":{
                  "id":"2",
                  "value":true,
                  "note":"",
                  "name":"Display"
               },
               "3":{
                  "id":"3",
                  "value":true,
                  "note":"",
                  "name":"AdSpot/RichMedia/Sliding"
               },
               "4":{
                  "id":"4",
                  "value":true,
                  "note":"",
                  "name":"E-mail"
               },
               "5":{
                  "id":"5",
                  "value":true,
                  "note":"",
                  "name":"Social Networks"
               },
               "6":{
                  "id":"6",
                  "value":true,
                  "note":"",
                  "name":"Price-Comparison"
               },
               "7":{
                  "id":"7",
                  "value":true,
                  "note":"",
                  "name":"Coupon/Promo Codes"
               },
               "8":{
                  "id":"8",
                  "value":true,
                  "note":"",
                  "name":"Cashback"
               },
               "9":{
                  "id":"9",
                  "value":true,
                  "note":"",
                  "name":"Teaser advertisements"
               },
               "10":{
                  "id":"10",
                  "value":true,
                  "note":"",
                  "name":"Clickunder/Popunder"
               },
               "11":{
                  "id":"11",
                  "value":true,
                  "note":"",
                  "name":"Doorways"
               },
               "12":{
                  "id":"12",
                  "value":true,
                  "note":"",
                  "name":"Content Sites"
               },
               "14":{
                  "id":"14",
                  "value":true,
                  "note":"",
                  "name":"Incentive"
               },
               "15":{
                  "id":"15",
                  "value":true,
                  "note":"",
                  "name":"Messenger/SMS"
               },
               "17":{
                  "id":"17",
                  "value":true,
                  "note":"",
                  "name":"Toolbar"
               },
               "20":{
                  "id":"20",
                  "value":true,
                  "note":"",
                  "name":"Youtube"
               },
               "32":{
                  "id":"32",
                  "value":true,
                  "note":"",
                  "name":"Adult"
               },
               "33":{
                  "id":"33",
                  "value":true,
                  "note":"",
                  "name":"Autoredirect"
               },
               "34":{
                  "id":"34",
                  "value":true,
                  "note":"",
                  "name":"Branded context ads"
               },
               "35":{
                  "id":"35",
                  "value":true,
                  "note":"",
                  "name":"Retargeting"
               },
               "39":{
                  "id":"39",
                  "value":false,
                  "note":"",
                  "name":"Pop-up"
               },
               "40":{
                  "id":"40",
                  "value":true,
                  "note":"",
                  "name":"Mobile traffic in mobile version"
               }
            },
            "additional-traffic-types":null,
            "geo":[
               {
                  "id":"175",
                  "name":"Poland",
                  "country_id":"175",
                  "area_id":"0",
                  "code":"PL",
                  "iso":"",
                  "parent_id":"3",
                  "city_id":"0",
                  "timezone":"3600"
               }
            ],
            "items":null,
            "favicon":"/graph/s/30/767_kaspersky.ico",
            "shop_id":"3253",
            "cpl":"",
            "cpa":"4.15",
            "rating":"1",
            "epc7days":0,
            "epc90days":0,
            "cpo7days":0,
            "cpo90days":0,
            "screen":"http://cdn77.cityads.com/graph/n/30/767_kaspersky.png",
            "is_exclusive":"0"
         }
      }
   }
}`
