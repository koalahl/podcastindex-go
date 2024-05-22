package podcastindex

import (
	"errors"
	"fmt"
	"time"
)

// SearchPodcasts for podcasts, authors or owners
func (c *Client) SearchPodcasts(term string) ([]*Podcast, error) {
	return c.SearchPodcastsC(term, false, 0)
}

// SearchPodcastsC for searching with more options than Search
//
// - clean for non explicit feeds according to itunes:explicit
//
// - fullBody to return the more then 100 characters in the descriptions
//
// - max for the number of results, when set to 0 it uses the API default
func (c *Client) SearchPodcastsC(term string, clean bool, max int) ([]*Podcast, error) {
	url := fmt.Sprintf("search/byterm?q=\"%s\"&fulltext%s%s", term, addClean(clean), addMax(max))
	result := &PodcastArrayResult{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not find a podcast for that term")
	}
	return result.Feeds, err
}

/*This call returns all of the episodes where the specified person is mentioned.

It searches the following fields:

- Person tags
- Episode title
- Episode description
- Feed owner
- Feed author
*/
func (c *Client) SearchEpisodes(term string) ([]*Episode, error) {
	url := fmt.Sprintf("search/byperson?q=\"%s\"&fulltext", term)	
	return c.getEpisodes(url, errors.New("Could not find a episode for that term"))
}

func (c *Client) getPodcast(url string, notFound error) (*Podcast, error) {
	result := &PodcastResult{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, notFound
	}
	return &result.Feed, err
}

// PodcastByFeedURL returns general information about a podcast by its
// feed URL
func (c *Client) PodcastByFeedURL(url string) (*Podcast, error) {
	u := fmt.Sprintf("podcasts/byfeedurl?url=%s&fulltext", url)
	return c.getPodcast(u, errors.New("Could not find a podcast for that feed URL"))
}

// PodcastByFeedID returns general information about a podcast by its id
func (c *Client) PodcastByFeedID(id string) (*Podcast, error) {
	url := fmt.Sprintf("podcasts/byfeedid?id=%s&fulltext", id)
	return c.getPodcast(url, errors.New("Could not find a podcast for that id"))
}

// PodcastByITunesID returns general information about a podcast by its
// ITune id
func (c *Client) PodcastByITunesID(id string) (*Podcast, error) {
	url := fmt.Sprintf("podcasts/byitunesid?id=%s&fulltext", id)
	return c.getPodcast(url, errors.New("Could not find a podcast for that iTunes id"))
}

func (c *Client) getEpisodes(url string, notFound error) ([]*Episode, error) {
	result := &EpisodeArrayResponse{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, notFound
	}
	return result.Items, nil
}

// EpisodesByFeedID returns all episodes for a podcast by its id
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned
//
// - since = only return episodes since that time. Set time to zero to not filter
// by time
func (c *Client) EpisodesByFeedID(id string, max int, since time.Time) ([]*Episode, error) {
	url := fmt.Sprintf("episodes/byfeedid?id=%s&fulltext%s%s", id, addMax(max), addTime(since))
	return c.getEpisodes(url, errors.New("Could not get episodes by feed id"))
}

// EpisodesByFeedURL returns episodes for a podcast by its feed URL
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned
//
// - since = only return episodes since that time. Set time to zero to not filter
// by time
func (c *Client) EpisodesByFeedURL(feedURL string, max int, since time.Time) ([]*Episode, error) {
	url := fmt.Sprintf("episodes/byfeedurl?url=\"%s\"&fulltext%s%s", feedURL, addMax(max), addTime(since))
	return c.getEpisodes(url, errors.New("Could not get episodes by feed URL"))
}

// EpisodesByITunesID returns episodes for a podcast by its iTunes id
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned
//
// - since = only return episodes since that time. Set time to zero to not filter
// by time
func (c *Client) EpisodesByITunesID(id string, max int, since time.Time) ([]*Episode, error) {
	url := fmt.Sprintf("episodes/byitunesid?id=%s&fulltext%s%s", id, addMax(max), addTime(since))
	return c.getEpisodes(url, errors.New("Could not get episodes by iTunes id"))
}

