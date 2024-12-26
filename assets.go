package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func (cfg apiConfig) getBucketURL(key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", cfg.s3Bucket, cfg.s3Region, key)
}

func getAssetPath(mediaType string) (string, error) {
	ext := mediaTypeToExt(mediaType)
	size := 32
	src := make([]byte, size)
	_, err := rand.Read(src)
	if err != nil {
		return "", err
	}
	encoded := base64.RawURLEncoding.EncodeToString(src)

	return fmt.Sprintf("%s%s", encoded, ext), nil
}

func getAssetPathNew(mediaType, aspectRatio string) (string, error) {
	ext := mediaTypeToExt(mediaType)
	size := 32
	src := make([]byte, size)
	_, err := rand.Read(src)
	if err != nil {
		return "", err
	}
	encoded := base64.RawURLEncoding.EncodeToString(src)

	return fmt.Sprintf("%s/%s%s", aspectRatio, encoded, ext), nil
}

func (cfg apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

func mediaTypeToExt(mediaType string) string {
	parts := strings.Split(mediaType, "/")
	if len(parts) != 2 {
		return ".bin"
	}
	return "." + parts[1]
}
