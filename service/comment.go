package service

import (
	"backend/global"
	"backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BuildAPIComment(comment model.Comment) APIComment {
	return APIComment{
		ID:          comment.ID,
		Text:        comment.Text,
		CommentUser: GetCommentUser(&comment),
		CreatedAt:   comment.CreatedAt.Local(),
	}
}

func BuildAPIComments(comments []model.Comment) (apiComments []APIComment) {
	for _, comment := range comments {
		apiComment := BuildAPIComment(comment)
		apiComments = append(apiComments, apiComment)
	}
	return
}

// Create creates comment
func (service *CreateCommentService) Create(c *gin.Context) Response {
	comment := model.Comment{
		Text:    service.Text,
		TweetID: service.TweetID,
		UserID:  CurrentUser(c).ID,
	}

	// Create Comment
	if err := global.GORM_DB.Create(&comment).Error; err != nil {
		return Response{
			Code:  50001,
			Msg:   "creating comment failed",
			Error: err.Error(),
		}
	}
	apiComment := BuildAPIComment(comment)

	return Response{
		Data: apiComment,
		Msg:  "comment created",
	}
}

func (service *DeleteCommentService) Delete(id string) Response {
	var comment model.Comment

	commentId, _ := strconv.Atoi(id)
	err := global.GORM_DB.First(&comment, uint(commentId)).Error
	if err != nil {
		return Response{
			Code:  404,
			Msg:   "comment doesn't exist",
			Error: err.Error(),
		}
	}

	err = global.GORM_DB.Delete(&comment).Error
	if err != nil {
		return Response{
			Code:  404,
			Msg:   "deleting comment failed",
			Error: err.Error(),
		}
	}

	return Response{
		Msg: "comment deleted",
	}
}

func GetCommentUser(comment *model.Comment) APIUser {
	commentUser, _ := GetUser(comment.UserID)
	return BuildAPIUser(commentUser)
}
