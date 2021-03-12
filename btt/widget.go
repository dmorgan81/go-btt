package btt

import (
	"context"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// RefreshWidget triggers a refresh of the widget with the specified UUID.
func (b *BTT) RefreshWidget(ctx context.Context, uuid string) error {
	log.WithField("uuid", uuid).Debug("RefreshWidget")
	return b.execute(ctx, "refresh_widget", map[string]string{"uuid": uuid}, ioutil.Discard)
}
