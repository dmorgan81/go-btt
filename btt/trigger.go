package btt

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// GetTrigger returns the JSON representation of the trigger with the specified UUID.
func (b *BTT) GetTrigger(ctx context.Context, uuid string, w io.Writer) error {
	log.WithField("uuid", uuid).Debug("GetTrigger")
	return b.execute(ctx, "get_trigger", map[string]string{"uuid": uuid}, w)
}

// ExecuteTrigger executes assigned actions for the trigger with the specified UUID.
func (b *BTT) ExecuteTrigger(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("ExecuteTrigger")
	return b.execute(ctx, "execute_assigned_actions_for_trigger", map[string]string{"uuid": uuid}, ioutil.Discard)
}

// DeleteTrigger deletes the trigger with the specified UUID.
func (b *BTT) DeleteTrigger(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("DeleteTrigger")
	return b.execute(ctx, "delete_trigger", map[string]string{"uuid": uuid}, ioutil.Discard)
}

// Trigger the named trigger with the specified name.
func (b *BTT) Trigger(ctx context.Context, name string, async bool) error {
	log.WithField("name", name).Debug("Trigger")
	action := "trigger_named"
	if async {
		action = "trigger_named_async_without_response"
	}
	return b.execute(ctx, action, map[string]string{"trigger_name": name}, ioutil.Discard)
}

// UpdateTrigger updates the trigger with the specified UUID with JSON content from the specified reader.
func (b *BTT) UpdateTrigger(ctx context.Context, uuid string, r io.Reader) error {
	log.WithField("uuid", uuid).Debug("UpdateTrigger")
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return b.execute(ctx, "update_trigger", map[string]string{"uuid": uuid, "json": string(d)}, ioutil.Discard)
}

// AddTrigger adds a trigger with JSON content from the specified reader.
// Returns the UUID of the new trigger.
func (b *BTT) AddTrigger(ctx context.Context, r io.Reader) (string, error) {
	log.Debug("AddTrigger")
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := b.execute(ctx, "add_new_trigger", map[string]string{"json": string(d)}, buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
