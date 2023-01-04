package itunes

// CountryValue represents the two-letter country code for the store you want to search.
// The search uses the default store front for the specified country. For example: US. The default is US.
type CountryValue string

// List of available countries.
//
// Country.io's JSON file that maps ISO2 country codes to country names: http://country.io/names.json
const (
	Afghanistan                            CountryValue = "AF"
	AlandIslands                           CountryValue = "AX"
	Albania                                CountryValue = "AL"
	Algeria                                CountryValue = "DZ"
	AmericanSamoa                          CountryValue = "AS"
	Andorra                                CountryValue = "AD"
	Angola                                 CountryValue = "AO"
	Anguilla                               CountryValue = "AI"
	Antarctica                             CountryValue = "AQ"
	AntiguaandBarbuda                      CountryValue = "AG"
	Argentina                              CountryValue = "AR"
	Armenia                                CountryValue = "AM"
	Aruba                                  CountryValue = "AW"
	Australia                              CountryValue = "AU"
	Austria                                CountryValue = "AT"
	Azerbaijan                             CountryValue = "AZ"
	Bahamas                                CountryValue = "BS"
	Bahrain                                CountryValue = "BH"
	Bangladesh                             CountryValue = "BD"
	Barbados                               CountryValue = "BB"
	Belarus                                CountryValue = "BY"
	Belgium                                CountryValue = "BE"
	Belize                                 CountryValue = "BZ"
	Benin                                  CountryValue = "BJ"
	Bermuda                                CountryValue = "BM"
	Bhutan                                 CountryValue = "BT"
	Bolivia                                CountryValue = "BO"
	BonaireSaintEustatiusandSaba           CountryValue = "BQ"
	BosniaandHerzegovina                   CountryValue = "BA"
	Botswana                               CountryValue = "BW"
	BouvetIsland                           CountryValue = "BV"
	Brazil                                 CountryValue = "BR"
	BritishIndianOceanTerritory            CountryValue = "IO"
	BritishVirginIslands                   CountryValue = "VG"
	Brunei                                 CountryValue = "BN"
	Bulgaria                               CountryValue = "BG"
	BurkinaFaso                            CountryValue = "BF"
	Burundi                                CountryValue = "BI"
	Cambodia                               CountryValue = "KH"
	Cameroon                               CountryValue = "CM"
	Canada                                 CountryValue = "CA"
	CapeVerde                              CountryValue = "CV"
	CaymanIslands                          CountryValue = "KY"
	CentralAfricanRepublic                 CountryValue = "CF"
	Chad                                   CountryValue = "TD"
	Chile                                  CountryValue = "CL"
	China                                  CountryValue = "CN"
	ChristmasIsland                        CountryValue = "CX"
	CocosIslands                           CountryValue = "CC"
	Colombia                               CountryValue = "CO"
	Comoros                                CountryValue = "KM"
	CookIslands                            CountryValue = "CK"
	CostaRica                              CountryValue = "CR"
	Croatia                                CountryValue = "HR"
	Cuba                                   CountryValue = "CU"
	Curacao                                CountryValue = "CW"
	Cyprus                                 CountryValue = "CY"
	CzechRepublic                          CountryValue = "CZ"
	DemocraticRepublicoftheCongo           CountryValue = "CD"
	Denmark                                CountryValue = "DK"
	Djibouti                               CountryValue = "DJ"
	Dominica                               CountryValue = "DM"
	DominicanRepublic                      CountryValue = "DO"
	EastTimor                              CountryValue = "TL"
	Ecuador                                CountryValue = "EC"
	Egypt                                  CountryValue = "EG"
	ElSalvador                             CountryValue = "SV"
	EquatorialGuinea                       CountryValue = "GQ"
	Eritrea                                CountryValue = "ER"
	Estonia                                CountryValue = "EE"
	Ethiopia                               CountryValue = "ET"
	FalklandIslands                        CountryValue = "FK"
	FaroeIslands                           CountryValue = "FO"
	Fiji                                   CountryValue = "FJ"
	Finland                                CountryValue = "FI"
	France                                 CountryValue = "FR"
	FrenchGuiana                           CountryValue = "GF"
	FrenchPolynesia                        CountryValue = "PF"
	FrenchSouthernTerritories              CountryValue = "TF"
	Gabon                                  CountryValue = "GA"
	Gambia                                 CountryValue = "GM"
	Georgia                                CountryValue = "GE"
	Germany                                CountryValue = "DE"
	Ghana                                  CountryValue = "GH"
	Gibraltar                              CountryValue = "GI"
	Greece                                 CountryValue = "GR"
	Greenland                              CountryValue = "GL"
	Grenada                                CountryValue = "GD"
	Guadeloupe                             CountryValue = "GP"
	Guam                                   CountryValue = "GU"
	Guatemala                              CountryValue = "GT"
	Guernsey                               CountryValue = "GG"
	Guinea                                 CountryValue = "GN"
	GuineaBissau                           CountryValue = "GW"
	Guyana                                 CountryValue = "GY"
	Haiti                                  CountryValue = "HT"
	HeardIslandandMcDonaldIslands          CountryValue = "HM"
	Honduras                               CountryValue = "HN"
	HongKong                               CountryValue = "HK"
	Hungary                                CountryValue = "HU"
	Iceland                                CountryValue = "IS"
	India                                  CountryValue = "IN"
	Indonesia                              CountryValue = "ID"
	Iran                                   CountryValue = "IR"
	Iraq                                   CountryValue = "IQ"
	Ireland                                CountryValue = "IE"
	IsleofMan                              CountryValue = "IM"
	Israel                                 CountryValue = "IL"
	Italy                                  CountryValue = "IT"
	IvoryCoast                             CountryValue = "CI"
	Jamaica                                CountryValue = "JM"
	Japan                                  CountryValue = "JP"
	Jersey                                 CountryValue = "JE"
	Jordan                                 CountryValue = "JO"
	Kazakhstan                             CountryValue = "KZ"
	Kenya                                  CountryValue = "KE"
	Kiribati                               CountryValue = "KI"
	Kosovo                                 CountryValue = "XK"
	Kuwait                                 CountryValue = "KW"
	Kyrgyzstan                             CountryValue = "KG"
	Laos                                   CountryValue = "LA"
	Latvia                                 CountryValue = "LV"
	Lebanon                                CountryValue = "LB"
	Lesotho                                CountryValue = "LS"
	Liberia                                CountryValue = "LR"
	Libya                                  CountryValue = "LY"
	Liechtenstein                          CountryValue = "LI"
	Lithuania                              CountryValue = "LT"
	Luxembourg                             CountryValue = "LU"
	Macao                                  CountryValue = "MO"
	Macedonia                              CountryValue = "MK"
	Madagascar                             CountryValue = "MG"
	Malawi                                 CountryValue = "MW"
	Malaysia                               CountryValue = "MY"
	Maldives                               CountryValue = "MV"
	Mali                                   CountryValue = "ML"
	Malta                                  CountryValue = "MT"
	MarshallIslands                        CountryValue = "MH"
	Martinique                             CountryValue = "MQ"
	Mauritania                             CountryValue = "MR"
	Mauritius                              CountryValue = "MU"
	Mayotte                                CountryValue = "YT"
	Mexico                                 CountryValue = "MX"
	Micronesia                             CountryValue = "FM"
	Moldova                                CountryValue = "MD"
	Monaco                                 CountryValue = "MC"
	Mongolia                               CountryValue = "MN"
	Montenegro                             CountryValue = "ME"
	Montserrat                             CountryValue = "MS"
	Morocco                                CountryValue = "MA"
	Mozambique                             CountryValue = "MZ"
	Myanmar                                CountryValue = "MM"
	Namibia                                CountryValue = "NA"
	Nauru                                  CountryValue = "NR"
	Nepal                                  CountryValue = "NP"
	Netherlands                            CountryValue = "NL"
	NewCaledonia                           CountryValue = "NC"
	NewZealand                             CountryValue = "NZ"
	Nicaragua                              CountryValue = "NI"
	Niger                                  CountryValue = "NE"
	Nigeria                                CountryValue = "NG"
	Niue                                   CountryValue = "NU"
	NorfolkIsland                          CountryValue = "NF"
	NorthKorea                             CountryValue = "KP"
	NorthernMarianaIslands                 CountryValue = "MP"
	Norway                                 CountryValue = "NO"
	Oman                                   CountryValue = "OM"
	Pakistan                               CountryValue = "PK"
	Palau                                  CountryValue = "PW"
	PalestinianTerritory                   CountryValue = "PS"
	Panama                                 CountryValue = "PA"
	PapuaNewGuinea                         CountryValue = "PG"
	Paraguay                               CountryValue = "PY"
	Peru                                   CountryValue = "PE"
	Philippines                            CountryValue = "PH"
	Pitcairn                               CountryValue = "PN"
	Poland                                 CountryValue = "PL"
	Portugal                               CountryValue = "PT"
	PuertoRico                             CountryValue = "PR"
	Qatar                                  CountryValue = "QA"
	RepublicoftheCongo                     CountryValue = "CG"
	Reunion                                CountryValue = "RE"
	Romania                                CountryValue = "RO"
	Russia                                 CountryValue = "RU"
	Rwanda                                 CountryValue = "RW"
	SaintBarthelemy                        CountryValue = "BL"
	SaintHelena                            CountryValue = "SH"
	SaintKittsandNevis                     CountryValue = "KN"
	SaintLucia                             CountryValue = "LC"
	SaintMartin                            CountryValue = "MF"
	SaintPierreandMiquelon                 CountryValue = "PM"
	SaintVincentandtheGrenadines           CountryValue = "VC"
	Samoa                                  CountryValue = "WS"
	SanMarino                              CountryValue = "SM"
	SaoTomeandPrincipe                     CountryValue = "ST"
	SaudiArabia                            CountryValue = "SA"
	Senegal                                CountryValue = "SN"
	Serbia                                 CountryValue = "RS"
	Seychelles                             CountryValue = "SC"
	SierraLeone                            CountryValue = "SL"
	Singapore                              CountryValue = "SG"
	SintMaarten                            CountryValue = "SX"
	Slovakia                               CountryValue = "SK"
	Slovenia                               CountryValue = "SI"
	SolomonIslands                         CountryValue = "SB"
	Somalia                                CountryValue = "SO"
	SouthAfrica                            CountryValue = "ZA"
	SouthGeorgiaandtheSouthSandwichIslands CountryValue = "GS"
	SouthKorea                             CountryValue = "KR"
	SouthSudan                             CountryValue = "SS"
	Spain                                  CountryValue = "ES"
	SriLanka                               CountryValue = "LK"
	Sudan                                  CountryValue = "SD"
	Suriname                               CountryValue = "SR"
	SvalbardandJanMayen                    CountryValue = "SJ"
	Swaziland                              CountryValue = "SZ"
	Sweden                                 CountryValue = "SE"
	Switzerland                            CountryValue = "CH"
	Syria                                  CountryValue = "SY"
	Taiwan                                 CountryValue = "TW"
	Tajikistan                             CountryValue = "TJ"
	Tanzania                               CountryValue = "TZ"
	Thailand                               CountryValue = "TH"
	Togo                                   CountryValue = "TG"
	Tokelau                                CountryValue = "TK"
	Tonga                                  CountryValue = "TO"
	TrinidadandTobago                      CountryValue = "TT"
	Tunisia                                CountryValue = "TN"
	Turkey                                 CountryValue = "TR"
	Turkmenistan                           CountryValue = "TM"
	TurksandCaicosIslands                  CountryValue = "TC"
	Tuvalu                                 CountryValue = "TV"
	USVirginIslands                        CountryValue = "VI"
	Uganda                                 CountryValue = "UG"
	Ukraine                                CountryValue = "UA"
	UnitedArabEmirates                     CountryValue = "AE"
	UnitedKingdom                          CountryValue = "GB"
	UnitedStates                           CountryValue = "US"
	UnitedStatesMinorOutlyingIslands       CountryValue = "UM"
	Uruguay                                CountryValue = "UY"
	Uzbekistan                             CountryValue = "UZ"
	Vanuatu                                CountryValue = "VU"
	Vatican                                CountryValue = "VA"
	Venezuela                              CountryValue = "VE"
	Vietnam                                CountryValue = "VN"
	WallisandFutuna                        CountryValue = "WF"
	WesternSahara                          CountryValue = "EH"
	Yemen                                  CountryValue = "YE"
	Zambia                                 CountryValue = "ZM"
	Zimbabwe                               CountryValue = "ZW"
)

