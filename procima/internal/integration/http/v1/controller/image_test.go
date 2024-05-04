package controller

import (
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/integration/http/v1/controller/mocks"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestImageController_ProcessImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceMock := mocks.NewMockImageService(ctrl)
	logger := loggerPkg.SetupLogger("debug", "pretty", "test", "no-host")

	controller := NewImageController(serviceMock, logger, []string{"jpeg", "jpg"}, 10<<20)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{
		{Key: "image", Value: "1"},
	}

	serviceMock.EXPECT().ProcessImage(gomock.Any(), gomock.Any()).Return(nil, nil)

	controller.ProcessImage(c)
}
