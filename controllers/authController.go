package controllers

import (
	"jwt/database"
	"jwt/models"
	"jwt/tinkoff"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if err := database.DB.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.Id); err != nil {
		return err
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	if err := database.DB.QueryRow(
		"SELECT id, name, email, password FROM users WHERE email = $1",
		data["email"],
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	); err != nil {
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unathenticacted",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	if err := database.DB.QueryRow(
		"SELECT id, name, email, password FROM users WHERE id = $1",
		claims.Issuer,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	); err != nil {
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Tinkoff(c *fiber.Ctx) error {
	tinkoff.Rest()

	rows, err := database.DB.Query("SELECT price FROM candels")
	if err != nil {
		return err
	}
	defer rows.Close()

	candelPrices := make([]float64, 0)

	for rows.Next() {
		var price float64
		if err := rows.Scan(&price); err != nil {
			return err
		}
		candelPrices = append(candelPrices, price)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	return c.JSON(candelPrices)
}
