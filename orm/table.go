package main

import (
	"time"
)

const (
	DataTypeLatestVersion int8 = 1 + iota
	DataTypeMultiVersion
	DataTypeAlbumLatestVersion
	DataTypeAlbumMultiVersion
)

const (
	DataTypeStandardLatestVersion = DataTypeLatestVersion + 10 + iota
	DataTypeStandardMultiVersion
	DataTypeStandardAlbumLatestVersion
	DataTypeStandardAlbumMultiVersion
)

const (
	DataTypeStandardIaLatestVersion = DataTypeLatestVersion + 20 + iota
	DataTypeStandardIaMultiVersion
	DataTypeStandardIaAlbumLatestVersion
	DataTypeStandardIaAlbumMultiVersion
)

type Data struct {
	Id    int64
	Date  time.Time `xorm:"index notnull default '2000-01-01 00:00:00'"`
	Type  int8      `xorm:"index notnull default 0"`
	Key   string    `xorm:"notnull default ''"`
	Count int64     `xorm:"notnull default 0"`
	Sum   int64     `xorm:"notnull default 0"` // math.MaxInt64 最大值为 8192 PB
}

type Classify struct {
	Id   int64
	Type int8   `xorm:"index notnull default 0"`
	Key  string `xorm:"index notnull default ''"`
	Name string `xorm:"notnull default ''"`
}
