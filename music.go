package itunes

import (
	"net/http"
	"time"
)

// MusicService handles communication with the music related methods
// of the iTunes Search API.
type MusicService struct {
	client *Client
}

type MusicArtist struct {
	WrapperType      string  `json:"wrapperType"`
	ArtistType       string  `json:"artistType"`
	ArtistName       string  `json:"artistName"`
	ArtistLinkURL    string  `json:"artistLinkUrl"`
	ArtistID         int     `json:"artistId"`
	PrimaryGenreName *string `json:"primaryGenreName,omitempty"`
	PrimaryGenreID   *int    `json:"primaryGenreId,omitempty"`
	AmgArtistID      *int    `json:"amgArtistId,omitempty"`
}

type MusicTrack struct {
	Song
	MusicVideo
}

type Album struct {
	WrapperType            string    `json:"wrapperType"`
	CollectionType         string    `json:"collectionType"`
	ArtistID               int       `json:"artistId"`
	CollectionID           int       `json:"collectionId"`
	AmgArtistID            *int      `json:"amgArtistId,omitempty"`
	ArtistName             string    `json:"artistName"`
	CollectionName         string    `json:"collectionName"`
	CollectionCensoredName string    `json:"collectionCensoredName"`
	ArtistViewURL          *string   `json:"artistViewUrl,omitempty"`
	CollectionViewURL      string    `json:"collectionViewUrl"`
	ArtworkURL60           string    `json:"artworkUrl60"`
	ArtworkURL100          string    `json:"artworkUrl100"`
	CollectionPrice        *float64  `json:"collectionPrice,omitempty"`
	CollectionExplicitness string    `json:"collectionExplicitness"`
	TrackCount             int       `json:"trackCount"`
	Copyright              string    `json:"copyright"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	ReleaseDate            time.Time `json:"releaseDate"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	ContentAdvisoryRating  *string   `json:"contentAdvisoryRating,omitempty"`
}

type MusicVideo struct {
	WrapperType             string    `json:"wrapperType"`
	Kind                    string    `json:"kind"`
	ArtistID                int       `json:"artistId"`
	CollectionID            *int      `json:"collectionId,omitempty"`
	TrackID                 int       `json:"trackId"`
	ArtistName              string    `json:"artistName"`
	CollectionName          *string   `json:"collectionName,omitempty"`
	TrackName               string    `json:"trackName"`
	CollectionCensoredName  *string   `json:"collectionCensoredName,omitempty"`
	TrackCensoredName       string    `json:"trackCensoredName"`
	ArtistViewURL           string    `json:"artistViewUrl"`
	CollectionViewURL       *string   `json:"collectionViewUrl,omitempty"`
	TrackViewURL            string    `json:"trackViewUrl"`
	PreviewURL              *string   `json:"previewUrl,omitempty"`
	ArtworkURL30            string    `json:"artworkUrl30"`
	ArtworkURL60            string    `json:"artworkUrl60"`
	ArtworkURL100           string    `json:"artworkUrl100"`
	CollectionPrice         *int      `json:"collectionPrice,omitempty"`
	TrackPrice              *int      `json:"trackPrice,omitempty"`
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
	ContentAdvisoryRating   *string   `json:"contentAdvisoryRating,omitempty"`
	CollectionArtistID      *int      `json:"collectionArtistId,omitempty"`
	CollectionArtistViewURL *string   `json:"collectionArtistViewUrl,omitempty"`
}

type Mix struct {
}

