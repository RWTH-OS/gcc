// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// BSD library calls.

package syscall

import (
	//"internal/race"
	//"unsafe"
)

func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	return 0, EINVAL
	/*if race.Enabled {
		race.ReleaseMerge(unsafe.Pointer(&ioSync))
	}
	var soff Offset_t
	var psoff *Offset_t
	if offset != nil {
		soff = Offset_t(*offset)
		psoff = &soff
	}
	written, err = sendfile(outfd, infd, psoff, count)
	if offset != nil {
		*offset = int64(soff)
	}
	return*/
}
