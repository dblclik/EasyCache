package models

type Health struct {
	Status     string  `json:"status" xml:"status"`
	CacheDepth int     `json:"cacheDepth" xml:"cacheDepth"`
	CacheSize  uintptr `json:"cacheSize" xml:"cacheSize"`
	Timestamp  string  `json:"timestamp" xml:"timestamp"`
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
