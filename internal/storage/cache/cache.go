package cache

import (
	"encoding/json"
	"io"
	"os"
)

func loadCache() (map[string]interface{}, error) {
	cacheFile, err := os.Open("./cache.json")
	if err != nil {
		return nil, err
	}
	defer cacheFile.Close()
	var data map[string]interface{}
	cacheFileInfo, err := cacheFile.Stat()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, cacheFileInfo.Size())

	for {
		_, err := cacheFile.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func updateCache(data map[string]any) error {
	cacheFile, err := os.OpenFile("./cache.json", os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer cacheFile.Close()
	cacheData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = cacheFile.Write(cacheData)
	if err != nil {
		return err
	}
	return nil
}

func WriteDataToCache(k string, v any) error {
	cache, err := loadCache()
	if err != nil {
		return err
	}
	cache[k] = v

	err = updateCache(cache)
	if err != nil {
		return err
	}

	return nil

}

func GetDataFromCache(k string) any {
	cache, err := loadCache()
	if err != nil {
		return err
	}
	val, ok := cache[k]
	if ok {
		return val
	} else {
		return nil
	}
}
