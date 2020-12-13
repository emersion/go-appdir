package appdir

import (
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/sys/windows"
)

type dirs struct {
	name string
}

var initOnce sync.Once
var localAppData string
var roamingAppData string

func initFolders() {
	var err error
	localAppData, err = windows.KnownFolderPath(windows.FOLDERID_LocalAppData, 0)
	if err != nil {
		localAppData = os.Getenv("LOCALAPPDATA")
	}
	roamingAppData, err = windows.KnownFolderPath(windows.FOLDERID_RoamingAppData, 0)
	if err != nil {
		roamingAppData = os.Getenv("APPDATA")
	}
}

func (d *dirs) UserConfig() string {
	initOnce.Do(initFolders)
	return filepath.Join(roamingAppData, d.name)
}

func (d *dirs) UserCache() string {
	initOnce.Do(initFolders)
	return filepath.Join(localAppData, d.name)
}

func (d *dirs) UserLogs() string {
	initOnce.Do(initFolders)
	return filepath.Join(localAppData, d.name)
}

func (d *dirs) UserData() string {
	initOnce.Do(initFolders)
	return filepath.Join(localAppData, d.name)
}
