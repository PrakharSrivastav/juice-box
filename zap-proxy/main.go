package main

import (
	"crypto/tls"
	"github.com/zaproxy/zap-api-go/zap"
	"log"
	"os"
	"strconv"
	"time"
)

func setupClient() (string, zap.Interface, error) {
	proxy := "http://10.45.253.151:8090"
	target := "http://10.45.253.151:1301/"
	//target := "http://192.168.0.18:8080/WebGoat/"
	//target := "http://192.168.0.18:9090/WebWolf/"
	//target := "http://192.168.0.18:3000/"
	//target := "https://public-firing-range.appspot.com/"
	cfg := zap.Config{
		Proxy:     proxy,
		APIKey:    "5364864132243598723485",
		TLSConfig: tls.Config{InsecureSkipVerify: true},
	}
	client, err := zap.NewClient(&cfg)
	if err != nil {
		return "", nil, err
	}

	return target, client, err
}

// zap owasp scanner runs as a proxy that needs interaction with a client
// it provides a client that can be used to interact with the proxy
func main() {
	log.Print("test")
	target, client, err := setupClient()
	if err != nil {
		log.Fatalf("cannot connect to zap proxy :: %v", err)
	}
	//u := strconv.Quote("loginUrl=http://10.45.253.151:1301/login.php&loginRequestData=" + strconv.Quote("username=bee&password=bug"))
	//method, err := client.Authentication().SetAuthenticationMethod("tt", "formBasedAuthentication", u)
	/*scans, err := client.Ascan().RemoveAllScans()
	if err != nil {
		log.Fatalf("auth error :: %v\n", err)
	}
	log.Print("method :: %v", scans)
	*/
	scanID, err := regularSpider(target, client)
	if err != nil {
		log.Fatalf("error running regular spider :: %v", err)
	}

	if err = ajaxSpider(target, client); err != nil {
		log.Fatalf("error running ajax spider :: %v", err)
	}

	if err = runPassiveScan(client); err != nil {
		log.Fatalf("error running passive scan :: %v", err)
	}
	// Give the passive scanner a chance to complete
	time.Sleep(2000 * time.Millisecond)

	if err = runActiveScan(scanID, target, client); err != nil {
		log.Fatalf("error running active scan :: %v", err)
	}

	if err = generateFile(client); err != nil {
		log.Fatalf("error creating report :: %v", err)
	}

}

func ajaxSpider(target string, client zap.Interface) error {
	log.Println("starting ajax spider")
	_, err := client.AjaxSpider().Scan(target, "", "", "")
	if err != nil {
		return err
	}

	// run maximum for 2 minutes
	timeout := time.Now().Add(time.Minute * 1)
	for {
		status, err := client.AjaxSpider().Status()
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		if status["status"] != "running" {
			break
		}
		if time.Now().After(timeout) {
			log.Println("timeout exceeded for ajax spider")
			_, _ = client.AjaxSpider().Stop()
			break
		}
	}
	log.Println("completed ajax spider")
	return nil

	/*
		num, err := client.AjaxSpider().NumberOfResults()
		if err != nil {
			log.Fatal(err)
		}
		results, err := client.AjaxSpider().Results("0", num["numberOfResults"].(string))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("num :: %v", num)
		log.Printf("%v", results)
	*/
}

func regularSpider(target string, client zap.Interface) (string, error) {
	log.Println("starting regular spider")
	resp, err := client.Spider().Scan(target, "", "True", "", "")
	if err != nil {
		return "", err
	}

	scanID := resp["scan"].(string)
	for {
		time.Sleep(1000 * time.Millisecond)
		resp, err = client.Spider().Status(scanID)
		if err != nil {
			return "", err
		}
		progress, err := strconv.Atoi(resp["status"].(string))
		if err != nil {
			return "", err
		}
		if progress >= 100 {
			break
		}
	}
	status, err := client.Spider().Status(scanID)
	if err != nil {
		return "", err
	}

	log.Printf("regular spider completed :: %v , scanID :: %s", status, scanID)
	return scanID, nil
	/*
		results, err := client.Spider().Results(scanID)
		if err != nil {
			log.Fatal(err)
		}
		for _, items := range results["results"].([]interface{}) {
			log.Printf("urls : %+v", items)
		}
	*/
}

func generateFile(client zap.Interface) error {
	log.Println("generating report")
	j, err := client.Core().Jsonreport()
	if err != nil {
		return err
	}

	f, err := os.Create("./zap-proxy/bwapp-2.json")
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	if _, err = f.Write(j); err != nil {
		return err
	}
	log.Println("report generated")
	return nil
}

func runPassiveScan(client zap.Interface) error {
	log.Println("running passive scan")
	pScan := client.Pscan()
	for {
		scan, err := pScan.RecordsToScan()
		if err != nil {
			return err
		}

		time.Sleep(time.Second)
		i, err := strconv.Atoi(scan["recordsToScan"].(string))
		if err != nil {
			return nil
		}

		if i <= 0 {
			break
		}
	}
	log.Println("completed passive scan")
	return nil
}

func runActiveScan(scanID, target string, client zap.Interface) error {
	log.Println("running active scan")
	resp, err := client.Ascan().Scan(target, "True", "False", "", "", "", "")
	if err != nil {
		return err
	}
	resp, _ = client.Ascan().Status(scanID)
	log.Printf("%v\n", resp)
	for {
		time.Sleep(1000 * time.Millisecond)
		resp, _ = client.Ascan().Status(scanID)
		log.Printf("%v\n", resp)
		progress, err := strconv.Atoi(resp["status"].(string))
		if err != nil {
			return err
		}
		if progress >= 100 {
			break
		}
	}
	log.Println("completed active scan")
	return nil
}
