package handler

import (
	"fmt"
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/go-chi/chi/v5"
	"google.golang.org/api/iterator"

	"github.com/phimtorr/phimtor/server/admin/http/ui"
)

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	nextPageToken := r.URL.Query().Get("nextPageToken")
	pager := iterator.NewPager(h.authClient.Users(ctx, ""), pageSize, nextPageToken)
	var users []*auth.ExportedUserRecord
	nextPageToken, err := pager.NextPage(&users)
	if err != nil {
		return err
	}

	return ui.ListUsers(users, nextPageToken).Render(ctx, w)
}

func (h *Handler) ViewUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	uid := chi.URLParam(r, "uid")
	user, err := h.authClient.GetUser(ctx, uid)
	if err != nil {
		return err
	}

	return ui.ViewUser(user).Render(ctx, w)
}

func (h *Handler) UpdatePremium(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	uid := chi.URLParam(r, "uid")

	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("parse form: %w", err)
	}

	premiumUntilStr := r.Form.Get("premium_until")
	premiumUntil, err := time.Parse("2006-01-02T15:04", premiumUntilStr)

	if err != nil {
		return fmt.Errorf("parse premium until: %w", err)
	}

	user, err := h.authClient.GetUser(ctx, uid)
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	var claims map[string]interface{}
	if user.CustomClaims != nil {
		claims = user.CustomClaims
	} else {
		claims = make(map[string]interface{})
	}

	claims["premium_until"] = premiumUntil.Unix()

	if err := h.authClient.SetCustomUserClaims(ctx, uid, claims); err != nil {
		return fmt.Errorf("set custom user claims: %w", err)
	}

	http.Redirect(w, r, "/users/"+uid, http.StatusFound)
	return nil
}
