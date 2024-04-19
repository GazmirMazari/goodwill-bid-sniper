package internal

const (
	baseURL = "https://buyerapi.shopgoodwill.com/api"

	searchURL = baseURL + "/search/advancedsearch"

	favoritesURL  = baseURL + "/Favorite/GetAllFavoriteItemsByType?Type=al"
	encryptionKey = "6696D2E6F042FEC4D6E3F32AD541143B" // Example, replace with actual key
	iv            = "0000000000000000"
)
