package officialaccount

import (
	"net/http"
	"sync"

	"github.com/silenceper/wechat/officialaccount/basic"
	"github.com/silenceper/wechat/officialaccount/config"
	"github.com/silenceper/wechat/officialaccount/context"
	"github.com/silenceper/wechat/officialaccount/device"
	"github.com/silenceper/wechat/officialaccount/js"
	"github.com/silenceper/wechat/officialaccount/material"
	"github.com/silenceper/wechat/officialaccount/menu"
	"github.com/silenceper/wechat/officialaccount/message"
	"github.com/silenceper/wechat/officialaccount/oauth"
	"github.com/silenceper/wechat/officialaccount/server"
	"github.com/silenceper/wechat/officialaccount/user"
)

//OfficialAccount 微信公众号相关API
type OfficialAccount struct {
	ctx *context.Context
}

//NewOfficialAccount 实例化公众号API
func NewOfficialAccount(cfg *config.Config) *OfficialAccount {
	if cfg.Cache == nil {
		panic("cache未设置")
	}
	ctx := &context.Context{
		Config: cfg,
	}
	ctx.SetAccessTokenLock(new(sync.RWMutex))
	ctx.SetJsAPITicketLock(new(sync.RWMutex))
	return &OfficialAccount{ctx}
}

// GetContext get Context
func (officialAccount *OfficialAccount) GetContext() *context.Context {
	return officialAccount.ctx
}

// GetBasic qr/url 相关配置
func (officialAccount *OfficialAccount) GetBasic() *basic.Basic {
	return basic.NewBasic(officialAccount.ctx)
}

// GetMenu 菜单管理接口
func (officialAccount *OfficialAccount) GetMenu() *menu.Menu {
	return menu.NewMenu(officialAccount.ctx)
}

// GetServer 消息管理
func (officialAccount *OfficialAccount) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	srv := server.NewServer(officialAccount.ctx)
	srv.Request = req
	srv.Writer = writer
	return srv
}

//GetAccessToken 获取access_token
func (officialAccount *OfficialAccount) GetAccessToken() (string, error) {
	return officialAccount.ctx.GetAccessToken()
}

// GetOauth oauth2网页授权
func (officialAccount *OfficialAccount) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(officialAccount.ctx)
}

// GetMaterial 素材管理
func (officialAccount *OfficialAccount) GetMaterial() *material.Material {
	return material.NewMaterial(officialAccount.ctx)
}

// GetJs js-sdk配置
func (officialAccount *OfficialAccount) GetJs() *js.Js {
	return js.NewJs(officialAccount.ctx)
}

// GetUser 用户管理接口
func (officialAccount *OfficialAccount) GetUser() *user.User {
	return user.NewUser(officialAccount.ctx)
}

// GetTemplate 模板消息接口
func (officialAccount *OfficialAccount) GetTemplate() *message.Template {
	return message.NewTemplate(officialAccount.ctx)
}

// GetDevice 获取智能设备的实例
func (officialAccount *OfficialAccount) GetDevice() *device.Device {
	return device.NewDevice(officialAccount.ctx)
}