package controller

import (
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/gorm"
	"placio-app/models"
)

//ref
//https://blog.cloudnativefolks.org/oauth-20-implementation-in-golang
//
// https://github.com/gofiber/jwt
// https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j
// https://github.com/markbates/goth/tree/master/providers/twitterv2
// https://github.com/gofiber/fiber/issues/292

//users := []models.User{}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept */*
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/v1/users [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	//user.ID = len(users) + 1
	//users = append(users, *user)
	return c.JSON(user)
}

// GetAllUsers godoc
// @Summary Retrieve a list of users
// @Description Retrieve a list of users
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} []models.User
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users [get]
func GetAllUsers(c *fiber.Ctx) error {
	//return c.JSON(users)
	return c.JSON(fiber.Map{"status": "ok"})
}

// GetUserByID godoc
// @Summary Retrieve a user by ID
// @Description Retrieve a user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// UpdateUser godoc
// @Summary Update an existing user by ID
// @Description Update an existing user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for i, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		if err := c.BodyParser(&users[i]); err != nil {
	//			return err
	//		}
	//		return c.JSON(user)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for i, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		users = append(users[:i], users[i+1:]...)
	//		return c.SendStatus(fiber.StatusNoContent)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetMessagesSent godoc
// @Summary Retrieve a list of messages sent by a specific user by ID
// @Description Retrieve a list of messages sent by a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Message
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/messages_sent [get]
func GetMessagesSent(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.MessagesSent)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetMessagesReceived godoc
// @Summary Retrieve a list of messages received by a specific user by ID
// @Description Retrieve a list of messages received by a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Message
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/messages_received [get]
func GetMessagesReceived(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.MessagesReceived)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetConversations godoc
// @Summary Retrieve a list of conversations a specific user is a participant in by ID
// @Description Retrieve a list of conversations a specific user is a participant in by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Conversation
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/conversations [get]
func GetConversations(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Conversations)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetGroups godoc
// @Summary Retrieve a list of groups a specific user is a member of by ID
// @Description Retrieve a list of groups a specific user is a member of by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Group
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/groups [get]
func GetGroups(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Groups)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetVoiceNotesSent godoc
// @Summary Retrieve a list of voice notes sent by a specific user by ID
// @Description Retrieve a list of voice notes sent by a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.VoiceNote
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/voice_notes_sent [get]
func GetVoiceNotesSent(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotesSent)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetVoiceNotesReceived godoc
// @Summary Retrieve a list of voice notes received by a specific user by ID
// @Description Retrieve a list of voice notes received by a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.VoiceNote
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/voice_notes_received [get]
func GetVoiceNotesReceived(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotesReceived)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetVoiceNotes godoc
// @Summary Retrieve a list of voice notes a specific user is a participant in by ID
// @Description Retrieve a list of voice notes a specific user is a participant in by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.VoiceNote
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/voice_notes [get]
func GetVoiceNotes(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotes)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetVoiceNote godoc
// @Summary Retrieve a specific voice note by ID
// @Description Retrieve a specific voice note by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "Voice Note ID"
// @Success 200 {object} models.VoiceNote
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/voice_notes/{id} [get]
func GetVoiceNote(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, voiceNote := range voiceNotes {
	//	if strconv.Itoa(voiceNote.ID) == id {
	//		return c.JSON(voiceNote)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("Voice note not found")
}

// GetUserNotifications godoc
// @Summary Retrieve a list of notifications for a specific user by ID
// @Description Retrieve a list of notifications for a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Notification
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/notifications [get]
func GetUserNotifications(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Notifications)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetUserBookings godoc
// @Summary Retrieve a list of bookings for a specific user by ID
// @Description Retrieve a list of bookings for a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Booking
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/bookings [get]
func GetUserBookings(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Bookings)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

// GetUserPayments godoc
// @Summary Retrieve a list of payments for a specific user by ID
// @Description Retrieve a list of payments for a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Payment
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/payments [get]
func GetUserPayments(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Payments)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}