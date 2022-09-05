package dao

import (
	"redis-learn/entity"

	"gorm.io/gorm"
)

const TbBlog = "tb_blog"

func QueryHotBlog(current int, pageSize int) []entity.Blog {
	var blogs []entity.Blog
	db.Table(TbBlog).Offset((current - 1) * pageSize).Limit(pageSize).Find(&blogs)
	return blogs
}
func SaveBlog(blog *entity.Blog) error {
	return db.Table(TbBlog).Create(blog).Error
}
func SelectBlogById(id int) (entity.Blog, error) {
	var blog entity.Blog
	err := db.Table(TbBlog).Where("id = ?", id).First(&blog).Error
	return blog, err
}
func AddBlogLike(blogId int) int {
	rowsAffected, err := addBlogLike(blogId, 1)
	if err != nil {
		return 0
	} else {
		return rowsAffected
	}
}
func SubBlogLike(blogId int) int {
	rowsAffected, err := addBlogLike(blogId, -1)
	if err != nil {
		return 0
	} else {
		return rowsAffected
	}
}
func addBlogLike(blogId, icr int) (int, error) {
	result := db.Table(TbBlog).Where("id = ?", blogId).Update("liked", gorm.Expr("liked + ?", icr))
	return int(result.RowsAffected), result.Error
}
