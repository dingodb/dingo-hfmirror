package model

type SystemInfo struct {
	Id                string  `json:"id"`
	Name              string  `json:"name"`
	Version           string  `json:"version"`
	CollectTime       int64   `json:"-"`
	MemoryUsedPercent float64 `json:"-"`
}

func (s *SystemInfo) SetMemoryUsed(collectTime int64, usedPercent float64) {
	s.CollectTime = collectTime
	s.MemoryUsedPercent = usedPercent
}
