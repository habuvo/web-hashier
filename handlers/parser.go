package handlers

import (
	"log"
	"net/url"
)

// ParseHosts parse stringsinput and augments it with a schema
func ParseHosts(raw []string) (hosts []string) {
	for _, candidate := range raw {
		uri, err := url.ParseRequestURI(candidate)
		if err != nil || uri.Host == "" {
			uri, err = url.ParseRequestURI("http://" + candidate)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}
		hosts = append(hosts, uri.String())
	}
	return
}
