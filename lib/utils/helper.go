package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func getValidationErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return "must be at least " + fe.Param() + " characters long"
	}
	return "is not valid"
}

func ValidateStruct(s interface{}, validate *validator.Validate) error {
	if err := validate.Struct(s); err != nil {
		var sb strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString(fmt.Sprintf("%s %s; ", err.Field(), getValidationErrorMsg(err)))
		}
		return fmt.Errorf("validation failed: %s", sb.String())
	}
	return nil
}

// GeneratePasswordHash hashes the plain password using bcrypt
func GeneratePasswordHash(password string) (string, error) {
	// bcrypt.DefaultCost is 10; you can also use bcrypt.MinCost or bcrypt.MaxCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengatur header untuk mencegah MIME-sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		// Mengatur header untuk mencegah website dibungkus dalam frame
		c.Header("X-Frame-Options", "DENY")
		// Mengaktifkan proteksi XSS pada browser
		c.Header("X-XSS-Protection", "1; mode=block")
		// Mengatur Content-Security-Policy untuk membatasi sumber daya yang dapat dimuat
		c.Header("Content-Security-Policy", "default-src 'self'")
		// Mengatur Referrer-Policy agar referrer tidak dikirim secara berlebihan
		c.Header("Referrer-Policy", "no-referrer")
		// Jika menggunakan HTTPS, sebaiknya aktifkan Strict-Transport-Security
		// c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		// Lanjutkan eksekusi request
		c.Next()
	}
}

func CorsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set header CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Lanjutkan eksekusi request
		c.Next()
	}
}
