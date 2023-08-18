package response

var ErrorMessage = map[int]string{
	EDefaultCode:        "服务器错误",
	EUnauthorizedCode:   "请先登录",
	ParamErr:            "参数不合法",
	InsufficientBalance: "余额不足",
	TooManyRequest:      "请求太频繁",
	Waiting:             "继续等待",
}

func message(code int) (msg string) {
	msg, ok := ErrorMessage[code]
	if ok {
		return
	}

	if code < 50000 {
		msg = ErrorMessage[40000]
	} else {
		msg = ErrorMessage[50000]
	}
	return
}
