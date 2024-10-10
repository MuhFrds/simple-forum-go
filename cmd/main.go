package main

import (
	"log"

	"github.com/MuhFrds/simple-forum-go/internal/configs"
	"github.com/MuhFrds/simple-forum-go/internal/handlers/memberships"
	"github.com/MuhFrds/simple-forum-go/pkg/internalsql"
	"github.com/gin-gonic/gin"

	membershipRepo "github.com/MuhFrds/simple-forum-go/internal/repository/memberships"
	membershipSvc "github.com/MuhFrds/simple-forum-go/internal/service/memberships"
)

func main() {
    r := gin.Default()

    var (
        cfg *configs.Config
    )

    err := configs.Init(
        configs.WithConfigFolder(
            []string{"./internal/configs"},
        ),
        configs.WithConfigFile("config"),
        configs.WithConfigType("config.yaml"),
    )

    if err != nil {
        log.Fatal("Gagal inisiasi config", err)
    }

    cfg = configs.Get()

    db, err := internalsql.Connect(cfg.Database.DataSourceName)
    if err != nil {
        log.Fatal("Gagal inisiasi database", err)
    }

    membershipRepo := membershipRepo.NewRepository(db)
    membershipService := membershipSvc.NewService(cfg, membershipRepo)

    membershipsHandler := memberships.NewHandler(r, membershipService)
    membershipsHandler.RegisterRoute()

    r.Run(cfg.Service.Port)
}
