package handler

import (
	"fmt"
	"net/http"

	"github.com/friendsofgo/errors"

	"github.com/ncruces/zenity"

	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/server/uri"
	"github.com/phimtorr/phimtor/desktop/setting"
)

func (h *Handler) GetSettings(w http.ResponseWriter, r *http.Request) error {
	settings := h.settingsStorage.GetSettings()
	return ui.Settings(settings).Render(r.Context(), w)
}

func (h *Handler) UpdateSetting(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.Wrap(err, "parse form")
	}

	err := h.settingsStorage.UpdateSettings(func(s setting.Settings) (setting.Settings, error) {
		if r.Form.Has("locale") {
			if err := s.SetLocale(r.Form.Get("locale")); err != nil {
				return s, fmt.Errorf("set locale: %w", err)
			}
		}
		if r.Form.Has("deleteAfterClosed") {
			s.SetDeleteAfterClosed(r.Form.Get("deleteAfterClosed") == "on")
		}
		return s, nil
	})
	if err != nil {
		return errors.Wrap(err, "update setting")
	}

	fullyRedirect(w, r, uri.GetSettings())
	return nil
}

func (h *Handler) ChangeDataDir(w http.ResponseWriter, r *http.Request) error {
	err := h.settingsStorage.UpdateSettings(func(s setting.Settings) (setting.Settings, error) {
		newDataDir, err := zenity.SelectFile(
			zenity.Directory(),
			zenity.Title("Select data directory"),
			zenity.Filename(s.GetCurrentDataDir()),
		)
		if err != nil {
			return s, fmt.Errorf("select data directory: %w", err)
		}

		if err := s.SetDataDir(newDataDir); err != nil {
			return s, fmt.Errorf("set data directory: %w", err)
		}

		return s, nil
	})
	if errors.Is(err, zenity.ErrCanceled) {
		fullyRedirect(w, r, uri.GetSettings())
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "change data dir")
	}

	fullyRedirect(w, r, uri.GetSettings())
	return nil
}
