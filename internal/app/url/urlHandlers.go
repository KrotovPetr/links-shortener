package url

import (
	"github.com/google/uuid"
	"net/http"
)

func RedirectURLHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		url := request.URL.Path

		response.WriteHeader(http.StatusTemporaryRedirect)
		response.Header().Set("Location", request.URL.Path)
		response.Write([]byte("http://localhost:8080" + url))

	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Only get method type"))
	}
}

func ShortenURLHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		//На будущие спринты
		//body, _ := ioutil.ReadAll(request.Body)
		//fmt.Print(body)
		shortURL := uuid.New()

		response.WriteHeader(http.StatusCreated)
		response.Header().Set("content-type", "text/plain")
		response.Write([]byte("http://localhost:8080/" + shortURL.String()))

	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Only post method type"))
	}

}

func ErrorHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusBadRequest)
	response.Write([]byte("Incorrect method"))
}
