package controller

import (
	"context"
	"fmt"
	"net/http"
	Dto "placio-app/Dto"
	"placio-app/database"
	"placio-app/middleware"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
	"placio-pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService service.IAuth
	utility     utility.IUtility
}

func NewAuthController(authService service.IAuth, utility utility.IUtility) *AuthController {
	return &AuthController{AuthService: authService, utility: utility}
}

func (controller *AuthController) RegisterRoutes(app *gin.RouterGroup) {
	authGroup := app.Group("/auth")

	authGroup.GET("/auth-status", utility.Use(func(ctx *gin.Context) error {
		authorizationHeader := ctx.GetHeader("Authorization")
		headerParts := strings.Split(authorizationHeader, " ")
		loginData := middleware.Status(headerParts[1], ctx)
		ctx.JSON(200,gin.H{"user":loginData})
		// if loginData != nil {
		// 	ctx.Status(200)
		// }
		ctx.Abort()
		return nil
	}))

	authGroup.POST("/sign-in", utility.Use(controller.signIn))
	//authGroup.POST("/auth/sign-up", utility.Use(a.SignUp))
	authGroup.POST("/sign-out", middleware.Verify("user"), utility.Use(controller.signOut))
	authGroup.GET("/refresh", middleware.Verify("user"), utility.Use(controller.refreshToken))
	authGroup.POST("/verify/password", middleware.Verify("user"), utility.Use(controller.changePassword))
	authGroup.POST("/verify/email", middleware.Verify("user"), utility.Use(controller.verifyEmail))
	authGroup.POST("/reset", utility.Use(controller.resetPassword))
	authGroup.POST("/verify/phone", middleware.Verify("user"), utility.Use(controller.verifyPhone))
	authGroup.GET("/status", middleware.Verify("user"), utility.Use(controller.getAuthStatus))
	authGroup.POST("/otp", middleware.Verify("user"), utility.Use(controller.getOTP))
	authGroup.POST("/magic", middleware.Verify("user"), utility.Use(controller.GetMagicLink))
	authGroup.POST("/magic/verify", middleware.Verify("user"), utility.Use(controller.verifyMagicLink))
	authGroup.POST("/password/reset/request", middleware.Verify("user"), utility.Use(controller.requestPasswordReset))
	authGroup.POST("/switch-account", middleware.Verify("user"), utility.Use(controller.switchAccount))
	authGroup.POST("/impersonate", middleware.Verify("user"), utility.Use(controller.impersonateUser))

}

var (
	userData *models.User
	useEmail bool
	account  *models.Account
	login    *models.LoginModel
	token    models.Token
)

// SignOut godoc
// @Summary Log out
// @Description Log out from the server
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 403 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Security ApiKeyAuth
// @Router /auth/sign-out [post]
func (controller *AuthController) signOut(c *gin.Context) error {
	tokenId, ok := c.Get("tokenID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token id"})
		return nil
	}

	provider, ok := c.Get("provider")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing provider"})
		return nil
	}

	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user"})
		return nil
	}
	err := token.Delete((tokenId).(string), (provider).(string), (user).(string), database.DB)
	if err != nil {
		logger.Error(context.Background(), err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting token"})
		return nil
	}
	// remove all from context
	c.Set("tokenID", nil)
	c.Set("provider", nil)
	c.Set("user", nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
	return nil
}

// RefreshToken godoc
// @Summary Refresh token
// @Description Refresh token
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} Dto.TokenResponse
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 403 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /api/v1/auth/refresh [get]
func (controller *AuthController) refreshToken(c *gin.Context) error {
	// get the refresh token from the request

	refreshToken := c.GetHeader("refresh_token")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing refresh token"})
		return nil
	}
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty refresh token"})
		return nil
	}

	tokens, err := token.RefreshTokens(refreshToken, database.DB)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    tokens,
		"message": "Token refreshed successfully",
		"status":  "ok",
	})

	return nil

}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 403 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /auth/reset [get]
func (controller *AuthController) resetPassword(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return nil
}

// ChangePassword godoc
// @Summary Change password
// @Description Change password
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 403 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /auth/change [get]
func (controller *AuthController) changePassword(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return nil
}

// VerifyEmail godoc
// @Summary Verify email
// @Description Verify email
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 403 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /api/v1/auth/verify [get]
func (controller *AuthController) verifyEmail(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return nil
}

