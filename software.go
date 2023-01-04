package itunes

import (
	"net/http"
	"time"
)

// SoftwareService handles communication with the software related methods
// of the iTunes Search API.
type SoftwareService struct {
	client *Client
}

type SoftwareDeveloper struct {
	WrapperType   string `json:"wrapperType"`
	ArtistType    string `json:"artistType"`
	ArtistName    string `json:"artistName"`
	ArtistLinkURL string `json:"artistLinkUrl"`
	ArtistID      int    `json:"artistId"`
}

type Software struct {
	Features                           []string  `json:"features"`
	SupportedDevices                   []string  `json:"supportedDevices"`
	Advisories                         []string  `json:"advisories"`
	IsGameCenterEnabled                *bool     `json:"isGameCenterEnabled,omitempty"`
	ScreenshotUrls                     []string  `json:"screenshotUrls"`
	IpadScreenshotUrls                 []string  `json:"ipadScreenshotUrls"`
	AppletvScreenshotUrls              []string  `json:"appletvScreenshotUrls"`
	ArtworkURL60                       string    `json:"artworkUrl60"`
	ArtworkURL512                      string    `json:"artworkUrl512"`
	ArtworkURL100                      string    `json:"artworkUrl100"`
	ArtistViewURL                      *string   `json:"artistViewUrl,omitempty"`
	Kind                               string    `json:"kind"`
	ArtistID                           int       `json:"artistId"`
	ArtistName                         string    `json:"artistName"`
	Genres                             []string  `json:"genres"`
	Price                              *float64  `json:"price,omitempty"`
	Description                        string    `json:"description"`
	IsVppDeviceBasedLicensingEnabled   bool      `json:"isVppDeviceBasedLicensingEnabled"`
	BundleID                           string    `json:"bundleId"`
	PrimaryGenreName                   string    `json:"primaryGenreName"`
	PrimaryGenreID                     int       `json:"primaryGenreId"`
	TrackID                            int       `json:"trackId"`
	TrackName                          string    `json:"trackName"`
	SellerName                         string    `json:"sellerName"`
	CurrentVersionReleaseDate          time.Time `json:"currentVersionReleaseDate"`
	ReleaseNotes                       *string   `json:"releaseNotes,omitempty"`
	MinimumOsVersion                   string    `json:"minimumOsVersion"`
	TrackCensoredName                  string    `json:"trackCensoredName"`
	LanguageCodesISO2A                 []string  `json:"languageCodesISO2A"`
	FileSizeBytes                      *string   `json:"fileSizeBytes,omitempty"`
	SellerURL                          *string   `json:"sellerUrl,omitempty"`
	FormattedPrice                     *string   `json:"formattedPrice,omitempty"`
	ContentAdvisoryRating              string    `json:"contentAdvisoryRating"`
	AverageUserRatingForCurrentVersion float64   `json:"averageUserRatingForCurrentVersion"`
	UserRatingCountForCurrentVersion   int       `json:"userRatingCountForCurrentVersion"`
	AverageUserRating                  float64   `json:"averageUserRating"`
	TrackViewURL                       string    `json:"trackViewUrl"`
	TrackContentRating                 string    `json:"trackContentRating"`
	Version                            string    `json:"version"`
	WrapperType                        string    `json:"wrapperType"`
	Currency                           string    `json:"currency"`
	GenreIds                           []string  `json:"genreIds"`
	ReleaseDate                        time.Time `json:"releaseDate"`
	UserRatingCount                    int       `json:"userRatingCount"`
}

// SearchDeveloper searches software developers.
func (s *SoftwareService) SearchDeveloper(term string, opt *SearchOptions) ([]*SoftwareDeveloper, *http.Response, error) {
	return search[SoftwareDeveloper](s.client, term, SoftwareMedia, SoftwareDeveloperEntitySoftware, opt)
}

// LookupDeveloper looks up software developers.
func (s *SoftwareService) LookupDeveloper(id *string, opt *LookupOptions) ([]*SoftwareDeveloper, *http.Response, error) {
	return lookup[SoftwareDeveloper](s.client, id, nil, nil, nil, nil, nil, nil, opt)
}

// SearchSoftware searches software.
func (s *SoftwareService) SearchSoftware(term string, opt *SearchOptions) ([]*Software, *http.Response, error) {
	return search[Software](s.client, term, SoftwareMedia, SoftwareEntitySoftware, opt)
}

// LookupSoftware looks up software.
func (s *SoftwareService) LookupSoftware(id *string, bundleId *string, opt *LookupOptions) ([]*Software, *http.Response, error) {
	return lookup[Software](s.client, id, nil, nil, nil, nil, nil, bundleId, opt)
}

// SearchiPad searches iPad software.
func (s *SoftwareService) SearchiPad(term string, opt *SearchOptions) ([]*Software, *http.Response, error) {
	return search[Software](s.client, term, SoftwareMedia, IPadSoftwareEntitySoftware, opt)
}

// SearchMac searches Mac software.
func (s *SoftwareService) SearchMac(term string, opt *SearchOptions) ([]*Software, *http.Response, error) {
	return search[Software](s.client, term, SoftwareMedia, MacSoftwareEntitySoftware, opt)
}

// SearchTV searches TV software.
func (s *SoftwareService) SearchTV(term string, opt *SearchOptions) ([]*Software, *http.Response, error) {
	return search[Software](s.client, term, SoftwareMedia, TVSoftwareEntitySoftware, opt)
}
