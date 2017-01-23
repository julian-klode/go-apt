package goapt

// #cgo LDFLAGS: -lapt-pkg
// #include <stdlib.h>
// #include "apt.h"
import "C"
import "unsafe"

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

// Session represents the C++ type pkgCacheFile, a cache with a depcache, a
// policy, and so on.
type Session struct{}

// Cache is a collection of packages and package files
type Cache struct{}

// Policy determines the priority of versions of a package.
type Policy struct{}

// Package is a combination of a name and an architecture.
type Package struct{}

// Version represents a single version of a package (like a .deb file)
type Version struct{}

type Config struct {
	c *C.struct_Configuration
}

// Get the current configuration object
func GetConfig() *Config {
	return &Config{C._config}
}

func (cfg *Config) Find(key string) string {
	ckey := C.CString(key)
	cval := C.apt_config_find(cfg.c, ckey)
	val := C.GoString(cval)
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(cval))
	return val
}

// NewSession begins a new session
func NewSession() *Session

// Cache returns the cache of the session
func (s *Session) Cache() *Cache

// Policy returns the policy of the session.
func (s *Session) Policy() *Policy

// Lookup looks a package up by its name in the cache.
func (c *Cache) Lookup(Pkg string, Arch string) *Package

// CandidateVersion returns the candidate version of a package.
func (p *Policy) CandidateVersion(Pkg *Package) *Version

// Priority returns the priority of the given version.
func (p *Policy) Priority(Ver *Version) int

// VerStr returns the version as a string.
func (v *Version) VerStr() string
