[![Go Reference](https://pkg.go.dev/badge/github.com/TheHippo/podcastindex.svg)](https://pkg.go.dev/github.com/TheHippo/podcastindex)

# podcastindex
Go library for accessing the API of [PodcastIndex](https://podcastindex.org/). You will need to have
an account for actually using the API.

When unsure how certain parameters work take a look the [API documentation](https://podcastindex-org.github.io/docs-api/) from PodcastIndex.

## Example

This code will show you the latest episodes of a podcast you just searched:


```golang
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/koalahl/podcastindex-go"
)

func main() {
    fmt.Println("ok")

    c := podcastindex.NewClient("APIKEY", "APISECRET")
    podcasts, err := c.Search("Crime in sports")
    if err != nil {
        log.Fatal(err)
    }
    podcast := podcasts[0]
    episodes, err := c.EpisodesByFeedID(podcast.ID, 0, time.Time{})
    if err != nil {
        log.Fatal(err)
    }
    for _, episode := range episodes {
        fmt.Printf("%s: %s\n", podcast.Author, episode.Title)
        fmt.Printf("%s\n", episode.Duration)
    }
}
```

### Status

There are only two things missing:

* Soundbites
* Publishing API, because I cannot test it