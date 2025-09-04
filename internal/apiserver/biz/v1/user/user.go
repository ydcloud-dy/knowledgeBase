package user

import (
	"context"
	"log/slog"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/store/where"
	"golang.org/x/sync/errgroup"

	apiv1 "github.com/ydcloud-dy/knowledgeBase.git/api/apiserver"
	"github.com/ydcloud-dy/knowledgeBase.git/internal/apiserver/model"
	"github.com/ydcloud-dy/knowledgeBase.git/internal/apiserver/store"
	"github.com/ydcloud-dy/knowledgeBase.git/internal/pkg/conversion"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/contextx"
	"github.com/ydcloud-dy/knowledgeBase.git/pkg/known"
)

// UserBiz 定义处理用户请求所需的方法.
type UserBiz interface {
	Create(ctx context.Context, rq *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error)
	Update(ctx context.Context, rq *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error)
	Delete(ctx context.Context, rq *apiv1.DeleteUserRequest) (*apiv1.DeleteUserResponse, error)
	Get(ctx context.Context, rq *apiv1.GetUserRequest) (*apiv1.GetUserResponse, error)
	List(ctx context.Context, rq *apiv1.ListUserRequest) (*apiv1.ListUserResponse, error)

	UserExpansion
}

// UserExpansion 定义用户操作的扩展方法.
type UserExpansion interface {
}

// userBiz 是 UserBiz 接口的实现.
type userBiz struct {
	store store.IStore
}

// 确保 userBiz 实现了 UserBiz 接口.
var _ UserBiz = (*userBiz)(nil)

func New(store store.IStore) *userBiz {
	return &userBiz{store: store}
}

// Create 实现 UserBiz 接口中的 Create 方法.
func (b *userBiz) Create(ctx context.Context, rq *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error) {
	var userM model.User
	_ = copier.Copy(&userM, rq)

	if err := b.store.User().Create(ctx, &userM); err != nil {
		return nil, err
	}

	return &apiv1.CreateUserResponse{UserID: userM.UserID}, nil
}

// Update 实现 UserBiz 接口中的 Update 方法.
func (b *userBiz) Update(ctx context.Context, rq *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error) {
	userM, err := b.store.User().Get(ctx, where.F("userID", contextx.UserID(ctx)))
	if err != nil {
		return nil, err
	}

	if rq.Username != nil {
		userM.Username = *rq.Username
	}
	if rq.Email != nil {
		userM.Email = *rq.Email
	}
	if rq.Nickname != nil {
		userM.Nickname = *rq.Nickname
	}
	if rq.Phone != nil {
		userM.Phone = *rq.Phone
	}

	if err := b.store.User().Update(ctx, userM); err != nil {
		return nil, err
	}

	return &apiv1.UpdateUserResponse{}, nil
}

// Delete 实现 UserBiz 接口中的 Delete 方法.
func (b *userBiz) Delete(ctx context.Context, rq *apiv1.DeleteUserRequest) (*apiv1.DeleteUserResponse, error) {
	if err := b.store.User().Delete(ctx, where.F("userID", contextx.UserID(ctx))); err != nil {
		return nil, err
	}

	return &apiv1.DeleteUserResponse{}, nil
}

// Get 实现 UserBiz 接口中的 Get 方法.
func (b *userBiz) Get(ctx context.Context, rq *apiv1.GetUserRequest) (*apiv1.GetUserResponse, error) {
	userM, err := b.store.User().Get(ctx, where.F("userID", contextx.UserID(ctx)))
	if err != nil {
		return nil, err
	}

	return &apiv1.GetUserResponse{User: conversion.UserodelToUserV1(userM)}, nil
}

// List 实现 UserBiz 接口中的 List 方法.
func (b *userBiz) List(ctx context.Context, rq *apiv1.ListUserRequest) (*apiv1.ListUserResponse, error) {
	whr := where.P(int(rq.Offset), int(rq.Limit))
	count, userList, err := b.store.User().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)

	// 设置最大并发数量为常量 MaxConcurrency
	eg.SetLimit(known.MaxErrGroupConcurrency)

	// 使用 goroutine 提高接口性能
	for _, user := range userList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				count, _, err := b.store.Post().List(ctx, where.F("userID", contextx.UserID(ctx)))
				if err != nil {
					return err
				}

				converted := conversion.UserodelToUserV1(user)
				converted.PostCount = count
				m.Store(user.ID, converted)

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		slog.ErrorContext(ctx, "Failed to wait all function calls returned", "err", err)
		return nil, err
	}

	users := make([]*apiv1.User, 0, len(userList))
	for _, item := range userList {
		user, _ := m.Load(item.ID)
		users = append(users, user.(*apiv1.User))
	}

	slog.DebugContext(ctx, "Get users from backend storage", "count", len(users))

	return &apiv1.ListUserResponse{TotalCount: count, Users: users}, nil
}
