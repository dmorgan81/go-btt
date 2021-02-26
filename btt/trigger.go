package btt

import (
	"bytes"
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetTrigger returns the JSON representation of the trigger with the specified UUID
func (b *BTT) GetTrigger(ctx context.Context, uuid string) (string, error) {
	log.WithField("uuid", uuid).Debug("GetTrigger")
	var s string
	err := b.simple(ctx, "get_trigger", uuid, func(r *http.Response) error {
		buf := bytes.NewBuffer(make([]byte, 0, r.ContentLength))
		_, err := buf.ReadFrom(r.Body)
		if err == nil {
			s = buf.String()
		}
		return err
	})
	return s, err
}

// ExecuteTrigger executes assigned actions for the trigger with the specified UUID
func (b *BTT) ExecuteTrigger(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("ExecuteTrigger")
	return b.simple(ctx, "execute_assigned_actions_for_trigger", uuid, nil)
}

// DeleteTrigger deletes the trigger with the specified UUID
func (b *BTT) DeleteTrigger(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("DeleteTrigger")
	return b.simple(ctx, "delete_trigger", uuid, nil)
}
