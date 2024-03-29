// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "runtime.h"
#include "defs.h"

int get_num_cpus(void);

int32
getproccount(void)
{
	return get_num_cpus();
}
