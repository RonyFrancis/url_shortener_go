package shortener

import (
	"encoding/json"
	"net/http"

	"github.com/RonyFrancis/url_shortener_go/model"
)

type Response struct {
	StatusCode   string `json:"status_code"`
	ShortenedUrl string `json:"shortened_url"`
}

func AddUrlHandler(w http.ResponseWriter, r *http.Request) {
	shorten := url.CodeGenerator(8)
	response := Response{StatusCode: "200", ShortenedUrl: shorten}
	json.NewEncoder(w).Encode(response)
}
