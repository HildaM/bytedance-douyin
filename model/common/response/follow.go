package response

type FollowResponse struct {
	*BasicResponse
	UserList []*UserInfo `json:"user_list"`
}

type FollowerResponse struct {
	*FollowResponse
}
