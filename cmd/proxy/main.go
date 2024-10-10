package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func handleProxy(responseWriter http.ResponseWriter, request *http.Request) {
	proxyTarget := request.Header.Get("Proxy-Target")

	if proxyTarget == "" {
		http.Error(responseWriter, "Proxy-Target Header is required", http.StatusBadRequest)
		return
	}

	url, err := url.Parse(proxyTarget)
	if err != nil {
		http.Error(responseWriter, "Invalid Proxy-Target Header", http.StatusBadRequest)
		return
	}
	fmt.Printf("url: %+v\n", url)

	fmt.Printf("Proxy Request: %+v\n", request.URL)

	// 원격 서버 URL 설정
	request.URL.Scheme = "https"
	request.URL.Host = url.Host

	// 원격 서버로 요청 전달
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		http.Error(responseWriter, "Server Error", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	responseWriter.WriteHeader(response.StatusCode)
	io.Copy(responseWriter, response.Body)
}

func main() {
	http.HandleFunc("/ping", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Write([]byte("pong"))
	})
	http.HandleFunc("/", handleProxy)

	fmt.Println("Proxy Server Start")

	// 웹 서버 실행
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
