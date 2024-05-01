package upnp

import (
	"context"

	"github.com/friendsofgo/errors"
)

func (u *UPnP) Stop(ctx context.Context) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	client, err := u.state.GetSelectedClient()
	if err != nil {
		return errors.Wrap(err, "get selected client")
	}

	if err := client.StopCtx(ctx, 0); err != nil {
		return errors.Wrap(err, "do stop")
	}

	return nil
}