// Country is a helper routine that allocates a new CountryValue
// to store v and returns a pointer to it.
func Country(v CountryValue) *CountryValue {
	p := new(CountryValue)
	*p = v
	return p
}

// MediaValue represents the media type you want to search for. For example: movie. The default is all.
//
// iTunes Search API docs: https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/Searching.html#//apple_ref/doc/uid/TP40017632-CH5-SW2
type MediaValue string

// List of available medias.
const (
	MovieMedia   MediaValue = "movie"
	PodcastMedia MediaValue = "podcast"
	MusicMedia   MediaValue = "music"
	// MusicVideoMedia MediaValue = "musicVideo"
	AudiobookMedia MediaValue = "audiobook"
	ShortFilmMedia MediaValue = "shortFilm"
	TVShowMedia    MediaValue = "tvShow"
	SoftwareMedia  MediaValue = "software"
	EbookMedia     MediaValue = "ebook"
)

// Media is a helper routine that allocates a new MediaValue
// to store v and returns a pointer to it.
func Media(v MediaValue) *MediaValue {
	p := new(MediaValue)
	*p = v
	return p
}

// EntityValue represents the type of results you want returned, relative to the specified media type.
// For example: movieArtist for a movie media type search. The default is the track entity associated with the specified media type.
//
// iTunes Search API docs: https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/Searching.html#//apple_ref/doc/uid/TP40017632-CH5-SW2
type EntityValue string

