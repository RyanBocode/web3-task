package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	PostCount int
}

type Post struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	Content       string
	UserID        uint
	CommentStatus string
}

type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Content string
	PostID  uint
}

// 钩子：文章创建后更新用户的 post_count
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id=?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count+1")).Error
}

// 钩子：删除评论后，如果该文章没有评论，更新状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var cnt int64
	if err := tx.Model(&Comment{}).Where("post_id=?", c.PostID).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt == 0 {
		return tx.Model(&Post{}).Where("id=?", c.PostID).Update("comment_status", "无评论").Error
	}
	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func p(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

func main() {
	// 使用SQLite内存数据库，不需要外部服务
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	must(err)

	// 自动迁移
	must(gdb.AutoMigrate(&User{}, &Post{}, &Comment{}))

	// 初始化数据
	seed(gdb)

	// 删除评论后触发钩子
	gdb.Where("post_id=?", 1).Delete(&Comment{})

	p("钩子测试完成！")
}

func seed(db *gorm.DB) {
	db.Where("1=1").Delete(&Comment{})
	db.Where("1=1").Delete(&Post{})
	db.Where("1=1").Delete(&User{})

	u := User{Name: "Ryan"}
	db.Create(&u)

	p := Post{Title: "First Post", Content: "Hello!", UserID: u.ID, CommentStatus: "无评论"}
	db.Create(&p)

	db.Create(&Comment{Content: "Nice!", PostID: p.ID})
	db.Create(&Comment{Content: "Great!", PostID: p.ID})

	// 更新评论状态
	db.Model(&Post{}).Where("id=?", p.ID).Update("comment_status", "有评论")
}
