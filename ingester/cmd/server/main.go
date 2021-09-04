package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/andresusanto/aws-elastic-search-demo/ingester/internal/config"
	"github.com/andresusanto/aws-elastic-search-demo/ingester/internal/es"
	"github.com/andresusanto/aws-elastic-search-demo/ingester/internal/logging"
	"github.com/andresusanto/aws-elastic-search-demo/ingester/pkg/events"
	"github.com/andresusanto/aws-elastic-search-demo/ingester/pkg/healthcheck"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.New()
	r := gin.New()

	logging.Setup(true, true)
	logging.Attach(r)

	es, err := es.NewClient(cfg)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to elastic search")
	}

	h := healthcheck.RegisterHandler(r)
	events.RegisterHandler(r, events.NewService(events.NewRepository(es)))

	go func() {
		err := r.Run(fmt.Sprintf(":%d", cfg.Port))
		log.Fatal().Err(err).Msg("cannot start HTTP server")
	}()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Info().Msg("SIGTERM received. Performing graceful shutdown.")
	h.Shutdown()
}
