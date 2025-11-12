package repo

import (
	"context"
	"errors"
	"fmt"
	"shortLink/internal/config"
	"shortLink/internal/model"

	"gorm.io/gorm"
)

func SaveShortLink(code string, url string) error {
	var sl model.ShortLink
	err := DB.Where("code = ?", code).First(&sl).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		sl = model.ShortLink{Code: code, URL: url}
		if err := DB.Create(&sl).Error; err != nil {
			return err
		}
	} else if err == nil {
		sl.URL = url
		if err := DB.Save(&sl).Error; err != nil {
			return err
		}
	} else {
		return err
	}

	ctx := context.Background()

	if RDB.Set(ctx, fmt.Sprintf("%s%s", config.C.Redis.KeyPrefix, code), url, TTL) != nil {
	}
	return nil
}

func FindShortLink(code string) (*model.ShortLink, error) {
	var sl model.ShortLink
	if err := DB.Where("code = ?", code).First(&sl).Error; err != nil {
		return nil, err
	}
	return &sl, nil
}

func IncrClick(code string) error {
	return DB.Model(&model.ShortLink{}).
		Where("code = ?", code).
		UpdateColumn("clicks", gorm.Expr("clicks + 1")).Error
}
