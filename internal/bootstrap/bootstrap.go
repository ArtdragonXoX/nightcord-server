//go:build linux
// +build linux

package bootstrap

func Init() {
	initConf()
	initStorage()
	initServer()
}
