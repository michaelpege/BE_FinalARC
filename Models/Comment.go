package Models

import (
	"github.com/michaelpege/arc/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetCommentsByPostId(comments *[]Comment, id string)(err error){
	if err = Config.DB.Where("post_id = ?",id).Find(&comments).Error; err != nil {
  return err
 }
 return nil
}

func PostCommentByPostId(comment *Comment, id string)(err error){
	if err = Config.DB.Create(comment).Error; err != nil {
  return err
 }
 return nil
}