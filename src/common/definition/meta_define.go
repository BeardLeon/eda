package definition

import (
	"container/list"
)

// A special segment that represents the file. You can imagine it as an
// inode in file system. It's indexed and sharded by file id.
type FileMeta struct {
	// eg: /f1
	Name string
	// a unique ID
	Id string
	// File's owner can be parent folder, or tags.
	// A file shall only have 1 parent folder,
	// but can have many tags pointing to it.
	OwnerList *list.List
	// This is the reference to blobs meta. Blob is a heavy binary sitting
	// on distributed data storage.
	BlobId string
	// A list of RangeCode
	RngCodeList *list.List
}
