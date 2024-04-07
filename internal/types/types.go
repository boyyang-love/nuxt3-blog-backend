// Code generated by goctl. DO NOT EDIT.
package types

type ArticleInfo struct {
	Id uint `json:"id"`
}

type Base struct {
	Code int    `json:"code" default:1`
	Msg  string `json:"msg" default:"ok"`
}

type BlogComments struct {
	Id            uint                 `json:"id"`
	Created       int                  `json:"created"`
	Content       string               `json:"content"`
	ArticleId     uint                 `json:"article_id"`
	CommentId     uint                 `json:"comment_id"`
	WebsiteUserId uint                 `json:"website_user_id"`
	UserId        uint                 `json:"user_id"`
	Type          string               `json:"type"`
	User          BlogCommentsUserInfo `json:"user"`
}

type BlogCommentsUserInfo struct {
	Id       uint   `json:"id"`
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type BlogSearchByIdReq struct {
	Id uint `form:"id"`
}

type BlogSearchByIdRes struct {
	Base
	Data BlogSearchByIdResData `json:"data"`
}

type BlogSearchByIdResData struct {
	Info ListBlogItem `json:"info"`
}

type BlogTags struct {
	Id      uint   `json:"id" json:"id"`
	Uid     string `json:"uid"`
	TagName string `json:"tag_name""`
	Type    string `json:"type"`
}

type BlogUserInfo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Motto    string `json:"motto"`
}

type CommentCreateReq struct {
	Content       string `json:"content"`
	ArticleId     uint   `json:"article_id,optional"`
	CommentId     uint   `json:"comment_id,optional"`
	WebsiteUserId uint   `json:"website_user_id,optional"`
	Type          string `json:"type,options=[article,comment,website]"`
}

type CommentCreateRes struct {
	Base
}

type CommentDeleteReq struct {
	Id uint `json:"id"`
}

type CommentDeleteRes struct {
	Base
}

type CommentInfo struct {
	Id            uint            `json:"id"`
	Created       int             `json:"created"`
	Content       string          `json:"content"`
	ArticleId     uint            `json:"article_id"`
	CommentId     uint            `json:"comment_id"`
	WebsiteUserId uint            `json:"website_user_id"`
	UserId        uint            `json:"user_id"`
	Type          string          `json:"type"`
	User          CommentUserInfo `json:"user"`
}

type CommentListReq struct {
	Id   uint   `form:"id"`
	Type string `form:"type,options=[article,comment,website]"`
}

type CommentListRes struct {
	Base
	Data CommentListResData `json:"data"`
}

type CommentListResData struct {
	Count int64         `json:"count"`
	Info  []CommentInfo `json:"info"`
}

