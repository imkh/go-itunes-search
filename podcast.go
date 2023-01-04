package itunes

import (
	"net/http"
	"time"
)

// PodcastService handles communication with the podcast related methods
// of the iTunes Search API.
type PodcastService struct {
	client *Client
}

type PodcastAuthor struct {
	WrapperType      string `json:"wrapperType"`
	ArtistType       string `json:"artistType"`
	ArtistName       string `json:"artistName"`
	ArtistLinkURL    string `json:"artistLinkUrl"`
	ArtistID         int    `json:"artistId"`
	PrimaryGenreName string `json:"primaryGenreName"`
	PrimaryGenreID   int    `json:"primaryGenreId"`
}

type Podcast struct {
	WrapperType            string    `json:"wrapperType"`
	Kind                   string    `json:"kind"`
	ArtistID               *int      `json:"artistId,omitempty"`
	CollectionID           int       `json:"collectionId"`
	TrackID                int       `json:"trackId"`
	ArtistName             string    `json:"artistName"`
	CollectionName         string    `json:"collectionName"`
	TrackName              string    `json:"trackName"`
	CollectionCensoredName string    `json:"collectionCensoredName"`
	TrackCensoredName      string    `json:"trackCensoredName"`
	ArtistViewURL          *string   `json:"artistViewUrl,omitempty"`
	CollectionViewURL      string    `json:"collectionViewUrl"`
	FeedURL                *string   `json:"feedUrl,omitempty"`
	TrackViewURL           string    `json:"trackViewUrl"`
	ArtworkURL30           string    `json:"artworkUrl30"`
	ArtworkURL60           string    `json:"artworkUrl60"`
	ArtworkURL100          string    `json:"artworkUrl100"`
	CollectionPrice        float64   `json:"collectionPrice"`
	TrackPrice             float64   `json:"trackPrice"`
	CollectionHdPrice      int       `json:"collectionHdPrice"`
	ReleaseDate            time.Time `json:"releaseDate"`
	CollectionExplicitness string    `json:"collectionExplicitness"`
	TrackExplicitness      string    `json:"trackExplicitness"`
	TrackCount             int       `json:"trackCount"`
	TrackTimeMillis        *int      `json:"trackTimeMillis,omitempty"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	ContentAdvisoryRating  *string   `json:"contentAdvisoryRating,omitempty"`
	ArtworkURL600          string    `json:"artworkUrl600"`
	GenreIds               []string  `json:"genreIds"`
	Genres                 []string  `json:"genres"`
}

// SearchAuthors searches podcast authors.
func (s *PodcastService) SearchAuthors(term string, opt *SearchOptions) ([]*PodcastAuthor, *http.Response, error) {
	return search[PodcastAuthor](s.client, term, PodcastMedia, PodcastAuthorEntityPodcast, opt)
}

// LookupAuthors looks up podcast authors.
func (s *PodcastService) LookupAuthors(id *string, opt *LookupOptions) ([]*PodcastAuthor, *http.Response, error) {
	return lookup[PodcastAuthor](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}

// SearchPodcasts searches podcasts.
func (s *PodcastService) SearchPodcasts(term string, opt *SearchOptions) ([]*Podcast, *http.Response, error) {
	return search[Podcast](s.client, term, PodcastMedia, PodcastEntityPodcast, opt)
}

// LookupPodcasts looks up podcasts.
func (s *PodcastService) LookupPodcasts(id *string, opt *LookupOptions) ([]*Podcast, *http.Response, error) {
	return lookup[Podcast](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}
