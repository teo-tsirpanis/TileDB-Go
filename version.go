// Package tiledb is a idomatic go binding to tiledb's c_api. Go structs are
// used for object style access to tiledb types, such as Config and VFS.
// Tiledb c objects that are alloced are set to be freeded on garbage collection
// using `runtime.SetFinalizer`.
package tiledb

/*
#cgo LDFLAGS: -ltiledb
#include <tiledb/tiledb.h>
*/
import "C"

// Version returns the TileDB shared library version these bindings are linked
// against at runtime
func Version() (major int, minor int, rev int) {
	var cmajor C.int = -1
	var cminor C.int = -1
	var crev C.int = -1
	C.tiledb_version(&cmajor, &cminor, &crev)

	return int(cmajor), int(cminor), int(crev)

}
