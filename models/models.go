package models

type Health struct {
	InstanceId     uint32  `json:"_id" xml:"_id"`
	Status         string  `json:"status" xml:"status"`
	CacheDepth     int     `json:"cacheDepth" xml:"cacheDepth"`
	CacheSize      uintptr `json:"cacheSize" xml:"cacheSize"`
	CacheItemLimit int     `json:"cacheItemLimit" xml:"cacheItemLimit"`
	Timestamp      string  `json:"timestamp" xml:"timestamp"`
}

type HashResponse struct {
	InstanceId uint32 `json:"_id" xml:"_id"`
	Status     string `json:"status" xml:"status"`
	Timestamp  string `json:"timestamp" xml:"timestamp"`
}

type CacheResponse struct {
	Exists    bool   `json:"exists" xml:"exists"`
	Key       string `json:"key" xml:"key"`
	Value     string `json:"value" xml:"value"`
	Timestamp string `json:"timestamp" xml:"timestamp"`
}

type CacheLoad struct {
	Key   string `json:"key" xml:"key"`
	Value string `json:"value" xml:"value"`
}

type CacheLoadConfirmation struct {
	Key       string `json:"key" xml:"key"`
	Timestamp string `json:"timestamp" xml:"timestamp"`
}
