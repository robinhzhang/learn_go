package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) GetURLContent(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	response, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	return string(response)
}
