package main
import (
        "net/http"
        "net/http/httputil"
        "net/url"
      	"time"
      	"net"
        "log"
        "fmt"
        "crypto/tls"
)
func main() {
     go ReverseHttpsProxy(445,"https://127.0.0.1:443/","my.crt","my.key")
     ReverseHttpProxy(8081,"http://127.0.0.1:8080/")
}

func ReverseHttpsProxy(port int,dst string,crt string,key string) {
     u, e := url.Parse(dst)
     if e != nil {
     	log.Fatal("Bad destination.")
     } 
     h := httputil.NewSingleHostReverseProxy(u)
     //if your certificate signed by yourself,you need use this bypass secure verify
     var InsecureTransport http.RoundTripper = &http.Transport{
	        Dial: (&net.Dialer{
	                Timeout:   30 * time.Second,
	                KeepAlive: 30 * time.Second,
	        }).Dial,
	        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	        TLSHandshakeTimeout: 10 * time.Second,
	}
    h.Transport = InsecureTransport
    err := http.ListenAndServeTLS(fmt.Sprintf(":%d",port),crt, key ,h) 
    if err != nil {
    	log.Println("Error:",err)
    }       
}

func ReverseHttpProxy(port int,dst string) {
	 u, e := url.Parse(dst)
     if e != nil {
     	log.Fatal("Bad http destination.")
     } 
     h := httputil.NewSingleHostReverseProxy(u)
    
    err := http.ListenAndServe(fmt.Sprintf(":%d",port),h) 
    if err != nil {
    	log.Println("Error:",err)
    } 
}