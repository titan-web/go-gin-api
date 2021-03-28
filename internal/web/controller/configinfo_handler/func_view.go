package configinfo_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (h *handler) View() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("configinfo", configs.Get())
	}
}