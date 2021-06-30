package model

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

type Sharding interface {
	GetDbSuffix() string
	GetTableSuffix() string
}

type ModHashSharding struct {
	ShardingVal   int64
	ShardingDb    int64
	ShardingTable int64
}

func (h *ModHashSharding) GetDbSuffix() string {
	return strconv.FormatInt(h.ShardingVal%h.ShardingDb, 10)
}

func (h *ModHashSharding) GetTableSuffix() string {
	return strconv.FormatInt(h.ShardingVal%h.ShardingTable, 10)
}

type Md5HashSharding struct {
	ShardingVal   int64
	ShardingDb    int64
	ShardingTable int64
}

func (h *Md5HashSharding) GetDbSuffix() string {
	return string([]byte(fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(h.ShardingVal, 10)))))[0:1])
}

func (h *Md5HashSharding) GetTableSuffix() string {
	return string([]byte(fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(h.ShardingVal, 10)))))[0:2])
}

type MonthSharding struct {
}

func (h *MonthSharding) GetDbSuffix() string {
	return ""
}

func (h *MonthSharding) GetTableSuffix() string {
	return time.Now().Format("2006_01")
}