// List of available entities.
const (
	// Movie media type
	MovieArtistEntityMovie EntityValue = "movieArtist"
	MovieEntityMovie       EntityValue = "movie"

	// Podcast media type
	PodcastAuthorEntityPodcast EntityValue = "podcastAuthor"
	PodcastEntityPodcast       EntityValue = "podcast"

	// Music media type
	MusicArtistEntityMusic EntityValue = "musicArtist"
	MusicTrackEntityMusic  EntityValue = "musicTrack"
	AlbumEntityMusic       EntityValue = "album"
	MusicVideoEntityMusic  EntityValue = "musicVideo"
	MixEntityMusic         EntityValue = "mix"
	SongEntityMusic        EntityValue = "song"

	// // MusicVideo media type
	// MusicArtistEntityMusicVideo EntityValue = "musicArtist"
	// MusicVideoEntityMusicVideo  EntityValue = "musicVideo"

	// Audiobook media type
	AudiobookAuthorEntityAudiobook EntityValue = "audiobookAuthor"
	AudiobookEntityAudiobook       EntityValue = "audiobook"

	// ShortFilm media type
	ShortFilmArtistEntityShortFilm EntityValue = "shortFilmArtist"
	ShortFilmEntityShortFilm       EntityValue = "shortFilm"

	// TVShow media type
	TVEpisodeEntityTVShow EntityValue = "tvEpisode"
	TVSeasonEntityTVShow  EntityValue = "tvSeason"

	// Software media type
	SoftwareDeveloperEntitySoftware EntityValue = "softwareDeveloper"
	SoftwareEntitySoftware          EntityValue = "software"
	IPadSoftwareEntitySoftware      EntityValue = "iPadSoftware"
	MacSoftwareEntitySoftware       EntityValue = "macSoftware"
	TVSoftwareEntitySoftware        EntityValue = "tvSoftware"

	// Ebook media type
	EbookAuthorEntityEbook EntityValue = "ebookAuthor"
	EbookEntityEbook       EntityValue = "ebook"
)

