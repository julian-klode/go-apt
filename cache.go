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

package apt

// #cgo LDFLAGS: -lapt-pkg
// #include <stdlib.h>
// #include "apt.h"
import "C"

// Cache represents an APT cache file. An APT cache contains information about
// available and installed packages, the versions, and to some extend, also
// configured repositories.
type Cache struct {
	c *C.struct_apt_cache
}

// LookupWithArch looks a package up by its name in the cache.
func (c *Cache) LookupWithArch(Pkg string, Arch string) *Package

// Lookup a package without an extra architecture argument.
func (c *Cache) Lookup(Pkg string) *Package

// GetPackage gets the package with the specified ID.
func (c *Cache) GetPackage(id int) *Package

// PackageCount Return the number of packages in the cache.
func (c *Cache) PackageCount() int
