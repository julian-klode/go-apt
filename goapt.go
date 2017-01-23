/*
Copyright (C) 2016 Julian Andres Klode <jak@jak-linux.org>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
*/

/*
Package goapt provides go bindings for the APT package management tool.

The bindings allow exploring the APT cache and configuration. They use cgo
with a middle layer of C bindings to the underlying C++ code.
*/
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

// NewSession opens a new session. Note that Config() and System() really are
// global variables. They will be initialized here, but there can only really
// be one session at one time.
func NewSession() (*Session, error)

// NewSessionWithConfig opens a new session, with the specified configuration
// set as the global configuration object instead of the default one.
func NewSessionWithConfig(cfg *Config) (*Session, error)

// Config gets the configuration object associated with the session. This
// is currently a global variable in apt.
func (s *Session) Config() (*Config, error) {
	return &Config{C.apt_config_get_default()}, nil
}

func (s *Session) System() (*System, error)
func (s *Session) Cache() (*Cache, error)
func (s *Session) DepCache() (*DepCache, error)
func (s *Session) Policy() (*Policy, error)

// DepCache is a collection of marked package states
type DepCache struct{}

// Policy determines the priority of versions of a package.
type Policy struct{}

// CandidateVersion returns the candidate version of a package.
func (p *Policy) CandidateVersion(Pkg *Package) *Version

// Priority returns the priority of the given version.
func (p *Policy) Priority(Ver *Version) int

// Package is a combination of a name and an architecture.
type Package struct{}

// System is a system-level package manager, like dpkg
type System struct{}

// Version represents a single version of a package (like a .deb file)
type Version struct{}

type Config struct {
	c *C.struct_apt_config
}

func (cfg *Config) Find(key string) string {
	ckey := C.CString(key)
	cval := C.apt_config_find(cfg.c, ckey)
	val := C.GoString(cval)
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(cval))
	return val
}

// VerStr returns the version as a string.
func (v *Version) VerStr() string
