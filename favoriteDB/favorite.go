package favoriteDB

import (
	// "log"

	"github.com/karry-almond/packages/model"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)
// status 返回0——成功，返回1——失败
// err 返回nil——成功，返回其他——失败原因
func NewFavorite(user_id int64, video_id int64) (status int32, err error) {

	// 创建一条favorite数据
	favorite := model.Favorite{
		//TODO:ID这里不是逐主键
		ID:      1,
		UserId:  user_id,
		VideoId: video_id}
	if err := Db.Create(&favorite).Error; err != nil{
		return 1, err
	}

	//更改对应video的favorite_count
	var video model.Video

	if err:= Db.Select("*").First(&model.Video{ID: video_id}).Scan(&video).Error; err != nil{
		return 1, err
	}

	if err := Db.Model(&model.Video{ID: video_id}).Update("favorite_count", video.FavoriteCount+1).Error; err != nil{
		return 1, err
	}

	return 0, nil

}

func CancelFavorite(user_id int64, video_id int64) (status int32, err error) {
	//先根据user_id和video_id寻找到id，再根据id软删除
	var favorite model.Favorite

	if err := Db.Select("*").First(&model.Favorite{UserId: user_id, VideoId: video_id}).Scan(&favorite).Error; err !=nil{
		return 1, err
	}

	if err := Db.Delete(&favorite).Error; err != nil{
		return 1, err
	}

	return 0, nil
}
