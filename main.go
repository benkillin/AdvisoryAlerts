package main

import (
	"net/http"
	"net/url"
)
import "io/ioutil"
import "fmt"
import "strings"
import "./RSSConsumer"

func getConfiguration() {

}

func main() {
	
	content := RSSConsumer.Fetch("http://www.google.com")
	/* https://dlintw.github.io/gobyexample/public/http-client.html */

	// Some random stuff to just get started with imports
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("get:\n", string(body))

	resp, err = http.PostForm("http://duckduckgo.com",
		url.Values{"q": {"github"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println("post:\n", string(body))
	strings.Contains(content, "tastic")


}
