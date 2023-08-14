package backend

import (
	"net/http"
	"os"
	"strings"
	"time"
)

func Connect(url string, delay int, retry int) bool {
	client := &http.Client{}

	if os.Getenv("NEXT_PUBLIC_DRUPAL_BASE_URL") != "" {
		if !strings.HasSuffix(url, "/") {
			url += "/"
		}
		url += "jsonapi"
	}

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	for i := 0; i < retry; i++ {
		r, err := client.Get(url)
		if err == nil && r.StatusCode == http.StatusOK {
			return true
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}

	return false
}
