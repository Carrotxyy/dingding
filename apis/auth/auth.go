// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package auth 身份验证
package auth

import (
	"net/url"

	"github.com/fastwego/dingding"
)

const (
	apiGetUserInfo    = "/user/getuserinfo"
	apiGetJSApiTicket = "/get_jsapi_ticket"
)

/*
获取用户userid



See: https://ding-doc.dingtalk.com/doc#/serverapi2/clotub

GET https://oapi.dingtalk.com/user/getuserinfo?access_token=access_token&code=code
*/
func GetUserInfo(ctx *dingding.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetUserInfo + "?" + params.Encode())
}

/*
获取jsapi_ticket



See: https://ding-doc.dingtalk.com/doc#/dev/uwa7vs

GET https://oapi.dingtalk.com/get_jsapi_ticket?access_token=ACCESS_TOKEN
*/
func GetJSApiTicket(ctx *dingding.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetJSApiTicket)
}
