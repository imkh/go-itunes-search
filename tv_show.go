package itunes

import (
	"net/http"
	"time"
)

// TVShowService handles communication with the TV show related methods
// of the iTunes Search API.
type TVShowService struct {
	client *Client
}

type TVEpisode struct {
	WrapperType            string    `json:"wrapperType"`
	Kind                   string    `json:"kind"`
	ArtistID               int       `json:"artistId"`
	CollectionID           int       `json:"collectionId"`
	TrackID                int       `json:"trackId"`
	ArtistName             string    `json:"artistName"`
	CollectionName         string    `json:"collectionName"`
	TrackName              string    `json:"trackName"`
	CollectionCensoredName string    `json:"collectionCensoredName"`
	TrackCensoredName      string    `json:"trackCensoredName"`
	ArtistViewURL          string    `json:"artistViewUrl"`
	CollectionViewURL      string    `json:"collectionViewUrl"`
	TrackViewURL           string    `json:"trackViewUrl"`
	PreviewURL             string    `json:"previewUrl"`
	ArtworkURL30           string    `json:"artworkUrl30"`
	ArtworkURL60           string    `json:"artworkUrl60"`
	ArtworkURL100          string    `json:"artworkUrl100"`
	CollectionPrice        int       `json:"collectionPrice"`
	TrackPrice             float64   `json:"trackPrice"`
	ReleaseDate            time.Time `json:"releaseDate"`
	CollectionExplicitness string    `json:"collectionExplicitness"`
	TrackExplicitness      string    `json:"trackExplicitness"`
	DiscCount              int       `json:"discCount"`
	DiscNumber             int       `json:"discNumber"`
	TrackCount             int       `json:"trackCount"`
	TrackNumber            int       `json:"trackNumber"`
	TrackTimeMillis        int       `json:"trackTimeMillis"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	ContentAdvisoryRating  *string   `json:"contentAdvisoryRating,omitempty"`
	ShortDescription       string    `json:"shortDescription"`
	LongDescription        string    `json:"longDescription"`
	CollectionHdPrice      *float64  `json:"collectionHdPrice,omitempty"`
	TrackHdPrice           *float64  `json:"trackHdPrice,omitempty"`
}

type TVSeason struct {
	WrapperType            string    `json:"wrapperType"`
	CollectionType         string    `json:"collectionType"`
	ArtistID               int       `json:"artistId"`
	CollectionID           int       `json:"collectionId"`
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
	Copyright              string    `json:"copyright"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	ReleaseDate            time.Time `json:"releaseDate"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	LongDescription        string    `json:"longDescription"`
	CollectionHdPrice      *float64  `json:"collectionHdPrice,omitempty"`
	ContentAdvisoryRating  *string   `json:"contentAdvisoryRating,omitempty"`
}

// SearchEpisode searches TV show episodes.
func (s *TVShowService) SearchEpisodes(term string, opt *SearchOptions) ([]*TVEpisode, *http.Response, error) {
	return search[TVEpisode](s.client, term, TVShowMedia, TVEpisodeEntityTVShow, opt)
}

// LookupEpisodes looks up TV show episodes.
func (s *TVShowService) LookupEpisodes(id *string, opt *LookupOptions) ([]*TVEpisode, *http.Response, error) {
	return lookup[TVEpisode](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}

// SearchTVShows searches TV show seasons.
func (s *TVShowService) SearchSeasons(term string, opt *SearchOptions) ([]*TVSeason, *http.Response, error) {
	return search[TVSeason](s.client, term, TVShowMedia, TVSeasonEntityTVShow, opt)
}

// LookupSeasons looks up TV show seasons.
func (s *TVShowService) LookupSeasons(id *string, opt *LookupOptions) ([]*TVSeason, *http.Response, error) {
	return lookup[TVSeason](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}
