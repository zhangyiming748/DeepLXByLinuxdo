package model

import (
	"DeepLXByLinuxdo/storage"
	"fmt"
	"time"
)

type TranslateCache struct {
	Id          int64     `xorm:"comment('主键id') INT(11)"`
	Src         string    `xorm:"varchar(255) comment(原文)"`
	Dst         string    `xorm:"varchar(255) comment(译文)"`
	Source_lang string    `xorm:"varchar(255) comment(源语言)"`
	Target_lang string    `xorm:"varchar(255) comment(目标语言)"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	DeletedAt   time.Time `xorm:"deleted"`
}

func init() {

}
func (t *TranslateCache) CreateOne() error {
	if _, err := storage.GetMysql().Insert(t); err != nil {
		return err
	}
	return nil
}
func (t *TranslateCache) FindCache() (bool, error) {
	fmt.Println(t.Src)
	return storage.GetMysql().Where("src=?", t.Src).And("source_lang=?", t.Source_lang).Get(t)
}
