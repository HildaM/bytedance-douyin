package vo

/**
 * @Author: 1999single
 * @Description:
 * @File: comment
 * @Version: 1.0.0
 * @Date: 2022/5/11 1:29
 */
type Comment struct {
	Id         int64     `json:"id"`
	User       *UserInfo `json:"user"`
	Content    string    `json:"content"`
	CreateDate string    `json:"create_date" time_format:"01-02"`
}