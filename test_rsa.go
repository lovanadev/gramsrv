package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func checkKey(name, keyPEM string) {
	block, _ := pem.Decode([]byte(keyPEM))
	if block == nil {
		fmt.Printf("%s: failed to decode PEM\n", name)
		return
	}
	_, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("%s: invalid RSA key (%v)\n", name, err)
	} else {
		fmt.Printf("%s: VALID RSA KEY!\n", name)
	}
}

func main() {
	keyI := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2SHOd76kkwf+IFjPqKop
rWaboVUjClkGCZ9GACJZypDHnIWmaNp25qkaaV4E7pTcOZgyC7JBONSpxIP5NvpM
s8KftbOg+K9Qp4a3AzQnJX03KSGWrZrYuL8svT012Z+MLVyy6waO/xJRFVuPNPGK
ZDcbcJHy/TpbYeiiJJMzDnuwQY10dpBPTI7/rsn0h22KQ+TPnL0GeHo9OOpIbC7d
Bm2P6X+aBFXOXH/FKLi1tbnDwOtDv6VsHljociwSSDGaCwZ7IP4tQX1RWQX1fB8L
iwhvGHK54d4Fa5r0rP6ARWs/bKtRFhTZ2rW4jnd9mYqGy6+uOXgeo+lcIUmS+mNo
+QIDAQAB
-----END PUBLIC KEY-----`

	keyL := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2SHOd76kkwf+IFjPqKop
rWaboVUjClkGCZ9GACJZypDHnlWmaNp25qkaaV4E7pTcOZgyC7JBONSpXlP5NvpM
s8KftbOg+K9Qp4a3AzQnJX03KSGWrZrYuL8svT012Z+MLVyy6waO/xJRFVuPNPGK
ZDcbcJHy/TpbYeiiJJMzDnuwQY10dpBPTI7/rsn0h22KQ+TPnL0GeHo9OOpIbC7d
Bm2P6X+aBFXOXH/FKLi1tbnDwOtDv6VsHljociwSSDGaCwZ7IP4tQX1RWQX1fB8L
iwhvGHK54d4Fa5r0rP6ARWs/bKtRFhTZ2rW4jnd9mYqGy6+uOXgeo+lcIUmS+mNo
+QIDAQAB
-----END PUBLIC KEY-----`

	checkKey("Key with I", keyI)
	checkKey("Key with l", keyL)
}
