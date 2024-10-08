package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	const getURL = "https://httpbin.org/get"
	const postURL = "https://httpbin.org/post"

	httpGet(getURL)

	httpPost(postURL)

	httpPostForm(postURL)
}

func httpGet(theURL string) {
	resp, err := http.Get(theURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", body)

}

func httpPost(theURL string) {

	postBody := []byte(`{"name" : "gabe", "age" : 43}`)
	postReader := bytes.NewReader(postBody)
	//http.Post()
	req, err := http.NewRequest("POST", theURL, postReader)
	if err != nil {
		fmt.Printf("client: could not create post request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: could not POST: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	//strings.NewReader()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: POST response body: %s\n", body)
}

func httpPostForm(theURL string) {

	resp, err := http.PostForm(theURL, url.Values{
		"Gabe": {"Anderson"},
	})
	if err != nil {
		fmt.Printf("client: could not POST: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: POST response body: %s\n", body)
}
