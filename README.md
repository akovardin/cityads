# Cityads API

A Go wrapper around Admitad API

## Install

```
go get github.com/horechek/cityads
```

## Method

There are 1 common method to communicate with api:

```go
Call(url, method string, params url.Values, result interface{}) error
```

## Example

Create client:

```go
var (
    base64Header = "xxx"
    clientId = "xxx"
)

client := cityads.NewClient(
    "https://cityads.com/api/rest/webmaster/json",
    "xxx",
)
```

`xxx` - it is remote auth key

Fetch banners from API

```go
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
```

`Banner`, `Websites`, `Campaigns` - it's types from current application, not from package

## Credentials

You can get `base64Header` and `clientId` here: https://www.admitad.com/ru/webmaster/account/settings/credentials/