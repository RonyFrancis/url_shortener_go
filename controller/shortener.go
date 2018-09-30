package shortener

import (
	"encoding/json"
	"github.com/RonyFrancis/url_shortener_go/model"
	"github.com/gorilla/mux"
	"net/http"
)

// Response json struct
type Response struct {
	StatusCode   string `json:"status_code"`
	ShortenedURL string `json:"shortened_url"`
}

// GetAllURLResponse response block
type GetAllURLResponse struct {
	StatusCode string          `json:"status_code"`
	URLS       []model.LinkURL `json:"urls"`
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

// GetAllUrls fetches all urls from database
func GetAllUrls(w http.ResponseWriter, r *http.Request) {
	response := GetAllURLResponse{StatusCode: "200", URLS: model.Geturls()}
	json.NewEncoder(w).Encode(response)
}

// RedirectURL redirect to original url
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := model.GetoriURL(vars["url"])
	http.Redirect(w, r, url, http.StatusSeeOther)
}
