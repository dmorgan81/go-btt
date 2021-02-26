package btt

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

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

// GetStringVariable is shorthand for GetVariable(ctx, name, false, false)
func (b *BTT) GetStringVariable(ctx context.Context, name string) (string, error) {
	return b.GetVariable(ctx, name, false, false)
}

// GetNumberVariable is shorthand for GetVariable(ctx, name, false, true)
func (b *BTT) GetNumberVariable(ctx context.Context, name string) (float64, error) {
	s, err := b.GetVariable(ctx, name, false, true)
	if err != nil {
		return -1, err
	}
	return strconv.ParseFloat(s, 64)
}

// GetPersistentStringVariable is shorthand for GetVariable(ctx, name, true, false)
func (b *BTT) GetPersistentStringVariable(ctx context.Context, name string) (string, error) {
	return b.GetVariable(ctx, name, true, false)
}

// GetPersistentNumberVariable is shorthand for GetVariable(ctx, name, true, true)
func (b *BTT) GetPersistentNumberVariable(ctx context.Context, name string) (float64, error) {
	s, err := b.GetVariable(ctx, name, true, true)
	if err != nil {
		return -1, err
	}
	return strconv.ParseFloat(s, 64)
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
	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}

// SetStringVariable is shorthand for SetVariable(ctx, name, value, false, false)
func (b *BTT) SetStringVariable(ctx context.Context, name, value string) error {
	return b.SetVariable(ctx, name, value, false, false)
}

// SetNumberVariable is shorthand for SetVariable(ctx, name, value, false, true)
func (b *BTT) SetNumberVariable(ctx context.Context, name string, value float64) error {
	return b.SetVariable(ctx, name, fmt.Sprintf("%f", value), false, true)
}

// SetPersistentStringVariable is shorthand for SetVariable(ctx, name, value, true, false)
func (b *BTT) SetPersistentStringVariable(ctx context.Context, name, value string) error {
	return b.SetVariable(ctx, name, value, true, false)
}

// SetPersistentNumberVariable is shorthand for SetVariable(ctx, name, value, true, true)
func (b *BTT) SetPersistentNumberVariable(ctx context.Context, name string, value float64) error {
	return b.SetVariable(ctx, name, fmt.Sprintf("%f", value), true, true)
}
