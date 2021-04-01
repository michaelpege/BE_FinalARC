package Models

import (
	"github.com/michaelpege/arc/Config"

	_ "github.com/go-sql-driver/mysql"
)

//get all post
func GetAllPosts(posts *[]Post) (err error) {
 if err = Config.DB.Find(posts).Error; err != nil {
  return err
 }
 return nil
}

func GetPostById(post *Post, id string) (err error) {
 if err = Config.DB.Where("id = ?",id).First(post).Error; err != nil {
  return err
 }
 return nil
}