// Entity is a helper routine that allocates a new EntityValue
// to store v and returns a pointer to it.
func Entity(v EntityValue) *EntityValue {
	p := new(EntityValue)
	*p = v
	return p
}

// AttributeValue represents the attribute you want to search for in the stores, relative to the specified media type.
// For example, if you want to search for an artist by name specify `entity=allArtist&attribute=allArtistTerm`.
// In this example, if you search for `term=maroon`, iTunes returns “Maroon 5” in the search results, instead of all artists who have ever recorded a song with the word “maroon” in the title.
// The default is all attributes associated with the specified media type.
//
// iTunes Search API docs: https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/Searching.html#//apple_ref/doc/uid/TP40017632-CH5-SW3
type AttributeValue string

const (
	// Movie media type
	ActorTermAttributeMovie       AttributeValue = "actorTerm"
	GenreIndexAttributeMovie      AttributeValue = "genreIndex"
	ArtistTermAttributeMovie      AttributeValue = "artistTerm"
	ShortFilmTermAttributeMovie   AttributeValue = "shortFilmTerm"
	ProducerTermAttributeMovie    AttributeValue = "producerTerm"
	RatingTermAttributeMovie      AttributeValue = "ratingTerm"
	DirectorTermAttributeMovie    AttributeValue = "directorTerm"
	ReleaseYearTermAttributeMovie AttributeValue = "releaseYearTerm"
	FeatureFilmTermAttributeMovie AttributeValue = "featureFilmTerm"
	MovieArtistTermAttributeMovie AttributeValue = "movieArtistTerm"
	MovieTermAttributeMovie       AttributeValue = "movieTerm"
	RatingIndexAttributeMovie     AttributeValue = "ratingIndex"
	DescriptionTermAttributeMovie AttributeValue = "descriptionTerm"

	// Podcast media type
	TitleTermAttributePodcast       AttributeValue = "titleTerm"
	LanguageTermAttributePodcast    AttributeValue = "languageTerm"
	AuthorTermAttributePodcast      AttributeValue = "authorTerm"
	GenreIndexAttributePodcast      AttributeValue = "genreIndex"
	ArtistTermAttributePodcast      AttributeValue = "artistTerm"
	RatingIndexAttributePodcast     AttributeValue = "ratingIndex"
	KeywordsTermAttributePodcast    AttributeValue = "keywordsTerm"
	DescriptionTermAttributePodcast AttributeValue = "descriptionTerm"

	// Music media type
	MixTermMusicAttribute      AttributeValue = "mixTerm"
	GenreIndexMusicAttribute   AttributeValue = "genreIndex"
	ArtistTermMusicAttribute   AttributeValue = "artistTerm"
	ComposerTermMusicAttribute AttributeValue = "composerTerm"
	AlbumTermMusicAttribute    AttributeValue = "albumTerm"
	RatingIndexMusicAttribute  AttributeValue = "ratingIndex"
	SongTermMusicAttribute     AttributeValue = "songTerm"

	// // MusicVideo media type
	// GenreIndexMusicVideoAttribute  AttributeValue = "genreIndex"
	// ArtistTermMusicVideoAttribute  AttributeValue = "artistTerm"
	// AlbumTermMusicVideoAttribute   AttributeValue = "albumTerm"
	// RatingIndexMusicVideoAttribute AttributeValue = "ratingIndex"
	// SongTermMusicVideoAttribute    AttributeValue = "songTerm"

	// Audiobook media type
	TitleTermAudiobookAttribute   AttributeValue = "titleTerm"
	AuthorTermAudiobookAttribute  AttributeValue = "authorTerm"
	GenreIndexAudiobookAttribute  AttributeValue = "genreIndex"
	RatingIndexAudiobookAttribute AttributeValue = "ratingIndex"

	// ShortFilm media type
	GenreIndexShortFilmAttribute      AttributeValue = "genreIndex"
	ArtistTermShortFilmAttribute      AttributeValue = "artistTerm"
	ShortFilmTermShortFilmAttribute   AttributeValue = "shortFilmTerm"
	RatingIndexShortFilmAttribute     AttributeValue = "ratingIndex"
	DescriptionTermShortFilmAttribute AttributeValue = "descriptionTerm"

	// Software media type
	SoftwareDeveloperSoftwareAttribute AttributeValue = "softwareDeveloper"

	// TVShow media type
	GenreIndexTVShowAttribute      AttributeValue = "genreIndex"
	TVEpisodeTermTVShowAttribute   AttributeValue = "tvEpisodeTerm"
	ShowTermTVShowAttribute        AttributeValue = "showTerm"
	TVSeasonTermTVShowAttribute    AttributeValue = "tvSeasonTerm"
	RatingIndexTVShowAttribute     AttributeValue = "ratingIndex"
	DescriptionTermTVShowAttribute AttributeValue = "descriptionTerm"
)

