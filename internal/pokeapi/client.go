package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/zzwsec/pokedexcli/internal/pokecache"
)

func GetLocationAreas(pageURL *string, cache *pokecache.Cache) (*LocationAreasResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}
	cacheKey, err := normalizeURL(url)
	if err != nil {
		cacheKey = url // 如果标准化失败,使用原始URL
	}
	// fmt.Printf("\n[DEBUG] Original URL: %s\n", url)
	// fmt.Printf("[DEBUG] Cache key: %s\n", cacheKey)

	// 尝试从缓存获取数据
	if cacheData, ok := cache.Get(cacheKey); ok {
		las := &LocationAreasResponse{}
		if err := json.Unmarshal(cacheData, las); err == nil {
			// fmt.Printf("[DEBUG] ✅ CACHE HIT! Using cached data\n")
			return las, nil
		}
	}

	// fmt.Printf("[DEBUG] ❌ CACHE MISS - Fetching from API...\n")
	// 缓存未命中或解析失败则发起http请求
	resp, err := http.Get(cacheKey)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	las := &LocationAreasResponse{}
	err = json.Unmarshal(data, las)
	if err != nil {
		return nil, fmt.Errorf("JSON parsing failed: %w", err)
	}

	// 添加缓存
	cache.Add(cacheKey, data)
	return las, nil
}

// 标准化URL,确保参数顺序一致
func normalizeURL(rawURL string) (string, error) {
	// 解析URL字符串为结构体
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	// 提取查询参数
	params := u.Query()
	// 重新编码查询参数
	u.RawQuery = params.Encode()
	return u.String(), nil
}
