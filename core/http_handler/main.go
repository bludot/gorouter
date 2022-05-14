package http_handler

import "net/http"

type HTTPResponse struct {
	Header     http.Header
	StatusCode int
	Body       []byte
}

func Handler(w http.ResponseWriter, r *http.Request, response HTTPResponse) {

}
