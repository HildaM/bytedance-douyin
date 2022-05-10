package repository

/**
 * @Author: 1999single
 * @Description:
 * @File: enter
 * @Version: 1.0.0
 * @Date: 2022/5/11 0:26
 */

type Group struct {
	UserDao UserDao
}

var GroupApp = new(Group)