// VerifyPhone godoc
// @Summary Verify phone
// @Description Verify phone
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} Dto.ErrorDTO "Bad request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 403 {object} Dto.ErrorDTO "Forbidden"
// @Failure 404 {object} Dto.ErrorDTO "Not found"
// @Failure 500 {object} Dto.ErrorDTO "Internal server error"
// @Router /api/v1/auth/verify [get]
func (controller *AuthController) verifyPhone(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return nil
}

// Signin authenticates a user via email/password or social network
// Signin authenticates a user and returns an access token.
// The function performs the following steps:
// 1. Parses the incoming request body into a SigninRequest.
// 2. Validates the input data.
// 3. Checks if the user exists.
// 4. Verifies the user's password.
// 5. Retrieves the user's account data.
// 6. Logs the sign-in attempt and checks for suspicious activity.
// 7. Generates an access token for the authenticated user.
// 8. Returns the access token in a SigninResponse.
//
// @Summary Authenticate a user
// @Description Authenticate a user and return an access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param SigninRequest body Dto.SigninRequest true "Sign In Data"
// @Success 200 {object} Dto.UserResponse "Successfully signed in"
// @Failure 400 {object} fiber.Error "Bad Request"
// @Failure 401 {object} fiber.Error "Unauthorized"
// @Failure 403 {object} fiber.Error "Forbidden"
// @Failure 404 {object} fiber.Error "Not Found"
// @Router /api/v1/auth/sign-in [post]
func (controller *AuthController) signIn(c *gin.Context) error {
	//defer sentry.Recover()
	data := new(Dto.SigninRequest)
	//logger.Info(context.Background(), string(c.BindJSON(data)))
	if err := c.BindJSON(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	var user models.User

	if data.Email != "" {
		useEmail = true
		data.Provider = "app"
		if err := controller.utility.Validate(data.ToJson(), []string{"email", "password"}); err != nil {
			return err
		}
	} else {
		if err := controller.utility.Validate(data.ToJson(), []string{"token"}); err != nil {
			return err
		}
		decode, err := token.VerifyToken("app", data.Token)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		data.Provider = decode.Provider
		// data.ProviderID = decode.ProviderID
		data.Email = decode.Email
	}

	// check user exists
	var err error
	if useEmail {
		logger.Info(context.Background(), "use email")
		userData, err = user.GetByEmail(data.Email, database.DB)
	} else {
		//userData, err = user.GET(nil, "", nil, map[string]string{
		//	"provider": data.Provider,
		//	"id":       data.ProviderID,
		//})
	}
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Please enter the correct login details")
	}

	// verify password
	if useEmail {
		verified, err := userData.VerifyPassword(data.Password)
		if err != nil {
			return err
		}
		//logger.Info(context.Background(), fmt.Sprintf("verified: %v", verified))
		if !verified {
			return fiber.NewError(fiber.StatusUnauthorized, "Please enter the correct login details")
		}
	}

	userAgent := c.Request.Header.Get("User-Agent")
	device := c.Request.Header.Get("Device")

	// log the sign in and check if it's suspicious
	log, err := login.Create(userData.ID, c.ClientIP(), userAgent, device, database.DB)
	if err != nil {
		return err
	}
	loginVerification, err := login.Verify(userData.ID, log, database.DB)
	if err != nil {
		return err
	}
	if loginVerification.Suspicious {
		return fiber.NewError(fiber.StatusUnauthorized, "Your account has been flagged for suspicious activity. Please contact support.")
	}

	// generate the token
	tokenData, err := user.GenerateToken(*userData)

	var newData = &models.Token{
		Provider: "app",
		Jwt:      tokenData.Access,
		Access:   tokenData.Access,
		ID:       tokenData.TokenID,
		TokenID:  tokenData.TokenID,
		//AccessTokenExpiry: tokenData.AccessExpiresIn,
		Refresh:          tokenData.Refresh,
		UserID:           tokenData.UserID,
		CodeCreateAt:     time.Time{},
		CodeExpiresIn:    tokenData.CodeExpiresIn,
		AccessCreateAt:   time.Time{},
		AccessExpiresIn:  tokenData.AccessExpiresIn,
		RefreshCreateAt:  time.Time{},
		RefreshExpiresIn: tokenData.RefreshExpiresIn,
	}

	logger.Info(context.Background(), fmt.Sprintf("newData: %v", newData))

	err = newData.Save(database.DB)
	if err != nil {
		return err
	}

	// return the token
	c.JSON(http.StatusOK, userData.GenerateUserResponse(newData))
	return nil
}

