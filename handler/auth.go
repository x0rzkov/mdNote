package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"mdNote/model"
	"net/http"
	"os"
	"strconv"
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

func (this *UserClaim) Ensure() error {
	if this.Id == "" || this.Name == "" {
		return errors.New("UserClaim.Ensure(): UserClaim validation error")
	}
	now := time.Now()

	this.IssuedAt = now.Unix()
	this.ExpiresAt = now.Add(time.Hour * 24).Unix()

	return nil
}

// GET /auth/callback/:provider
func (h Handler) Auth(c echo.Context) error {
	jsonUser := new(UserClaim)

	switch provider := c.Param("provider"); provider {
	case "github":
		bindData := echo.Map{}
		if err := c.Bind(&bindData); err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		client := &http.Client{}
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(echo.Map{
			"client_id":     bindData["clientId"].(string),
			"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
			"code":          bindData["code"].(string),
		})
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		req, err := http.NewRequest(echo.POST, "https://github.com/login/oauth/access_token", buf)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Accept", "application/json")

		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		resBody := echo.Map{}
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if err, exist := resBody["error"]; exist {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		accessToken, exist := resBody["access_token"]
		if !exist {
			log.Println(errors.New("Cannot find access_token"))
			return echo.NewHTTPError(http.StatusBadRequest, "Cannot find access_token")
		}

		req, err = http.NewRequest(echo.GET, "https://api.github.com/user", nil)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		query := req.URL.Query()
		query.Add("access_token", accessToken.(string))
		req.URL.RawQuery = query.Encode()

		res, err = client.Do(req)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		resBody = echo.Map{}
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if err, exist := resBody["error"]; exist {
			log.Println(errors.New(err.(string)))
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		id, exist := resBody["id"]
		if !exist {
			log.Println(errors.New("Cannot find id"))
			return echo.NewHTTPError(http.StatusBadRequest, "Cannot find id")
		}

		name, exist := resBody["login"]
		if !exist {
			log.Println(errors.New("Cannot find login"))
			return echo.NewHTTPError(http.StatusBadRequest, "Cannot find login")
		}

		jsonUser.Id = strconv.Itoa(int(id.(float64)))
		jsonUser.Name = name.(string)
	}

	dbUser := &model.User{
		ID: jsonUser.Id,
	}

	httpStatus := http.StatusOK

	if result := h.DB.First(dbUser); result.Error != nil {
		log.Println(result.Error)
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	} else if result.RecordNotFound() {
		if err := h.DB.Create(dbUser); err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		httpStatus = http.StatusCreated
	}

	if err := jsonUser.Ensure(); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jsonUser)

	tokenString, err := token.SignedString(h.SecretKey)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "JWT"
	cookie.Value = tokenString
	cookie.Expires = time.Unix(jsonUser.ExpiresAt, 0)
	c.SetCookie(cookie)

	return c.NoContent(httpStatus)
}

func (h Handler) AuthRequired() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &UserClaim{},
		SigningKey: h.SecretKey,
		AuthScheme: "JWT",
	}
	return middleware.JWTWithConfig(config)
}
