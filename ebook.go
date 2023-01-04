package itunes

import (
	"net/http"
	"time"
)

// EbookService handles communication with the ebook related methods
// of the iTunes Search API.
type EbookService struct {
	client *Client
}

type EbookAuthor struct {
	WrapperType      string `json:"wrapperType"`
	ArtistType       string `json:"artistType"`
	ArtistName       string `json:"artistName"`
	ArtistLinkURL    string `json:"artistLinkUrl"`
	ArtistID         int    `json:"artistId"`
	PrimaryGenreName string `json:"primaryGenreName"`
	PrimaryGenreID   int    `json:"primaryGenreId"`
	AmgArtistID      *int   `json:"amgArtistId,omitempty"`
}

type Ebook struct {
	Currency          string    `json:"currency"`
	ArtistIds         []int     `json:"artistIds"`
	ArtistID          int       `json:"artistId"`
	ArtistName        string    `json:"artistName"`
	Genres            []string  `json:"genres"`
	Price             float64   `json:"price"`
	Description       string    `json:"description"`
	GenreIds          []string  `json:"genreIds"`
	ReleaseDate       time.Time `json:"releaseDate"`
	TrackID           int       `json:"trackId"`
	TrackName         string    `json:"trackName"`
	ArtworkURL60      string    `json:"artworkUrl60"`
	ArtworkURL100     string    `json:"artworkUrl100"`
	ArtistViewURL     string    `json:"artistViewUrl"`
	TrackCensoredName string    `json:"trackCensoredName"`
	FileSizeBytes     int       `json:"fileSizeBytes"`
	FormattedPrice    string    `json:"formattedPrice"`
	TrackViewURL      string    `json:"trackViewUrl"`
	Kind              string    `json:"kind"`
	UserRatingCount   *int      `json:"userRatingCount,omitempty"`
	AverageUserRating *float64  `json:"averageUserRating,omitempty"`
}

// SearchAuthors searches ebook authors.
func (s *EbookService) SearchAuthors(term string, opt *SearchOptions) ([]*EbookAuthor, *http.Response, error) {
	return search[EbookAuthor](s.client, term, EbookMedia, EbookAuthorEntityEbook, opt)
}

// LookupAuthors looks up ebook authors.
func (s *EbookService) LookupAuthors(id *string, amgArtistId *string, opt *LookupOptions) ([]*EbookAuthor, *http.Response, error) {
	return lookup[EbookAuthor](s.client, id, amgArtistId, nil, nil, nil, nil, nil, opt)
}

// SearchEbook searches ebooks.
func (s *EbookService) SearchEbooks(term string, opt *SearchOptions) ([]*Ebook, *http.Response, error) {
	return search[Ebook](s.client, term, EbookMedia, EbookEntityEbook, opt)
}

// LookupEbooks looks up ebooks.
func (s *EbookService) LookupEbooks(id *string, isbn *string, opt *LookupOptions) ([]*Ebook, *http.Response, error) {
	return lookup[Ebook](s.client, id, nil, nil, nil, nil, isbn, nil, opt)
}
