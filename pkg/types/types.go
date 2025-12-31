package types

type ListDir []ListFile

type ListFile struct {
	Name string `json:"name"` // 文件名
	Size int64  `json:"size"` // 文件大小
	Time int64  `json:"time"` // 修改时间
}
