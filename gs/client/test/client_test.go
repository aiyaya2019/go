package test

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

var path = "http://127.0.0.1/api"

//get请求
func TestGet(t *testing.T) {
	if response, err := ghttp.Get(path); err != nil {
		panic(err)
		fmt.Println("1")
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
		fmt.Println("2")
	}

	if response, err := ghttp.Post(path); err != nil {
		panic(err)
		fmt.Println("3")
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
		fmt.Println("4")
	}
}

//GET请求带参数
func TestGetParam(t *testing.T) {
	if response, err := ghttp.Get(path + "/hello?city=广州"); err != nil {
		panic(err)
		fmt.Println("1")
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
		fmt.Println("2")
	}
}

//post请求
func testPost(t *testing.T) {
	if response, err := ghttp.Post(path+"/test", "name=名称&age=18"); err != nil {
		panic(err)
		fmt.Println("1")
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
		fmt.Println("2")
	}
}

//post json
func TestPostJson(t *testing.T) {
	if response, err := ghttp.Post(path+"/test2", `{"passport":"john","password":"123456"}`); err != nil {
		panic(err)
		fmt.Println("1")
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
		fmt.Println("2")
	}
}

//post header头
func TestPostHeader(t *testing.T) {
	c := ghttp.NewClient()
	c.SetHeader("Cookie", "name=john; score=100")

	if r, e := c.Post(path + "/test3"); e != nil {
		panic(e)
		fmt.Println("1")
	} else {
		fmt.Println(r.ReadAllString())
		fmt.Println("2")
	}
}

//post header头
func TestPostHeader2(t *testing.T) {
	c := ghttp.NewClient()
	c.SetHeaderRaw(`
		accept-encoding: gzip, deflate, br
		accept-language: zh-CN, zh; q = 0.9, en; q = 0.8
		referer: https://idonottell.you
		cookie: name=john; score=100
		user-agent: my test http client
	`)

	if r, e := c.Post(path + "/test4"); e != nil {
		panic(e)
		fmt.Println("1")
	} else {
		fmt.Println(r.ReadAllString())
		fmt.Println("2")
	}

}
