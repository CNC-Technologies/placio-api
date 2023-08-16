package middleware

import (
	"context"
	"github.com/getsentry/sentry-go"
	"log"
	"net/http"
	"net/url"
	"os"
	"placio-pkg/errors"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	UpdatedAt     string `json:"updated_at"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Scope         string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken() gin.HandlerFunc {
	issuerURL, err := url.Parse(os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE"), "KpDGogGXqWeuGQfZ4Wu30neiHS79hGiU", "Gv4QCgbya8fTxZACFpMdrElFhkARloMl"},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Failed to set up the jwt validator")
	}

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
			c.Abort()
			return
		}

		// remove the Bearer prefix
		tokenString = tokenString[7:]
		tokenInterface, err := jwtValidator.ValidateToken(context.Background(), tokenString)
		if err != nil {
			log.Printf("Encountered error while validating JWT: %v", err)
			sentry.CaptureException(err)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to validate JWT."})
			c.Abort()
			return
		}

		if validatedClaims, ok := tokenInterface.(*validator.ValidatedClaims); ok {
			c.Set("user", strings.Split(validatedClaims.RegisteredClaims.Subject, "|")[1])
			c.Set("auth0_id", validatedClaims.RegisteredClaims.Subject)
			c.Next()
		} else {
			sentry.CaptureException(err)
			// handle error, the assertion failed
		}
	}
}

func EnsureValidTokenButAllowAccess() gin.HandlerFunc {
	issuerURL, err := url.Parse(os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE"), "KpDGogGXqWeuGQfZ4Wu30neiHS79hGiU"},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		errors.LogAndReturnError(err)
	}

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// If no Authorization header present, let the request pass through.
		// UserID will be empty in the context.
		if tokenString == "" {
			c.Next()
			return
		}

		// remove the Bearer prefix
		tokenString = tokenString[7:]
		tokenInterface, err := jwtValidator.ValidateToken(context.Background(), tokenString)
		if err != nil {
			// if JWT is invalid, let the request pass through.
			// UserID will be empty in the context.
			log.Printf("Encountered error while validating JWT: %v", err)
			c.Next()
			return
		}

		if validatedClaims, ok := tokenInterface.(*validator.ValidatedClaims); ok {
			c.Set("user", strings.Split(validatedClaims.RegisteredClaims.Subject, "|")[1])
			c.Set("auth0_id", validatedClaims.RegisteredClaims.Subject)
		}

		c.Next()
	}
}