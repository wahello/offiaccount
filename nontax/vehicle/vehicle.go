package vehicle

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiQuerystate     = "/nontax/vehicle/querystate"
	apiEntranceNotify = "/nontax/vehicle/entrancenotify"
	apiPayApply       = "/nontax/vehicle/payapply"
)

/*
用户状态查询

在停车场、高速、加油等场景下，商户需获取用户车主服务状态/需要关联车主服务。本接口，会查询用户是否开通、授权、有欠费或黑名单用户情况，并将对应的用户状态进行返回

See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html

POST https://api.weixin.qq.com/nontax/vehicle/querystate?access_token=$AccessToken
*/
func Querystate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQuerystate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
用户入场通知

在停车场场景下，如用户已加入车主平台，则进行入场通知；如用户已经欠费，会发送用户欠费入场通知。本接口，会查询用户是否有欠费或黑名单用户情况，并将对应的用户状态进行返回

See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html

POST https://api.weixin.qq.com/nontax/vehicle/entrancenotify?access_token=$AccessToken
*/
func EntranceNotify(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiEntranceNotify, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
申请扣款

委托代扣可应用于定期扣款或需事后扣款以期提高效率的场景。例如高速，停车场等通过用户授权给商户，进行委托扣款的场景

See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html

POST https://api.weixin.qq.com/nontax/vehicle/payapply?access_token=$AccessToken
*/
func PayApply(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPayApply, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}