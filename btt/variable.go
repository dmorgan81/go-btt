package btt

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// GetVariable gets the specified variable from BTT
func (b *BTT) GetVariable(ctx context.Context, name string, persistent, number bool) (string, error) {
	log.WithFields(log.Fields{
		"variable":   name,
		"persistent": persistent,
		"number":     number,
	}).Debug("GetVariable")

	p, n := "", "_string"
	if persistent {
		p = "_persistent"
	}
	if number {
		n = "_number"
	}

	req, err := b.newRequest(ctx, fmt.Sprintf("get%s%s_variable", p, n))
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("variableName", name)
	req.URL.RawQuery = q.Encode()
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug(req.URL.String())
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("btt GetVariable error: %w", err)
	}
	defer resp.Body.Close()

	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("btt GetVariable error: %w", err)
	}
	return buf.String(), nil
}

// SetVariable sets the specified variable to the specified value in BTT
func (b *BTT) SetVariable(ctx context.Context, name, value string, persistent, number bool) error {
	log.WithFields(log.Fields{
		"variable":   name,
		"value":      value,
		"persistent": persistent,
		"number":     number,
	}).Debug("SetVariable")

	p, n := "", "_string"
	if persistent {
		p = "_persistent"
	}
	if number {
		n = "_number"
	}

	req, err := b.newRequest(ctx, fmt.Sprintf("set%s%s_variable", p, n))
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("variableName", name)
	q.Add("to", value)
	req.URL.RawQuery = q.Encode()
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug(req.URL.String())
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("btt SetVariable error: %w", err)
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	return nil
}
