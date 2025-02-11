package core

import (
	"github.com/rigdev/rig/internal/client"
	"github.com/rigdev/rig/internal/config"
	"github.com/rigdev/rig/internal/gateway"
	"github.com/rigdev/rig/internal/oauth2"
	"github.com/rigdev/rig/internal/repository"
	"github.com/rigdev/rig/internal/service"
	pkg_service "github.com/rigdev/rig/pkg/service"
	"github.com/rigdev/rig/pkg/telemetry"
	"go.uber.org/fx"
)

func GetModule(cfg config.Config) fx.Option {
	return fx.Options(
		fx.Supply(cfg),
		client.GetModule(cfg),
		pkg_service.Module,
		oauth2.Module,
		service.Module,
		repository.Module,
		gateway.Module,
		telemetry.Module,
	)
}
