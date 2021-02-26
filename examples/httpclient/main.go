package main

import (
	"eureka-client/httpclient"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	getRow()
	getText()
	postForm()
	postJson()
	postMultipart()
	handler()
	save()
	getJson()
	customHttp()
	requestInterceptor()
}

func getRow() {
	raw, err := httpclient.Get("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Send().
		Raw()
	if err != nil {
		panic(err)
	}
	fmt.Println(raw)
}

func getText() {
	text, err := httpclient.Get("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func postForm() {
	text, err := httpclient.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Form(url.Values{
			"form1": {"value1"},
			"form2": {"123"},
		}).
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func postJson() {
	text, err := httpclient.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Json(map[string]interface{}{
			"json1": "value1",
			"json2": 2,
		}).
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func postMultipart() {
	text, err := httpclient.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Multipart(httpclient.FileForm{
			Value: url.Values{
				"form1": {"value1"},
				"form2": {"value2"},
			},
			File: map[string]string{
				"file1": "./examples/main.go",
				"file2": "./examples/main.go",
			},
		}).
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func save() {
	err := httpclient.Get("https://www.cnblogs.com/bener/p/10683404.html").
		Send().
		Save("./10683404.html")
	if err != nil {
		panic(err)
	}
}

func getJson() {
	var v map[string]interface{}
	err := httpclient.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Json(map[string]interface{}{
			"json1": "value1",
			"json2": 2,
		}).
		Send().
		Json(&v)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func handler() {
	result := httpclient.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Json(map[string]interface{}{
			"json1": "value1",
			"json2": 2,
		}).
		Send()
	if result.Err != nil {
		panic(result.Err)
	}

	b, err := ioutil.ReadAll(result.Resp.Body)
	if err != nil {
		panic(err)
	}
	defer result.Resp.Body.Close()

	fmt.Println(string(b))
}

func customHttp() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	text, err := httpclient.Request("https://github.com/xuanbo", "OPTIONS", client).
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func requestInterceptor() {
	logRequestInterceptor := func(request *http.Request) error {
		fmt.Println(request.URL)
		return nil
	}
	httpclient.AddRequestInterceptors(logRequestInterceptor)

	text, err := httpclient.Get("https://github.com/xuanbo").
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}
