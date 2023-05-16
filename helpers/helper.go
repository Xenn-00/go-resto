package helpers

import (
	"log"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	if err != nil {
		return false, "incorect email or password"
	}
	return true, ""
}

func ToFixed(num float64, precision int) float64 {
	return math.Round(num/float64(precision)) * float64(precision)
}

func ItemByOrder(id string) (OrderItem []primitive.M, err error) {
	return nil, nil
}

func InTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}