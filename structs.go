package OreGo

import (
	"github.com/valyala/fasthttp"
)

type Ore struct {
	Client fasthttp.Client
}

//users
type User struct {
	ID int64 `json:"id"`
	CreatedAt string `json:"createdAt"`
	Username string `json:"username"`
	Roles []string `json:"roles"`
	Starred []string `json:"starred"`
	AvatarTemplate string `json:"avatarTemplate"`
	Projects []Project `json:"projects"`
}

//projects
type Project struct {
	PluginID string `json:"pluginId"`
	CreatedAt string `json:"createdAt"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	Description string `json:"description"`
	Href string `json:"href"`
	Members []ProjectMember `json:"members"`
	Channels []Channel `json:"channels"`
	RecommendedVersion Version `json:"recommended"`
	Category Category `json:"category"`
	Views int64 `json:"views"`
	Downloads int64 `json:"downloads"`
	Stars int64 `json:"stars"`
}

type Category struct {
	Title string `json:"title"`
	Icon string `json:"icon"`
}

type Version struct {
	ID int64 `json:"id"`
	CreatedAt string `json:"createdAt"`
	Name string `json:"name"`
	Dependencies []Dependency `json:"dependencies"`
	PluginID string `json:"pluginId"`
	Channel Channel `json:"channel"`
	FileSize int64 `json:"fileSize"`
	MD5 string `json:"md5"`
	StaffApproved bool `json:"staffApproved"`
	HRef string `json:"href"`
}

type Dependency struct {
	PluginID string `json:"pluginID"`
	Version string `json:"version"`
}

type Channel struct {
	Name string `json:"name"`
	Color string `json:"color"`
}

type ProjectMember struct {
	UserID int64 `json:"userId"`
	Name string `json:"name"`
	Roles []string `json:"roles"`
	HeadRole string `json:"headRole"`
}