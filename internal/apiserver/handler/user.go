package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	v1 "github.com/ydcloud-dy/knowledgeBase.git/api/apiserver"
	"github.com/ydcloud-dy/knowledgeBase.git/internal/pkg/core"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/errorsx"
)

// CreateUser 创建新用户.
func (h *Handler) CreateUser(c *gin.Context) {
	slog.Info("Create user function called")

	var rq v1.CreateUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateCreateUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Create(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// UpdateUser 更新用户信息.
func (h *Handler) UpdateUser(c *gin.Context) {
	slog.Info("Update user function called")

	var rq v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateUpdateUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage(err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Update(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// DeleteUser 删除用户.
func (h *Handler) DeleteUser(c *gin.Context) {
	slog.Info("Delete user function called")

	var rq v1.DeleteUserRequest
	if err := c.ShouldBindUri(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().Delete(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// GetUser 获取用户信息.
func (h *Handler) GetUser(c *gin.Context) {
	slog.Info("Get user function called")

	var rq v1.GetUserRequest
	if err := c.ShouldBindUri(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().Get(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// ListUser 列出用户信息.
func (h *Handler) ListUser(c *gin.Context) {
	slog.Info("List user function called")

	var rq v1.ListUserRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().List(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}
