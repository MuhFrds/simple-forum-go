package main

import (
	"github.com/MuhFrds/go-101/tree/main/fastcampus/fastcampuss/internal/handler/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  
  membershipsHandler := memberships.NewHandler(r)
  membershipsHandler.RegisterRoute()

  r.Run(":9999")
}