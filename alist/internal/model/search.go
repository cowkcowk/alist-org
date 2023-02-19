package model

import (
	"time"
)

type IndexProgress struct {
	ObjCount     uint64     `json:"obj_count"`
	IsDone       bool       `json:"is_done"`
	LastDoneTime *time.Time `json:"last_done_time"`
	Error        string     `json:"error"`
}

type SearchNode struct {
	Parent string `json:"parent" gorm:"index"`
	Name   string `json:"name"`
	IsDir  bool   `json:"is_dir"`
	Size   int64  `json:"size"`
}
