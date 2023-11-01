package opts

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

type TaskStat int

const (
	Ok   TaskStat = 4
	Fail TaskStat = 5
)

func getParams(r *http.Request) (param Params, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return param, err
	}
	if err = json.Unmarshal(body, &param); err != nil {
		return param, err
	}
	//todo basic validator
	return param, nil
}
func jsonEncode(code int, data interface{}) []byte {
	tr := new(JsonResult)
	tr.Code, tr.Data = code, data
	result, _ := json.Marshal(tr)
	return result
}

func requestDeser(r *http.Request) (param interface{}, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return param, err
	}
	if err = json.Unmarshal(body, &param); err != nil {
		return param, err
	}
	//todo basic validator
	return param, nil

}

func respSer(code int, data interface{}) []byte {
	resp, _ := json.Marshal(JsonResult{
		Code: code,
		Data: data,
	})

	return resp

}

func JsonPost(url string, reqData string) (bool, string) {
	contentType := "application/json"
	fmt.Printf("request:\n")
	fmt.Println(reqData)
	resp, err := http.Post(url, contentType, strings.NewReader(reqData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return false, "req fail"
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return false, "resp fail"
	}
	fmt.Printf("resposne:\n")
	fmt.Println(string(b))
	return true, string(b)
}

func GenUUID(seed ...uint8) string {
	v4 := uuid.NewV4()
	return v4.String()
}
