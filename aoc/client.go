package aoc

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Client interface {
	GetInput(year int, day int, part int) (Input, error)
	SendSolution(year int, day int, part int, solution string) error
}

const URL = "https://adventofcode.com"

type DefaultClient struct {
	client http.Client
}

func NewDefaultClient(session string) (DefaultClient, error) {
	sessionCookie := http.Cookie{
		Name:     "session",
		Value:    session,
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	}
	url, err := url.Parse(URL)
	if err != nil {
		return DefaultClient{}, err
	}
	cookieJar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return DefaultClient{}, err
	}
	cookieJar.SetCookies(url, []*http.Cookie{&sessionCookie})

	return DefaultClient{
		client: http.Client{
			Jar:     cookieJar,
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (c DefaultClient) GetInput(year int, day int, part int) (Input, error) {
	url := fmt.Sprintf("%s/%d/day/%d/input", URL, year, day)
	resp, err := c.client.Get(url)
	if err != nil {
		return Input{}, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Input{}, err
	}

	return NewInput(string(bytes)), nil
}

func (c DefaultClient) SendSolution(year int, day int, part int, solution string) error {
	return nil
}
