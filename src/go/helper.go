package swagger

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func getSecretHash(req *http.Request, prefix string) (string, error) {
	url := strings.TrimPrefix(req.URL.Path, prefix)
	params := strings.Split(url, "/")
	return params[0], nil
}

func formatResponse(v interface{}) []byte {
	marsharlJson, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return marsharlJson
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
