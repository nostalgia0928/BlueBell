package controller

import (
	"BlueBell/dao/mysql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 社区

// CommunityHandler 社区列表
func CommunityHandler(c *gin.Context) {
	communityList, err := mysql.GetCommunityList() //没走logic层
	if err != nil {
		zap.L().Error("mysql.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	communityID := c.Param("id")
	communityList, err := mysql.GetCommunityDetailByID(communityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityID() failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeSuccess, err.Error())
		return
	}
	ResponseSuccess(c, communityList)
}
