// SPDX-FileCopyrightText: 2025 Shun Sakai
//
// SPDX-License-Identifier: CC0-1.0

package main

import (
	"os"
	"os/exec"
	"runtime"
	"testing"

	"github.com/google/go-cmdtest"
)

func TestCLI(t *testing.T) {
	t.Parallel()

	if runtime.GOOS == "windows" {
		t.SkipNow()
	}

	if err := exec.Command("go", "build").Run(); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.Remove("pagen"); err != nil {
			t.Fatal(err)
		}
	}()

	ts, err := cmdtest.Read("testdata")
	if err != nil {
		t.Fatal(err)
	}

	ts.Commands["pagen"] = cmdtest.Program("pagen")
	ts.Run(t, false)
}
