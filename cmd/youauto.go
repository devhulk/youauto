package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type NotionPage struct {
	ID string `json:"id"`
}

var pages []NotionPage

type NotionPageData struct {
	Title string `json:"title"`
}

var pageTitle NotionPageData
var pageTitles []NotionPageData

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "We automate stuff.\n")
}
func syncNotion(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("got /hello request\n")
	//io.WriteString(w, "Hello, HTTP!\n")
	client := &http.Client{}
	url := "https://notion-proxy.devhulk.workers.dev/"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Custom-PSK", os.Getenv("CLOUDFLARE_CUSTOM_AUTH_TOKEN"))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	defer resp.Body.Close()
	fmt.Printf("response Status: %v\n", resp.Status)
	fmt.Printf("response Headers: %v\n", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	err1 := json.Unmarshal(body, &pages)
	fmt.Printf("%+v", pages)
	if err1 != nil {
		fmt.Printf("error: %v\n", err)
	}

	for _, page := range pages {
		fmt.Printf("%+v\n", page.ID)
		getPageData(page.ID)
	}

	b, err2 := json.Marshal(pageTitles)
	if err2 != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("%+v\n", pageTitles)
	io.WriteString(w, string(b))
}

func getPageData(id string) {
	url := fmt.Sprintf("https://notion-pages.devhulk.workers.dev?id=%s", id)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	defer resp.Body.Close()
	fmt.Printf("response Status: %v\n", resp.Status)
	fmt.Printf("response Headers: %v\n", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err1 := json.Unmarshal(body, &pageTitle)
	if err1 != nil {
		fmt.Printf("error: %v\n", err)
	}
	pageTitles = append(pageTitles, pageTitle)

}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/sync/notion", syncNotion)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("ListenAndServe: %v\n", err)
	}
}
