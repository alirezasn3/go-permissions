package goPermissions

type Permissions struct {
	Read    bool
	Write   bool
	Execute bool
}

// This function only works on linux
func GetPermissions(path string) (*Permissions, error) {
	return &Permissions{}, nil
}
