package btt

import (
	"context"

	log "github.com/sirupsen/logrus"
)

// RefreshWidget triggers a refresh of the widget with the specified UUID
func (b *BTT) RefreshWidget(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("RefreshWidget")
	return b.simple(ctx, "refresh_widget", uuid, nil)
}
