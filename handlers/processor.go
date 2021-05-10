package handlers

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type ProcessResult struct {
	URI  string
	Hash string
	Err  error
}

//ProcessRequest get uri for the request and returns ProcessResult with hash and error of the request
func ProcessRequest(wg *sync.WaitGroup, in chan string, result chan ProcessResult) {
	for uri := range in {
		result <- doRequest(uri)
	}
	wg.Done()
}

func doRequest(uri string) (res ProcessResult) {
	resp, err := http.Get(uri)
	if err != nil {
		res.Err = err
		return
	}

	defer resp.Body.Close()

	var buf bytes.Buffer

	err = resp.Write(&buf)
	if err != nil {
		res.Err = err
		return
	}

	md5Hash := md5.Sum(buf.Bytes())
	res.URI = uri
	res.Hash = hex.EncodeToString(md5Hash[:])

	return
}

func ProcessResponse(cancel context.CancelFunc, res chan ProcessResult) {
	for result := range res {
		if result.Err != nil {
			log.Println(result.Err.Error())
			continue
		}
		fmt.Printf("%s %s\n", result.URI, result.Hash)
	}
	cancel()
}
