package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func QueryAllCommunity() ([]*models.Community, error) {
	data, err := mysql.QueryAllCommunity()
	return data, err
}

func QueryCommunityDetail(id int64) (*models.CommunityDetail, error) {
	data, err := mysql.QueryCommunityDetail(id)
	return data, err
}
