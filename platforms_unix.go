// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package main

import (
	"syscall"
)

func Umask(mask int) int {
	return syscall.Umask(m)
}
