package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)


/*
*	Reverse Proxy without using the premade httputil, but instead writing behavior on its own
*	Just to get a feel for it and play arround with the basics
*/

func init(){
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// http proxy, 
func main(){
	const port = "8888"
	fmt.Println("running on Port:",port)
	
	proxyURL, err := url.Parse("http://vprobst.de")		// the addr we proxy to
	if err != nil{
		log.Fatal(err)
	}

	proxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request){
		// build a new req to the server
		req.Host = proxyURL.Host
		req.URL.Host = proxyURL.Host
		req.URL.Scheme = proxyURL.Scheme
		req.RequestURI = ""

		// copy the IP and forward it instead of the ip of the proxy-server
		remoteAddr, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil{
			fmt.Print(rw, err)
			return
		}
		req.Header.Set("X-Forwarded-For", remoteAddr)

		// send req, receive a response to server
		response, err := http.DefaultClient.Do(req)
		if err != nil{
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Print(rw, err)
			return
		}

		// copy response headers and paste them into client-response
		for key, values := range response.Header{
			for _, value := range values{
				rw.Header().Set(key,value)
			}
		}

		// -> 1/3 optional streaming-data, to enable streaming of data
		// otherwise it would wait for all flushes, then send one big chunk.
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-time.Tick(10*time.Millisecond):
					rw.(http.Flusher).Flush()
				case <-done:
					return
				}
			}
		}()
		// <- 2/3 optional streaming-data 

		// trailers in the request
		trailerKeys := []string{}
		for key := range response.Trailer{
			trailerKeys = append(trailerKeys, key)
		}
		rw.Header().Set("Trailer", strings.Join(trailerKeys, ","))

		//create the response body
		rw.WriteHeader(response.StatusCode)
		io.Copy(rw, response.Body)

		// fill the trailers in the response
		for key, values := range response.Trailer{
			for _,value := range values{
				rw.Header().Set(key, value)
			}
		}

		//-> 3/3 optional streaming-data closes the channel		
		close(done)								
	})
	http.ListenAndServe(":"+port, proxy)
}
