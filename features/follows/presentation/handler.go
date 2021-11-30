package presentation

import (
	"net/http"
	"prog/features/follows"
	"prog/features/follows/presentation/response"
	"prog/middlewares"

	"strconv"

	"github.com/labstack/echo/v4"
)

type FollowsHandler struct {
	FollowsBusiness follows.Business
}

func NewArticleLikesHandler(followsBusiness follows.Business) *FollowsHandler {
	return &FollowsHandler{FollowsBusiness: followsBusiness}
}

func (uh *FollowsHandler) FollowUser(e echo.Context) error {

	followingUserId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not follow user",
			"err":     "userId must be integer",
		})
	}
	followerUserId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not follow user",
			"err":     err.Error(),
		})
	}
	err = uh.FollowsBusiness.FollowUser(followingUserId, followerUserId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not follow user",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "follow user",
	})
}

func (uh *FollowsHandler) UnfollowUser(e echo.Context) error {

	followingUserId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not unfollow user",
			"err":     "userId must be integer",
		})
	}
	followerUserId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not unfollow user",
			"err":     err.Error(),
		})
	}
	err = uh.FollowsBusiness.UnfollowUser(followingUserId, followerUserId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not unfollow user",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "unfollow user",
	})
}

func (alh *FollowsHandler) GetFollowingUsers(e echo.Context) error {
	userId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get following user",
			"err":     err.Error(),
		})
	}
	data, err := alh.FollowsBusiness.GetFollowingUsers(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get following user",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToUserResponseList(data),
	})

}

func (alh *FollowsHandler) GetFollowersUsers(e echo.Context) error {
	userId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get followers user",
			"err":     err.Error(),
		})
	}
	data, err := alh.FollowsBusiness.GetFollowersUsers(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get followers user",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToUserResponseList(data),
	})

}
