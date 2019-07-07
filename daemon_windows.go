// Copyright 2019 Janos Guljas. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package resenje.org/daemon provides functionality to execute binaries
// in the background. It requires no external dependencies.

// +build windows

package daemon

import "errors"

func setSid() (s int, err error) {
	return 0, errors.New("unable to daemonize")
}
