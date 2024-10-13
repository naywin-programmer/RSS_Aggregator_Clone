package utils

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

func ExtractUrl(text string) (string, error) {
	pattern := regexp.MustCompile(`(?s)(https?://)([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)([a-zA-Z0-9_@=./&?\-%]+)\b([^\W]*)\b`)
	value := pattern.FindString(text)
	if value == "" {
		return value, errors.New("invalid url or url not found")
	}
	return value, nil
}

func ExtractUrls(text string) ([]string, error) {
	pattern := regexp.MustCompile(`(?m)(https?://)([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)([a-zA-Z0-9_@=./&?\-%]+)\b([^\W]*)\b`)
	values := pattern.FindAllString(text, -1)
	if values == nil {
		return values, errors.New("invalid url or url not found")
	}
	return values, nil
}

func Slug(str string, seperator rune) string {
	pattern := regexp.MustCompile(`(?s)\W+`)
	convertedUrl := pattern.ReplaceAllLiteralString(str+"-"+time.Now().UTC().Format(time.RFC3339), string(seperator))
	return strings.Trim(strings.ToLower(convertedUrl), string(seperator))
}

func RFC1123ZDateFormatVariant() string {
	return "Mon, _2 Jan 2006 15:04:05 -0700"
}
