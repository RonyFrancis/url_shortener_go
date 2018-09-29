package shortener

import (
	"encoding/json"
	"net/http"

	"github.com/RonyFrancis/url_shortener_go/model"
)

// Response json struct
type Response struct {
	StatusCode   string `json:"status_code"`
	ShortenedURL string `json:"shortened_url"`
}

// AddURLHandler store url to database
func AddURLHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()          // Parses the request body
	x := r.Form.Get("url") // x will be "" if parameter is not set
	// TODO: valiate x is empty or not a valid url
	shorten := model.CodeGenerator(8, x)
	response := Response{StatusCode: "200", ShortenedURL: shorten}
	json.NewEncoder(w).Encode(response)
}
