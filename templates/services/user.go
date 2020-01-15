package services

import (
	"encoding/json"
	base "github.com/grestful/core"
	"{{package}}/models"
)

type UserService struct {
	base.Controller
}

type UserIdService struct {
	base.Controller
}

type UserSaveRequest struct {
	*models.DefaultModel `json:"account"`
	Type                string `json:"type"`
}
type UserSaveService struct {
	base.Controller
	Account *UserSaveRequest
}

func (u *UserIdService) Process() base.IError {
	info := &models.DefaultModel{}
	info.GetById(u.Ctx.Params.ByName("id"))
	u.Data = info
	return nil
}

func (u *UserSaveService) Decode() base.IError {
	u.Account = &UserSaveRequest{&models.DefaultModel{}, ""}
	if bt, err := u.Ctx.GetRawData(); err == nil {
		if err := json.Unmarshal(bt, u.Account); err != nil {
			return base.NewError(err)
		}
	} else {
		return base.NewError(err)
	}

	return nil
}

func (u *UserSaveService) Process() base.IError {
	//account := u.Account

	//if account.UserName == "" || account.Password == "" || account.Type == "" {
	//	return base.NewErrorStr("注册失败，缺少账户名密码")
	//}
	//
	//if account.Type == "1" || account.Type == "2" {
	//	if account.IsRegister() {
	//		return base.NewErrorStr("手机已被注册")
	//	}
	//}
	//
	//account.UserId = 12312412
	//now := time.Now().Format("2006-01-02 15:04:05")
	//account.LastLoginTime = now
	//account.FirstLoginTime = account.LastLoginTime
	//account.RegTime = account.LastLoginTime
	//e := account.Create()

	//if len(e) > 0 {
	//	u.Code = "500"
	//	u.SetError(base.NewError(e[0]))
	//}

	return nil
}
