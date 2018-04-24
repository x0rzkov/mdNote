package handler

import (
	"github.com/labstack/echo"
	"mdNote/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"fmt"
)

func (h Handler) GetNote(c echo.Context) error {
	claim := c.Get("user").(*jwt.Token).Claims.(*UserClaim)
	note := model.Note{
		UserID: claim.Token,
	}

	if id := c.Param("id"); id == "" {
		return c.NoContent(http.StatusBadRequest)
	} else {
		note.ID = id
	}

	if result := h.DB.First(&note); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	} else if result.RecordNotFound() {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": &note,
	})
}

func (h Handler) GetNotes(c echo.Context) error {
	claim := c.Get("user").(*jwt.Token).Claims.(*UserClaim)

	if category := c.Param("category"); category == "" {
		notes := []model.Note{}

		if result := h.DB.Select("id, user_id, category, title, created_at").Order("created_at desc").Where(&model.Note{UserID: claim.Token}).Find(&notes); result.Error != nil {
			return echo.NewHTTPError(http.StatusBadRequest, result.Error)
		} else if result.RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"data": &notes,
		})
	} else {
		notes := []model.Note{}

		if result := h.DB.Select("id, user_id, category, title, created_at").Order("created_at desc").Where(&model.Note{UserID: claim.Token, Category: category}).Find(&notes); result.Error != nil {
			return echo.NewHTTPError(http.StatusBadRequest, result.Error)
		} else if result.RecordNotFound() {
			return c.NoContent(http.StatusNoContent)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"data": &notes,
		})
	}
}

func (h Handler) SaveNote(c echo.Context) error {
	claim := c.Get("user").(*jwt.Token).Claims.(*UserClaim)

	note := model.Note{}
	if err := c.Bind(&note); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if note.ID == "" {
		note.UserID = claim.Token

		for {
			note.ID = fmt.Sprintf("%x", GenerateRandomKey(2))
			if h.DB.First(&model.Note{}, "id = ?", note.ID).RecordNotFound() {
				break
			}
		}

		if err := h.DB.Create(&note).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"data": &note,
		})
	} else if h.DB.First(&model.Note{}, "id = ?", note.ID).RecordNotFound() {
		return c.NoContent(http.StatusNotFound)
	}

	if err := h.DB.Save(&note).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": &note,
	})
}

func (h Handler) DeleteNote(c echo.Context) error {
	claim := c.Get("user").(*jwt.Token).Claims.(*UserClaim)
	note := model.Note{
		UserID: claim.Token,
	}

	if id := c.Param("id"); id == "" {
		return c.NoContent(http.StatusBadRequest)
	} else {
		note.ID = id
	}

	if result := h.DB.Delete(&note); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	} else if result.RecordNotFound() {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}