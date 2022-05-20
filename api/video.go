package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/global"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"time"
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
		response.FailWithMessage(c, err.Error())
		return
	}
	
	bc := c.Keys["claims"]
	claims := bc.(vo.BaseClaims)
	userId := claims.Id
	username := claims.Name
	
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	
	// 通过token获取用户名及用户id
	path := global.GVA_CONFIG.File.VideoOutput
	videoName := utils.GenerateFilename(username, userId)
	fn := path + videoName + ".mp4"
	
	// 上传file文件到指定的文件路径fn
	if err := c.SaveUploadedFile(file, fn); err != nil {
		fmt.Println("存储视频失败, filePath,", fn)
		response.FailWithMessage(c, err.Error())
		return
	}
	
	// FIXME: utils.ReadFrameAsJpeg 在本地环境下出错
	// 读取视频文件的第一帧作为视频封面
	// reader := utils.ReadFrameAsJpeg(fn)
	// img, err := imaging.Decode(reader)
	// if err != nil {
	// 	response.FailWithMessage(c, err.Error())
	// 	return
	// }
	// imagePath := global.GVA_CONFIG.File.ImageOutput
	
	// replace .mp4 to .jpg
	imageName := videoName[:len(videoName)-4]
	// url := imagePath + imageName + ".jpg"
	// if err := imaging.Save(img, url); err != nil {
	// 	response.FailWithMessage(c, err.Error())
	// 	return
	// }
	
	// @Author: jtan
	// @Description: 存储上传的 video
	playUrl := videoName
	coverUrl := imageName
	videoPost := bo.VideoPost{
		AuthorId: userId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	err = videoService.PostVideo(videoPost)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	
	response.OkWithMessage(c, "上传成功")
}

// VideoFeed 视频流接口
func (api *VideoApi) VideoFeed(c *gin.Context) {
	videoList := make([]*vo.Video, 1, 1)
	user := vo.Author{Id: 1, Name: "charon", FollowCount: 0, FollowerCount: 1, IsFollow: false}
	videoList[0] = &vo.Video{Id: 1, Author: &user, PlayUrl: "http://220.243.147.162:8080/videos/sss.mp4", CoverUrl: "http://220.243.147.162:8080/images/sss.jpg", FavoriteCount: 0, CommentCount: 0, IsFavorite: false}
	data := vo.FeedResponseVo{NextTime: time.Now().Unix(), VideoList: videoList}
	response.OkWithData(c, data)
}

func (api *VideoApi) VideoList(c *gin.Context) {
	bc := c.Keys["claims"]
	claims := bc.(vo.BaseClaims)
	userId := claims.Id
	list, err := videoService.GetVideoList(userId)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	fmt.Println(list)
	response.OkWithData(c, vo.PublishResponseVo{VideoList: list})
}