// EpisodeByID return a single episode by its id
func (c *Client) EpisodeByID(id string) (*Episode, error) {
	url := fmt.Sprintf("episodes/byid?id=%s&fulltext", id)
	result := &EpisodeResponse{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not find episode")
	}
	return result.Episode, nil
}

// RandomEpisodes returns a random episode
//
// categories and notCategories can be combined
//
// - languages = the languages the episodes should be in. "unknown" for when the language is not known.
// Leave empty if languages does not matter
//
// - categories = name of the category or categories the episodes should be in.
// Leave empty if categories do not matter
//
// - notCategories = name of the category or categories the episodes should not be in.
// Leave empty if categories do not matter
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned, the default is 1
func (c *Client) RandomEpisodes(languages, categories, notCategories []string, max int) ([]*Episode, error) {
	url := fmt.Sprintf("episodes/random?fulltext%s%s%s%s", addMax(max), addFilter("lang", languages), addFilter("cat", categories), addFilter("notcat", notCategories))
	result := &RandomEpisodesResponse{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not get random episodes")
	}
	return result.Items, nil
}

// RecentEpisodes returns the last episodes across the entire database
//
// - before = only return episodes that are older than the episode with this id. set to zero
// to ignore
//
// - excludeString = exclude episodes with this string in title or url. Leave empty for no
// filter
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned, the default is 10
func (c *Client) RecentEpisodes(before int, max int, exclude string) ([]*Episode, error) {
	url := fmt.Sprintf("recent/episodes?fulltext%s%s%s", addMax(max), addExclude(exclude), addBefore(before))
	return c.getEpisodes(url, errors.New("Could not get recent episodes"))
}

// RecentPodcasts returns the last updated podcasts
//
// - languages = the languages the podcast should be in. "unknown" for when the language is not known.
// Leave empty if languages does not matter
//
// - categories = name of the category or categories the podcast should be in.
// Leave empty if categories do not matter
//
// - notCategories = name of the category or categories the podcast should not be in.
// Leave empty if categories do not matter
//
// - max = number of podcasts to return, if max is 0 the default number of episodes will be
// returned, the default is 40
//
// - since = only return episodes since that time. Set time to zero to not filter
// by time
func (c *Client) RecentPodcasts(languages, categories, notCategories []string, max int, since time.Time) ([]*RecentPodcast, error) {
	url := fmt.Sprintf("recent/feeds?fulltext%s%s%s%s%s",
		addMax(max), addFilter("lang", languages), addFilter("cat", categories),
		addFilter("notcat", notCategories), addTime(since),
	)
	result := &RecentPodcastsResponse{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not find the recently updated podcasts")
	}
	return result.Feeds, err
}

// NewPodcasts return up to 1000 podcasts that have been added to the database over the last week
func (c *Client) NewPodcasts() ([]*NewPodcast, error) {
	url := fmt.Sprintf("recent/newfeeds")
	result := &NewPodcastResponse{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not find the newest podcasts")
	}
	return result.Feeds, err
}

func (c *Client) Categories() ( []*Category,error) {
	url := fmt.Sprintf("categories/list")
	result := &CategoryArrayResponse{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not find the newest podcasts")
	}

	return result.Feeds, err
}

// PodcastsTrending returns the top max podcasts by their popularity
func (c *Client) PodcastsTrending(languages, categories, notCategories []string, max int, since time.Time) ([]*Podcast, error) {
	url := fmt.Sprintf("podcasts/trending?fulltext%s%s%s%s%s",
	addMax(max), addFilter("lang", languages), addFilter("cat", categories),
	addFilter("notcat", notCategories), addTime(since))

	result := &PodcastsTrendingResponse{}
	err := c.request(url, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not find the trending podcasts")
	}
	return result.Feeds, err
}

func (c *Client) AddByFeedURL(feedURL string) error {
	url := fmt.Sprintf("add/byfeedurl?url=%s", feedURL)

	result := &AddByFeedURLResponse{}
	err := c.request(url, result)
	if err != nil {
		return err
	}
	if result.Status == "false" {
		return errors.New("Could not add podcast by feed URL")
	}

	return nil
}
