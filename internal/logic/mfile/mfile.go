package mfile

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"one2.3/internal/consts"
	"one2.3/internal/dao"
	"one2.3/internal/model"
	"one2.3/internal/model/entity"
	"one2.3/internal/service"
	"time"
)

type sMFile struct{}

func init() {
	service.RegisterMFile(New())
}

func New() *sMFile {
	return &sMFile{}
}

// 同一上传文件
func (s *sMFile) Upload(ctx context.Context, in model.FileUploadInput) (*model.FileUploadOutput, error) {
	uploadPath := g.Cfg().MustGet(ctx, "mupload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("上传文件路径配置不存在")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	// 同一用户1分钟之内只能上传10张图片
	count, err := dao.Mfile.Ctx(ctx).
		Where(dao.Mfile.Columns().MemberId, service.Context().MGet(ctx).Id).
		WhereGTE(dao.Mfile.Columns().CreatedAt, gtime.Now().Add(time.Minute)).
		Count()
	if err != nil {
		return nil, err
	}
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("您上传得太频繁，请稍后再操作")
	}
	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)

	if err != nil {
		return nil, err
	}
	// 记录到数据表
	data := entity.Mfile{
		Name:     fileName,
		Src:      gfile.Join(uploadPath, dateDirName, fileName),
		Url:      "/upload/" + dateDirName + "/" + fileName,
		MemberId: service.Context().MGet(ctx).Id,
	}
	result, err := dao.Mfile.Ctx(ctx).Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &model.FileUploadOutput{
		Id:   uint64(id),
		Name: in.File.Filename,
		Path: data.Src,
		Url:  data.Url,
	}, nil
}
