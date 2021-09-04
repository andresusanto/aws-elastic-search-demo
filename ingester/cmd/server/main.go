package main

import (
	"fmt"

	"github.com/andresusanto/aws-elastic-search-demo/ingester/internal/config"
	"github.com/andresusanto/aws-elastic-search-demo/ingester/internal/es"
	"github.com/andresusanto/aws-elastic-search-demo/ingester/internal/logging"
	"github.com/andresusanto/aws-elastic-search-demo/ingester/pkg/events"

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

	events.RegisterHandler(r, events.NewService(events.NewRepository(es)))

	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
