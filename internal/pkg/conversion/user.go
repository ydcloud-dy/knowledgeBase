package conversion

import (
	"github.com/onexstack/onexstack/pkg/core"

	apiv1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
	"github.com/ydcloud-dy/knowledgeBase.git/internal/apiserver/model"
)

// UserodelToUserV1 将模型层的 User（用户模型对象）转换为 Protobuf 层的 User（v1 用户对象）.
func UserodelToUserV1(userModel *model.User) *apiv1.User {
	var protoUser apiv1.User
	_ = core.CopyWithConverters(&protoUser, userModel)
	return &protoUser
}

// UserV1ToUserodel 将 Protobuf 层的 User（v1 用户对象）转换为模型层的 User（用户模型对象）.
func UserV1ToUserodel(protoUser *apiv1.User) *model.User {
	var userModel model.User
	_ = core.CopyWithConverters(&userModel, protoUser)
	return &userModel
}
