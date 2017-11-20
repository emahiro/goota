package model

import (
	"time"
)

type Qiita struct {
	RenderBody     string    `json:"render_body"`
	Body           string    `json:"body"`
	Coediting      bool      `json:"coediting"`
	CommentsCount  int64     `json:"comments_count"`
	Group          struct{}  `json:"group"`
	ID             int64     `json:"id"`
	LikesCount     int64     `json:"likes_count"`
	Private        bool      `json:"private"`
	ReactionsCount int64     `json:"reactions_count"`
	Tags           []Tag     `json:"tags"`
	Title          string    `json:"title"`
	UpdatedAt      time.Time `json:"updated_at"`
	Url            string    `json:"url"`
	User           User      `json:"user"`
}

type Tag struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type User struct {
	Description       string `json:"description"`
	FacebookID        string `json:"facebook_id"`
	FolloweesCount    int64  `json:"followees_count"`
	FollowersCount    int64  `json:"followers_count"`
	GithubLoginName   string `json:"github_login_name"`
	ID                string `json:"id"`
	ItemsCount        int64  `json:"items_count"`
	LinkedinID        string `json:"linkedin_id"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	Organization      string `json:"organization"`
	PermanentID       int64  `json:"permanent_id"`
	ProfileImageUrl   string `json:"profile_image_url"`
	TwitterScreenName string `json:"twitter_screen_name"`
	WebSiteUrl        string `json:"web_site_url"`
}

type Goota struct {
	Qiita
}
