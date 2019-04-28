package main 

// import (
// 	"fmt"
// 	"net/http"
// 	"bytes"
// 	"io/ioutil"
// )
// func main() {
//     url := "http://restapi3.apiary.io/notes"
// 	fmt.Println("URL:>", url)
	
// 	// transforms json into byte array for networking transport
// 	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
// 	// method, url, body
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
// 	// Headers for the request
//     req.Header.Set("X-Custom-Header", "myvalue")
//     req.Header.Set("Content-Type", "application/json")

// 	// creates an HTTP client
// 	client := &http.Client{}
	
// 	// send a http request, response and error
//     resp, err := client.Do(req)
//     if err != nil {
//         panic(err)
// 	}
// 	// close the body of the response no matter what
// 	defer resp.Body.Close()
	
// 	// print status
// 	fmt.Println("response Status:", resp.Status)
// 	// print header
// 	fmt.Println("response Headers:", resp.Header)
// 	//read in the body using io tools
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	// cast the body to a string and print it.
//     fmt.Println("response Body:", string(body))
// }