package opts

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

type InitOptions struct {
	Mux *http.ServeMux
}
type Opts struct {
	mux *http.ServeMux
}

func WakeOpts(initOpts InitOptions) *Opts {
	opt := new(Opts)
	opt.mux = initOpts.Mux
	opt.mallocRegHandler()
	return opt
}
func (se *Opts) mallocRegHandler() {
	se.mux.HandleFunc("/go-test", se.test)
	se.mux.HandleFunc("/login", se.login)
	se.mux.HandleFunc("/usrinfo", se.usrinfo)
	se.mux.HandleFunc("/logout", se.logout)
	se.mux.HandleFunc("/getTracks", se.getTracks)
}

// just a test Handler
func (se *Opts) test(w http.ResponseWriter, r *http.Request) {
	params, err := requestDeser(r)
	if err != nil {
		w.Write(respSer(1, map[string]interface{}{"error": "testerror"}))
		return
	}
	res, _ := json.Marshal(params)
	fmt.Println(string(res))
	w.Write(respSer(0, params))

}

func (se *Opts) login(w http.ResponseWriter, r *http.Request) {
	p, done := valid(w, r)
	if done {
		return
	}
	hash := sha256.New()

	password := p.Password
	hash.Write([]byte(password))

	sha256Hash := hash.Sum(nil)
	var user User
	entry, err := user.QueryByEntry(p.LoginName, hex.EncodeToString(sha256Hash))
	if err != nil {
		w.Write(respSer(1, "密码不正确"))
	} else {
		uuid := GenUUID()
		entry.Token = uuid
		entry.UpdateTokenByUser()
		w.Write(respSer(0, entry))
	}

}

func (se *Opts) usrinfo(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	token := vars["token"][0]
	var user User
	entry, err := user.QueryByToken(token)
	//token
	if err != nil {
		w.Write(respSer(1, "重新登录试试"))
	} else {
		w.Write(respSer(0, entry))
	}

}

func (se *Opts) logout(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query()["token"][0]
	var user User
	_, err := user.QueryByToken(token)
	if err == nil {
		user.ClearTokenByUser()
	}
	w.Write(respSer(0, "clear!"))

}

func (se *Opts) getTracks(w http.ResponseWriter, r *http.Request) {
	var foodTrucks FoodTrucks
	trucks, err := foodTrucks.GetFoodTrucks()
	if err != nil {
		w.Write(respSer(1, "未知错误"))
	} else {
		w.Write(respSer(0, trucks))
	}
}

func valid(w http.ResponseWriter, r *http.Request) (Params, bool) {
	params, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"error": "param valid error"}))
		return Params{}, true
	}
	return params, false
}
