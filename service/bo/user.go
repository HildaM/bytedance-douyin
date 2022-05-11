package bo

/**
 * @Author: 1999single
 * @Description:
 * @File: user
 * @Version: 1.0.0
 * @Date: 2022/5/11 1:14
 */
type UserInfoBo struct {
	ID            int64
	Name          string
	FollowCount   int
	FollowerCount int
	Follow        bool
}

type UserBo struct {
	Name string
	Pwd  string
}

type UserRegisterBo struct {
	Id    int64
	Token string
}

type UserLoginBo struct {
	Id int64
}
