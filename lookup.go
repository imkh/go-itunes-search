package itunes

import (
	"net/http"
)

// SearchOptions represents the available options for all lookup methods.
//
// iTunes Lookup API docs: https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/Searching.html
type LookupOptions struct {
	// The two-letter country code for the store you want to search.
	// The search uses the default store front for the specified country. For example: US. The default is US.
	Country *CountryValue `url:"country,omitempty"`

	// The number of search results you want the iTunes Store to return. For example: 25. The default is 50.
	// Can be 1 to 200.
	Limit *int `url:"limit,omitempty"`

	// The language, English or Japanese, you want to use when returning search results.
	// Specify the language using the five-letter codename. For example: en_us. The default is en_us (English).
	Lang *LanguageValue `url:"lang,omitempty"`

	// A flag indicating whether or not you want to include explicit content in your search results. The default is Yes.
	Explicit *ExplicitValue `url:"explicit,omitempty"`

	// // Undocumented parameter. The default is popular.
	// Sort *SortValue `url:"limit,omitempty"`
}

type lookupOptions struct {
	LookupOptions

	ID          *string `url:"id,omitempty"`
	AMGArtistID *string `url:"amgArtistId,omitempty"`
	UPC         *string `url:"upc,omitempty"`
	AMGAlbumID  *string `url:"amgAlbumId,omitempty"`
	AMGVideoID  *string `url:"amgVideoId,omitempty"`
	ISBN        *string `url:"isbn,omitempty"`
	BundleID    *string `url:"bundleId,omitempty"`
	// Entity      *EntityValue `url:"entity,omitempty"`
}

func lookup[T any](client *Client, id *string, amgArtistId *string, upc *string, amgAlbumId *string, amgVideoId *string, isbn *string, bundleId *string, opt *LookupOptions) ([]*T, *http.Response, error) {
	if opt == nil {
		opt = &LookupOptions{}
	}
	opts := &lookupOptions{
		LookupOptions: *opt,
		ID:            id,
		AMGArtistID:   amgArtistId,
		UPC:           upc,
		AMGAlbumID:    amgAlbumId,
		AMGVideoID:    amgVideoId,
		ISBN:          isbn,
	}

	req, err := client.NewRequest(http.MethodGet, "lookup", opts)
	if err != nil {
		return nil, nil, err
	}

	var gr GenericResponse[[]*T]
	resp, err := client.Do(req, &gr)
	if err != nil {
		return nil, resp, err
	}

	return gr.Results, resp, err
}
