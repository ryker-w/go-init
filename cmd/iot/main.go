package main

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/ryker-w/go-init/cmd"
	"github.com/ryker-w/go-init/internal/api"
	"github.com/ryker-w/go-init/internal/db/model"
	"github.com/ryker-w/go-init/internal/etc"
	"github.com/ryker-w/go-init/internal/setup"
	"time"
)

import _ "github.com/lib/pq"

func main() {
	orm.Debug = true

	// 捕获异常 panic
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
	time.Sleep(time.Second * 2)
}

func _main() (err error) {
	// 配置文件的名称
	configName := "config"
	// 初始化应用程序
	application := app.New()
	// 启动应用程序
	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		// 加载配置文件
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
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
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}

		builder.
			EnableDatabase(dbConfig.Build(), // 连接数据库
				new(model.User), // 初始化数据库表
			).
			EnableWeb(etc.Config.Web.Listen, api.Route). // 启动web监听(端口号, router)
			ComponentBefore(setup.SomeWork)	// 初始化任务
		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
