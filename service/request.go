package service

// UserLoginService login service
type UserLoginService struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserRegisterService register service
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Email           string `form:"email" json:"email" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

type ListUserService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

type UpdateUserService struct {
	Nickname   string `form:"nickname" json:"nickname"`
	Avatar     string `form:"avatar" json:"avatar"`
	Bio        string `form:"bio" json:"bio"`
	Website    string `form:"website" json:"website"`
	CoverPhoto string `form:"coverPhoto" json:"coverPhoto"`
	Location   string `form:"location" json:"location"`
}

// Tweet list service
type ListTweetService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

type CreateTweetService struct {
	Text       string `form:"text" json:"text" binding:"required,max=220"`
	Tags       string `form:"tags" json:"tags"`
	TweetFiles []struct {
		Url string `form:"url" json:"url"`
	} `form:"files" json:"files"`
}

type DeleteTweetService struct{}

// CreateCommentService
type CreateCommentService struct {
	Text    string `form:"text" json:"text" binding:"required,max=220"`
	TweetID uint   `form:"tweet_id" json:"tweet_id" binding:"required"`
}

type DeleteCommentService struct{}

// FollowService
type FollowService struct {
	Following uint `form:"following" json:"following" binding:"required"`
}

// LikeService
type LikeService struct {
	TweetID uint `form:"tweet_id" json:"tweet_id" binding:"required"`
}

// RetweetService
type RetweetService struct {
	TweetID uint `form:"tweet_id" json:"tweet_id" binding:"required"`
}