type Song struct {
	WrapperType             string    `json:"wrapperType"`
	Kind                    string    `json:"kind"`
	ArtistID                int       `json:"artistId"`
	CollectionID            int       `json:"collectionId"`
	TrackID                 int       `json:"trackId"`
	ArtistName              string    `json:"artistName"`
	CollectionName          string    `json:"collectionName"`
	TrackName               string    `json:"trackName"`
	CollectionCensoredName  string    `json:"collectionCensoredName"`
	TrackCensoredName       string    `json:"trackCensoredName"`
	CollectionArtistID      *int      `json:"collectionArtistId,omitempty"`
	CollectionArtistName    *string   `json:"collectionArtistName,omitempty"`
	ArtistViewURL           string    `json:"artistViewUrl"`
	CollectionViewURL       string    `json:"collectionViewUrl"`
	TrackViewURL            string    `json:"trackViewUrl"`
	PreviewURL              string    `json:"previewUrl"`
	ArtworkURL30            string    `json:"artworkUrl30"`
	ArtworkURL60            string    `json:"artworkUrl60"`
	ArtworkURL100           string    `json:"artworkUrl100"`
	CollectionPrice         *float64  `json:"collectionPrice,omitempty"`
	TrackPrice              *float64  `json:"trackPrice,omitempty"`
	ReleaseDate             time.Time `json:"releaseDate"`
	CollectionExplicitness  string    `json:"collectionExplicitness"`
	TrackExplicitness       string    `json:"trackExplicitness"`
	DiscCount               int       `json:"discCount"`
	DiscNumber              int       `json:"discNumber"`
	TrackCount              int       `json:"trackCount"`
	TrackNumber             int       `json:"trackNumber"`
	TrackTimeMillis         int       `json:"trackTimeMillis"`
	Country                 string    `json:"country"`
	Currency                string    `json:"currency"`
	PrimaryGenreName        string    `json:"primaryGenreName"`
	IsStreamable            bool      `json:"isStreamable"`
	ContentAdvisoryRating   *string   `json:"contentAdvisoryRating,omitempty"`
	CollectionArtistViewURL *string   `json:"collectionArtistViewUrl,omitempty"`
}

// SearchArtists searches music artists.
func (s *MusicService) SearchArtists(term string, opt *SearchOptions) ([]*MusicArtist, *http.Response, error) {
	return search[MusicArtist](s.client, term, MusicMedia, MusicArtistEntityMusic, opt)
}

// LookupArtists looks up music artists.
func (s *MusicService) LookupArtists(id *string, amgArtistId *string, opt *LookupOptions) ([]*MusicArtist, *http.Response, error) {
	return lookup[MusicArtist](s.client, id, amgArtistId, nil, nil, nil, nil, nil, opt)
}

// SearchTracks searches music tracks.
func (s *MusicService) SearchTracks(term string, opt *SearchOptions) ([]*MusicTrack, *http.Response, error) {
	return search[MusicTrack](s.client, term, MusicMedia, MusicTrackEntityMusic, opt)
}

// LookupTracks looks up music tracks.
func (s *MusicService) LookupTracks(id *string, opt *LookupOptions) ([]*MusicTrack, *http.Response, error) {
	return lookup[MusicTrack](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}

// SearchAlbums searches albums.
func (s *MusicService) SearchAlbums(term string, opt *SearchOptions) ([]*Album, *http.Response, error) {
	return search[Album](s.client, term, MusicMedia, AlbumEntityMusic, opt)
}

// LookupAlbums looks up albums.
func (s *MusicService) LookupAlbums(id *string, upc *string, amgAlbumId *string, opt *LookupOptions) ([]*Album, *http.Response, error) {
	return lookup[Album](s.client, id, nil, upc, amgAlbumId, nil, nil, nil, opt)
}

// SearchArtists searches music videos.
func (s *MusicService) SearchVideos(term string, opt *SearchOptions) ([]*MusicVideo, *http.Response, error) {
	return search[MusicVideo](s.client, term, MusicMedia, MusicVideoEntityMusic, opt)
}

// LookupVideos looks up music videos.
func (s *MusicService) LookupVideos(id *string, upc *string, amgVideoId *string, opt *LookupOptions) ([]*MusicVideo, *http.Response, error) {
	return lookup[MusicVideo](s.client, id, nil, upc, nil, amgVideoId, nil, nil, opt)
}

// SearchMixes searches mixes.
func (s *MusicService) SearchMixes(term string, opt *SearchOptions) ([]*Mix, *http.Response, error) {
	return search[Mix](s.client, term, MusicMedia, MixEntityMusic, opt)
}

// LookupMixes looks up mixes.
func (s *MusicService) LookupMixes(id *string, opt *LookupOptions) ([]*Mix, *http.Response, error) {
	return lookup[Mix](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}

// SearchSongs searches songs.
func (s *MusicService) SearchSongs(term string, opt *SearchOptions) ([]*Song, *http.Response, error) {
	return search[Song](s.client, term, MusicMedia, SongEntityMusic, opt)
}

// LookupSongs looks up songs.
func (s *MusicService) LookupSongs(id *string, opt *LookupOptions) ([]*Song, *http.Response, error) {
	return lookup[Song](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}
