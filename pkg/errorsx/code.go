package errorsx

import "net/http"

// errorsx 预定义标准的错误.
var (
	// OK 代表请求成功.
	OK = &ErrorX{Code: http.StatusOK, Message: ""}

	// ErrInternal 表示所有未知的服务器端错误.
	ErrInternal = &ErrorX{Code: http.StatusInternalServerError, Reason: "InternalError", Message: "Internal server error."}

	// ErrNotFound 表示资源未找到.
	ErrNotFound = &ErrorX{Code: http.StatusNotFound, Reason: "NotFound", Message: "Resource not found."}

	// ErrBind 表示请求体绑定错误.
	ErrBind = &ErrorX{Code: http.StatusBadRequest, Reason: "BindError", Message: "Error occurred while binding the request body to the struct."}

	// ErrDBWrite 标识数据库写入错误
	ErrDBWrite = &ErrorX{Code: http.StatusBadRequest, Reason: "ErrDBWrite", Message: "Write DB Error."}

	// ErrDBRead 标识数据库读取错误
	ErrDBRead = &ErrorX{Code: http.StatusBadRequest, Reason: "ErrDBRead", Message: "Read DB Error."}

	// ErrUserNotFound 找不到用户
	ErrUserNotFound = &ErrorX{Code: http.StatusBadRequest, Reason: "ErrUserNotFound", Message: "User Not Found."}

	// ErrPostNotFound 找不到文章
	ErrPostNotFound = &ErrorX{Code: http.StatusBadRequest, Reason: "ErrPostNotFound", Message: "Post Not Found."}

	// ErrInvalidArgument 表示参数验证失败.
	ErrInvalidArgument = &ErrorX{Code: http.StatusBadRequest, Reason: "InvalidArgument", Message: "Argument verification failed."}

	// ErrUnauthenticated 表示认证失败.
	ErrUnauthenticated = &ErrorX{Code: http.StatusUnauthorized, Reason: "Unauthenticated", Message: "Unauthenticated."}

	// ErrPermissionDenied 表示请求没有权限.
	ErrPermissionDenied = &ErrorX{Code: http.StatusForbidden, Reason: "PermissionDenied", Message: "Permission denied. Access to the requested resource is forbidden."}

	// ErrOperationFailed 表示操作失败.
	ErrOperationFailed = &ErrorX{Code: http.StatusConflict, Reason: "OperationFailed", Message: "The requested operation has failed. Please try again later."}
)
