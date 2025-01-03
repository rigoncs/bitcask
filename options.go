package bitcask

type Options struct {
	// 数据库数据目录
	DirPath string

	// 数据文件的大小
	DataFileSize int64

	// 数据文件合并的阈值
	DataFileMergeRatio float32
}
