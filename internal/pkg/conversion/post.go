package conversion

import (
	"github.com/jinzhu/copier"

	apiv1 "github.com/ydcloud-dy/knowledgeBase.git/api/apiserver"
	"github.com/ydcloud-dy/knowledgeBase.git/internal/apiserver/model"
)

// PostodelToPostV1 将模型层的 Post（博客模型对象）转换为 Protobuf 层的 Post（v1 博客对象）.
func PostodelToPostV1(postModel *model.Post) *apiv1.Post {
	var protoPost apiv1.Post
	_ = copier.Copy(&protoPost, postModel)
	return &protoPost
}

// PostV1ToPostodel 将 Protobuf 层的 Post（v1 博客对象）转换为模型层的 Post（博客模型对象）.
func PostV1ToPostodel(protoPost *apiv1.Post) *model.Post {
	var postModel model.Post
	_ = copier.Copy(&postModel, protoPost)
	return &postModel
}
