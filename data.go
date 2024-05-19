package podcastindex

type PodcastArrayResult struct {
	Status      string     `json:"status"`
	Feeds       []*Podcast `json:"feeds"`
	Count       int        `json:"count"`
	Query       string     `json:"query"`
	Description string     `json:"description"`
}

type PodcastResult struct {
	Status string `json:"status"`
	Query  struct {
		URL string `json:"url"`
	} `json:"query"`
	Feed        Podcast `json:"feed"`
	Description string  `json:"description"`
}

// Podcast contains all informations about a podcast returned from the podcastindex API
type Podcast struct {
	ID                     uint            `json:"id"`
	Title                  string          `json:"title"`
	URL                    string          `json:"url"`
	OriginalURL            string          `json:"originalUrl"`
	Link                   string          `json:"link"`
	Description            string          `json:"description"`
	Author                 string          `json:"author"`
	OwnerName              string          `json:"ownerName"`
	Image                  string          `json:"image"`
	Artwork                string          `json:"artwork"`
	LastUpdateTime         Time            `json:"lastUpdateTime"`
	LastCrawlTime          Time            `json:"lastCrawlTime"`
	LastParseTime          Time            `json:"lastParseTime"`
	LastGoodHTTPStatusTime Time            `json:"lastGoodHttpStatusTime"`
	LastHTTPStatus         int             `json:"lastHttpStatus"`
	ContentType            string          `json:"contentType"`
	ItunesID               int             `json:"itunesId"`
	Generator              string          `json:"generator"`
	Language               string          `json:"language"`
	Type                   int             `json:"type"`
	Dead                   int             `json:"dead"`
	CrawlErrors            int             `json:"crawlErrors"`
	ParseErrors            int             `json:"parseErrors"`
	Categories             map[uint]string `json:"categories"`
}

type EpisodeArrayResponse struct {
	Status      string     `json:"status"`
	Items       []*Episode `json:"items"`
	Count       int        `json:"count"`
	Query       string     `json:"query"`
	Description string     `json:"description"`
}

type RandomEpisodesResponse struct {
	Status      string     `json:"status"`
	Items       []*Episode `json:"episodes"`
	Count       int        `json:"count"`
	Query       string     `json:"query"`
	Description string     `json:"description"`
}

type EpisodeResponse struct {
	Status      string   `json:"status"`
	ID          string   `json:"id"`
	Episode     *Episode `json:"episode"`
	Description string   `json:"description"`
}

// Episode contains all information about a single podcast episode returned from
// the podcastindex API
type Episode struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	Link            string   `json:"link"`
	Description     string   `json:"description"`
	GUID            string   `json:"guid"`
	DatePublished   Time     `json:"datePublished"`
	DateCrawled     Time     `json:"dateCrawled"`
	EnclosureURL    string   `json:"enclosureUrl"`
	EnclosureType   string   `json:"enclosureType"`
	EnclosureLength int      `json:"enclosureLength"`
	Duration        Duration `json:"duration"`
	Explicit        int      `json:"explicit"`
	Episode         int      `json:"episode"`
	EpisodeType     string   `json:"episodeType"`
	Season          int      `json:"season"`
	Image           string   `json:"image"`
	FeedItunesID    int      `json:"feedItunesId"`
	FeedImage       string   `json:"feedImage"`
	FeedID          int      `json:"feedId"`
	FeedLanguage    string   `json:"feedLanguage"`
	ChaptersURL     string   `json:"chaptersUrl"`
	TranscriptURL   string   `json:"transcriptUrl"`
}

type RecentPodcastsResponse struct {
	Status      string           `json:"status"`
	Feeds       []*RecentPodcast `json:"feeds"`
	Count       int              `json:"count"`
	Max         interface{}      `json:"max"`
	Since       interface{}      `json:"since"`
	Description string           `json:"description"`
}

// RecentPodcast contains all information about an recently
// updated podcast
type RecentPodcast struct {
	ID                    int    `json:"id"`
	URL                   string `json:"url"`
	Title                 string `json:"title"`
	NewestItemPublishTime Time   `json:"newestItemPublishTime"`
	Description           string `json:"description"`
	Image                 string `json:"image"`
	ItunesID              int    `json:"itunesId"`
	Language              string `json:"language"`
}

type NewPodcastResponse struct {
	Status      string        `json:"status"`
	Feeds       []*NewPodcast `json:"feeds"`
	Count       int           `json:"count"`
	Max         string        `json:"max"`
	Description string        `json:"description"`
}

// NewPodcast contains data for a newly added podcast
type NewPodcast struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	TimeAdded   Time   `json:"timeAdded"`
	Status      string `json:"status"`
	ContentHash string `json:"contentHash"`
	Language    string `json:"language"`
}
