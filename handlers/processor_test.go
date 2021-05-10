package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "test")
	}))
	defer ts.Close()

	var testcases = []struct {
		uri   string
		isErr bool
	}{
		{"hrh", true},
		{ts.URL, false},
	}

	for i, cs := range testcases {
		res := doRequest(cs.uri)
		if (res.Err != nil) != cs.isErr {
			t.Errorf("case # %d wait for %t got %t", i, cs.isErr, res.Err != nil)
		}
	}
}