type CommentUserInfo struct {
	Id       uint   `json:"id"`
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type CreateBlogReq struct {
	Title   string `json:"title"`
	Des     string `json:"des"`
	Cover   string `json:"cover"`
	Content string `json:"content"`
	Tags    []uint `json:"tags,optional"`
}

type CreateBlogRes struct {
	Base
	Data CreateBlogResData `json:"data"`
}

type CreateBlogResData struct {
	Id uint `json:"id"`
}

type CreateTagReq struct {
	Name string `json:"name"`
	Type string `json:"type,options=[article,image]"`
}

type CreateTagRes struct {
	Base
}

type DeleteBlogReq struct {
	Id uint `json:"id"`
}

type DeleteBlogRes struct {
	Base
}

type DeleteTagReq struct {
	Id uint `json:"id"`
}

type DeleteTagRes struct {
	Base
}

type EmailSendCodeReq struct {
	Email string `json:"email"`
}

type EmailSendCodeRes struct {
	Base
}

type FileDeleteReq struct {
	Id       uint   `json:"id"`
	FilePath string `json:"file_path"`
	Type     string `json:"type,options=[blog,images,avatar,bg]"`
}

type FileDeleteRes struct {
	Base
}

type FileInfo struct {
	Id       uint   `json:"id"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

type FileListReq struct {
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
	Type  string `form:"type,options=[blog,images,avatar,bg]"`
}

type FileListRes struct {
	Base
	Data FileListResdata `json:"data"`
}

type FileListResdata struct {
	Count int64      `json:"count"`
	Infos []FileInfo `json:"infos"`
}

type FileUploadReq struct {
	Hash     string `json:"hash,optional"`
	FileName string `json:"file_name,optional"`
	Ext      string `json:"ext,optional"`
	Size     int64  `json:"size,optional"`
	FilePath string `json:"file_path,optional"`
}

type FileUploadRes struct {
	Base
	Data FileUploadResdata `json:"data"`
}

type FileUploadResdata struct {
	FileName string `json:"file_name"`
	Path     string `json:"path"`
}

type ListBlogItem struct {
	Id      uint           `json:"id"`
	Uid     string         `json:"uid"`
	Created int64          `json:"created"`
	Updated int64          `json:"updated"`
	Title   string         `json:"title"`
	Des     string         `json:"des"`
	Cover   string         `json:"cover"`
	Content string         `json:"content"`
	UserId  uint           `json:"user_id"`
	User    BlogUserInfo   `json:"user" gorm:"column:User;reference:UserId"`
	Tag     []*BlogTags    `json:"tag" gorm:"column:Tag;many2many:article_tag"`
	Comment []BlogComments `json:"comment" gorm:"column:Comment"`
}

type ListBlogReq struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Type  string `form:"type,optional,options=[top,recently,created]"`
}

type ListBlogRes struct {
	Base
	Data ListBlogResData `json:"data"`
}

type ListBlogResData struct {
	Count int64          `json:"count"`
	List  []ListBlogItem `json:"list"`
}

type ListTagReq struct {
	Type string `form:"type,options=[image,article]"`
}

type ListTagRes struct {
	Base
	Data ListTagResData `json:"data"`
}

type ListTagResData struct {
	Tags []TagInfo `json:"tags"`
}

type SearchReq struct {
	Keyword string `form:"keyword"`
	Type    string `form:"type,options=[article,images]"`
}

type SearchRes struct {
	Base
	Data SearchResData `json:"data"`
}

type SearchResData struct {
	Infos []SearchResDataInfo `json:"infos"`
}

type SearchResDataInfo struct {
	Id      uint   `json:"id"`
	Uid     string `json:"uid"`
	Created int64  `json:"created"`
	Title   string `json:"title"`
	Des     string `json:"des"`
	Cover   string `json:"cover"`
}

type SignInReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInRes struct {
	Base
	Data SignInResData `json:"data"`
}

type SignInResData struct {
	UserInfo UserInfo `json:"user_info"`
	Token    string   `json:"token"`
}

type SignUpReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type SignUpRes struct {
	Base
}

type TagInfo struct {
	Id      uint           `json:"id" form:"id" gorm:"primaryKey"`
	Uid     string         `json:"uid" form:"uid"`
	TagName string         `json:"tag_name" form:"tag_name"`
	Type    string         `json:"type" form:"type" gorm:"type:enum('image', 'article')"`
	Article []*ArticleInfo `json:"articles" gorm:"column:Article;many2many:article_tag"`
}

type UpdateBlogReq struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Des     string `json:"des"`
	Cover   string `json:"cover"`
	Content string `json:"content"`
	Tags    []uint `json:"tags,optional"`
}

type UpdateBlogRes struct {
	Base
}

type UpdateTagReq struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type UpdateTagRes struct {
	Base
}

type UpdateUserPasswordReq struct {
	Password string `json:"password"`
}

type UpdateUserPasswordRes struct {
	Base
}

type UpdateUserReq struct {
	Id       uint   `json":"id"`
	Username string `json:"username,optional"`
	Avatar   string `json:"avatar,optional"`
	Motto    string `json:"motto,optional"`
	Address  string `json:"address,optional"`
	Tel      int    `json:"tel,optional"`
	Email    string `json:"email,optional"`
	QQ       int    `json:"qq,optional"`
	Wechat   string `json:"wechat,optional"`
	GitHub   string `json:"git_hub,optional"`
	Cover    string `json:"cover,optional"`
}

type UpdateUserRes struct {
	Base
}

type UserInfo struct {
	Id       uint   `json:"id"`
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Cover    string `json:"cover"`
	Avatar   string `json:"avatar"`
	Account  string `json:"account"`
	Motto    string `json:"motto"`
	Address  string `json:"address"`
	Tel      int    `json:"tel"`
	Email    string `json:"email"`
	QQ       int    `json:"qq"`
	Wechat   string `json:"wechat"`
	GitHub   string `json:"git_hub"`
}
