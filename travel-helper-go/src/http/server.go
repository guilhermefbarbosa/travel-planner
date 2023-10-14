package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Server struct {
	app     *fiber.App
	routers map[string]Router
	config  Config
}

type Config struct {
	Port             int
	AppName          string
	TimeoutInSeconds time.Duration
	Routers          []Router
	Handlers         []Handler
}

type Router struct {
	Path        string
	Middlewares []func(*fiber.Ctx) error
}

type Handler interface {
	Routes(routers map[string]Router)
}

func NewServer(sc Config) Server {
	return Server{
		config: sc,
	}
}

func (s Server) Start() {
	s.app = fiber.New(
		fiber.Config{
			AppName:      s.config.AppName,
			IdleTimeout:  s.config.TimeoutInSeconds,
			WriteTimeout: s.config.TimeoutInSeconds,
			ReadTimeout:  s.config.TimeoutInSeconds,
			Immutable:    true,
		})

	routers := make(map[string]Router, len(s.config.Routers))
	for _, eachRouter := range s.config.Routers {
		group := s.app.Group(eachRouter.Path)
		for _, eachMiddleware := range eachRouter.Middlewares {
			group.Use(eachMiddleware)
		}
		routers[eachRouter.Path] = eachRouter
	}

	for _, eachHandler := range s.config.Handlers {
		eachHandler.Routes(s.routers)
	}

	if err := s.app.Listen(fmt.Sprintf(":%d", s.config.Port)); err != nil {
		panic(err)
	}
}
