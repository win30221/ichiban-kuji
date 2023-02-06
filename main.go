package main

import (
	"fmt"

	"github.com/win30221/core/basic"
	"github.com/win30221/core/config"
	commonDelivery "github.com/win30221/core/http/delivery"
	"github.com/win30221/core/http/middleware"
	"github.com/win30221/core/storage"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	_ "github.com/win30221/ichiban-kuji/docs"

	"github.com/gin-gonic/gin"
	"github.com/win30221/ichiban-kuji/service/delivery"
	"github.com/win30221/ichiban-kuji/service/repository"
	"github.com/win30221/ichiban-kuji/service/usecase"
)

const (
	ServerName = "ichiban-kuji"
)

// build version
// go build 時使用 -ldflags 傳入
var (
	VERSION   string
	COMMIT    string
	BUILDTIME string
)

// @title ichiban-kuji 模組
// @description 使用者管理
// @securityDefinitions.apikey Systoken
// @in header
// @name Systoken

func main() {
	basic.Version = VERSION
	basic.Commit = COMMIT
	basic.BuildTime = BUILDTIME

	basic.Init(ServerName)

	// 建立 HTTP server
	r := gin.New()

	// 設定 http server Middleware
	r.Use(gin.Recovery())
	r.Use(middleware.Log()...)

	// 設定基礎路由
	_, privateGroup := commonDelivery.SetBasicRouter(r)

	mongoM := storage.GetMongoDB("/storage/mongo/master/celluloid-picket", nil)
	mongoS := storage.GetMongoDB("/storage/mongo/slave/celluloid-picket", readpref.SecondaryPreferred())
	redis := storage.GetRedis("/storage/redis/celluloid-picket", basic.ServerName)

	// ttl
	userIchibanKujiTTL, _ := config.GetSeconds("/service/ichiban-kuji/user_ichiban_kuji_ttl", true)
	ichibanKujiTTL, _ := config.GetSeconds("/service/ichiban-kuji/ichiban_kuji_ttl", true)
	ichibanKujiRewardsTTL, _ := config.GetSeconds("/service/ichiban-kuji/ichiban_kuji_rewards_ttl", true)
	ichibanKujiRewardTTL, _ := config.GetSeconds("/service/ichiban-kuji/ichiban_kuji_reward_ttl", true)

	// repo
	mongoRepo := repository.NewMongoRepo(mongoM, mongoS)
	redisRepo := repository.NewRedisRepo(
		redis,
		userIchibanKujiTTL,
		ichibanKujiTTL,
		ichibanKujiRewardsTTL,
		ichibanKujiRewardTTL,
	)

	// use case
	useCase := usecase.New(
		mongoRepo,
		redisRepo,
	)

	delivery.New(privateGroup, useCase)

	r.Run(fmt.Sprintf("%s:%s", basic.Host, basic.Port))
}
