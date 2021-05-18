package service

import (
	"backend/global"
	"backend/model"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	// PassWordCost
	PassWordCost = 12
	// Active user
	Active string = "active"
	// Inactive user
	Inactive string = "inactive"
	// Suspended user
	Suspend string = "suspend"
)

// setSession set session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login
func (service *UserLoginService) Login(c *gin.Context) Response {
	var user model.User

	if err := global.GORM_DB.Where("email = ?", service.Email).First(&user).Error; err != nil {
		return ParamErr("wrong email or password", nil)
	}

	if !CheckPassword(&user, service.Password) {
		return ParamErr("wrong email or password", nil)
	}
	// set session
	service.setSession(c, user)

	apiUser := BuildAPIUser(user)

	return Response{
		Data: apiUser,
		Msg:  "Logged in",
	}
}

// valid form
func (service *UserRegisterService) valid() *Response {
	if service.PasswordConfirm != service.Password {
		return &Response{
			Code: 40001,
			Msg:  "passwords don't match",
		}
	}

	var count int64
	global.GORM_DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &Response{
			Code: 40001,
			Msg:  "nickname not available",
		}
	}

	count = 0
	global.GORM_DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &Response{
			Code: 40001,
			Msg:  "username not available",
		}
	}

	count = 0
	global.GORM_DB.Model(&model.User{}).Where("email = ?", service.Email).Count(&count)
	if count > 0 {
		return &Response{
			Code: 40001,
			Msg:  "email not available",
		}
	}

	return nil
}

// Register
func (service *UserRegisterService) Register() Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Email:    service.Email,
		Status:   Active,
	}

	// validate
	if err := service.valid(); err != nil {
		return *err
	}

	// encrypt
	if err := SetPassword(&user, service.Password); err != nil {
		return Err(
			CodeEncryptError,
			"encryption failed",
			err,
		)
	}

	// Create user
	if err := global.GORM_DB.Create(&user).Error; err != nil {
		return ParamErr("register failed", err)
	}

	return Response{
		Msg: "Registered Ok",
	}
}

func (service *ListUserService) List(c *gin.Context) Response {
	users := []model.User{}
	apiListUsers := []APIListUser{}
	var total int64

	if service.Limit == 0 {
		service.Limit = 12
	}

	if err := global.GORM_DB.Model(model.User{}).Count(&total).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "database connection error",
			Error: err.Error(),
		}
	}

	if err := global.GORM_DB.Limit(service.Limit).Offset(service.Start).Find(&users).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "database connection error",
			Error: err.Error(),
		}
	}

	for _, user := range users {
		apiListUser := BuildAPIListUser(c, user)
		apiListUsers = append(apiListUsers, apiListUser)
	}

	return Response{
		Data: DataList{Items: apiListUsers, Total: uint(total)},
		Msg:  "Got users",
	}
}

func (service *UpdateUserService) Update(c *gin.Context) Response {
	user := model.User{
		Nickname: service.Nickname,
		Avatar:   service.Avatar,
		Bio:      service.Bio,
		Website:  service.Website,
		CoverPhoto: service.CoverPhoto,
		Location: service.Location,
	}
	
	var count int64
	global.GORM_DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return Response{
			Code: 40001,
			Msg:  "nickname not available",
		}
	}

	if err := global.GORM_DB.Where("id = ?", CurrentUser(c).ID).Updates(&user).Error; err != nil {
		return ParamErr("updating user failed", err)
	}

	apiUser := BuildAPIProfile(c, user)

	return Response{
		Data: apiUser,
		Msg:  "user updated",
	}
}

// CurrentUser get current user
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// GetUser get user by id
func GetUser(ID interface{}) (model.User, error) {
	var user model.User
	result := global.GORM_DB.First(&user, ID)
	return user, result.Error
}

func BuildAPIUser(user model.User) APIUser {
	return APIUser{
		ID:       user.ID,
		UserName: user.UserName,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
}

func BuildAPIProfile(c *gin.Context, user model.User) APIProfile {
	currentUser := CurrentUser(c)
	return APIProfile{
		ID:             user.ID,
		UserName:       user.UserName,
		Nickname:       user.Nickname,
		Avatar:         user.Avatar,
		Status:         user.Status,
		Email:          user.Email,
		Bio:            user.Bio,
		Website:        user.Website,
		Location:       user.Location,
		Dob:            user.Dob,
		CoverPhoto:     user.CoverPhoto,
		Tweets:         GetTweets(c, &user),
		TweetCount:     GetTweetCount(&user),
		Followers:      GetFollowers(&user),
		Followings:     GetFollowings(&user),
		FollowerCount:  GetFollowerCount(&user),
		FollowingCount: GetFollowingCount(&user),
		IsFollowing:    IsFollowing(currentUser.ID, &user),
		IsMe:           IsMe(&user, currentUser.ID),
	}
}

func BuildAPIUsers(users []model.User) (apiUsers []APIUser) {
	for _, user := range users {
		apiUser := BuildAPIUser(user)
		apiUsers = append(apiUsers, apiUser)
	}
	return
}

func BuildAPIListUser(c *gin.Context, user model.User) APIListUser {
	currentUser := CurrentUser(c)
	return APIListUser{
		ID:          user.ID,
		UserName:    user.UserName,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		IsFollowing: IsFollowing(currentUser.ID, &user),
		IsMe:        IsMe(&user, currentUser.ID),
	}
}

// SetPassword set password
func SetPassword(user *model.User, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword check password
func CheckPassword(user *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func GetTweetCount(user *model.User) int64 {
	return global.GORM_DB.Model(&user).Association("Tweets").Count()
}

func GetTweets(c *gin.Context, user *model.User) (apiTweets []APITweet) {
	var tweets []model.Tweet
	global.GORM_DB.Model(&user).Order("created_at desc").Association("Tweets").Find(&tweets)
	return BuildAPITweets(c, tweets)
}

func GetFollowers(user *model.User) (apiUsers []APIUser) {
	var follows []model.Follow
	var user_ids []uint
	var users []model.User
	global.GORM_DB.Where("following=?", user.ID).Find(&follows)
	for _, follow := range follows {
		user_ids = append(user_ids, follow.Follower)
	}
	global.GORM_DB.Where(user_ids).Find(&users)
	return BuildAPIUsers(users)
}

func GetFollowings(user *model.User) (apiUsers []APIUser) {
	var follows []model.Follow
	var user_ids []uint
	var users []model.User
	global.GORM_DB.Where("follower=?", user.ID).Find(&follows)
	for _, follow := range follows {
		user_ids = append(user_ids, follow.Following)
	}
	global.GORM_DB.Where(user_ids).Find(&users)
	return BuildAPIUsers(users)
}

func GetFollowingCount(user *model.User) (count int64) {
	global.GORM_DB.Where("follower=?", user.ID).Find(&model.Follow{}).Count(&count)
	return
}

func GetFollowerCount(user *model.User) (count int64) {
	global.GORM_DB.Where("following=?", user.ID).Find(&model.Follow{}).Count(&count)
	return
}

func IsFollowing(follower uint, user *model.User) bool {
	results := global.GORM_DB.Where("follower=? And following=?", follower, user.ID).Find(&model.Follow{})
	return results.RowsAffected > 0
}

func IsMe(user *model.User, otherID uint) bool {
	return user.ID == otherID
}

func GetProfile(c *gin.Context, user_id string) Response {
	id, _ := strconv.Atoi(user_id)
	user, _ := GetUser(uint(id))
	return Response{
		Data: BuildAPIProfile(c, user),
		Msg:  "got profile",
	}
}
