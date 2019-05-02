package vk_bmstu_schedule_status

import (
	"strconv"
	"strings"
)

type Title struct {
	Prefix   string
	WeekName string
	WeekNum  int
	WeekAbbr string
	replacer strings.Replacer
}

func (title Title) Format(pattern string) string {
	replacer := strings.NewReplacer(
		"p", title.Prefix,
		"s", title.WeekName,
		"d", strconv.Itoa(title.WeekNum),
		"a", title.WeekAbbr)
	return replacer.Replace(pattern)
}
