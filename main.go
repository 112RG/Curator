package main

import (
	"net/http"
	"os"
	"time"

	"curator/injection"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().
		Timestamp().Logger()
	router, err := injection.Inject()
	configureLogHandler(router)
	if err != nil {
		log.Error().Err(err)
	}
	log.Info().Msg("Startup complete. Serving requests")
	http.ListenAndServe(":5000", handlers.RecoveryHandler()(router))
}

func configureLogHandler(r *mux.Router) {
	logHandler := hlog.NewHandler(log.Logger)
	r.Use(logHandler)
	r.Use(hlog.AccessHandler((func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Dur("duration", duration).
			Msg("")
	})))
	r.Use(hlog.RemoteAddrHandler("ip"))
	r.Use(hlog.UserAgentHandler("user_agent"))
	r.Use(hlog.RefererHandler("referer"))
	//r.Use(hlog.RequestIDHandler("req_id", "X-Request-Id"))
}
