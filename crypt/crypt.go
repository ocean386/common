package crypt

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"sort"
	"strings"
)

// RsaPrivateSign RSA 签名-私钥
func RsaPrivateSign(originalData, strPrivateKey string) (string, error) {

	bytePrivateKey, err := base64.StdEncoding.DecodeString(strPrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "base64 DecodeString private key failed")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(bytePrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "Parse pkcs8 private key failed")
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("Convert private key failed")
	}

	hashed := sha256.Sum256([]byte(originalData))
	// RSA私钥对哈希数据-签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", errors.Wrap(err, "Sign private key failed")
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// RsaPublicVerySign RSA 验签-公钥
func RsaPublicVerySign(originalData, strSign, strPublicKey string) (bool, error) {

	// 解码Base64编码的签名
	bytePublicKey, err := base64.StdEncoding.DecodeString(strPublicKey)
	if err != nil {
		return false, errors.Wrap(err, "Base64 decode public key failed")
	}

	// 解析公钥
	rsaPublicKey, err := x509.ParsePKIXPublicKey(bytePublicKey)
	if err != nil {
		return false, errors.Wrap(err, "Parse public key failed")
	}

	byteSign, err := base64.StdEncoding.DecodeString(strSign)
	if err != nil {
		return false, errors.Wrap(err, "Base64 decode sign key failed")
	}

	hashed := sha256.Sum256([]byte(originalData))
	// RSA公钥-验签
	err = rsa.VerifyPKCS1v15(rsaPublicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], byteSign)
	if err != nil {
		return false, errors.Wrap(err, "Verify sign failed")
	}

	return true, nil
}

// EncryptAES AES加密-CBC模式(密钥长度32个字节)
func EncryptAES(originData []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.Wrap(err, "Aes encrypt new cipher failed")
	}

	blockSize := block.BlockSize()
	padding := blockSize - len(originData)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	originData = append(originData, paddingText...)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encryptData := make([]byte, len(originData))
	blockMode.CryptBlocks(encryptData, originData)
	return hex.EncodeToString(encryptData), nil
}

// DecryptAES AES解密-CBC模式(密钥长度32个字节)
func DecryptAES(originData string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "Aes decrypt new cipher failed")
	}

	byteDecodeData, err := hex.DecodeString(originData)
	if err != nil {
		return nil, errors.Wrap(err, "Aes decrypt decode string failed")
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decryptData := make([]byte, len(byteDecodeData))
	blockMode.CryptBlocks(decryptData, byteDecodeData)
	size := len(decryptData)
	unPadding := int(decryptData[size-1])
	decryptData = decryptData[:(size - unPadding)]
	return decryptData, nil
}

// GetSortFieldAsc 结构体字段进行ASCII排序
func GetSortFieldAsc(obj interface{}) (byteParam []byte) {

	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	data := make(map[string]interface{})

	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldName := objType.Field(i).Tag.Get("json")
		if fieldName == "" {
			continue
		}

		// 如果字段值为零值，则不添加到 map 中
		if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			continue
		}

		data[fieldName] = field.Interface()
	}

	// 对字段按名称进行排序
	var fieldNames []string
	for fieldName := range data {
		fieldNames = append(fieldNames, fieldName)
	}
	sort.Strings(fieldNames)

	// 按排序后的字段名拼接字段值
	var filteredFields []string
	for _, fieldName := range fieldNames {
		value := data[fieldName]
		valueStr := ""
		switch v := value.(type) {
		case string:
			valueStr = v
		default: // 转换字段值为字符串
			valueStr = fmt.Sprintf("%v", v)
		}

		// 只添加非空字段值
		if valueStr != "" {
			filteredFields = append(filteredFields, valueStr)
		}
	}

	// 将过滤后的字段值拼接成一个字符串
	strParam := strings.Join(filteredFields, "")
	byteParam = []byte(strParam)
	return
}
