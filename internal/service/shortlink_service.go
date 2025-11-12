package service

import (
	"errors"
	"shortLink/internal/pkg"
	"shortLink/internal/repo"
	"time"
)

func CreateShortLink(originalUrl string) (string, error) {
	code := pkg.GenerateCode(originalUrl)
	error := repo.SaveShortLink(code, originalUrl)
	return code, error
}

func GetOriginalURL(code string) (string, error) {
	sl, err := repo.FindShortLink(code)
	if err != nil {
		return "", errors.New("not found")
	}

	// 业务规则：过期即视为无效
	if sl.ExpireAt != nil && sl.ExpireAt.Before(time.Now()) {
		return "", errors.New("link expired")
	}

	go repo.IncrClick(code)
	return sl.URL, nil
}
