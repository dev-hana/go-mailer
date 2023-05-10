package routers

import (
	"context"
	"fmt"
	"time"

	"github.com/dev-hana/go-mailer/conf"
	"github.com/dev-hana/go-mailer/docs"
	"github.com/dev-hana/go-mailer/services"
	"github.com/gin-gonic/gin"
	"github.com/procyon-projects/chrono"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

func RunAPI() error {
	// Config
	port, mode, dbInit, err := conf.GetServerConfig()
	if err != nil {
		return err
	}

	if mode {
		gin.SetMode(gin.ReleaseMode) //DEFAULT: DebugMode
	}

	dbms, dsn, err := conf.GetDBConfig()
	if err != nil {
		return err
	}

	r := gin.Default()
	r.Use(LoggerMiddleware())
	r.Use(gin.Recovery())

	docs.SwaggerInfo.Title = "[MSA Setting] 메일 전송 자동화"
	docs.SwaggerInfo.Description = "각 프로젝트에서 메일 전송을 자동화 시키기 위한 샘플입니다."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/v1"

	h, err := services.NewHandler(dbms, dsn)
	if err != nil {
		return err
	}

	if dbInit {
		if err := h.InitTable(); err != nil {
			return err
		}
	}

	v1Group := r.Group("/v1")
	v1Group.Use(h.CheckServerConnection)
	{
		v1Group.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1Group.GET("/ping", h.CheckPing)

		mailGroup := v1Group.Group("/mails")
		{
			mailGroup.POST("", h.CreateMail)
		}
	}

	second, err := conf.GetSchedulerConfig()
	task := chrono.NewDefaultTaskScheduler()
	if _, err := task.ScheduleAtFixedRate(func(ctx context.Context) {
		h.SendMailScheduler()
	}, time.Duration(second)*time.Second); err != nil {
		return err
	}

	r.Run(fmt.Sprintf(":%d", port))

	return nil
}
