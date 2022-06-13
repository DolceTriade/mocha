package middleware

import (
	"github.com/vitorsalgado/mocha/internal/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRecover(t *testing.T) {
	msg := "error test"
	fn := func(w http.ResponseWriter, r *http.Request) {
		panic(msg)
	}

	ts := httptest.NewServer(Recovery(http.HandlerFunc(fn)))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, http.StatusTeapot, res.StatusCode)
	assert.Equal(t, "text/plain", res.Header.Get("content-type"))
	assert.True(t, strings.Contains(string(body), msg))
}
