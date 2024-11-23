package jobs

import (
	"context"
	"errors"
	"time"

	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/iterator"
)

func RunSetPremiumForNewUsers(ctx context.Context, authClient *firebaseAuth.Client) {
	timer := time.NewTicker(1 * time.Minute)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			ctx := log.Ctx(ctx).With().
				Str("job", "set_premium_for_new_users").
				Time("time", time.Now()).
				Logger().WithContext(ctx)
			if err := runSetPremiumForNewUsers(ctx, authClient); err != nil {
				log.Ctx(ctx).Error().Err(err).Msg("Failed to set premium for new users")
			}
		}
	}
}

func runSetPremiumForNewUsers(ctx context.Context, authClient *firebaseAuth.Client) error {
	iter := authClient.Users(ctx, "")
	for {
		user, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Failed to get next user")
			break
		}

		var claims map[string]interface{}
		if user.CustomClaims != nil {
			claims = user.CustomClaims
		} else {
			claims = make(map[string]interface{})
		}

		if _, ok := claims["premium_until"]; ok {
			log.Ctx(ctx).Debug().Str("uid", user.UID).Msg("User already has premium")
			continue
		}

		claims["premium_until"] = time.Now().AddDate(0, 1, 0).Unix()

		if err := authClient.SetCustomUserClaims(ctx, user.UID, claims); err != nil {
			log.Ctx(ctx).Error().Err(err).Str("uid", user.UID).Msg("Failed to set premium")
		} else {
			log.Ctx(ctx).Info().Str("uid", user.UID).Msg("Set premium")
		}
	}

	return nil
}
