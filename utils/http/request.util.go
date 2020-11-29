package reqjson

import (
	"encoding/json"
	"net/http"
	"time"
)

func Get(url string, target interface{}) error {
	client := http.Client{Timeout: 60 * time.Second}

	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}