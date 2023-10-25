package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/amqp"
	"github.com/lishimeng/app-starter/cache"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/mqtt"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-init/cmd"
	"github.com/ryker-w/go-init/cmd/go-init/ddd"
	"github.com/ryker-w/go-init/cmd/go-init/static"
	"github.com/ryker-w/go-init/internal/db"
	"github.com/ryker-w/go-init/internal/etc"
	"github.com/ryker-w/go-init/internal/process"
	"github.com/ryker-w/go-init/internal/setup"

	_ "github.com/lib/pq"
)

func main() {
	// 捕获panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 打印启动信息
	fmt.Println(cmd.AppName)
	fmt.Println(cmd.Version)

	// 主程序
	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	// 等待退出
	time.Sleep(time.Second * 2)
}

func _main() (err error) {
	// 配置文件的名称
	configName := "config"
	// 初始化应用程序
	application := app.New()
	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) (e error) {
		// 加载配置文件、环境变量
		e = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}

		// 数据库连接配置
		dbConfig := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true, // 初始化表
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}
		// 缓存配置
		redisOpts := cache.RedisOptions{
			Addr:     etc.Config.Redis.Addr,
			Password: etc.Config.Redis.Password,
		}
		cacheOpts := cache.Options{
			MaxSize: 10000,
			Ttl:     time.Hour * 24,
		}
		// Web filter、Token jwt
		if etc.Config.Token.Enable {
			issuer := etc.Config.Token.Issuer
			tokenKey := []byte(etc.Config.Token.Key)
			builder = builder.EnableTokenValidator(func(inject app.TokenValidatorInjectFunc) {
				provider := token.NewJwtProvider(token.WithIssuer(issuer),
					token.WithKey(tokenKey, tokenKey), // hs256的秘钥必须是[]byte
					token.WithAlg("HS256"),
					token.WithDefaultTTL(time.Duration(etc.Config.Token.Ttl)*time.Hour),
				)
				storage := token.NewLocalStorage(provider)
				factory.Add(provider)
				inject(storage)
			})
		}

		// 启动服务（按需添加）
		builder.
			// 数据库服务
			EnableDatabase(dbConfig.Build(), db.RegisterTables()...).
			// 打印数据库操作log
			EnableOrmLog().
			// Web 服务
			EnableWeb(etc.Config.Web.Listen, ddd.Router).
			// Web 日志等级
			SetWebLogLevel("debug").
			// UI 编译资源。前后端打包部署。
			EnableStaticWeb(func() http.FileSystem {
				return http.FS(static.Static)
			}).
			// Amqp 服务
			EnableAmqp(amqp.Connector{Conn: etc.Config.Amqp.Conn}).
			// 注册Amqp handler
			RegisterAmqpHandlers(&process.AmqpConsumer{}).
			// Mqtt 服务
			EnableMqtt(mqtt.WithBroker(etc.Config.Mqtt.Broker),
				mqtt.WithAuth(etc.Config.Mqtt.UserName, etc.Config.Mqtt.Password),
				mqtt.WithRandomClientId(),
				mqtt.WithOnConnectHandler(process.OnConnectHandler), // 连接成功时
				mqtt.WithOnLostHandler(process.OnLostHandler),       // 断开连接时
			).
			// Redis 服务
			EnableCache(redisOpts, cacheOpts).
			// 启动时执行任务。Web服务开启前
			ComponentBefore(setup.SomeWork).
			// 启动时执行任务。Web服务开启后
			ComponentAfter(setup.SomeWork)
		return
	}, func(s string) {
		// 程序退出时
		log.Info(s)
	})

	return
}
