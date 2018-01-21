package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/zhanglindeng/goweb/config"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/route"
	"github.com/zhanglindeng/goweb/schedule"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// true: goweb.exe -migrate
	migrate := flag.Bool("migrate", false, "migrate database")
	flag.Parse()

	r, err := route.Create()
	if err != nil {
		log.Fatalln(err)
	}

	// migrate database
	if *migrate {
		log.Println("database migrating...")
		if err := model.Migrate(); err != nil {
			log.Fatalln(err)
		}
		log.Println("database migrated")
	}

	// schedule
	if err := schedule.Create(); err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Addr:    config.AppPort,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
		log.Println("server listen port:", config.AppPort)
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 30 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
