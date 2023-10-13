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

package profile

import (
	"fmt"
	"github.com/Nicknamezz00/mercury/internal/constants"
	"github.com/Nicknamezz00/mercury/server/version"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Profile struct {
	Mode    string `json:"mode"`
	Addr    string `json:"-"`
	Port    int    `json:"-"`
	Data    string `json:"-"`
	DSN     string `json:"-"`
	Driver  string `json:"driver"`
	Version string `json:"version"`
}

func (p *Profile) IsDev() bool {
	return p.Mode != constants.PRODUCTION
}

// GetProfile will return a profile for dev or production.
func GetProfile() (*Profile, error) {
	profile := Profile{}
	err := viper.Unmarshal(&profile)
	if err != nil {
		return nil, err
	}

	if profile.Mode != constants.DEMO && profile.Mode != constants.DEV && profile.Mode != constants.PRODUCTION {
		profile.Mode = constants.DEMO
	}

	if profile.Mode == constants.PRODUCTION && profile.Data == "" {
		if runtime.GOOS == "windows" {
			profile.Data = filepath.Join(os.Getenv("ProgramData"), "mercury")

			if _, err := os.Stat(profile.Data); os.IsNotExist(err) {
				if err := os.MkdirAll(profile.Data, 0770); err != nil {
					fmt.Printf("Failed to create data directory: %s, err: %+v\n", profile.Data, err)
					return nil, err
				}
			}
		} else {
			profile.Data = "/var/opt/mercury"
		}
	}

	dataDir, err := checkDataDir(profile.Data)
	if err != nil {
		fmt.Printf("Failed to check dsn: %s, err: %+v\n", dataDir, err)
		return nil, err
	}

	profile.Data = dataDir
	if profile.Driver == constants.SQLITE && profile.DSN == "" {
		dbFile := fmt.Sprintf("mercury_%s.db", profile.Mode)
		profile.DSN = filepath.Join(dataDir, dbFile)
	}
	profile.Version = version.GetCurrentVersion(profile.Mode)

	return &profile, nil
}

func checkDataDir(dataDir string) (string, error) {
	// Convert to absolute path if relative path is supplied.
	if !filepath.IsAbs(dataDir) {
		relativeDir := filepath.Join(filepath.Dir(os.Args[0]), dataDir)
		absDir, err := filepath.Abs(relativeDir)
		if err != nil {
			return "", err
		}
		dataDir = absDir
	}
	// Trim trailing \ or / in case user supplies
	dataDir = strings.TrimRight(dataDir, "\\/")
	if _, err := os.Stat(dataDir); err != nil {
		return "", errors.Wrapf(err, "unable to access data folder %s", dataDir)
	}
	return dataDir, nil
}
