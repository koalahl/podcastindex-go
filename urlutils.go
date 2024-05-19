package podcastindex

import (
	"fmt"
	"strings"
	"time"
)

func addMax(max int) string {
	if max != 0 {
		return fmt.Sprintf("&max=%d", max)
	}
	return ""
}

func addBefore(before int) string {
	if before != 0 {
		return fmt.Sprintf("&before=%d", before)
	}
	return ""
}

func addExclude(exclude string) string {
	if len(exclude) != 0 {
		return fmt.Sprintf("&excludeString=%s", exclude)
	}
	return ""
}

func addClean(clean bool) string {
	if clean {
		return "&clean"
	}
	return ""
}

func addTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("&since=%d", t.Unix())
}

func addFilter(name string, filter []string) string {
	if len(filter) == 0 {
		return ""
	}
	return fmt.Sprintf("&%s=%s", name, strings.Join(filter, ","))
}
