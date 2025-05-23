package handler

import (
	"dingospeed/internal/model"
	"dingospeed/internal/service"
	"dingospeed/pkg/app"
	"dingospeed/pkg/config"
	"dingospeed/pkg/util"

	"github.com/labstack/echo/v4"
)

type SysHandler struct {
	sysService *service.SysService
}

func NewSysHandler(sysService *service.SysService) *SysHandler {
	return &SysHandler{
		sysService: sysService,
	}
}

func (s *SysHandler) Info(c echo.Context) error {
	info := &model.SystemInfo{}
	if appInfo, ok := app.FromContext(c.Request().Context()); ok {
		info.Id = appInfo.ID()
		info.Name = appInfo.Name()
		info.Version = appInfo.Version()
		info.StartTime = appInfo.StartTime()
	}
	info.HfNetLoc = config.SysConfig.GetHfNetLoc()
	return util.ResponseData(c, info)
}
