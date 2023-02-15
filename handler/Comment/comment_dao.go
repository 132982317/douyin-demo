package Comment

import (
	"douyin.core/Model"
	user "douyin.core/handler/User"
	"errors"
	"time"
)

type Comment struct {
	Content    string     `json:"content"`     // 评论内容
	CreateDate string     `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64      `json:"id"`          // 评论id
	User       *user.User `json:"user"`        // 评论用户信息
}

type CommentDao struct{}

func NewCommentDao() *CommentDao {
	return &CommentDao{}
}

func (cmt *CommentDao) AddComment(userid int64, content string, user *user.UserInfoDao) (error, Comment) {
	userInfo, err := user.GetUserByuserID(userid)
	if err != nil {
		return errors.New("user Not found"), Comment{}
	}
	newCmt := Comment{
		Content:    content,
		CreateDate: time.Now().String(),
		User:       userInfo,
	}
	return Model.DB.Create(newCmt).Error, newCmt
}

func (cmt *CommentDao) DeleteComment(cmtid string) error {
	comment := Comment{}
	return Model.DB.Where("CommentID = ?", cmtid).Delete(&comment).Error
}

func (cmt *CommentDao) GetCommentList(videoID string) ([]Comment, error) {
	var comments []Comment
	err := Model.DB.Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}