package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/parag08/go-errors/pkg/errors"
	"github.com/parag08/go-errors/pkg/logger"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "weight tracker API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreateSandwich() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		ingredients, err := s.sanwdwitchMakerService.New()

		if err != nil {
			fmt.Println("here", err)
			logger.SystemErr(err)
			if intErr, ok := err.(*errors.Error); ok {
				c.JSON(int(errors.GetKind(*intErr)), nil)
				return
			}
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		ingredientsString := strings.Join(ingredients, ",")

		fmt.Println("here", ingredients)
		response := map[string]string{
			"status":      "success",
			"data":        "new sandwich created",
			"ingredients": ingredientsString,
		}

		c.JSON(http.StatusOK, response)
		return
	}
}
