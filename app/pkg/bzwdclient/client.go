package bzwdclient

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type TranslateLevel int

const (
	LV1 TranslateLevel = 1
	LV2 TranslateLevel = 2
)

func (tl TranslateLevel) String() string {
	return strconv.Itoa(int(tl))
}

func Translate(text string, level TranslateLevel) (string, error) {
	resp, err := post(text, level)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := getResult(resp.Body)
	if err != nil {
		return "", nil
	}

	return result, nil
}

func post(text string, level TranslateLevel) (*http.Response, error) {
	vals := url.Values{}
	vals.Add("before", text)
	vals.Add("level", level.String())
	vals.Add("transbtn", "翻訳")

	req, err := http.NewRequest(
		"POST",
		"https://bizwd.net/",
		strings.NewReader(vals.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	return client.Do(req)
}

func getResult(body io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return "", err
	}
	return doc.Find("textarea[name=\"after\"]").Text(), nil
}
