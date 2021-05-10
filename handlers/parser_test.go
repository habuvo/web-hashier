package handlers_test

import (
	"testing"

	"github.com/habuvo/web-hashier/handlers"
)

var testcases = []struct {
	raw   []string
	hosts []string
}{
	{nil, nil},
	{[]string{"1"}, []string{"http://1"}},
	{[]string{"www.google.com"}, []string{"http://www.google.com"}},
	{[]string{"http://www.google.com"}, []string{"http://www.google.com"}},
	{[]string{"www.r1.com", "www.r2.com", "www.r3.com"}, []string{"http://www.r1.com", "http://www.r2.com", "http://www.r3.com"}},
}

func TestParseHosts(t *testing.T) {
	for i, cs := range testcases {
		hosts := handlers.ParseHosts(cs.raw)
		if len(hosts) != len(cs.hosts) {
			t.Errorf("case # %d wait for %d got %d", i, len(cs.hosts), len(hosts))
		}
		if len(hosts) != 0 && hosts[0] != cs.hosts[0] {
			t.Errorf("case # %d wait for %s got %s", i, cs.hosts[0], hosts[0])
		}
	}
}
