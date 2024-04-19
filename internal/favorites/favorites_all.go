package favorites

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	baseURL      = "https://buyerapi.shopgoodwill.com/api"
	favoritesURL = baseURL + "/Favorite/GetAllFavoriteItemsByType?Type=all"
)

type IFavoritesResponse interface {
	FetchAll() ([]interface{}, error)
}

type Favorites struct {
	Client *http.Client
}

func (f *Favorites) FetchAll() ([]interface{}, error) {
	// Create an http request with context.
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", favoritesURL, nil)
	if err != nil {
		return nil, err
	}

	// Send the request.
	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read body here before closing it.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	// Unmarshal response to the "favorites" variable.
	var favorites []interface{}
	err = json.Unmarshal(body, &favorites)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}
