package blog

import (
	"context"
	"errors"
	"fmt"
	"github.com/caijunduo/go-scaffold/internal/blog/store"
	"github.com/caijunduo/go-scaffold/internal/pkg/config"
	"github.com/caijunduo/go-scaffold/internal/pkg/helper"
	"github.com/caijunduo/go-scaffold/internal/pkg/log"
	"github.com/caijunduo/go-scaffold/pkg/driver"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile string

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "blog",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			//verflag.PrintAndExitIfRequested()

			log.Init(helper.GetLogOption())
			defer log.Sync()

			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	cobra.OnInitialize(func() {
		config.New(config.Option{
			File:       configFile,
			EnvPrefix:  "BLOG",
			ConfigName: "blog.yaml",
		})
	})

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

func run() error {
	if err := initStore(); err != nil {
		return err
	}

	// 设置 Gin 模式
	gin.SetMode(viper.GetString("mode"))

	// 创建 Gin 引擎
	g := gin.New()

	mws := []gin.HandlerFunc{gin.Recovery()}
	g.Use(mws...)

	if err := installRouters(g); err != nil {
		return err
	}

	// 创建 HTTP Server 实例
	httpsrv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}

	// 运行 HTTP 服务器
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排障
	log.Info("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())
		}
	}()

	// 等待中断信号优雅地关闭服务器（10 秒超时)。
	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Info("Shutting down server ...")

	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Error("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Info("Server exiting")
	return nil
}

func initStore() error {
	db, err := driver.NewMySQL(helper.GetDBOption())
	if err != nil {
		return err
	}

	store.Init(db)

	return nil
}
