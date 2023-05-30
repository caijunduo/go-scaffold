package v1

import (
	"github.com/caijunduo/go-scaffold/internal/pkg/api"
	"github.com/caijunduo/go-scaffold/pkg/codec"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"
)

// CreateUserRequest 指定了 `POST /v1/users` 接口的请求参数.
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (req *CreateUserRequest) Validate(c *gin.Context) (err error) {
	_ = c.ShouldBindJSON(req)
	if err = validation.Validate(&req.Username,
		validation.Required.Error("请输入用户名"),
		is.Alphanumeric.Error("用户名格式错误"),
		validation.Length(1, 255).Error("用户名格式错误")); err != nil {
		return
	}
	if err = validation.Validate(&req.Password,
		validation.Required.Error("请输入密码"),
		is.Base64.Error("请输入密码")); err != nil {
		return
	}
	if req.Password, err = codec.Base64StdDecodeStd(req.Password); err != nil {
		return
	}
	if err = validation.Validate(&req.Password,
		validation.Length(6, 10).Error("密码格式错误")); err != nil {
		return
	}
	if err = validation.Validate(&req.Email,
		validation.Required.Error("请输入邮箱"),
		is.Email.Error("请输入符合规范的邮箱")); err != nil {
		return err
	}
	if err = validation.Validate(&req.Phone,
		validation.Required.Error("请输入手机号"),
		validation.Length(11, 11).Error("手机号格式错误")); err != nil {
		return err
	}
	return
}

type ShowUserResponse Userinfo

type Userinfo struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUserRequest struct {
	api.Paginate
}

func (req *GetUserRequest) Validate(c *gin.Context) error {
	if err := req.Paginate.Validate(c); err != nil {
		return err
	}
	return nil
}
