package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aerogo/aero"
	"github.com/klausklapper/twitchberry/internal/stream"
	"github.com/klausklapper/twitchberry/internal/web"
)

func ListenAndServe(ctx context.Context) {
	s := stream.New()
	defer s.Stop()

	app := aero.New()
	defer app.Shutdown()
	app.Config.GZip = true

	app.Get("/", func(ctx aero.Context) error {
		fl, err := web.Content.ReadFile("index.html")
		if err != nil {
			panic(fmt.Sprintf("failed to load index: %v", err))
		}
		return ctx.Bytes(fl)
	})

	app.Get("/*file", func(ctx aero.Context) error {
		fl, err := web.Content.ReadFile(ctx.Get("file"))
		if err != nil {
			return ctx.Error(404)
		}

		return ctx.Bytes(fl)
	})

	app.Get("/api/v1/twitch", func(ctx aero.Context) error {
		return ctx.JSON(map[string]interface{}{
			"isStreaming": false,
			"started":     time.Now(),
		})
	})

	app.Post("/api/v1/twitch", func(ctx aero.Context) error {
		// todo
		return s.Start()
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Shutdown requested, stopping web server ...")
				app.Shutdown()
			}
		}
	}()

	app.Run()
}
