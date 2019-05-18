package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/TruthHun/BookStack/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Token string
}

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type APIUser struct {
	Uid         int    `json:"uid"`
	Token       string `json:"token"`
	Account     string `json:"username"` //对应 member.account
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
	Description string `json:"intro"`
}

type APIBookList struct {
	BookId      int       `json:"book_id"`
	BookName    string    `json:"book_name"`
	Identify    string    `json:"identify"`
	OrderIndex  string    `json:"sort"`
	Description string    `json:"description"`
	Label       string    `json:"tags"`
	DocCount    int       `json:"doc_count"` // 章节数量
	Vcnt        int       `json:"view"`      // 阅读
	Star        int       `json:"star"`      // 收藏
	Lang        string    `json:"lang"`
	Cover       string    `json:"cover"`
	Score       int       `json:"score"`                // 文档项目评分，默认40，即4.0星
	CntScore    int       `json:"cnt_score"`            // 评分个数
	CntComment  int       `json:"cnt_comment"`          // 评论人数
	ModifyTime  time.Time `json:"updated_at"`           // 更新时间
	CreateTime  time.Time `json:"created_at"`           // 新建时间
	Sharer      string    `json:"sharer,omitempty"`     // 分享人
	Author      string    `json:"author,omitempty"`     // 原作者
	AuthorURL   string    `json:"author_url,omitempty"` // 原作者连接地址
}

//###################################//

const (
	messageInternalServerError     = "服务内部错误，请联系管理员"
	messageUsernameOrPasswordError = "用户名或密码不正确"
	messageLoginSuccess            = "登录成功"
	messageRequiredLogin           = "您未登录或者您的登录已过期，请重新登录"
	messageLogoutSuccess           = "退出登录成功"
	messageSuccess                 = "操作成功"
	messageNotFound                = "资源不存在"
)

//###################################//

func (this *BaseController) Response(httpStatus int, message string, data ...interface{}) {
	this.Ctx.ResponseWriter.Header().Set("Powered By", "BookChat")
	this.Ctx.Output.SetStatus(httpStatus)
	resp := APIResponse{Message: message}
	if len(data) > 0 {
		resp.Data = data[0]
	}

	// support gzip
	if strings.ToLower(this.Ctx.Request.Header.Get("content-encoding")) == "gzip" {
		// TODO
	}
	this.Data["json"] = resp
	this.ServeJSON()
	this.StopRun()
}

// 验证access token
func (this *BaseController) Prepare() {
	this.Token = this.Ctx.Request.Header.Get("Authorization")
	if beego.AppConfig.String("runmode") == "dev" {
		beego.Debug("auth data: ", fmt.Sprintf("%+v", models.NewAuth().AllFromCache()))
		time.Sleep(1 * time.Second)
	}
}

func (this *BaseController) isLogin() {

}

func (this *BaseController) completeImage() {

}