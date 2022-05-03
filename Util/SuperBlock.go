package util

type SuperBlock struct {
	Filesystem_type   int64
	Inodes_count      int64
	Blocks_count      int64
	Free_blocks_count int64
	Free_inodes_count int64
	Mnt_count         int64
	M_time            [16]byte
	Magic             int64
	Inode_size        int64
	Block_size        int64
	First_ino         int64
	First_blo         int64
	Bm_inode_start    int64
	Bm_block_start    int64
	Inode_start       int64
	Block_start       int64
}

func NewSuperBlock() SuperBlock {
	return SuperBlock{Filesystem_type: 0, Inodes_count: 0, Blocks_count: 0, Free_blocks_count: 0, Free_inodes_count: 0, Magic: 0xEF53}
}