package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	mu.Lock()
	defer mu.Unlock()
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// GenerateAPIKey 生成 API Key
func GenerateAPIKey() string {
	timestamp := time.Now().UnixNano()       // 使用纳秒级时间戳
	uuID := uuid.New().String()              // 生成 UUID
	randomString := GenerateRandomString(32) // 增加随机字符串长度
	rawKey := fmt.Sprintf("%s%d%s", uuID, timestamp, randomString)

	hash := sha256.New()
	hash.Write([]byte(rawKey))
	apiKey := hex.EncodeToString(hash.Sum(nil))
	return apiKey
}

// GenerateAPISecret 生成 API Secret
func GenerateAPISecret(userID string) string {
	timestamp := time.Now().UnixNano()       // 使用纳秒级时间戳
	randomString := GenerateRandomString(64) // 增加随机字符串长度
	rawKey := fmt.Sprintf("%s%d%s", userID, timestamp, randomString)

	hash := sha512.New()
	hash.Write([]byte(rawKey))
	apiSecret := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return apiSecret
}
