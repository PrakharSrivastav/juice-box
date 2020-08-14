package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"time"
)

func main() {
		cmd := exec.Command("docker", "run",
			"-v", "/Users/prakharsrivastav/Workspace/juice-box/nikto:/tmp",
			"--rm", "prakharsrivastav/nikto:latest",
			"-h", "http://10.45.253.151:1301/portal.php",
			"-o", "xml", "--output", "/tmp/bWapp.xml",
			//"-Tuning", "923457",
		)
		//,
		//"-o", "json", "--output", "/tmp/resp.json",
		//"-Tuning", "923457")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		fmt.Println("Result: " + out.String())
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}

	//loadOSVDB()
	//ExampleScrape()
	//readDatabase()
}

type result struct {
	Result string `json:"result"`
	Data   struct {
		Search []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Lastseen       string `json:"lastseen"`
				BulletinFamily string `json:"bulletinFamily"`
				Description    string `json:"description"`
				Modified       string `json:"modified"`
				Published      string `json:"published"`
				Href           string `json:"href"`
				ID             string `json:"id"`
				Title          string `json:"title"`
				Type           string `json:"type"`
				Cvss           struct {
					Score  float64 `json:"score"`
					Vector string  `json:"vector"`
				} `json:"cvss"`
				Vhref string `json:"vhref"`
			} `json:"_source"`
			Highlight struct {
				ID   []string `json:"id"`
				Type []string `json:"type"`
			} `json:"highlight"`
			Sort            []float64 `json:"sort"`
			FlatDescription string    `json:"flatDescription"`
		} `json:"search"`
		ExactMatch    interface{} `json:"exactMatch"`
		References    interface{} `json:"references"`
		Total         int         `json:"total"`
		MaxSearchSize int         `json:"maxSearchSize"`
	} `json:"data"`
}

func readDatabase() {
	//str := "https://vulners.com/api/v3/search/id/?id=OSVDB:*&references=True&apiKey=<YOUR_API_KEY>"
	str := "https://vulners.com/api/v3/search/lucene/"
	u, err := url.Parse(str)
	if err != nil {
		log.Fatalf("bad url : %v", err)
	}
	client := &http.Client{Timeout: time.Second * 10}

	req := struct {
		Query  string `json:"query"`
		APIKey string `json:"apiKey"`
	}{
		APIKey: "<YOUR_API_KEY>",
		Query:  "type:osvdb AND id:630",
	}
	b, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
	}
	buffer := bytes.NewBuffer(b)

	reqPost, err := http.NewRequest(http.MethodPost, u.String(), buffer)
	if err != nil {
		log.Fatalf("bad request :: %v", err)
	}

	resp, err := client.Do(reqPost)
	if err != nil {
		log.Fatalf("could not send request :: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	var rr result
	err = json.Unmarshal(body, &rr)
	if err != nil {
		log.Fatalf("error unmarshalling response :: %v", err)
	}

	log.Printf("response is %+v", rr)
}

func loadOSVDB() {
	// Make HTTP request
	response, err := http.Get("https://vulners.com/api/v3/search/id/?id=OSVDB:631&references=True&apiKey=<YOUR_API_KEY>")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = response.Body.Close() }()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("#body").Each(func(i int, s *goquery.Selection) {
		log.Printf("%+v", s.Text())
	})
}

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}
