package handler

import(
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/iamstivgo/integrador-back-go.git/internal/domain"
	"github.com/iamstivgo/integrador-back-go.git/internal/dentist"
	"github.com/iamstivgo/integrador-back-go.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	service dentist.Service
}

func NewDentistHandler(service dentist.Service) *DentistHandler {
	return &DentistHandler{service: service}
}

func (h *DentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentist, err := h.service.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		c.JSON(200, dentist)
	}
}

func (h *DentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid body"))
			return
		}
		dentist, err = h.service.Create(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, dentist)
	}
}

func (h *DentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.service.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		var dentist domain.Dentist
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid body"))
			return
		}
		p, err := h.service.Update(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *DentistHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.service.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}
