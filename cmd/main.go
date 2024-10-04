package main

import (
	"log"

	"github.com/MuhFrds/simple-forum-go/internal/configs"
	"github.com/MuhFrds/simple-forum-go/internal/handlers/memberships"
	"github.com/MuhFrds/simple-forum-go/pkg/internalsql"
	"github.com/gin-gonic/gin"

	membershipRepo "github.com/MuhFrds/simple-forum-go/internal/repository/memberships"
)

func main() {
  r := gin.Default()

  var (
    cfg * configs.Config
  )

  err:= configs.Init(
    configs.WithConfigFolder(
      []string {"./internal/configs"},
    ),
    configs.WithConfigFile("config"),
    configs.WithConfigType("config.yaml"),
  )

  if err != nil {
    log.Fatal("Gagal inisiasi config", err)
  }

  cfg = configs.Get()
  log.Println("config", cfg)

  db, err:= internalsql.Connect(cfg.Database.DataSourceName)
  if err != nil {
    log.Fatal("Gagal inisiasi database", err)
  }

  _ = membershipRepo.NewRepository(db)
  
  membershipsHandler := memberships.NewHandler(r)
  membershipsHandler.RegisterRoute()

  r.Run(cfg.Service.Port)
}