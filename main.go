package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	clip "github.com/atotto/clipboard"
	"github.com/skanehira/clipboard-image"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	http.HandleFunc("/getimg", GetImg)
	http.HandleFunc("/setclip", SetClip)
	http.HandleFunc("/openurl", OpenUrl)
	http.ListenAndServe(":8377", nil)
}

func GetImg(w http.ResponseWriter, req *http.Request) {
	buf, err := clipboard.ReadFromClipboard()
	if err != nil {
		w.Write([]byte("noimg"))
		return
	}
	img, err := ioutil.ReadAll(buf)
	if err != nil {
		w.Write([]byte("noimg"))
		return
	}
	md5Byte := md5.Sum(img)
	md5str := fmt.Sprintf("%x", md5Byte)
	w.Header().Add("Md5", md5str)
	w.Write([]byte(img))

}

/*
curl -H "Content-Type:text/plain" --data-binary @./main.go http://192.168.56.1:8377/setclip
*/
func SetClip(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("write begin "))
	bodyText, _ := ioutil.ReadAll(req.Body)
	err := clip.WriteAll(string(bodyText))
	if err != nil {
		w.Write([]byte("write err"))
		return
	}
	w.Write([]byte(bodyText))

}

/*
curl http://192.168.56.1:8377/openurl -d '{"url":"https://www.baidu.com"}' -X POST -H "Content-Type:application/json"
*/
func OpenUrl(w http.ResponseWriter, req *http.Request) {
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	t := struct {
		Url string `json:"url"`
	}{}
	err := d.Decode(&t)
	if err != nil {
		w.Write([]byte("write err"))
		return
	}
	url := t.Url
	w.Write([]byte(url))
	go open.Run(url)
}
