package sanhe

import (
	"fmt"
	"time"

	"codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/alipay"
	"codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RequestParam struct {
	ToUserId   string
	Mobile     string
	PageUrl    string
	TemplateId string
	Year       int
	Month      int
	SendTime   time.Time
}

// 发送小程序模板账单
func (c *client) SendMiniTemplateMessageBill(param RequestParam) bool {
	// param.TemplateId = ``
	// return false
	data := map[string]interface{}{
		"keyword1": map[string]interface{}{"value": param.Mobile},
		"keyword2": map[string]interface{}{"value": fmt.Sprintf(`%d年%d月`, param.Year, param.Month)},
	}
	ReParam := alipay.MiniTemplateMessageReq{
		ToUserId:       param.ToUserId,   //接收模板消息的用户 user_id，一般为2088开头的16为数字。
		FormId:         ``,               //支付消息模板.订阅消息模板无需传入本参数
		UserTemplateId: param.TemplateId, //商家在商家自运营中心选用的消息模板ID，详情参见 选用消息模板 。
		Page:           param.PageUrl,    //小程序的跳转页面。用于用户点击模板消息 进入小程序查看 按钮后，跳转至商家小程序对应页面
		Data:           data,
	}
	resp, err := c.AliClient.MiniTemplateMessage(ReParam)
	if err != `` {
		logger.Use(`bill`).Info(`SendMiniTemplateMessageBill`, zap.Any("params", gin.H{"param": param}), zap.Any("data", gin.H{"resp": resp, "err": err}))
		return false
	}
	return resp
}

// 生活号发送单一模板消息
func (c *client) SendMessageSingleBill(param RequestParam) bool {
	// return true
	context := make(map[string]interface{}, 7)
	context[`url`] = param.PageUrl
	context[`action_name`] = `点击查看详情`
	context[`keyword1`] = map[string]string{"value": param.Mobile, "color": "#000000"}
	context[`keyword2`] = map[string]string{"value": fmt.Sprintf(`%d年%d月`, param.Year, param.Month), "color": "#000000"}
	context[`first`] = map[string]string{"value": `您好,` + fmt.Sprintf(`%d`, param.Month) + `月话费账单已出`, "color": "#000000"}
	context[`remark`] = map[string]string{"value": `点击查看详情看账单及了解首月0元享90G等更多优惠`, "color": "#000000"}
	context[`head_color`] = `#000000`

	reParam := alipay.MessageSingleSendReq{
		ToUserId:   param.ToUserId, //消息接收用户的支付宝用户id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成
		TemplateId: param.TemplateId,
		Context:    context, //消息模板ID
	}
	resp, err := c.AliClient.MessageSingleSend(reParam)
	if err != `` {
		logger.Use(`bill`).Info(`SendMessageSingleBill`, zap.Any("params", gin.H{"param": param}), zap.Any("data", gin.H{"resp": resp, "err": err}))
		return false
	}
	return resp
}
