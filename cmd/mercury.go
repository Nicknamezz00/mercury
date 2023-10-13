/*
 * MIT License
 *
 * Copyright (c) 2023 Runze Wu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nicknamezz00/mercury/common/log"
	"github.com/Nicknamezz00/mercury/server"
	_profile "github.com/Nicknamezz00/mercury/server/profile"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/Nicknamezz00/mercury/store/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	profile *_profile.Profile
	mode    string
	addr    string
	port    int
	data    string
	driver  string
	dsn     string
)

var rootCmd = &cobra.Command{
	Use:   "mercury",
	Short: "A light-weight note taking service written in Go & TypeScript",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		dbDriver, err := db.NewDBDriver(profile)
		if err != nil {
			log.Error("failed to create db driver", zap.Error(err))
			return
		}
		if err := dbDriver.Migrate(ctx); err != nil {
			log.Error("failed to migrate db driver", zap.Error(err))
			return
		}

		s := store.New(dbDriver, profile)
		svr, err := server.New(ctx, profile, s)
		if err != nil {
			log.Error("failed to create server", zap.Error(err))
			return
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			sig := <-c
			log.Info(fmt.Sprintf("receive signal: %s", sig.String()))
			svr.Shutdown(ctx)
			cancel()
		}()

		greeting()

		if err := svr.Start(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Error("failed to start server", zap.Error(err))
				return
			}
		}

		// Wait for CTRL + C
		<-ctx.Done()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "demo", `mode of server, can be "prod" or "dev" or "demo"`)
	rootCmd.PersistentFlags().StringVarP(&addr, "addr", "a", "", "address of server")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8081, "port of server")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "data directory")
	rootCmd.PersistentFlags().StringVarP(&driver, "driver", "", "", "database driver")
	rootCmd.PersistentFlags().StringVarP(&dsn, "dsn", "", "", "database source name(aka. DSN)")

	err := viper.BindPFlag("mode", rootCmd.PersistentFlags().Lookup("mode"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("data", rootCmd.PersistentFlags().Lookup("data"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("driver", rootCmd.PersistentFlags().Lookup("driver"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("dsn", rootCmd.PersistentFlags().Lookup("dsn"))
	if err != nil {
		panic(err)
	}

	viper.SetDefault("mode", "demo")
	viper.SetDefault("driver", "sqlite")
	viper.SetDefault("addr", "")
	viper.SetDefault("port", 8081)
	viper.SetEnvPrefix("mercury")
}

func initConfig() {
	viper.AutomaticEnv()
	var err error
	profile, err = _profile.GetProfile()
	if err != nil {
		fmt.Printf("failed to get profile, error: %+v\n", err)
		return
	}

	println("---")
	println("Server profile")
	println("data:", profile.Data)
	println("dsn:", profile.DSN)
	println("addr:", profile.Addr)
	println("port:", profile.Port)
	println("mode:", profile.Mode)
	println("driver:", profile.Driver)
	println("version:", profile.Version)
	println("---")
}

func Execute() error {
	defer log.Sync()
	return rootCmd.Execute()
}

func greeting() {
	if len(profile.Addr) == 0 {
		fmt.Printf("Version %s has been started on port %d\n", profile.Version, profile.Port)
	} else {
		fmt.Printf("Version %s has been started on address '%s' and port %d\n", profile.Version, profile.Addr, profile.Port)
	}
	println("---")
	println("Hello! Server started!")
}
