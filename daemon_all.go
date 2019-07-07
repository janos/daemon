// Copyright 2019 Janos Guljas. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package resenje.org/daemon provides functionality to execute binaries
// in the background. It requires no external dependencies.

// +build !windows

package daemon

import (
	"fmt"
	"syscall"
)

func setSid() (s int, err error) {
	s, err = syscall.Setsid()
	if err != nil {
		return 0, fmt.Errorf("setsid syscall: %s", err)
	}
	return s, nil
}
