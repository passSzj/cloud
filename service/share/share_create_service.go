package share

import (
	"context"
	"github.com/ChenMiaoQiu/go-cloud-disk/rpc/types/rpc"
	shorturlservice "github.com/ChenMiaoQiu/go-cloud-disk/rpc_client"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
)

type ShareCreateService struct {
	FileId string `json:"fileid" form:"fileid" binding:"required"`
	Title  string `json:"title" form:"title" binding:"required"`
}

type createShareSuccessResponse struct {
	ShareId   string `json:"shareid"`
	ShareLink string `json:"sharelink"`
}

func (service *ShareCreateService) CreateShare(userId string) serializer.Response {
	// check file owner
	var shareFile model.File
	if err := model.DB.Where("uuid = ? and owner = ?", service.FileId, userId).Find(&shareFile).Error; err != nil {
		logger.Log().Error("[ShareCreateService.CreateShare] Fail to find file info: ", err)
		return serializer.DBErr("", err)
	}

	// create share and save to database
	newShare := model.Share{
		Owner:       userId,
		FileId:      service.FileId,
		Title:       service.Title,
		Size:        shareFile.Size,
		FileName:    shareFile.FileName + "." + shareFile.FilePostfix,
		SharingTime: time.Unix(time.Now().Unix(), 0).Format(utils.DefaultTimeTemplate),
	}
	if err := model.DB.Create(&newShare).Error; err != nil {
		logger.Log().Error("[ShareCreateService.CreateShare] Fail to create share: ", err)
		return serializer.DBErr("", err)
	}

	//调用短链接rpc服务生成短链接  赋值到sharelink  失败则赋值为空  如果为空前端则根据ShareId使用长链接
	shortUrl, err := shorturlservice.ShortUrlClient.Convert(context.Background(), &rpc.ConvertRequest{
		LongUrl: "https://space.bilibili.com/4638193/pugv",
	})

	if err != nil {
		shortUrl.ShortUrl = ""
	}

	return serializer.Success(createShareSuccessResponse{
		ShareId:   newShare.Uuid,
		ShareLink: shortUrl.ShortUrl,
	})
}
