package controller

import "github.com/BaiZe1998/douyin-simple-demo/dto"

type Response struct {
	StatusCode int32  `json:"StatusCode"`
	StatusMsg  string `json:"StatusMsg,omitempty"`
	NextTime   int64  `json:"next_time,omitempty"`
}

type Video struct {
	Id            int64    `json:"id,omitempty"`
	Author        dto.User `json:"author"`
	PlayUrl       string   `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64    `json:"id,omitempty"`
	User       dto.User `json:"dto.user"`
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`
}

type User struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
