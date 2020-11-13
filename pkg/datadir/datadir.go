package datadir

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/xiaods/k8e/pkg/version"
)

var (
	DefaultDataDir     = "/var/lib/" + version.Program
	DefaultHomeDataDir = "${HOME}/.k8e/" + version.Program
	HomeConfig         = "${HOME}/.kube/" + version.Program + ".yaml"
	GlobalConfig       = "/etc/" + version.Program + "/" + version.Program + ".yaml"
)

func LocalHome(dataDir string, forceLocal bool) (string, error) {
	if dataDir == "" {
		if os.Getuid() == 0 && !forceLocal {
			dataDir = DefaultDataDir
		} else {
			dataDir = DefaultHomeDataDir
		}
	}

	dataDir, err := Resolve(dataDir)
	if err != nil {
		return "", errors.Wrapf(err, "resolving %s", dataDir)
	}

	return filepath.Abs(dataDir)
}

func HomeKubeConfig(write, rootless bool) (string, error) {
	if write {
		if os.Getuid() == 0 && !rootless {
			return GlobalConfig, nil
		}
		return Resolve(HomeConfig)
	}

	if _, err := os.Stat(GlobalConfig); err == nil {
		return GlobalConfig, nil
	}

	return Resolve(HomeConfig)
}

var (
	homes = []string{"$HOME", "${HOME}", "~"}
)

func Resolve(s string) (string, error) {
	for _, home := range homes {
		if strings.Contains(s, home) {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return "", errors.Wrap(err, "determining current user")
			}
			s = strings.Replace(s, home, homeDir, -1)
		}
	}

	return s, nil
}
