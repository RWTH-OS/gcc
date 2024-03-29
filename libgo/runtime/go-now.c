// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stddef.h>
#include <stdint.h>
#include <sys/time.h>

#include "runtime.h"

// Return current time.  This is the implementation of time.now().

struct time_now_ret
now()
{
#ifdef __hermit__
  struct timeval tv;
  struct time_now_ret ret;

  gettimeofday (&tv, NULL);
  ret.sec = tv.tv_sec;
  ret.nsec = tv.tv_usec * 1000;
  return ret;
#else
  struct timespec ts;
  struct time_now_ret ret;

  clock_gettime (CLOCK_REALTIME, &ts);
  ret.sec = ts.tv_sec;
  ret.nsec = ts.tv_nsec;
  return ret;
#endif
}
