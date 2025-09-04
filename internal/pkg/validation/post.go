package validation

import (
	"context"

	v1 "github.com/ydcloud-dy/knowledgeBase.git/api/apiserver"
)

func (v *Validator) ValidateCreatePostRequest(ctx context.Context, rq *v1.CreatePostRequest) error {
	return nil
}

func (v *Validator) ValidateUpdatePostRequest(ctx context.Context, rq *v1.UpdatePostRequest) error {
	return nil
}
