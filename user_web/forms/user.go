package forms

type PasswordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=15"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,len=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=15"`
	Code     string `form:"code" json:"code" binding:"required,len=6"`
}
type UpdateUserForm struct {
	Name     string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Gender   string `form:"gender" json:"gender" binding:"required,oneof=female male"`
	Birthday string `form:"birthday" json:"birthday" binding:"required,datetime=2006-01-02"`
}
