package apiserver

import "time"

// Article 表示博客文章
type Article struct {
	// PostID 文章ID
	ArticleID string `json:"articleID"`
	// UserID 作者ID
	UserID string `json:"userID"`
	// Title 标题
	Title string `json:"title"`
	// Content 正文
	Content string `json:"content"`
	// Abstract 摘要
	Abstract string `json:"abstract"`
	// CategoryID 分类ID
	CategoryID string `json:"categoryID"`
	// CategoryName 分类名（冗余字段，返回时展示）
	CategoryName string `json:"categoryName"`
	// Tags 标签（多个标签 ID）
	Tags []string `json:"tags"`
	// Cover 封面图
	Cover string `json:"cover"`
	// View 浏览数
	View int64 `json:"view"`
	// Like 点赞数
	Like int64 `json:"like"`
	// Collect 收藏数
	Collect int64 `json:"collect"`
	// Comment 评论数
	Comment int64 `json:"comment"`
	// IsRecommend 是否推荐
	IsRecommend bool `json:"isRecommend"`
	// IsRelease 是否发布
	IsRelease bool `json:"isRelease"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt 最后更新时间
	UpdatedAt time.Time `json:"updatedAt"`
	// PublishedAt 发布时间
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
}

//
// ----------- 创建文章 -----------
//

// CreateArticleRequest 创建文章请求
type CreateArticleRequest struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Abstract    string   `json:"abstract"`
	CategoryID  string   `json:"categoryID"`
	Tags        []string `json:"tags"`
	Cover       string   `json:"cover"`
	IsRecommend bool     `json:"isRecommend"`
	IsRelease   bool     `json:"isRelease"`
}

// CreateArticleResponse 创建文章响应
type CreateArticleResponse struct {
	articleID string `json:"articleID"`
}

//
// ----------- 更新文章 -----------
//

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	ArticleID   string    `json:"articleID" uri:"articleID"`
	Title       *string   `json:"title"`
	Content     *string   `json:"content"`
	Abstract    *string   `json:"abstract"`
	CategoryID  *string   `json:"categoryID"`
	Tags        *[]string `json:"tags"`
	Cover       *string   `json:"cover"`
	IsRecommend *bool     `json:"isRecommend"`
	IsRelease   *bool     `json:"isRelease"`
}

// UpdateArticleResponse 更新文章响应
type UpdateArticleResponse struct {
	Success bool `json:"success"`
}

//
// ----------- 删除文章 -----------
//

// DeleteArticleRequest 删除文章请求
type DeleteArticleRequest struct {
	ArticleIDs []string `json:"articleIDs" binding:"required"`
}

// DeleteArticleResponse 删除文章响应
type DeleteArticleResponse struct {
	Success bool `json:"success"`
	Count   int  `json:"count"`
}

//
// ----------- 获取单篇文章 -----------
//

// GetArticleRequest 获取文章请求
type GetArticleRequest struct {
	ArticleID string `json:"articleID" uri:"articleID"`
}

// GetArticleResponse 获取文章响应
type GetArticleResponse struct {
	Article *Article `json:"article"`
}

//
// ----------- 获取文章列表 -----------
//

// ListArticleRequest 获取文章列表请求
type ListArticleRequest struct {
	Offset   int64   `form:"offset"`
	Limit    int64   `form:"limit"`
	Title    *string `form:"title"`
	Category *string `form:"categoryID"`
	Tag      *string `form:"tagID"`
	Release  *bool   `form:"isRelease"`
}

// ListArticleResponse 获取文章列表响应
type ListArticleResponse struct {
	TotalCount int64      `json:"total_count"`
	Articles   []*Article `json:"articles"`
}
