package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type confirmationMailInput struct {
	Email string `json:"email" binding:"required"`
}

func (h *Handler) sendCode(c *gin.Context) {
	var input confirmationMailInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf(input.Email)

	code := generateConfirmationCode()
	if err := sendConfirmationCode(input.Email, code); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to send confirmation code")
		return
	}

	if err := h.services.Authorization.SetCode(input.Email, code); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, "Confirmation code sent to email")
}

func generateConfirmationCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	code := r.Intn(999999)
	return fmt.Sprintf("%06d", code)
}

func sendConfirmationCode(email, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Код подтверждения для входа")
	m.SetBody("text/plain", "Ваш код подтверждения: "+code)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
