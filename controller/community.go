package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.QueryAllCommunity()
	if err != nil {
		zap.L().Error("logic.QueryAllCommunity failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "社区id"
// @Param id query int false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} []models.CommunityDetail
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {

	communityID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.QueryCommunityDetail(communityID)
	if err != nil {
		zap.L().Error("logic.QueryCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
