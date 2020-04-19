package handler

import (
	"cloud_storage/db"
	"cloud_storage/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var pwd_salt = "#089"

// SingupHandler：处理用户注册请求
func SingupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	r.ParseForm()

	username := r.Form.Get("username")
	pwd := r.Form.Get("password")
	if len(username) < 3 || len(pwd) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}

	enc_pwd := util.Sha1([]byte(pwd + pwd_salt))
	suc := db.UserSingnup(username, enc_pwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}

}

// SiginInHandler:登录接口
func SiginInHandler(w http.ResponseWriter, r *http.Request) {
	var (
		username, pwd string
		encPwd        string
	)

	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/user/toSignin", http.StatusFound)
		return
	}

	r.ParseForm()
	//1.校验用户名及密码
	username = r.Form.Get("username")
	pwd = r.Form.Get("password")
	fmt.Println("username:", username, "password:", pwd)
	encPwd = util.Sha1([]byte(pwd + pwd_salt))

	checkPwd := db.UserSignin(username, encPwd)
	if !checkPwd {
		w.Write([]byte("FAILED"))
		return
	}

	//2.生成访问凭证(token)

	//3.登录成功后重定向到首页
}

// RedirectSignin：跳转到登录页面
func RedirectSignin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		datas, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(datas))
	}
}
