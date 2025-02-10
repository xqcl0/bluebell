package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

func QueryAllCommunity() ([]*models.Community, error) {
	var data []*models.Community
	sqlStr := "select community_id , community_name from community"
	err := db.Select(&data, sqlStr)
	if errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("no community in db")
		err = nil
	}
	return data, err
}

func QueryCommunityDetail(id int64) (*models.CommunityDetail, error) {
	var data *models.CommunityDetail
	sqlStr := "select community_id , community_name, introduction from community where community.community_id = ?"
	err := db.Get(&data, sqlStr, id)
	if errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("no communityDetail in db")
		err = ErrorInvalidID
	}
	return data, err
}
