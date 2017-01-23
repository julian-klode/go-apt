package goapt

// #cgo LDFLAGS: -lapt-pkg
// #include "apt.h"
import "C"

// InitConfig (re-)initializes the configuration
func InitConfig() {
    C.apt_init_config()
}

// InitSystem initializes the packaging system like the "deb" one
func InitSystem() {
    C.apt_init_system()
}

// Init initializes both the configuration and the system.
func init() {
    InitConfig()
    InitSystem()
}
