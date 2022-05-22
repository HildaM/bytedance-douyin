package api

import (
	"bytedance-douyin/api/response"
	"bytedance-douyin/api/vo"
	"bytedance-douyin/exceptions"
	"bytedance-douyin/global"
	"bytedance-douyin/service/bo"
	"bytedance-douyin/utils"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
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

	title, ok := c.GetPostForm("title")
	if !ok {
		response.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}

	_, err = os.Executable()
	if err != nil {
		panic(err)
	}
	//exPath := filepath.Dir(ex)
	//fmt.Println(exPath)

	// 通过token获取用户名及用户id
	path := global.GVA_CONFIG.File.VideoOutput
	videoName := utils.GenerateFilename(username, userId) + ".mp4"
	fn := path + videoName

	// 上传file文件到指定的文件路径fn
	if err := c.SaveUploadedFile(file, fn); err != nil {
		fmt.Println("存储视频失败, filePath,", fn)
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
	url := imagePath + imageName
	if err := imaging.Save(img, url); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	// @Author: jtan
	// @Description: 存储上传的 video
	playUrl := videoName
	coverUrl := imageName
	videoPost := bo.VideoPost{
		AuthorId: userId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
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
	// 这个接口没有经过鉴权，所以要解析token
	token := c.Query("token")
	tokenId := int64(0)
	// 登录了
	if token != "" {
		j := utils.NewJWT()
		claims, err := j.ParseTokenRedis(token)
		if err != nil {
			response.FailWithMessage(c, err.Error())
		}
		tokenId = claims.BaseClaims.Id
	}

	var feed vo.FeedVo
	if err := c.ShouldBind(&feed); err != nil {
		response.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}

	t := feed.LatestTime
	if t == 0 {
		t = time.Now().Unix()
	}

	// 获取视频流
	videos, err := videoService.GetVideoFeed(tokenId, t)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	nextTime := videos[len(videos)-1].CreatedAt.Unix()

	data := vo.FeedResponseVo{NextTime: nextTime, VideoList: videos}
	response.OkWithData(c, data)
}

// VideoList 视频发布列表
func (api *VideoApi) VideoList(c *gin.Context) {
	tokenId, ok := c.Get("tokenId")
	if !ok {
		response.FailWithMessage(c, exceptions.ParamValidationError.Error())
		return
	}
	userIdStr := c.Query("user_id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	list, err := videoService.GetVideoList(int64(userId), tokenId.(int64))
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	response.OkWithData(c, vo.PublishResponseVo{VideoList: list})
}
