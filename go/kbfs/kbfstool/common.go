// Copyright 2016 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/keybase/client/go/kbfs/idutil"
	"github.com/keybase/client/go/kbfs/tlf"
)

const (
	topName     = "keybase"
	publicName  = "public"
	privateName = "private"
)

const publicSuffix = tlf.ReaderSep + idutil.PublicUIDName

func byteCountStr(n int) string {
	if n == 1 {
		return "1 byte"
	}
	return fmt.Sprintf("%d bytes", n)
}

func printError(prefix string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", prefix, err)
}
