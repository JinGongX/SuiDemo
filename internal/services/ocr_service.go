package services

import (
	"changeme/platform"
)

type OCRService struct{}

// 识别Base64图片（data:image/png;base64,...）
func (o *OCRService) RecognizeImageBase64(base64str string) (string, error) {
	return platform.RecognizeImageBase64(base64str)
}