// Attribute is a helper routine that allocates a new AttributeValue
// to store v and returns a pointer to it.
func Attribute(v AttributeValue) *AttributeValue {
	p := new(AttributeValue)
	*p = v
	return p
}

// LanguageValue represents the language, English or Japanese, you want to use when returning search results.
// Specify the language using the five-letter codename. For example: en_us. The default is en_us (English).
type LanguageValue string

// List of available languages.
//
// Commonly used IETF language tags: https://gist.github.com/traysr/2001377
const (
	English             LanguageValue = "en"
	EnglishUnitedStates LanguageValue = "en_us"
	EnglishGreatBritain LanguageValue = "en_gb"
	French              LanguageValue = "fr"
	German              LanguageValue = "de"
	Polish              LanguageValue = "pl"
	Dutch               LanguageValue = "nl"
	Finnish             LanguageValue = "fi"
	Swedish             LanguageValue = "sv"
	Italian             LanguageValue = "it"
	SpanishSpain        LanguageValue = "es"
	PortuguesePortugal  LanguageValue = "pt"
	Russian             LanguageValue = "ru"
	PortugueseBrazil    LanguageValue = "pt_br"
	SpanishMexico       LanguageValue = "es_mx"
	ChinesePRC          LanguageValue = "zh_cn"
	ChineseTaiwan       LanguageValue = "zh_tw"
	Japanese            LanguageValue = "ja"
	Korean              LanguageValue = "ko"
)

// Language is a helper routine that allocates a new LanguageValue
// to store v and returns a pointer to it.
func Language(v LanguageValue) *LanguageValue {
	p := new(LanguageValue)
	*p = v
	return p
}

// ExplicitValue represents a flag indicating whether or not you want to include explicit content in your search results. The default is Yes.
type ExplicitValue string

// List of available explicit values.
const (
	YesExplicit ExplicitValue = "Yes"
	NoExplicit  ExplicitValue = "No"
)

// Explicit is a helper routine that allocates a new ExplicitValue
// to store v and returns a pointer to it.
func Explicit(v ExplicitValue) *ExplicitValue {
	p := new(ExplicitValue)
	*p = v
	return p
}

// SortValue represents the sort
type SortValue string

// List of available sort values.
const (
	PopularSort SortValue = "popular"
	RecentSort  SortValue = "recent"
)

// Sort is a helper routine that allocates a new SortValue
// to store v and returns a pointer to it.
func Sort(v SortValue) *SortValue {
	p := new(SortValue)
	*p = v
	return p
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}
