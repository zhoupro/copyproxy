package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	clip "github.com/atotto/clipboard"
	"github.com/skanehira/clipboard-image"
	"github.com/skratchdot/open-golang/open"
)

var syncRemoteHost string

func main() {

	InputParams()
	if syncRemoteHost != "" {

		go syncCopyToRemote()
	}
	http.HandleFunc("/getimg", GetImg)
	http.HandleFunc("/setclip", SetClip)
	http.HandleFunc("/openurl", OpenUrl)
	http.ListenAndServe(":8377", nil)
}

func InputParams() {
	flag.StringVar(&syncRemoteHost, "h", "", "host")
	//解析命令行参数
	flag.Parse()
}

func syncCopyToRemote() {
	lastSendedMd5Str := ""
	for {
		time.Sleep(2000 * time.Millisecond)
		curCopyStr, _ := clip.ReadAll()
		if curCopyStr == "" {
			continue
		}
		fmt.Println(curCopyStr)
		curCopyStrMd5 := Md5Str([]byte(curCopyStr))
		if lastSendedMd5Str != "" && lastSendedMd5Str == curCopyStrMd5 {
			fmt.Println(fmt.Sprintf("equal %v", time.Now().Second()))
			continue
		}
		go sendRealCopy(curCopyStr)
		lastSendedMd5Str = curCopyStrMd5
	}
}

func sendRealCopy(copyStr string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	fmt.Printf("%s\n", copyStr)

	f := strings.NewReader(copyStr)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:8377/setclip", syncRemoteHost), f)
	if err != nil {
		//todo
	}
	req.Header.Set("Content-Type", "text/plain")
	_, err = http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		//todo
	}
	return Md5Str([]byte(copyStr))
}

func Md5Str(byteArr []byte) string {

	md5Byte := md5.Sum(byteArr)
	md5str := fmt.Sprintf("%x", md5Byte)
	return md5str
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
	fmt.Println(string(bodyText))
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
