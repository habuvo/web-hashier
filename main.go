package main

import (
	"context"
	"flag"
	"log"
	"sync"

	"github.com/habuvo/web-hashier/handlers"
)

func main() {
	var parallel int

	flag.IntVar(&parallel, "parallel", 3, "number of the concurrent requests")

	flag.Parse()

	// simple logic collision check
	if parallel == 0 {
		log.Fatal("could not process requests without processors")
	}

	// start processors pool and reporter instance
	wg := new(sync.WaitGroup)
	uri := make(chan string, parallel)
	res := make(chan handlers.ProcessResult, parallel)

	ctx, cancel := context.WithCancel(context.Background())

	go handlers.ProcessResponse(cancel, res)

	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go handlers.ProcessRequest(wg, uri, res)
	}

	for _, host := range handlers.ParseHosts(flag.Args()) {
		uri <- host
	}

	close(uri)
	wg.Wait()

	close(res)
	<-ctx.Done()

	log.Println("finished")
}
