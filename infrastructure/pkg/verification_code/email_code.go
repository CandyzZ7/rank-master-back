package verification_code

import (
	"crypto/tls"
	"net/smtp"
	"rank-master-back/internal/config"
	"strconv"

	"github.com/jordan-wright/email"
)

func SendEmailCode(config config.Config, toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "RankMaster <rank_master@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte(`<p>您正在注册RankMaster，以下是您的验证码，验证码将在` + strconv.Itoa(CodeValidityTime) + `分钟后过期<br><br>
	<span style="
	font-size: 24px;
	font-weight: 600;
	">` + code + `</span>
	<br>如果这不是您的邮件请忽略，很抱歉打扰您，请原谅。
	</p>`)
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "rank_master@163.com", config.Email.AuthorizationPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}
