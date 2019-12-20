package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
)

func main() {
	bytes, err := ioutil.ReadFile("./1.json")
	if err != nil {
		fmt.Printf("File read Failed:%v\n", err)
		return
	}

	v := &[]string{}
	err2 := json.Unmarshal(bytes, v)
	if err2 != nil {
		fmt.Printf("JSON read Failed:%v\n", err2)
		return
	}
	sum := 0
	for key, val := range *v {
		if sum >= 200 {
			break
		}
		res, err := regexp.MatchString(`gif`, val)
		if err != nil || !res {
			continue
		}
		sum++
		err2 := download(val)
		if err2 != nil {
			fmt.Printf("第%d张下载失败:%v\n", key, err2)
			continue
		}
		fmt.Printf("第%d张下载成功\n", key+1)
	}
}

func download(url string) error {
	fileName := path.Base(url)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	defer res.Body.Close()

	file, err2 := os.Create("./file/" + fileName)
	if err2 != nil {
		return err2
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	_, err3 := io.Copy(writer, reader)
	if err3 != nil {
		return err3
	}
	return nil
}
