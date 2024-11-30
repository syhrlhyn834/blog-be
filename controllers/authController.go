package controllers

import (
	"be-golang/connection"
	"be-golang/models"
	"be-golang/resources"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret key untuk menandatangani JWT
var secretKey = []byte("KontolKuda123") // Gantilah dengan secret key yang lebih aman

// Struct untuk validasi login
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Fungsi login
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := resources.ProcessValidationErrors(err)
		if errors != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
	}

	// Cari user berdasarkan email
	var user models.User
	if err := connection.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Cek password langsung tanpa bcrypt
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	// Menandatangani token dengan secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Kirimkan token sebagai respon
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successful",
		"data":    user,
		"token":   tokenString,
	})
}

// Fungsi untuk mendapatkan data user yang sedang login
func GetUser(c *gin.Context) {
	// Ambil user berdasarkan token yang valid
	userId, _ := c.Get("userId")
	var user models.User
	if err := connection.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    user,
	})
}

// Fungsi untuk me-refresh token
func RefreshToken(c *gin.Context) {
	// Ambil token dari header Authorization
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]

	// Verifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token method")
		}
		return secretKey, nil
	})

	// Cek apakah token valid
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Ambil ID user dari claims token
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["sub"].(float64)

	// Masukkan token yang lama ke dalam blacklist
	blacklistedToken := models.TokenBlacklist{
		Token:     tokenString,
		CreatedAt: time.Now(),
		ExpiredAt: time.Unix(int64(claims["exp"].(float64)), 0),
	}
	connection.DB.Create(&blacklistedToken)

	// Generate token baru
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Menandatangani token baru
	newTokenString, err := newToken.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new token"})
		return
	}

	// Kirimkan token baru sebagai respon
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   newTokenString,
	})
}

// Fungsi untuk logout
func Logout(c *gin.Context) {
	// Ambil token dari header Authorization
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]

	// Masukkan token ke blacklist
	blacklistedToken := models.TokenBlacklist{
		Token:     tokenString,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Hour), // misalnya token kadaluarsa 1 jam
	}
	connection.DB.Create(&blacklistedToken)

	// Respon logout berhasil
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged out successfully",
	})
}
