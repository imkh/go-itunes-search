package itunes

import (
	"net/http"
	"time"
)

// AudiobookService handles communication with the audiobook related methods
// of the iTunes Search API.
type AudiobookService struct {
	client *Client
}

type AudiobookAuthor struct {
	EbookAuthor
}

type Audiobook struct {
	WrapperType            string    `json:"wrapperType"`
	ArtistID               int       `json:"artistId"`
	CollectionID           int       `json:"collectionId"`
	AmgArtistID            *int      `json:"amgArtistId,omitempty"`
	ArtistName             string    `json:"artistName"`
	CollectionName         string    `json:"collectionName"`
	CollectionCensoredName string    `json:"collectionCensoredName"`
	ArtistViewURL          string    `json:"artistViewUrl"`
	CollectionViewURL      string    `json:"collectionViewUrl"`
	ArtworkURL60           string    `json:"artworkUrl60"`
	ArtworkURL100          string    `json:"artworkUrl100"`
	CollectionPrice        float64   `json:"collectionPrice"`
	CollectionExplicitness string    `json:"collectionExplicitness"`
	TrackCount             int       `json:"trackCount"`
	Copyright              *string   `json:"copyright,omitempty"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	ReleaseDate            time.Time `json:"releaseDate"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	PreviewURL             string    `json:"previewUrl"`
	Description            string    `json:"description"`
}

// SearchAuthors searches audiobook authors.
func (s *AudiobookService) SearchAuthors(term string, opt *SearchOptions) ([]*AudiobookAuthor, *http.Response, error) {
	return search[AudiobookAuthor](s.client, term, AudiobookMedia, AudiobookAuthorEntityAudiobook, opt)
}

// LookupAuthors looks up audiobook authors.
func (s *AudiobookService) LookupAuthors(id *string, amgArtistId *string, opt *LookupOptions) ([]*AudiobookAuthor, *http.Response, error) {
	return lookup[AudiobookAuthor](s.client, id, amgArtistId, nil, nil, nil, nil, nil, opt)
}

// SearchAudiobooks searches audiobooks.
func (s *AudiobookService) SearchAudiobooks(term string, opt *SearchOptions) ([]*Audiobook, *http.Response, error) {
	return search[Audiobook](s.client, term, AudiobookMedia, AudiobookEntityAudiobook, opt)
}

// LookupAudiobooks looks up audiobooks.
func (s *AudiobookService) LookupAudiobooks(id *string, opt *LookupOptions) ([]*Audiobook, *http.Response, error) {
	return lookup[Audiobook](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}
