package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
		client := &http.Client{}
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(echo.Map{
			"client_id":     os.Getenv("GITHUB_CLIENT_ID"),
			"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
			"code":          c.QueryParam("code"),
		})
		if err != nil {
			err = fmt.Errorf("access_token request json encoding: %v", err)
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		req, err := http.NewRequest(echo.POST, "https://github.com/login/oauth/access_token", buf)
		if err != nil {
			err = fmt.Errorf("access_token request making: %v", err)
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Accept", "application/json")

		res, err := client.Do(req)
		if err != nil {
			err = fmt.Errorf("access_token requesting: %v", err)
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		resBody := echo.Map{}
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			err = fmt.Errorf("access_token response json decoding: %v", err)
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		log.Println("access_token:", resBody)

		if err, exist := resBody["error"]; exist {
			err = fmt.Errorf("access token response has error: %v", err)
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
			err = fmt.Errorf("get user data request making: %v", err)
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		query := req.URL.Query()
		query.Add("access_token", accessToken.(string))
		req.URL.RawQuery = query.Encode()

		res, err = client.Do(req)
		if err != nil {
			err = fmt.Errorf("get user data requesting: %v", err)
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		resBody = echo.Map{}
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			err = fmt.Errorf("get user data response json decoding: %v", err)
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		log.Println("user data:", resBody)

		if err, exist := resBody["error"]; exist {
			err = fmt.Errorf("get user data response has error: %v", err)
			log.Println(err)
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

	if result := h.DB.First(dbUser); result.Error != nil {
		if result.RecordNotFound() {
			if err := h.DB.Create(dbUser).Error; err != nil {
				err = fmt.Errorf("create user: %v", err)
				log.Println(err)
				return echo.NewHTTPError(http.StatusBadRequest, err)
			}
		} else {
			err := fmt.Errorf("find user: %v", result.Error)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}

	if err := jsonUser.Ensure(); err != nil {
		err = fmt.Errorf("jwt user claim ensure: %v", err)
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jsonUser)

	tokenString, err := token.SignedString(h.SecretKey)
	if err != nil {
		err = fmt.Errorf("jwt token to string: %v", err)
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "JWT"
	cookie.Value = tokenString
	cookie.Expires = time.Unix(jsonUser.ExpiresAt, 0)
	cookie.Path = "/"

	c.SetCookie(cookie)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h Handler) AuthRequired() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &UserClaim{},
		SigningKey: h.SecretKey,
		AuthScheme: "JWT",
	}
	return middleware.JWTWithConfig(config)
}
