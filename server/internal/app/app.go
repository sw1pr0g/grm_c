package app

import (
	"github.com/sw1pr0g/amazing-people-calendar/server/config"
	"github.com/sw1pr0g/amazing-people-calendar/server/pkg/postgres"
	"log"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatalf("app - Run - postgres.New: %s", err)
	}
	defer pg.Close()

	//handler := gin.New()
	//v1.NewRouter(handler, l, translationUseCase)
	//httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	//
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	//
	//select {
	//case s := <-interrupt:
	//	l.Info("app - Run - signal: " + s.String())
	//case err = <-httpServer.Notify():
	//	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	//
	//err = httpServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	//}
}
