package main

import (
	"context"
	"github.com/PrakharSrivastav/swagger/swagger"
	"github.com/antihax/optional"
	"log"
	"net/http"
	"time"
)

func main() {

	cfg := swagger.NewConfiguration()

	cfg.BasePath = "http://10.45.253.151:8090"
	cfg.HTTPClient = &http.Client{
		Timeout: time.Second * 20,
	}
	cfg.AddDefaultHeader("X-ZAP-API-Key", "5364864132243598723485")
	cfg.AddDefaultHeader("apiKey", "5364864132243598723485")
	cfg.AddDefaultHeader("Accept", "application/json")

	client := swagger.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), swagger.ContextAPIKey, swagger.APIKey{
		Key: "5364864132243598723485",
	})

	scanners, resp, err := client.SpiderApi.SpiderActionScan(ctx, &swagger.SpiderApiSpiderActionScanOpts{
		Url:         optional.NewString("http://10.45.253.151:1301/"),
		MaxChildren: optional.NewInt32(2),
		Recurse:     optional.NewBool(true),
		ContextName: optional.NewString("context-1"),
	})
	if err != nil {
		log.Println(err)
	}
	log.Printf("resp :: %v", resp)

	log.Printf("scanners :: %v", scanners)
}
