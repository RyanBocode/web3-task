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

	// 初始化数据
	seed(gdb)

	// 查询
	assocQueries(gdb)
}

func seed(db *gorm.DB) {
	// 删除已有数据
	db.Where("1=1").Delete(&Comment{})
	db.Where("1=1").Delete(&Post{})
	db.Where("1=1").Delete(&User{})

	// 创建用户、文章、评论
	u := User{Name: "Ryan"}
	db.Create(&u)

	p1 := Post{Title: "第一篇", Content: "Hello", UserID: u.ID, CommentStatus: "无评论"}
	p2 := Post{Title: "第二篇", Content: "World", UserID: u.ID, CommentStatus: "无评论"}
	db.Create(&p1)
	db.Create(&p2)

	// 评论
	db.Create(&Comment{Content: "赞", PostID: p1.ID})
	db.Create(&Comment{Content: "学到", PostID: p1.ID})
	db.Create(&Comment{Content: "继续", PostID: p2.ID})

	// 更新评论状态
	db.Model(&Post{}).Where("id IN ?", []uint{p1.ID, p2.ID}).Update("comment_status", "有评论")
}

func assocQueries(db *gorm.DB) {
	var u User
	db.Where("name=?", "Ryan").First(&u)

	var posts []Post
	db.Preload("Comments").Where("user_id=?", u.ID).Order("id").Find(&posts)

	// 输出用户文章和评论
	fmt.Println("该用户的文章及评论：")
	for _, p := range posts {
		fmt.Printf("#%d %s (%s)\n", p.ID, p.Title, p.CommentStatus)
		for _, c := range p.Comments {
			fmt.Println("  -", c.Content)
		}
	}

	// 查询评论最多的文章
	type PostWithCnt struct {
		Post
		CommentCount int64
	}
	var postWithCnt PostWithCnt
	db.Table("posts").
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&postWithCnt)

	fmt.Printf("\n评论最多的文章: %s (评论数: %d)\n", postWithCnt.Title, postWithCnt.CommentCount)
}
