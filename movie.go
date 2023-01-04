package itunes

import (
	"net/http"
	"time"
)

// MovieService handles communication with the movie related methods
// of the iTunes Search API.
type MovieService struct {
	client *Client
}

type MovieArtist struct {
	WrapperType      string `json:"wrapperType"`
	ArtistType       string `json:"artistType"`
	ArtistName       string `json:"artistName"`
	ArtistLinkURL    string `json:"artistLinkUrl"`
	ArtistID         int    `json:"artistId"`
	AmgArtistID      *int   `json:"amgArtistId,omitempty"`
	PrimaryGenreName string `json:"primaryGenreName"`
	PrimaryGenreID   int    `json:"primaryGenreId"`
}

type Movie struct {
	WrapperType             string    `json:"wrapperType"`
	Kind                    string    `json:"kind"`
	CollectionID            *int      `json:"collectionId,omitempty"`
	TrackID                 int       `json:"trackId"`
	ArtistName              string    `json:"artistName"`
	CollectionName          *string   `json:"collectionName,omitempty"`
	TrackName               string    `json:"trackName"`
	CollectionCensoredName  *string   `json:"collectionCensoredName,omitempty"`
	TrackCensoredName       string    `json:"trackCensoredName"`
	CollectionArtistID      *int      `json:"collectionArtistId,omitempty"`
	CollectionArtistViewURL *string   `json:"collectionArtistViewUrl,omitempty"`
	CollectionViewURL       *string   `json:"collectionViewUrl,omitempty"`
	TrackViewURL            string    `json:"trackViewUrl"`
	PreviewURL              *string   `json:"previewUrl,omitempty"`
	ArtworkURL30            string    `json:"artworkUrl30"`
	ArtworkURL60            string    `json:"artworkUrl60"`
	ArtworkURL100           string    `json:"artworkUrl100"`
	CollectionPrice         *float64  `json:"collectionPrice,omitempty"`
	TrackPrice              *float64  `json:"trackPrice,omitempty"`
	TrackRentalPrice        *float64  `json:"trackRentalPrice,omitempty"`
	CollectionHdPrice       *float64  `json:"collectionHdPrice,omitempty"`
	TrackHdPrice            *float64  `json:"trackHdPrice,omitempty"`
	TrackHdRentalPrice      *float64  `json:"trackHdRentalPrice,omitempty"`
	ReleaseDate             time.Time `json:"releaseDate"`
	CollectionExplicitness  string    `json:"collectionExplicitness"`
	TrackExplicitness       string    `json:"trackExplicitness"`
	DiscCount               *int      `json:"discCount,omitempty"`
	DiscNumber              *int      `json:"discNumber,omitempty"`
	TrackCount              *int      `json:"trackCount,omitempty"`
	TrackNumber             *int      `json:"trackNumber,omitempty"`
	TrackTimeMillis         *int      `json:"trackTimeMillis,omitempty"`
	Country                 string    `json:"country"`
	Currency                string    `json:"currency"`
	PrimaryGenreName        string    `json:"primaryGenreName"`
	ContentAdvisoryRating   string    `json:"contentAdvisoryRating"`
	LongDescription         string    `json:"longDescription"`
	ShortDescription        *string   `json:"shortDescription,omitempty"`
	HasITunesExtras         *bool     `json:"hasITunesExtras,omitempty"`
	ArtistID                *int      `json:"artistId,omitempty"`
	ArtistViewURL           *string   `json:"artistViewUrl,omitempty"`
}

// SearchArtists searches movie artists.
func (s *MovieService) SearchArtists(term string, opt *SearchOptions) ([]*MovieArtist, *http.Response, error) {
	return search[MovieArtist](s.client, term, MovieMedia, MovieArtistEntityMovie, opt)
}

// LookupArtists looks up movie artists.
func (s *MovieService) LookupArtists(id *string, amgArtistId *string, opt *LookupOptions) ([]*MovieArtist, *http.Response, error) {
	return lookup[MovieArtist](s.client, id, amgArtistId, nil, nil, nil, nil, nil, opt)
}

// SearchMovies searches movies.
func (s *MovieService) SearchMovies(term string, opt *SearchOptions) ([]*Movie, *http.Response, error) {
	return search[Movie](s.client, term, MovieMedia, MovieEntityMovie, opt)
}

// LookupMovies looks up movies.
func (s *MovieService) LookupMovies(id *string, amgVideoId *string, opt *LookupOptions) ([]*Movie, *http.Response, error) {
	return lookup[Movie](s.client, id, nil, nil, nil, amgVideoId, nil, nil, opt)
}
