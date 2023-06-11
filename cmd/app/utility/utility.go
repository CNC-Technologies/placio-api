package utility

import (
	"encoding/json"
	"fmt"
	"strings"

	sentry "github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

type IUtility interface {
	Validate(form map[string]interface{}, fields []string) error
	//ValidateEmail(email string) error
	//ValidatePassword(password string) error
	//ValidatePhone(phone string) error
	//ValidateName(name string) error
	//ValidateAddress(address string) error
	//ValidateDate(date string) error
	//ValidateTime(time string) error
	//ValidatePasswordMatch(password string, confirmPassword string) error
	//ValidatePasswordStrength(password string) error
}

type Utility struct {
}

func NewUtility() *Utility {
	return &Utility{}
}
func (u *Utility) Validate(form map[string]interface{}, fields []string) error {
	// sanitise the input
	for f, v := range form {
		// sanitise
		if s, ok := v.(string); ok && strings.Contains(s, "<script>") {
			form[f] = strings.ReplaceAll(strings.ReplaceAll(s, "<script>", ""), "</script>", "")
		}
	}

	// check for required fields
	if len(fields) > 0 {
		for _, f := range fields {
			if _, ok := form[f]; !ok || form[f] == nil {
				// field is required
				return fmt.Errorf("%s field is required", f)
			}
		}
	}

	return nil

}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf("invalid email address")
	}
	return nil
}
func ValidatePassword(password string) error {
	// check password length
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	// check password complexity
	if !strings.ContainsAny(password, "0123456789") {
		return fmt.Errorf("password must contain at least one number")
	}

	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	return nil
}

func ValidatePasswordMatch(password, confirmPassword string) error {
	if password != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}

func Assert(data interface{}, err string, input map[string]interface{}) bool {
	if data == nil {
		m := map[string]interface{}{"message": err}
		if input != nil {
			for k, v := range input {
				m[k] = v
			}
		}
		panic(m)
	}
	return true
}

//
//func Use(fn func(*fiber.Ctx) error) fiber.Handler {
//	//fmt.Println("Entering utility.Use function")
//	//logger.Info(context.Background(), "middleware.Use")
//	defer sentry.Recover()
//	//defer sentry.Flush(2 * time.Second)
//	return func(c *fiber.Ctx) error {
//		err := fn(c)
//		if err != nil {
//			sentry.CaptureException(err)
//			return c.Next()
//		}
//		return nil
//	}
//}

func Use(fn func(*gin.Context) error) gin.HandlerFunc {
	//fmt.Println("Entering utility.Use function")
	//logger.Info(context.Background(), "middleware.Use")
	defer sentry.Recover()
	//defer sentry.Flush(2 * time.Second)
	return func(c *gin.Context) {
		err := fn(c)
		if err != nil {
			sentry.CaptureException(err)
			c.Next()
		}
	}
}

// StructToMap structToMap converts a struct to a map, omitting fields that are nil.
func StructToMap(v interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// MergeMaps mergeMaps merges two map[string]interface{} objects.
func MergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}
	return merged
}

// ProcessResponse processes the response from a service.
func ProcessResponse(data interface{}, status string, message string) gin.H {
	return gin.H{
		"data":    data,
		"status":  status,
		"message": message,
	}
}

// SplitString
func SplitString(s string, sep string) []string {
	return strings.Split(s, sep)
}
