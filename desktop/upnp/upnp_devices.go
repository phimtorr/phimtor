package upnp

import (
	"context"

	"github.com/huin/goupnp/dcps/av1"
	"github.com/rs/zerolog/log"
)

func (u *UPnP) Scan(ctx context.Context) error {
	clients, errs, err := av1.NewAVTransport1ClientsCtx(ctx)
	if err != nil {
		return err
	}
	if len(errs) > 0 {
		log.Ctx(ctx).Error().Errs("errs", errs).Msg("Failed to scan devices")
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	u.state.AvailableClients = clients
	return nil
}

func (u *UPnP) GetAvailableClients() []*av1.AVTransport1 {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.state.AvailableClients
}

func (u *UPnP) SelectDevice(ctx context.Context, udn string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.state.SelectedUDN = udn
	return nil
}

func (u *UPnP) GetSelectedDeviceUDN() string {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.state.SelectedUDN
}
