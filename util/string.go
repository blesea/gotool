package util

import (
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
)

// ToJSONStr converts an object to a JSON string representation
func ToJSONStr(obj interface{}) string {
	str := ""
	resBytes, _ := json.Marshal(obj)
	str = string(resBytes)
	return str
}

// CalMD5 calculates the MD5 hash of the given data and returns it as a hexadecimal string
func CalMD5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}