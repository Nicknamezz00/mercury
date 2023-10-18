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
	"github.com/Nicknamezz00/mercury/common/util"
	"github.com/Nicknamezz00/mercury/internal/types"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/Nicknamezz00/mercury/store/db/sqlite"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	setupCmdFlagHostUsername = "username"
	setupCmdFlagHostPassword = "password"
	setupCmd                 = &cobra.Command{
		Use:   "setup",
		Short: "Initial setup",
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
			driver, err := sqlite.NewDB(profile)
			if err != nil {
				fmt.Printf("failed to create db driver, error: %+v\n", err)
				return
			}
			if err := driver.Migrate(ctx); err != nil {
				fmt.Printf("failed to migrate db, error: %+v\n", err)
				return
			}
			s := store.New(driver, profile)
			if err := setup(ctx, s, hostUsername, hostPassword); err != nil {
				fmt.Printf("failed to setup, error: %v\n", err)
				return
			}
		},
	}
)

func init() {
	setupCmd.Flags().String(setupCmdFlagHostUsername, "", "Owner username")
	setupCmd.Flags().String(setupCmdFlagHostPassword, "", "Owner password")
	rootCmd.AddCommand(setupCmd)
}

type setupService struct {
	store *store.Store
}

func (s setupService) setup(ctx context.Context, username string, password string) error {
	if err := s.makeSureHostUserNotExists(ctx); err != nil {
		return err
	}
	if err := s.createUser(ctx, username, password); err != nil {
		return errors.Wrap(err, "create user error")
	}
	return nil
}

func (s setupService) createUser(ctx context.Context, username string, password string) error {
	userCreate := &store.User{
		Username: username,
		// The new signup user should be normal user by default.
		Role:     types.RoleHost,
		Nickname: username,
	}

	if len(userCreate.Username) < 3 {
		return errors.New("username is too short, minimum length is 3")
	}
	if len(userCreate.Username) > 32 {
		return errors.New("username is too long, maximum length is 32")
	}
	if len(password) < 3 {
		return errors.New("password is too short, minimum length is 3")
	}
	if len(password) > 512 {
		return errors.New("password is too long, maximum length is 512")
	}
	if len(userCreate.Nickname) > 64 {
		return errors.New("nickname is too long, maximum length is 64")
	}
	if userCreate.Email != "" {
		if len(userCreate.Email) > 256 {
			return errors.New("email is too long, maximum length is 256")
		}
		if !util.ValidateEmail(userCreate.Email) {
			return errors.New("invalid email format")
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}
	userCreate.PasswordHash = string(passwordHash)
	if _, err := s.store.CreateUser(ctx, userCreate); err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (s setupService) makeSureHostUserNotExists(ctx context.Context) error {
	r := types.RoleHost
	exist, err := s.store.ListUsers(ctx, &store.FindUser{
		Role: &r,
	})
	if err != nil {
		return errors.Wrap(err, "find user list")
	}
	if len(exist) != 0 {
		return errors.New("host user already exists")
	}
	return nil
}

func setup(ctx context.Context, s *store.Store, username string, password string) error {
	svc := setupService{
		store: s,
	}
	return svc.setup(ctx, username, password)
}
