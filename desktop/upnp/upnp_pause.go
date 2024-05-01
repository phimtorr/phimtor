package upnp

import (
	"context"

	"github.com/friendsofgo/errors"
)

func (u *UPnP) Pause(ctx context.Context) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	client, err := u.state.GetSelectedClient()
	if err != nil {
		return errors.Wrap(err, "get selected client")
	}

	if err := client.PauseCtx(ctx, 0); err != nil {
		return errors.Wrap(err, "do pause")
	}

	return nil
}
