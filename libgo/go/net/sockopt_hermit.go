// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"os"
	"syscall"
)

func setDefaultSockopts(s, family, sotype int, ipv6only bool) error {
	if family == syscall.AF_INET6 && sotype != syscall.SOCK_RAW {
		// Allow both IP versions even if the OS default
		// is otherwise.  Note that some operating systems
		// never admit this option.
		syscall.SetsockoptInt(s, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, boolint(ipv6only))
	}
	// Allow broadcast.
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1))
}

func setDefaultListenerSockopts(s int) error {
	// Allow reuse of recently-used addresses.
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
}

func setDefaultMulticastSockopts(s int) error {
	// Allow multicast UDP and raw IP datagram sockets to listen
	// concurrently across multiple listeners.
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
}

// Boolean to int.
func boolint(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ipv4AddrToInterface(ip IP) (*Interface, error) {
	ift, err := Interfaces()
	if err != nil {
		return nil, err
	}
	for _, ifi := range ift {
		ifat, err := ifi.Addrs()
		if err != nil {
			return nil, err
		}
		for _, ifa := range ifat {
			switch v := ifa.(type) {
			case *IPAddr:
				if ip.Equal(v.IP) {
					return &ifi, nil
				}
			case *IPNet:
				if ip.Equal(v.IP) {
					return &ifi, nil
				}
			}
		}
	}
	if ip.Equal(IPv4zero) {
		return nil, nil
	}
	return nil, errNoSuchInterface
}

func interfaceToIPv4Addr(ifi *Interface) (IP, error) {
	if ifi == nil {
		return IPv4zero, nil
	}
	ifat, err := ifi.Addrs()
	if err != nil {
		return nil, err
	}
	for _, ifa := range ifat {
		switch v := ifa.(type) {
		case *IPAddr:
			if v.IP.To4() != nil {
				return v.IP, nil
			}
		case *IPNet:
			if v.IP.To4() != nil {
				return v.IP, nil
			}
		}
	}
	return nil, errNoSuchInterface
}

func setIPv4MreqToInterface(mreq *syscall.IPMreq, ifi *Interface) error {
	if ifi == nil {
		return nil
	}
	ifat, err := ifi.Addrs()
	if err != nil {
		return err
	}
	for _, ifa := range ifat {
		switch v := ifa.(type) {
		case *IPAddr:
			if a := v.IP.To4(); a != nil {
				copy(mreq.Interface[:], a)
				goto done
			}
		case *IPNet:
			if a := v.IP.To4(); a != nil {
				copy(mreq.Interface[:], a)
				goto done
			}
		}
	}
done:
	if bytesEqual(mreq.Multiaddr[:], IPv4zero.To4()) {
		return errNoSuchMulticastInterface
	}
	return nil
}

func setReadBuffer(fd *netFD, bytes int) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd.sysfd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, bytes))
}

func setWriteBuffer(fd *netFD, bytes int) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd.sysfd, syscall.SOL_SOCKET, syscall.SO_SNDBUF, bytes))
}

func setKeepAlive(fd *netFD, keepalive bool) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd.sysfd, syscall.SOL_SOCKET, syscall.SO_KEEPALIVE, boolint(keepalive)))
}

func setLinger(fd *netFD, sec int) error {
	var l syscall.Linger
	if sec >= 0 {
		l.Onoff = 1
		l.Linger = int32(sec)
	} else {
		l.Onoff = 0
		l.Linger = 0
	}
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return os.NewSyscallError("setsockopt", syscall.SetsockoptLinger(fd.sysfd, syscall.SOL_SOCKET, syscall.SO_LINGER, &l))
}
