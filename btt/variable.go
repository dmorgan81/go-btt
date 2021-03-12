package btt

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// GetVariable gets the specified variable from BTT.
func (b *BTT) GetVariable(ctx context.Context, name string, number bool) (string, error) {
	log.WithFields(log.Fields{
		"variable": name,
		"number":   number,
	}).Debug("GetVariable")

	n := "_string"
	if number {
		n = "_number"
	}

	action := fmt.Sprintf("get%s_variable", n)
	buf := &bytes.Buffer{}
	if err := b.execute(ctx, action, map[string]string{"variableName": name}, buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// GetStringVariable is shorthand for GetVariable(ctx, name, false).
func (b *BTT) GetStringVariable(ctx context.Context, name string) (string, error) {
	return b.GetVariable(ctx, name, false)
}

// GetNumberVariable is shorthand for GetVariable(ctx, name, true).
func (b *BTT) GetNumberVariable(ctx context.Context, name string) (float64, error) {
	s, err := b.GetVariable(ctx, name, true)
	if err != nil {
		return -1, err
	}
	return strconv.ParseFloat(s, 64)
}

// SetVariable sets the specified variable to the specified value in BTT.
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

	action := fmt.Sprintf("set%s%s_variable", p, n)
	return b.execute(ctx, action, map[string]string{"variableName": name, "to": value}, ioutil.Discard)
}

// SetStringVariable is shorthand for SetVariable(ctx, name, value, false, false).
func (b *BTT) SetStringVariable(ctx context.Context, name, value string) error {
	return b.SetVariable(ctx, name, value, false, false)
}

// SetNumberVariable is shorthand for SetVariable(ctx, name, value, false, true).
func (b *BTT) SetNumberVariable(ctx context.Context, name string, value float64) error {
	return b.SetVariable(ctx, name, fmt.Sprintf("%f", value), false, true)
}

// SetPersistentStringVariable is shorthand for SetVariable(ctx, name, value, true, false).
func (b *BTT) SetPersistentStringVariable(ctx context.Context, name, value string) error {
	return b.SetVariable(ctx, name, value, true, false)
}

// SetPersistentNumberVariable is shorthand for SetVariable(ctx, name, value, true, true).
func (b *BTT) SetPersistentNumberVariable(ctx context.Context, name string, value float64) error {
	return b.SetVariable(ctx, name, fmt.Sprintf("%f", value), true, true)
}
