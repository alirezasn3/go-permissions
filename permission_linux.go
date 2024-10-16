package goPermissions

import (
	"errors"
	"io/fs"
	"os"
	"syscall"
)

type Permissions struct {
	Read    bool
	Write   bool
	Execute bool
}

func GetPermissions(path string) (*Permissions, error) {
	// get info
	info, e := os.Stat(path)
	if errors.Is(e, fs.ErrNotExist) {
		return nil, fs.ErrNotExist
	}

	// get stat
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, errors.New("wrong os")
	}

	// get file mode
	mode := uint32(info.Mode())

	// check if current user is owner or is in file group
	if uint32(os.Getuid()) == stat.Uid {
		mode = mode >> 6
	} else if uint32(os.Getgid()) == stat.Gid {
		mode = mode >> 3
	}

	return &Permissions{
		Read:    mode&4 == 4,
		Write:   mode&2 == 2,
		Execute: mode&1 == 1,
	}, nil
}
