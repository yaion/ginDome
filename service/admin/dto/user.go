package dto

type LoginParams struct {
	Name      string `json:"name" validate:"required,min=3,max=32"`
	Password  string `json:"password" validate:"required,min=6,max=32"`
	CaptchaId string `json:"captcha_id" validate:"required"`
	Verify    string `json:"verify" validate:"required"`
}

type LoginBack struct {
	Token string `json:"token"`
	Msg   string `json:"msg"`
}

type RegisterParams struct {
	UserName string `json:"user_name" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	NickName string `json:"nick_name" validate:"required,min=3,max=128"`
	Phone    string ` json:"phone" validate:"required,len=11"`
	//Avatar string `json:"avatar" `
	Sex    uint8  `json:"sex" `
	Email  string `json:"email" validate:"required,email"`
	Remark string `json:"remark"`
}

type UserListParams struct {
	UserName string   `json:"user_name" validate:"omitempty,min=3,max=32"`
	NickName string   `json:"nick_name" validate:"omitempty,min=3,max=128"`
	Phone    string   ` json:"phone" validate:"omitempty,len=11"`
	PageSize uint64   `json:"page_size" validate:"required,numeric"`
	Sort     UserSort `json:"sort" validate:"required"`
	Page     uint64   `json:"page" validate:"required,numeric"`
}

type UpdateUserParams struct {
	UserName string `json:"user_name" validate:"min=3,max=32"`
	Password string `json:"password" validate:"min=6,max=32"`
	NickName string `json:"nick_name" validate:"min=3,max=128"`
	Phone    string ` json:"phone" validate:"len=11"`
	Avatar   string `json:"avatar" validate:""`
	Sex      uint8  `json:"sex" validate:""`
	Status   uint8  `json:"status" validate:""`
	Email    string `json:"email" validate:"email"`
	Remark   string `json:"remark" validate:""`
}

type UserSort struct {
	UserName string `json:"user_name" validate:"omitempty,oneof=asc desc"`
	NickName string `json:"nick_name" validate:"omitempty,oneof=asc desc"`
	Phone    string ` json:"phone" validate:"omitempty,oneof=asc desc"`
	Id       string `json:"id" validate:"omitempty,oneof=asc desc"`
}

// JsonResult json回参
type JsonResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg" example:"success"`
	Data interface{} `json:"data,omitempty"`
}
