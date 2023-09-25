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
	"fmt"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/Nicknamezz00/mercury/store/db"
	"github.com/spf13/cobra"
	"time"
)

var (
	setupCmdFlagHostUsername = ""
	setupCmdFlagHostPassword = ""
	setupCmd                 = &cobra.Command{
		Use:   "setup",
		Short: "Setup for mercury",
		Run: func(cmd *cobra.Command, _ []string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			hostUsername, err := cmd.Flags().GetString(setupCmdFlagHostUsername)
			if err != nil {
				fmt.Printf("failed to get host username, error: %v\n", err)
				return
			}
			hostPassword, err := cmd.Flags().GetString(setupCmdFlagHostPassword)
			if err != nil {
				fmt.Printf("failed to get host password, error: %v\n", err)
				return
			}

			db := db.New(profile)
			if err := db.Open(); err != nil {
				fmt.Printf("failed to open db, error: %v\n", err)
				return
			}
			if err := db.Migrate(ctx); err != nil {
				fmt.Printf("failed to migrate db, error: %v\n", err)
				return
			}
			s := store.New(db.DBInstance, profile)
			if err := setup(ctx, s, hostUsername, hostPassword); err != nil {
				fmt.Printf("failed to setup, error: %v\n", err)
				return
			}
		},
	}
)

func setup(ctx context.Context, s *store.Store, username string, password string) error {
	panic("implement me")
}
