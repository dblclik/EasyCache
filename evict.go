package main

func evict() {
	keyToDelete := LRUCache.Tail.Data
	LRUCache.RemoveTailDLL()

	delete(CacheMap, keyToDelete)
}
