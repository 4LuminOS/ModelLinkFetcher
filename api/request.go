package api

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func makeRequest(ctx context.Context, method string, requestURL *url.URL, headers http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, requestURL.String(), body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		req.Header = headers
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:129.0) Gecko/20100101 Firefox/129.0")

	if s := req.Header.Get("Content-Length"); s != "" {
		contentLength, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}

		req.ContentLength = contentLength
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
