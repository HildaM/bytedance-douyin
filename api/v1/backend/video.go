package backend

import (
	"bytedance-douyin/global"
	r "bytedance-douyin/model/data"
	"bytedance-douyin/utils"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: vedio
 * @Version: 1.0.0
 * @Date: 2022/5/6 18:33
 */
type VideoApi struct{}

func (api *VideoApi) PostVideo(c *gin.Context) {
	file, err := c.FormFile("data")

	if err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	// todo 生成文件名
	path := global.GVA_CONFIG.File.VideoOutput
	videoName := utils.GenerateFilename(file.Filename)
	fn := path + videoName

	// 上传file文件到指定的文件路径fn
	if err := c.SaveUploadedFile(file, fn); err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	// 读取视频文件的第一帧作为视频封面
	reader := utils.ReadFrameAsJpeg(fn)
	img, err := imaging.Decode(reader)
	if err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}
	imagePath := global.GVA_CONFIG.File.ImageOutput

	// replace .mp4 to .jpg
	imageName := videoName[:len(videoName)-4] + ".jpg"
	if err := imaging.Save(img, imagePath+imageName); err != nil {
		r.FailWithMessage(c, fmt.Sprintf("%s", err))
	}

	r.OkWithMessage(c, "上传成功")
}

func (api *VideoApi) VideoFeed(c *gin.Context) {

}

func (api *VideoApi) VideoList(c *gin.Context) {

}
