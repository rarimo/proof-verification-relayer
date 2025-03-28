package cli

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/alecthomas/kingpin"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/service/api"
	"github.com/rarimo/proof-verification-relayer/internal/service/checker"
	"github.com/rarimo/proof-verification-relayer/internal/service/listener"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

func Run(args []string) bool {
	defer func() {
		if rvr := recover(); rvr != nil {
			logan.New().WithRecover(rvr).Error("app panicked")
		}
	}()

	var (
		cfg            = config.New(kv.MustFromEnv())
		log            = cfg.Log()
		app            = kingpin.New("proof-verification-relayer", "")
		runCmd         = app.Command("run", "run command")
		serviceCmd     = runCmd.Command("service", "run service")
		migrateCmd     = app.Command("migrate", "migrate command")
		migrateUpCmd   = migrateCmd.Command("up", "migrate db up")
		migrateDownCmd = migrateCmd.Command("down", "migrate db down")
		checkerCmd     = runCmd.Command("checker", "run checker")
	)

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	run := func(f func(context.Context, config.Config)) {
		wg.Add(1)
		go func() {
			f(ctx, cfg)
			wg.Done()
		}()
	}

	switch cmd {
	case serviceCmd.FullCommand():
		run(api.Run)
		run(listener.Run)
		run(checker.Run)
	case migrateUpCmd.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(cfg)
	// handle any custom commands here in the same way
	case checkerCmd.FullCommand():
		run(api.Run)
		run(checker.Run)
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}
	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	wgch := make(chan struct{})
	go func() {
		wg.Wait()
		close(wgch)
	}()

	select {
	case <-ctx.Done():
		cfg.Log().WithError(ctx.Err()).Info("Interrupt signal received")
		stop()
		<-wgch
	case <-wgch:
		cfg.Log().Warn("all services stopped")
	}

	return true
}
