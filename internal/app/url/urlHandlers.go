package url

import (
	"encoding/json"
	"net/http"
)

type URLData struct {
	URL string `json:"url"`
}

type ErrorData struct {
	ERROR string `json:"error"`
}

func RedirectURLHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		fullURL := "http://localhost:8080" + request.URL.Path
		urlData := URLData{
			URL: fullURL,
		}

		jsonData, _ := json.Marshal(urlData)

		response.WriteHeader(http.StatusTemporaryRedirect)
		response.Header().Set("Location", request.URL.Path)
		response.Write([]byte(jsonData))

	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Only get method type"))
	}
}

func ShortenURLHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		shortURL := "fgdggd"
		fullURL := "http://localhost:8080/" + shortURL

		urlData := URLData{
			URL: fullURL,
		}

		jsonData, _ := json.Marshal(urlData)

		//На будущие спринты
		//body, _ := ioutil.ReadAll(request.Body)
		//fmt.Print(body)
		//shortURL := uuid.New()

		response.WriteHeader(http.StatusCreated)
		response.Header().Set("Content-type", "text/plain")
		response.Write([]byte(jsonData))

	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Only post method type"))
	}

}

func ErrorHandler(response http.ResponseWriter, request *http.Request) {
	errorData := ErrorData{
		ERROR: request.Method + " is incorrect method",
	}

	jsonData, _ := json.Marshal(errorData)

	response.WriteHeader(http.StatusBadRequest)
	response.Write([]byte(jsonData))
}
