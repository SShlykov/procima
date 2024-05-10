package images

import (
	v1 "github.com/SShlykov/procima/procima/internal/app/http/v1"
	"github.com/gin-gonic/gin"
)

func (ic *imageController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST(v1.ImageUploadURL, ic.ProcessImage)
}