//func (controller *AuthController) authenticate(c *gin.Context, userData models.User) error {
//	accountData, err := account.GetAccount(userData.AccountID, database.DB)
//	if err != nil {
//		return err
//	}
//
//	subscription, err := account.GetSubscription(userData.AccountID)
//	if err != nil {
//		return err
//	}
//
//	userAccounts, err := account.GetUserAccount(userData.ID)
//	if err != nil {
//		return err
//	}
//
//	// create & store the token
//	//jwt, err := token.Generate(userData.ID, userData.AccountID, "app", userData.AccountID, userData.Email)
//	if err != nil {
//		return err
//	}
//	//
//	//err = token.Save("app", map[string]string{
//	//	"provider_id":  userData.AccountID,
//	//	"email":        userData.Email,
//	//	"accessToken":  jwt.Access,
//	//	"refreshToken": jwt.Refresh,
//	//	//"expires":     jwt.AccessExpiresIn,
//	//
//	//}, userData.ID)
//	if err != nil {
//		return err
//	}
//
//	userData.UpdateUser(userData.ID, userData.AccountID, time.Now(), false)
//
//	// return user to client
//	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
//		//"token":        jwt,
//		"subscription": subscription.Status,
//		"plan":         accountData.Plan,
//		"permission":   userData.Permission,
//		"name":         userData.Name,
//		"accounts":     userAccounts,
//		"account_id":   userData.AccountID,
//		"has_password": userData.HasPassword,
//		"onboarded":    userData.Onboarded,
//	})
//}

//func Signout(c *gin.Context) error {
//	// destroy social tokens
//	err := token.Delete(nil, c.Locals("provider").(string), c.Locals("user").(int))
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to sign out",
//		})
//	}
//	return c.SendStatus(fiber.StatusOK)
//}

// SwitchAccount godoc
// @Summary Switch account
// @Description Lets a user switch account
// @Tags Auth
// @Accept json
// @Produce json
// @Param account path string true "Account ID"
// @Success 200 {object} Dto.UserResponse "User"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 403 {object} Dto.ErrorDTO "Forbidden"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /auth/switch-account/{account} [post]
func (controller *AuthController) switchAccount(c *gin.Context) error {
	// GET the user and account ID from the request
	//userID := c.Locals("user").(string)
	//accountID := c.Params("account")

	// Check if the user belongs to this account
	//userData, err := userData.GetUserById(userID, database.DB)
	//if err != nil {
	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//		"message": "You don't belong to this account.",
	//	})
	//}

	//return Authenticate(c, userData)
	return nil
}

// ImpersonateUser godoc
// @Summary Impersonate user
// @Description Lets an admin impersonate another user
// @Tags Auth
// @Accept json
// @Produce json
// @Param account path string true "Account ID"
// @Param user path string true "User ID"
// @Success 200 {object} Dto.UserResponse "User"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 403 {object} Dto.ErrorDTO "Forbidden"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /auth/impersonate/{account}/{user} [post]
func (controller *AuthController) impersonateUser(c *gin.Context) error {
	//	// GET the token from the request body
	//	token := new(struct {
	//		Token string `json:"token"`
	//	})
	//	if err := c.BodyParser(token); err != nil {
	//		return errors.Wrap(err, "failed to parse request body")
	//	}
	//
	//	// Verify the authorization token
	//	data, err := auth.VerifyToken(token.Token)
	//	if err != nil {
	//		return errors.Wrap(err, "failed to verify token")
	//	}
	//
	//	// Check if the token is valid and has master permission
	//	if data.UserID == "" || data.Permission != "master" {
	//		return utility.NewError(fiber.StatusUnauthorized, "Invalid token")
	//	}
	//
	//	// GET the user data
	//	userData, err := user.GET(data.UserID)
	//	if err != nil {
	//		return errors.Wrap(err, "failed to get user data")
	//	}
	//
	//	// Authenticate the user and generate a new token
	//	jwt, err := auth.GenerateToken(userData.AccountID, userData.ID, userData.Permission, "app")
	//	if err != nil {
	//		return errors.Wrap(err, "failed to generate token")
	//	}
	//
	//	// Save the token to the database
	//	if err := auth.SaveToken("app", jwt.AccessToken, userData.ID); err != nil {
	//		return errors.Wrap(err, "failed to save token to database")
	//	}
	//
	//	// Update user's last active timestamp and disabled status
	//	if err := user.Update(userData.ID, userData.AccountID, user.UpdateParams{
	//		LastActive: time.Now(),
	//		Disabled:   false,
	//	}); err != nil {
	//		return errors.Wrap(err, "failed to update user")
	//	}
	//
	//	// Return the response
	//	return c.JSON(fiber.Map{
	//		"token":        jwt.AccessToken,
	//		"subscription": "",
	//		"plan":         "",
	//		"permission":   userData.Permission,
	//		"name":         userData.Name,
	//		"accounts":     nil,
	//		"account_id":   userData.AccountID,
	//		"has_password": false,
	//		"onboarded":    userData.Onboarded,
	//	})
	return nil
}

