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
	r.ParseForm()          // Parses the request body
	x := r.Form.Get("url") // x will be "" if parameter is not set
	// TODO: valiate x is empty or not a valid url
	shorten := model.CodeGenerator(8, x)
	response := Response{StatusCode: "200", ShortenedUrl: shorten}
	json.NewEncoder(w).Encode(response)
}
