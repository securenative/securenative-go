package utils

import (
	"os/exec"
	"strings"
)

type VersionUtils struct {}

func NewVersionUtils() *VersionUtils {
	return &VersionUtils{}
}

func (u *VersionUtils) GetVersion() string {
	app := "git"

	arg0 := "tag"

	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()

	if err != nil {
		return "0.0.1"
	}

	return strings.Replace(string(stdout), "v", "", -1)
}
