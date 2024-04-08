package auth

import (
	"net/http"
	"time"
)

func (a Auth) SessionLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "unable to parse form", http.StatusInternalServerError)
		return
	}

	idToken := r.Form.Get("idToken")
	if idToken == "" {
		http.Error(w, "idToken is required", http.StatusBadRequest)
		return
	}

	decoded, err := a.authClient.VerifyIDToken(r.Context(), idToken)
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}
	// Return error if the sign-in is older than 5 minutes.
	if time.Now().Unix()-int64(decoded.Claims["auth_time"].(float64)) > 5*60 {
		http.Error(w, "Recent sign-in required", http.StatusUnauthorized)
		return
	}

	expiresIn := time.Hour * 24 * 5
	cookie, err := a.authClient.SessionCookie(r.Context(), idToken, expiresIn)
	if err != nil {
		http.Error(w, "Failed to create a session cookie", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    cookie,
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		Secure:   true,
	})

	w.Write([]byte(`Logged in`))
}

func (a Auth) SessionLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/login", http.StatusFound)
}
