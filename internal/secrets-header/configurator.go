// Package kucoinheader contain logic to create secret header and signature.
package kucoinheader

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

// CreateSecretsHeaders creates the secret values needed for https requests to kucoin-client.
func CreateSecretsHeaders(method, url, body, secret, passPhrase, key, version string, now time.Time) map[string]string {
	headers := make(map[string]string)

	dur := now.Add(1 * time.Second).UnixMilli()
	timeStamp := strconv.FormatInt(dur, 10)

	signature := strings.Join([]string{timeStamp, method, url, body}, "")

	sign := encode(secret, signature)
	passPhraseEncoded := encode(secret, passPhrase)

	headers["KC-API-KEY"] = key
	headers["KC-API-SIGN"] = sign
	headers["KC-API-TIMESTAMP"] = timeStamp
	headers["KC-API-PASSPHRASE"] = passPhraseEncoded
	headers["KC-API-KEY-VERSION"] = version

	headers["Content-Type"] = "application/json"

	return headers
}

func encode(key, message string) string {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(message))
	res := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(res)
}
