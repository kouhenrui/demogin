package global

import (
	"text/template"
)

/**
 * @ClassName captcha
 * @Description 图片验证器
 * @Author khr
 * @Date 2023/5/6 16:01
 * @Version 1.0
 */
const formTemplateSrc = `
<!doctype html>
<head><title>验证码</title></head>
<body>
<script>
function reload() {
    setSrcQuery(document.getElementById('image'), "reload=" + (new Date()).getTime());
    return false;
}
</script>
<form action="/process" method=post>
<p>输入你在下面的图片中看到的数字:</p>
<p><img id=image src="/captcha/{{.CaptchaId}}.png" alt="Captcha image"></p>
<input type=hidden name=captchaId value="{{.CaptchaId}}"><br>
<input name=captchaSolution>
<input type=submit value=提交>
</form>
`

var FormTemplate = template.Must(template.New("example").Parse(formTemplateSrc))

//func showFormHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		http.NotFound(w, r)
//		return
//	}
//	d := struct {
//		CaptchaId string
//	}{
//		captcha.New(),
//	}
//	if err := formTemplate.Execute(w, &d); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func processFormHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	if !captcha.VerifyString(r.FormValue("captchaId"), r.FormValue("captchaSolution")) {
//		io.WriteString(w, "Wrong captcha solution! No robots allowed!\n")
//	} else {
//		io.WriteString(w, "Great job, human! You solved the captcha.\n")
//	}
//	io.WriteString(w, "<br><a href='/'>Try another one</a>")
//}
