package database

import (
	"blog/internal/config"
	"blog/internal/models"
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(cfg *config.Config) error {
	db, err := ConnectAndMigrate(cfg)
	if err != nil {
		return err
	}
	DB = db
	log.Printf("数据库连接成功: %s", cfg.DBDriver)
	return nil
}

func ConnectAndMigrate(cfg *config.Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	if cfg.DBDriver == "sqlite" {
		// SQLite配置
		sqliteConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // 显示SQL日志
		}
		db, err = gorm.Open(sqlite.Open(cfg.SQLitePath), sqliteConfig)
		if err != nil {
			return nil, fmt.Errorf("SQLite连接失败: %v", err)
		}
		log.Printf("SQLite数据库文件: %s", cfg.SQLitePath)
	} else if cfg.DBDriver == "mysql" {
		if cfg.MySQLDSN == "" {
			return nil, errors.New("MYSQL_DSN is empty")
		}
		db, err = gorm.Open(mysql.Open(cfg.MySQLDSN), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("MySQL连接失败: %v", err)
		}
	} else {
		return nil, fmt.Errorf("unsupported DB_DRIVER: %s", cfg.DBDriver)
	}
	if err != nil {
		return nil, err
	}

	// 自动迁移
	log.Println("开始数据库迁移...")
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		return nil, fmt.Errorf("数据库迁移失败: %v", err)
	}
	log.Println("数据库迁移完成")

	return db, nil
}
