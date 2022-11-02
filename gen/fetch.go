package gen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchManga() error {

	fmt.Println("Fetching manga list...")
	r, err := http.Get("https://api.mangadex.org/manga/6670ee28-f26d-4b61-b49c-d71149cd5a6e")
	if err != nil {
		return err
	}

	var resp GetMangaResponse
	err = json.NewDecoder(r.Body).Decode(&resp)

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Printf("\n\n%+v\n\n", string(out))
	return nil

}