// GetAuthStatus godoc
// @Summary GET auth status
// @Description Returns the auth status of the user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} Dto.AuthStatusResponse "Auth status"
// @Failure 401 {object} fiber.Map "Unauthorized"
// @Failure 500 {object} fiber.Map "Internal Server Error"
// @Router /auth/status [get]
func (controller *AuthController) getAuthStatus(c *gin.Context) error {
	//	// Check if there's a valid account/user
	//	var hasJWT, hasSocialToken, usingSocialSignin bool
	//
	//	// Check if there's an active JWT
	//	if c.Provider() == "app" {
	//		usingSocialSignin = false
	//		hasJWT = token.Verify("app", c.UserID())
	//	}
	//
	//	// Check if there's an active access token if the user is
	//	// signed in via social network or was their account de-authed
	//	if c.Provider() != "app" {
	//		usingSocialSignin = true
	//		hasSocialToken = token.Verify(c.Provider(), c.UserID())
	//	}
	//
	//	// Check if the user has an active subscription
	//	subscription := account.Subscription(c.AccountID())
	//	userAccounts := user.Accounts(c.UserID())
	//	user.UpdateLastActive(c.UserID(), c.AccountID(), time.Now())
	//
	//	// Return the auth status
	//	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	//		"data": fiber.Map{
	//			"jwt_token":     hasJWT,
	//			"social_token":  hasSocialToken,
	//			"subscription":  subscription.Status,
	//			"accounts":      userAccounts,
	//			"account_id":    c.AccountID(),
	//			"authenticated": usingSocialSignin && hasSocialToken || !usingSocialSignin && hasJWT,
	//		},
	//	})
	return nil
}

