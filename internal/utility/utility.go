package utility

import (
	"net/url"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	log "github.com/sirupsen/logrus"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func FuzzyFind(list []string) string {
	idx, err := fuzzyfinder.Find(
		list,
		func(i int) string {
			clean := strings.Trim(list[i], "\"")
			return clean
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(list[idx], "\"")
}
