package itunes

import (
	"net/http"
)

// ShortFilmService handles communication with the short film related methods
// of the iTunes Search API.
type ShortFilmService struct {
	client *Client
}

type ShortFilmArtist struct {
	MovieArtist
}

type ShortFilm struct {
	Movie
}

// SearchArtists searches short film artists.
func (s *ShortFilmService) SearchArtists(term string, opt *SearchOptions) ([]*ShortFilmArtist, *http.Response, error) {
	return search[ShortFilmArtist](s.client, term, ShortFilmMedia, ShortFilmArtistEntityShortFilm, opt)
}

// LookupArtists looks up short film artists.
func (s *ShortFilmService) LookupArtists(id *string, amgArtistId *string, opt *LookupOptions) ([]*ShortFilmArtist, *http.Response, error) {
	return lookup[ShortFilmArtist](s.client, id, amgArtistId, nil, nil, nil, nil, nil, opt)
}

// SearchShortFilms searches short films.
func (s *ShortFilmService) SearchShortFilms(term string, opt *SearchOptions) ([]*ShortFilm, *http.Response, error) {
	return search[ShortFilm](s.client, term, ShortFilmMedia, ShortFilmEntityShortFilm, opt)
}

// LookupShortFilms looks up short films.
func (s *ShortFilmService) LookupShortFilms(id *string, amgVideoId *string, opt *LookupOptions) ([]*ShortFilm, *http.Response, error) {
	return lookup[ShortFilm](s.client, id, nil, nil, nil, amgVideoId, nil, nil, opt)
}
