package itunes

import (
	"net/http"
)

// SearchOptions represents the available options for all search methods.
//
// iTunes Search API docs: https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/Searching.html
type SearchOptions struct {
	// The two-letter country code for the store you want to search.
	// The search uses the default store front for the specified country. For example: US. The default is US.
	Country *CountryValue `url:"country,omitempty"`

	// AttributeValue represents the attribute you want to search for in the stores, relative to the specified media type.
	// For example, if you want to search for an artist by name specify `entity=allArtist&attribute=allArtistTerm`.
	// In this example, if you search for `term=maroon`, iTunes returns “Maroon 5” in the search results, instead of all artists who have ever recorded a song with the word “maroon” in the title.
	// The default is all attributes associated with the specified media type.
	Attribute *AttributeValue `url:"attribute,omitempty"`

	// The number of search results you want the iTunes Store to return. For example: 25. The default is 50.
	// Can be 1 to 200.
	Limit *int `url:"limit,omitempty"`

	// The language, English or Japanese, you want to use when returning search results.
	// Specify the language using the five-letter codename. For example: en_us. The default is en_us (English).
	Lang *LanguageValue `url:"lang,omitempty"`

	// A flag indicating whether or not you want to include explicit content in your search results. The default is Yes.
	Explicit *ExplicitValue `url:"explicit,omitempty"`

	// // Undocumented parameter that should be used with `Limit` for pagination purposes.
	// Offset *int `url:"offset,omitempty"`
}

type searchOptions struct {
	SearchOptions

	Term   *string      `url:"term,omitempty"`
	Media  *MediaValue  `url:"media,omitempty"`
	Entity *EntityValue `url:"entity,omitempty"`
}

func search[T any](client *Client, term string, media MediaValue, entity EntityValue, opt *SearchOptions) ([]*T, *http.Response, error) {
	if opt == nil {
		opt = &SearchOptions{}
	}
	opts := &searchOptions{
		SearchOptions: *opt,
		Term:          &term,
		Media:         &media,
		Entity:        &entity,
	}

	req, err := client.NewRequest(http.MethodGet, "search", opts)
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
