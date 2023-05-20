// Code generated by goctl. DO NOT EDIT.
package types

type ApplyNewVCodeRequest struct {
	Buid string `path:"buid"`
	Key  string `form:"key"` // 0 为 devloper 登录, 不可以使用
}

type ApplyNewVCodeResponse struct {
	Vcode string `json:"vcode"`
}

type AddOneRequest struct {
	Vcode string `path:"vcode"`
}

type AddOneResponse struct {
	Status int32 `json:"status"`
}

type VerifyRequest struct {
	Vcode string `path:"vcode"`
}

type VerifyResponse struct {
	Jwt string `json:"jwt"`
}

type CheckRequest struct {
}

type CheckResponse struct {
	Buid string `json:"buid"`
}

type AddKeyRequest struct {
	Key string `form:"key"`
}

type AddKeyResponse struct {
	Ok bool `json:"ok"`
}

type DeleteKeyRequest struct {
	Key string `form:"key"`
}

type DeleteKeyResponse struct {
	Ok bool `json:"ok"`
}

type GetKeyListRequest struct {
}

type GetKeyListResponse struct {
	Balance int64    `json:"balance"`
	Keys    []string `json:"keys"`
}

type RechargeRequest struct {
	Buid   string `json:"buid"`
	Amount int64  `json:"amount"`
}

type RechargeResponse struct {
	Ok bool `json:"ok"`
}
