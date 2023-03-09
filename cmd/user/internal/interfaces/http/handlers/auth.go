package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"

	"placio-api/cmd/user/internal/domain/user"
	"placio-api/pkg/commandbus"
	apperrors "placio-api/pkg/errors"
	httpjson "placio-api/pkg/http/response/json"
)

type requestBody struct {
	Email string `json:"email"`
}

const authCookieName = "oauthstate"

// BuildSocialAuthHandler wraps user gRPC client with http.Handler
func BuildSocialAuthHandler(config *oauth2.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		expiration := time.Now().Add(365 * 24 * time.Hour)

		b := make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, b); err != nil {
			return apperrors.Wrap(err)
		}

		state := base64.URLEncoding.EncodeToString(b)

		cookie := http.Cookie{Name: authCookieName, Value: state, Expires: expiration}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusTemporaryRedirect)

		return nil
	}

	return httpjson.HandlerFunc(fn)
}

// BuildAuthCallbackHandler wraps user gRPC client with http.Handler
func BuildAuthCallbackHandler(authConfig *oauth2.Config, apiURL string, cb commandbus.CommandBus, commandName string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		oauthState, _ := r.Cookie(authCookieName)
		if r.FormValue("state") != oauthState.Value {
			return apperrors.Wrap(fmt.Errorf("invalid oauth state"))
		}

		oauthToken, err := authConfig.Exchange(r.Context(), r.FormValue("code"))
		if err != nil {
			return apperrors.Wrap(err)
		}

		profileData, err := getProfile(oauthToken.AccessToken, apiURL)
		if err != nil {
			return apperrors.Wrap(err)
		}

		var emailData requestBody
		if err := json.Unmarshal(profileData, &emailData); err != nil {
			return apperrors.Wrap(err)
		}

		c, err := user.NewCommandFromPayload(commandName, profileData)
		if err != nil {
			return apperrors.Wrap(err)
		}

		if err := cb.Publish(r.Context(), c); err != nil {
			return apperrors.Wrap(err)
		}

		w.WriteHeader(http.StatusNoContent)

		return nil
	}

	return httpjson.HandlerFunc(fn)
}

func getProfile(accessToken, apiURL string) ([]byte, error) {
	resp, e := http.Get(apiURL + "?access_token=" + url.QueryEscape(accessToken))
	if e != nil {
		return nil, e
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return body, apperrors.Wrap(err)
	}

	return body, nil
}
