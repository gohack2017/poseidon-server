package sms

type SMSOutput struct {
	Code    int    `json:"code"`  // 返回值为 2 时，表示提交成功
	ID      string `json:"smsid"` // 当提交成功后，此字段为流水号，否则为 0
	Message string `json:"msg"`   //提交结果描述
}

func (this *SMSOutput) IsOK() bool {
	return this.Code == 2
}
