package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func CreateJob(context echo.Context) error {
	return context.String(http.StatusOK, "Job Created")
}
