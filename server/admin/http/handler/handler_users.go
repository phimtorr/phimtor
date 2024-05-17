package handler

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
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
