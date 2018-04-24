package handler

import (
	"net/http"
	"mdNote/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type UserClaim struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

func (this *UserClaim) Set(u model.User) {
	now := time.Now()

	this.Token = u.ID
	this.Name = u.Name
	this.IssuedAt = now.Unix()
	this.ExpiresAt = now.Add(time.Hour * 24).Unix()
}

func (h Handler) Login(c echo.Context) error {
	jsonUser := new(UserClaim)
	if err := c.Bind(jsonUser); err != nil {
		return err
	}

	dbUser := &model.User{
		ID: jsonUser.Token,
	}

	if result := h.DB.First(dbUser); result.Error != nil {
		return result.Error
	} else if result.RecordNotFound() {
		return echo.NewHTTPError(http.StatusNoContent)
	} else {
		jsonUser.Set(*dbUser)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jsonUser)

		t, err := token.SignedString(h.SecretKey)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}

func (h Handler) SignUp(c echo.Context) error {
	jsonUser := new(UserClaim)
	c.Bind(jsonUser)

	dbUser := &model.User{
		ID:   jsonUser.Token,
		Name: jsonUser.Name,
	}

	if result := h.DB.Create(dbUser); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusCreated)
}

func (h Handler) AuthRequired() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &UserClaim{},
		SigningKey: h.SecretKey,
		AuthScheme: "JWT",
	}
	return middleware.JWTWithConfig(config)
}
