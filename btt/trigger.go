package btt

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// GetTrigger returns the JSON representation of the trigger with the specified UUID
func (b *BTT) GetTrigger(ctx context.Context, uuid string) (string, error) {
	log.WithField("uuid", uuid).Debug("GetTrigger")
	req, err := b.newRequest(ctx, "get_trigger")
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("uuid", uuid)
	req.URL.RawQuery = q.Encode()
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug(req.URL.String())
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("btt GetTrigger error: %w", err)
	}
	defer resp.Body.Close()

	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("btt GetTrigger error: %w", err)
	}
	return buf.String(), nil
}

// ExecuteTrigger executes assigned actions for the trigger with the
// specified UUID in BTT
func (b *BTT) ExecuteTrigger(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("ExecuteTrigger")
	req, err := b.newRequest(ctx, "execute_assigned_actions_for_trigger")
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
		return fmt.Errorf("btt ExecuteTrigger error: %w", err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}
