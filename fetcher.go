package umbeluzi

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

type Fetcher interface {
	Fetch() (io.Reader, error)
}

type bytesFetcher struct {
	b []byte
}

func (b bytesFetcher) Fetch() (io.Reader, error) {
	return bytes.NewReader(b.b), nil
}

func NewBytesFetcher(b []byte) bytesFetcher {
	return bytesFetcher{
		b: b,
	}
}

type fsFetcher struct {
	path string
}

func (h fsFetcher) Fetch() (io.Reader, error) {
	f, err := os.Open(h.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, f); err != nil {
		return nil, err
	}

	return buf, nil
}

func NewFSFetcher(path string) fsFetcher {
	return fsFetcher{
		path: path,
	}
}

type urlFetcher struct {
	timeout int
	url     string
	client  http.Client
}

func (h urlFetcher) Fetch() (io.Reader, error) {
	req, err := http.NewRequest(http.MethodGet, h.url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return nil, err
	}

	return buf, nil
}

func NewURLFetcher(url string, opts ...URLFetcherOption) urlFetcher {
	f := urlFetcher{
		url: url,
	}

	for _, opt := range opts {
		opt.Apply(&f)
	}

	return f
}

type URLFetcherOption interface {
	Apply(*urlFetcher)
}

type withTimeout int

func (w withTimeout) Apply(h *urlFetcher) {
	h.timeout = int(w)
}

func WithFetchTimeout(t int) URLFetcherOption {
	return withTimeout(t)
}

type withClient http.Client

func (w withClient) Apply(h *urlFetcher) {
	h.client = http.Client(w)
}

func WithFetchClient(c http.Client) URLFetcherOption {
	return withClient(c)
}