// VerifyMagicLink godoc
// @Summary Verify magic link
// @Description Verifies a magic link
// @Tags Auth
// @Accept json
// @Produce json
// @Param token body string true "Magic link token"
// @Success 200 {object} Dto.UserResponse "User"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /auth/magic/verify [post]
func (controller *AuthController) verifyMagicLink(c *gin.Context) error {
	//	data := new(struct {
	//		Token string `json:"token"`
	//	})
	//	if err := c.BodyParser(data); err != nil {
	//		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	//	}
	//
	//	magicToken, err := Verify(data.Token)
	//	if err != nil {
	//		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	//	}
	//
	//	userData, err := user.GET(magicToken.UserID)
	//	if err != nil {
	//		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	//	}
	//
	//	// log the sign in and check if it's suspicious
	//	log, err := login.Create(userData.ID, c)
	//	if err != nil {
	//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create login record")
	//	}
	//	risk, err := login.Verify(userData.ID, log)
	//	if err != nil {
	//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to verify login risk")
	//	}
	//
	//	// notify the user of suspicious logins
	//	if risk.Level > 0 {
	//		err = mail.Send(mail.Options{
	//			To:        userData.Email,
	//			Subject:   "New sign-in on your account",
	//			Template:  "new_signin",
	//			Variables: mail.Variables{"ip": risk.Flag.IP, "time": risk.Time, "device": risk.Flag.Device, "browser": risk.Flag.Browser},
	//		})
	//		if err != nil {
	//			return c.Status(fiber.StatusInternalServerError).SendString("Failed to send email")
	//		}
	//	}
	//
	//	// 2fa is required
	//	if userData.TwoFactorEnabled {
	//		// notify the client and use email to identify the user when sending otp
	//		// send a token so the otp password screen can't be accessed directly without a password
	//		jwt, err := Token(TokenData{Email: userData.Email, Provider: "app"}, nil, time.Minute*5)
	//		if err != nil {
	//			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate JWT token")
	//		}
	//		return c.JSON(fiber.Map{
	//			"2fa_required": true,
	//			"token":        jwt,
	//		})
	//	}
	//
	//	// authenticate the user
	//	userAccounts, err := user.Account(userData.ID)
	//	if err != nil {
	//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve user accounts")
	//	}
	//
	//	err = TokenSave("app", userData.ID, TokenData{Access: Token(TokenData{AccountID: userData.AccountID, UserID: userData.ID, Permission: userData.Permission})})
	//	if err != nil {
	//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save token")
	//	}
	//
	//	err = user.Update(userData.ID, userData.AccountID, user.UpdateData{LastActive: time.Now(), Disabled: false})
	//	if err != nil {
	//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update user data")
	//	}
	//
	//	return c.JSON(fiber.Map{
	//		"token":        Token(TokenData{AccountID: userData.AccountID, UserID: userData.ID, Permission: userData.Permission}),
	//		"subscription": subscription.Status,
	//		"plan":         accountData.Plan,
	//		"permission":   userData.Permission,
	//		"name":         userData.Name,
	//		"accounts":     userAccounts,
	//		"account_id":   userData.AccountID,
	//		"has_password": userData.Password != "",
	//		"onboarded":    userData.Onboarded,
	//	})
	return nil
}

// GetOTP godoc
// @Summary GET OTP
// @Description Gets an OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param token header string true "JWT token"
// @Success 200 {object} fiber.Map "OTP"
// @Failure 401 {object} fiber.Map "Unauthorized"
// @Failure 500 {object} fiber.Map "Internal Server Error"
// @Router /auth/otp [get]
func (controller *AuthController) getOTP(c *gin.Context) error {
	return nil
}

// GetMagicLink godoc
// @Summary GET magic link
// @Description Gets a magic link
// @Tags Auth
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 200 {object} fiber.Map "Magic link"
// @Failure 401 {object} fiber.Map "Unauthorized"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 409 {object} Dto.ErrorDTO "Conflict"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /auth/magic [post]
func (controller *AuthController) GetMagicLink(c *gin.Context) error {
	return nil
}

// RequestPasswordReset godoc
// @Summary Request password reset
// @Description Requests a password reset
// @Tags Auth
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 200 {object} fiber.Map "Password reset"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /auth/password/reset/request [post]
func (controller *AuthController) requestPasswordReset(c *gin.Context) error {
	return nil
}

//func SocialAuthHandler(provider string, signInURL string, socialURL string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Use the passport-go middleware to authenticate the user
//		if err := passport.Authenticate(c, provider, passport.Options{
//			FailureRedirect: signInURL,
//		}); err != nil {
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape(err.Error())))
//			return
//		}
//
//		// GET the user's profile from the authenticated session
//		profile, ok := passport.ProfileFromContext(c)
//		if !ok {
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape("Unable to get user profile")))
//			return
//		}
//
//		// Authenticate the user
//		email := ""
//		if len(profile.Emails) > 0 {
//			email = profile.Emails[0].Value
//		}
//		userData, err := user.GET(nil, email, nil, map[string]interface{}{
//			"provider": provider,
//			"id":       profile.ID,
//		})
//		if err != nil {
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape("Unable to get user data")))
//			return
//		}
//
//		if userData != nil {
//			// Generate a JWT token and redirect the user to the social URL
//			jwt, err := auth.Token(map[string]interface{}{
//				"provider":    provider,
//				"provider_id": profile.ID,
//				"email":       email,
//			}, nil, 300)
//			if err != nil {
//				c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape("Unable to generate JWT token")))
//				return
//			}
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?provider=%s&token=%s", socialURL, provider, jwt))
//		} else {
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape("You're not registered")))
//		}
//	}
//}
