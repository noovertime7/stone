package services

import (
	"context"
	"github.com/noovertime7/stone/dao"
)

type BannerService struct{}

func (s *BannerService) List(ctx context.Context) ([]*dao.Banner, error) {
	model := dao.Banner{}
	return model.FindList(ctx)
}

func (s *BannerService) Save(ctx context.Context, image, url string, sortOrder int) error {
	model := dao.Banner{
		Image:     image,
		Url:       url,
		SortOrder: sortOrder,
	}
	return model.Save(ctx, &model)
}

func (s *BannerService) Delete(ctx context.Context, id int) error {
	model := dao.Banner{}
	return model.Delete(ctx, id)
}
