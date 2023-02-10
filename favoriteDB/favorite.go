package favoriteDB

import (
	"context"
	"log"

	"github.com/karry-almond/packages/model"
)

func NewFavorite(ctx context.Context, user_id int64, video_id int64) (status int32, err error) {

	// 创建一条favorite数据
	favorite := model.Favorite{
		//TODO:ID这里不是逐主键
		ID:      1,
		UserId:  user_id,
		VideoId: video_id}
	result := Db.Create(&favorite)

	//更改对应video的favorite_count
	var video model.Video

	Db.Select("*").First(&model.Video{ID: video_id}).Scan(&video)

	Db.Model(&model.Video{ID: video_id}).Update("favorite_count", video.FavoriteCount+1)

	log.Println(result)

	return 1, nil

}

func CancelFavorite(ctx context.Context, user_id int64, video_id int64) (status int32, err error) {
	//先根据user_id和video_id寻找到id，再根据id软删除
	var favorite model.Favorite

	Db.Select("*").First(&model.Favorite{UserId: user_id, VideoId: video_id}).Scan(&favorite)

	Db.Delete(&favorite)

	return 1, nil
}
