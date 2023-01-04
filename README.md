# go-itunes-search

A Go library for accessing the [iTunes Search API](https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/index.html). Inspired by [go-gitlab](https://github.com/xanzy/go-gitlab).

## Coverage

### Movie

- [x] Search and look up movie artists
- [x] Search and look up movies

### Podcast

- [x] Search and look up podcast authors
- [x] Search and look up podcasts

### Music

- [x] Search and look up music artists
- [x] Search and look up music tracks
- [x] Search and look up albums
- [x] Search and look up music videos
- [x] Search and look up mixes
- [x] Search and look up songs

### Audiobook

- [x] Search and look up audiobook authors
- [x] Search and look up audiobooks

### Short Film

- [x] Search and look up short film artists
- [x] Search and look up short films

### TV Show

- [x] Search and look up TV show episodes
- [x] Search and look up TV show seasons

### Software

- [x] Search and look up software developers
- [x] Search and look up software
- [x] Search and look up iPad software
- [x] Search and look up Mac software
- [x] Search and look up TV software

### Ebook

- [x] Search and look up ebook authors
- [x] Search and look up ebooks

## Installation

Inside your project directory, run:

```console
$ go get github.com/imkh/go-itunes-search
```

or import the module and run `go get` without parameters.

```go
import "github.com/imkh/go-itunes-search"
```

## Usage

```go
package main

import (
	"github.com/imkh/go-itunes-search"
)

func main() {
	client, err := itunes.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Search movies
	term := "spider-man"
	searchOpts := &itunes.SearchOptions{
		Country: itunes.Country(itunes.UnitedStates),
	}
	movies, _, err := client.Movie.SearchMovies(term, searchOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Search movies %q (%d results)\n", term, len(movies))
	for _, movie := range movies {
		fmt.Printf("  %s (%d)\n", movie.TrackName, movie.ReleaseDate.Year())
	}

	fmt.Println()

	// Look up albums
	id := "1548835870"
	lookupOpts := &itunes.LookupOptions{
		Country: itunes.Country(itunes.Japan),
		Lang:    itunes.Language(itunes.English),
	}
	albums, _, err := client.Music.LookupAlbums(&id, nil, nil, lookupOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Look up albums %q (%d results)\n", id, len(albums))
	for _, album := range albums {
		fmt.Printf("  %s / %s (%d)\n", album.CollectionName, album.ArtistName, album.ReleaseDate.Year())
	}
}
```