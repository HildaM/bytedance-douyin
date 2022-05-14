package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/global"
	"bytedance-douyin/utils"
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
	bc := c.Keys["claims"]

	claims := bc.(vo.BaseClaims)
	//fmt.Println(claims)
	userId := claims.Id
	username := claims.Name

	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	// 通过token获取用户名及用户id
	path := global.GVA_CONFIG.File.VideoOutput
	videoName := utils.GenerateFilename(username, userId)
	fn := path + videoName + ".mp4"

	// 上传file文件到指定的文件路径fn
	if err := c.SaveUploadedFile(file, fn); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	// 读取视频文件的第一帧作为视频封面
	reader := utils.ReadFrameAsJpeg(fn)
	img, err := imaging.Decode(reader)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	imagePath := global.GVA_CONFIG.File.ImageOutput

	// replace .mp4 to .jpg
	imageName := videoName[:len(videoName)-4] + ".jpg"
	if err := imaging.Save(img, imagePath+imageName); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	response.OkWithMessage(c, "上传成功")
}

// VideoFeed 视频流接口
func (api *VideoApi) VideoFeed(c *gin.Context) {
	//videoList := make([]*data.Video, 1, 1)
	//user := data.Author{Id: 1, Name: "charon", FollowCount: 0, FollowerCount: 1, IsFollow: false}
	//videoList[0] = &data.Video{Id: 1, Author: &user, PlayUrl: "http://220.243.147.162:8080/videos/sss.mp4", CoverUrl: "http://220.243.147.162:8080/images/sss.jpg", FavoriteCount: 0, CommentCount: 0, IsFavorite: false}
	//data := data.FeedData{NextTime: time.Now().Unix(), VideoList: videoList}
	//response.OkWithData(c, data)
}

func (api *VideoApi) VideoList(c *gin.Context) {

}
