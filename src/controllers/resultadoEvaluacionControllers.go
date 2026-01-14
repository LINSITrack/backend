package controllers

import (
	"net/http"
	"strconv"

	"github.com/LINSITrack/backend/src/models"
	"github.com/LINSITrack/backend/src/services"
	"github.com/gin-gonic/gin"
)

type ResultadoEvaluacionController struct {
	resultadoService *services.ResultadoEvaluacionService
}

func NewResultadoEvaluacionController(resultadoService *services.ResultadoEvaluacionService) *ResultadoEvaluacionController {
	return &ResultadoEvaluacionController{resultadoService: resultadoService}
}

func (c *ResultadoEvaluacionController) GetAllResultados(ctx *gin.Context) {
	resultados, err := c.resultadoService.GetAllResultados()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resultados)
}

func (c *ResultadoEvaluacionController) GetResultadoByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	resultado, err := c.resultadoService.GetResultadoByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resultado)
}

func (c *ResultadoEvaluacionController) GetResultadosByAlumnoID(ctx *gin.Context) {
	alumnoID, err := strconv.Atoi(ctx.Param("alumnoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumno ID"})
		return
	}

	resultados, err := c.resultadoService.GetResultadosByAlumnoID(alumnoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultados)
}

func (c *ResultadoEvaluacionController) GetResultadosByEvaluacionID(ctx *gin.Context) {
	evaluacionID, err := strconv.Atoi(ctx.Param("evaluacionId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid evaluacion ID"})
		return
	}

	resultados, err := c.resultadoService.GetResultadosByEvaluacionID(evaluacionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultados)
}

func (c *ResultadoEvaluacionController) CreateResultado(ctx *gin.Context) {
	var resultado models.ResultadoEvaluacion
	if err := ctx.ShouldBindJSON(&resultado); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	// Validaciones requeridas
	if resultado.AlumnoID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "el ID del alumno es obligatorio"})
		return
	}
	if resultado.EvaluacionID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "el ID de la evaluación es obligatorio"})
		return
	}

	if err := c.resultadoService.CreateResultado(&resultado); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resultado)
}

func (c *ResultadoEvaluacionController) UpdateResultado(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateRequest models.ResultadoEvaluacionUpdateRequest
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos",
			"details": err.Error(),
		})
		return
	}

	// Validaciones para campos que no pueden estar vacíos si se envían
	if updateRequest.Nota != nil && (*updateRequest.Nota < 0 || *updateRequest.Nota > 10) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "la nota debe estar entre 0 y 10"})
		return
	}
	if updateRequest.Devolucion != nil && *updateRequest.Devolucion == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "la devolución no puede estar vacía"})
		return
	}
	if updateRequest.Observaciones != nil && *updateRequest.Observaciones == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "las observaciones no pueden estar vacías"})
		return
	}

	resultado, err := c.resultadoService.UpdateResultado(id, &updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resultado)
}

func (c *ResultadoEvaluacionController) DeleteResultado(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.resultadoService.DeleteResultado(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
