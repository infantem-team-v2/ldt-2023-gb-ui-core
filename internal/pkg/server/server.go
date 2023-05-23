package server

import (
	"fmt"
	"gb-auth-gate/config"
	_ "gb-auth-gate/docs"
	"gb-auth-gate/internal/pkg/dependency"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp"
	"gb-auth-gate/pkg/thttp/server"
	"gb-auth-gate/pkg/tlogger"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	mwLogger "github.com/gofiber/fiber/v2/middleware/logger"
	mwRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"os"
)

type Server struct {
	container *dependency.TDependencyContainer
	App       *fiber.App      `di:"app"`
	Config    *config.Config  `di:"config"`
	Logger    tlogger.ILogger `di:"logger"`
}

// NewServer (Fabric) Builds server with DI container which contains main pkg Singletons,
// which we use to build other entities
func NewServer() *Server {
	ctn := dependency.
		NewDependencyContainer().
		BuildDependencies().
		BuildContainer()
	s := &Server{
		container: ctn,
	}
	ctn.Inject(s)
	terrors.Init()
	return s
}

func (s *Server) mapHandlers() {
	for _, hFunc := range emptyHandlers {
		h := hFunc(s.App)
		s.container.Inject(h)
		server.MapRoutes(h)
	}
}

// MapHandlers with middlewares and routers
func (s *Server) MapHandlers() *Server {
	// Getting dependencies from container
	sh := s.container.ContainerInstance().
		Get("stacktraceHandler").(*terrors.StacktraceHandler)

	// Make recover on top of app's stack
	s.App.Use(mwRecover.New(mwRecover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: sh.Handle,
	}))

	s.App.Use(requestid.New(requestid.Config{
		Header: thttp.RequestIdHeader,
	}))
	// Logging fiber's info about requests
	s.App.Use(mwLogger.New(mwLogger.Config{
		Output: os.Stdout,
		Format: fmt.Sprintf("${time} | ${magenta} [${respHeader:%s] ${white} | ${latency} | ${status} - ${method} ${path}\n",
			thttp.RequestIdHeader),
	}))

	// Swagger docs on /docs
	s.App.Get("/docs/*", swagger.HandlerDefault)

	// Prometheus for fiber app metrics
	pmth := fiberprometheus.New(s.Config.BaseConfig.Service.Name)
	pmth.RegisterAt(s.App, "/metrics")
	s.App.Use(pmth.Middleware)

	// Cross-Origin politics
	s.App.Use(cors.New(cors.Config{
		AllowOrigins: "https://ldt2023.infantem.tech",
		AllowHeaders: "*",
	}))

	//ah.SetRoutes()
	s.mapHandlers()
	return s
}

// Run app on tconfig.BaseConfig.System.Port
func (s *Server) Run() error {
	s.Logger.Infof("STARTED SERVER")
	return s.App.Listen(fmt.Sprintf("%s:%s", s.Config.BaseConfig.System.Host, s.Config.BaseConfig.System.Port))
}
