package database

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type GormModel struct {
	Id        int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResolverDSN struct {
	Sources   []string `yaml:"Sources"`
	Replicas  []string `yaml:"Replicas"`
	Secondary []string `yaml:"Secondary"`
}

var (
	db *gorm.DB
)

func OpenDB(dsn ResolverDSN, config *gorm.Config, maxIdleConnes, maxOpenConnes int) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if db, err = gorm.Open(mysql.Open(dsn.Sources[0]), config); err != nil {
		log.Errorf("opens database failed: %s", err.Error())
		return
	}

	if err = db.Use(
		dbresolver.Register(
			dbresolver.Config{
				Sources:  buildDialector(dsn.Sources),
				Replicas: buildDialector(dsn.Replicas),
				Policy:   dbresolver.RandomPolicy{},
			},
		).Register(
			dbresolver.Config{Replicas: buildDialector(dsn.Secondary)},
			"secondary",
		).SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(maxIdleConnes).
			SetMaxOpenConns(maxOpenConnes)); err != nil {
		log.Errorf("opens database failed: %s", err.Error())
		return
	}

	return
}

// 获取数据库链接
func DB() *gorm.DB {
	return db
}

func buildDialector(dsns []string) []gorm.Dialector {
	var res []gorm.Dialector
	for _, dsn := range dsns {
		res = append(res, mysql.Open(dsn))
	}
	return res
}

// AutoMigrate ...
func AutoMigrate(models ...interface{})  {
	if err := db.AutoMigrate(models...); nil != err {
		log.Errorf("auto migrate tables failed: %s", err.Error())
	}
}


