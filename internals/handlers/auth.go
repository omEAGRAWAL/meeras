package handlers

//
//import (
//	"context"
//	"go.mongodb.org/mongo-driver/mongo"
//	"meeras/internals/database"
//	"meeras/internals/models"
//	"net/http"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"golang.org/x/crypto/bcrypt"
//)
//
//// HashPassword hashes a given password using bcrypt
//func HashPassword(password string) (string, error) {
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		return "", err
//	}
//	return string(hashedPassword), nil
//}
//
//// SignupHandler handles user registration
//func SignupHandler(c *gin.Context) {
//	var user models.User
//	if err := c.ShouldBindJSON(&user); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
//		return
//	}
//
//	// Hash the password before storing
//	hashedPassword, err := HashPassword(user.Password)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
//		return
//	}
//	user.Password = hashedPassword
//	user.ID = primitive.NewObjectID()
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	collection := database.Client.Database("meeras").Collection("users")
//
//	// Check if user already exists
//	res := collection.FindOne(ctx, bson.M{"email": user.Email})
//	if res.Err() == nil { // User already exists
//		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
//		return
//	} else if res.Err() != mongo.ErrNoDocuments { // Other DB errors
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
//		return
//	}
//	res2 := collection.FindOne(ctx, bson.M{"mobile": user.Mobile})
//	if res2.Err() == nil { // User already exists
//		c.JSON(http.StatusConflict, gin.H{"error": "User already exists mobile"})
//		return
//	} else if res2.Err() != mongo.ErrNoDocuments { // Other DB errors
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
//		return
//	}
//
//	// Insert new user
//	_, err = collection.InsertOne(ctx, bson.M{
//		"_id":      user.ID,
//		"name":     user.Name,
//		"email":    user.Email,
//		"mobile":   user.Mobile,
//		"password": user.Password,
//		"role":     "customer", // Default role
//	})
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
//}
//package handlers

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"meeras/internals/database"
	"meeras/internals/models"
	"net/http"
	"time"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// HashPassword hashes a given password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plain text password
func CheckPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// GenerateToken creates a JWT token for authentication
func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// SignupHandler handles user registration
func SignupHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Hash the password before storing
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Client.Database("meeras").Collection("users")

	// Check if user already exists
	res := collection.FindOne(ctx, bson.M{"email": user.Email})
	if res.Err() == nil { // User already exists
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	} else if res.Err() != mongo.ErrNoDocuments { // Other DB errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Insert new user
	_, err = collection.InsertOne(ctx, bson.M{
		"_id":      user.ID,
		"name":     user.Name,
		"email":    user.Email,
		"mobile":   user.Mobile,
		"password": user.Password,
		"role":     "customer", // Default role
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginHandler handles user login and returns a JWT token
func LoginHandler(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.Client.Database("meeras").Collection("users")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verify password
	if !CheckPassword(user.Password, credentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
