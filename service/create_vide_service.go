package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 视频投稿的服务
type CreateVideoService struct {
	Title  string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info   string `form:"info" json:"info" binding:"max=3000"`
}

func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title:  service.Title,
		Info:   service.Info,
	}

	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:   50001,
			Msg:    "视频保存失败",
			Error:   err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
