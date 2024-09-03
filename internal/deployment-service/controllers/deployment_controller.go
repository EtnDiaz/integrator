package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/roneeSoft/integrator/pkg/shared/utils"
	"log"
	"net/http"
)

type DeploymentController struct{}

func (ctrl DeploymentController) Deploy(c *gin.Context) {

	releaseName := c.PostForm("releaseName")
	chartPath := c.PostForm("chartPath")

	// todo helm integration for deployment.
	log.Printf("Deploying release: %s using chart: %s", releaseName, chartPath)

	c.JSON(http.StatusOK, utils.Response{Message: "Deployment started", Data: releaseName})
}
