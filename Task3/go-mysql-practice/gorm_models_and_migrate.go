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
	Posts     []Post
}

type Post struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	Content       string
	UserID        uint
	Comments      []Comment
	CommentStatus string
}

type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Content string
	PostID  uint
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

	// 自动迁移模型
	must(gdb.AutoMigrate(&User{}, &Post{}, &Comment{}))
	p("AutoMigrate 完成 (User/Post/Comment)")

	// 创建一些测试数据
	user := User{Name: "张三", PostCount: 0}
	gdb.Create(&user)

	post := Post{Title: "第一篇博客", Content: "Hello World!", UserID: user.ID, CommentStatus: "无评论"}
	gdb.Create(&post)

	comment := Comment{Content: "很好的文章！", PostID: post.ID}
	gdb.Create(&comment)

	p("测试数据创建完成")
	p("用户: %s", user.Name)
	p("文章: %s", post.Title)
	p("评论: %s", comment.Content)
}
