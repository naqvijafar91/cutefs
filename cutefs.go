package cutefs

// SuperBlock represents the first block of the storage space. It is 4 bytes(size of each field) * 4 = 16 bytes, but it takes a complete block on the disk, we can
// use it for future requirements of the file system
type SuperBlock struct {
	MagicNumber uint32
	TotalBlocks uint32
	InodeBlocks uint32
	TotalInodes uint32
}

// BlockPointer is a block pointer of 2 bytes
type BlockPointer uint16

// Inode is total 8 * 2 bytes(size of each field) = 16 bytes on disk
type Inode struct {
	// size of valid is 1 byte in memory but, we are reserving 1 byte for the future,
	Valid bool
	// SizeInBytes represents the actual size of the data in bytes and not the on disk size which will be multiples of block size
	SizeInBytes uint16
	// DirectBlockPointers is an array of block pointers which point to an actual data block on the block device.
	DirectBlockPointers [5]BlockPointer
	// IndirectBlockPointer is an indirect block pointer which points to an indirect block which itself will be an array of pointers
	IndirectBlockPointer BlockPointer
}

// DataBlock is a block containing the actual data
type DataBlock struct {
	Data [BlockSizeBytes]byte
}

// IndirectPointerBlock is a special type of block which contains pointers to the actual data block.
type IndirectPointerBlock struct {
	DirectBlockPointers [MaxBlockPointersPerBlock]BlockPointer
}

// InodeBlock is a block which contains inode entries
type InodeBlock struct {
	Inodes []Inode
}

// BlockSizeBytes is the block size of the fs
const BlockSizeBytes = 4096

// MagicNumber is the magic number of the fs. It could be any number which identifies our fs.
const MagicNumber = 786786786

// MaxBlockPointersPerBlock represents the max number of block pointers in each block. It can be calculated as block size / size of each pointer
const MaxBlockPointersPerBlock = BlockSizeBytes / 2
