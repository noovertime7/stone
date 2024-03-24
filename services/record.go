package services

import (
	"context"
	"fmt"
	"github.com/noovertime7/stone/dao"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/runtime"
	"github.com/noovertime7/stone/utils"
)

type RecordService struct {
}

func (s *RecordService) Get(ctx context.Context, id int) (dao.Record, error) {
	model := &dao.Record{}
	data, ok, err := model.Find(ctx, dao.Record{Id: id})
	if !ok {
		return dao.Record{}, fmt.Errorf("分类不存在")
	}
	return data, err
}

func (s *RecordService) List(ctx context.Context) ([]*dao.Record, error) {
	model := dao.Record{}

	return model.FindList(ctx, dao.Record{})
}

func (s *RecordService) Delete(ctx context.Context, id int) error {
	model := dao.Record{}
	return model.Delete(ctx, dao.Record{Id: id}, false)
}

func (s *RecordService) Page(ctx context.Context, params runtime.DataBasePager) (dto.PagerRecordOut, error) {
	model := dao.Record{}
	data, total, err := model.PageList(ctx, params)
	return dto.PagerRecordOut{
		Total:    total,
		List:     data,
		Page:     params.GetPage(),
		PageSize: params.GetPageSize(),
	}, err
}

//func (s *RecordService) Update(ctx context.Context, id int, name string) error {
//	model := dao.Record{}
//	return model.Updates(ctx, common.WithIDOption(id, common.Equal), &dao.Record{Name: name})
//}

func (s *RecordService) Save(ctx context.Context, in *dto.CreateRecordInput) error {
	stone := dao.Stone{Id: in.StoneId}
	st, _, err := stone.Find(ctx, stone)
	if err != nil {
		return err
	}
	if utils.IsStrEmpty(in.Description) {
		in.Description = fmt.Sprintf("在%s施工,石材:%s", in.Location, st.Name)
	}

	var model = dao.Record{
		StoneId:          in.StoneId,
		StoneName:        st.Name,
		Video:            in.Video,
		Images:           in.Images,
		Location:         in.Location,
		Description:      in.Description,
		Longitude:        in.Longitude,
		Latitude:         in.Latitude,
		DetailedLocation: in.DetailedLocation,
	}
	return model.Save(ctx, &model)
}
