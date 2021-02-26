package btt

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// BTT represents a connection to a BTT webserver
type BTT struct {
	addr, secret string
	client       *http.Client
}

// New returns a BTT pointed at addr
func New(addr string) *BTT {
	return &BTT{addr: addr, client: http.DefaultClient}
}

// WithSecret returns a BTT that uses the supplied secret
func (b *BTT) WithSecret(s string) *BTT {
	b.secret = s
	return b
}

// WithHTTPClient returns a BTT that uses the supplied http.Client
func (b *BTT) WithHTTPClient(c *http.Client) *BTT {
	b.client = c
	return b
}

func (b *BTT) newRequest(ctx context.Context, action string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s/", b.addr, action), nil)
	if err != nil {
		return nil, fmt.Errorf("btt error: %w", err)
	}

	if b.secret != "" {
		q := req.URL.Query()
		q.Add("shared_secret", b.secret)
		req.URL.RawQuery = q.Encode()
	}
	return req, nil
}

func (b *BTT) simple(ctx context.Context, action, uuid string, f func(r *http.Response) error) error {
	log.WithFields(log.Fields{
		"action": action,
		"uuid":   uuid,
	}).Debug("simple")

	req, err := b.newRequest(ctx, action)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("uuid", uuid)
	req.URL.RawQuery = q.Encode()
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug(req.URL.String())
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("btt error: %w", err)
	}
	defer resp.Body.Close()

	if f == nil {
		_, err := io.Copy(ioutil.Discard, resp.Body)
		return err
	}

	if err := f(resp); err != nil {
		return fmt.Errorf("btt error: %w", err)
	}
	return nil
}
