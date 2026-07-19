package main

import (
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	var kq int
	var err error
	if kq, err = unix.Kqueue(); err != nil {
		err = os.NewSyscallError("kqueue", err)
		return
	}

	_, err = unix.Kevent(kq, []unix.Kevent_t{{
		Ident:  0,
		Filter: unix.EVFILT_USER,
		Flags:  unix.EV_ADD | unix.EV_CLEAR,
	}}, nil, nil)
	if err != nil {
		err = os.NewSyscallError("kevent | pipe2", err)
		return
	}

}
