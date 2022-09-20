package handler

import(
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/iamstivgo/integrador-back-go.git/internal/domain"
	"github.com/iamstivgo/integrador-back-go.git/internal/patient"
	"github.com/iamstivgo/integrador-back-go.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service patient.Service
}

func NewPatientHandler(service patient.Service) *PatientHandler {
	return &PatientHandler{service: service}
}

func (h *PatientHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.service.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		c.JSON(200, patient)
	}
}

func (h *PatientHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		err := c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid body"))
			return
		}
		patient, err = h.service.Create(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, patient)
	}
}

func (h *PatientHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.service.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		var patient domain.Patient
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid body"))
			return
		}
		patient, err = h.service.Update(id, patient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, patient)
	}
}

func (h *PatientHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.service.Delete(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 204, nil)
	}
}