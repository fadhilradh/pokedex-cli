package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetMap(url *string) Maps {
	reqUrl := baseURL
	if url != nil {
		reqUrl = *url
	}

	fmt.Println(reqUrl)
	res, err := http.Get(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("response: %q", body)
	}

	maps := Maps{}
	err = json.Unmarshal(body, &maps)
	if err != nil {
		log.Fatal(err)
	}

	return maps
}
