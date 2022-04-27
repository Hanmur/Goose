package validator

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetArticleRequest struct {
	Title     string `form:"title" binding:"required,min=1,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=1,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=1,max=100"`
	Desc          string `form:"desc" binding:"required,min=1,max=250"`
	Content       string `form:"content" binding:"required"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=250"`
	Content       string `form:"content" binding:""`
	CoverImageUrl string `form:"cover_image_url" binding:"max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
