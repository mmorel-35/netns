//go:build !linux
// +build !linux

package netns

import (
	"errors"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

// Setns sets namespace using golang.org/x/sys/unix.Setns on Linux. It
// is not implemented on other platforms.
//
// Deprecated: Use golang.org/x/sys/unix.Setns instead.
func Setns(_ NsHandle, _ int) (err error) {
	return ErrNotImplemented
}

func Set(_ NsHandle) (err error) {
	return ErrNotImplemented
}

func New() (ns NsHandle, err error) {
	return -1, ErrNotImplemented
}

func NewNamed(_ string) (NsHandle, error) {
	return -1, ErrNotImplemented
}

func DeleteNamed(_ string) error {
	return ErrNotImplemented
}

func Get() (NsHandle, error) {
	return -1, ErrNotImplemented
}

func GetFromPath(_ string) (NsHandle, error) {
	return -1, ErrNotImplemented
}

func GetFromName(_ string) (NsHandle, error) {
	return -1, ErrNotImplemented
}

func GetFromPid(_ int) (NsHandle, error) {
	return -1, ErrNotImplemented
}

func GetFromThread(_ int, _ int) (NsHandle, error) {
	return -1, ErrNotImplemented
}

func GetFromDocker(_ string) (NsHandle, error) {
	return -1, ErrNotImplemented
}
