package rpc

import (
	. "Go-SDK/type"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

const (
	RPCCLIENT_TIMEOUT = 300
)

// 连接参数
type RpcClient struct {
	serverAddr       string
	httpClient       *http.Client
	ws               *websocket.Conn
	Handler          *sync.Map
	SubscribeHandler *sync.Map
}

var Client *RpcClient

//初始化rpc客户端
func InitClient(host string, port int, useSSL bool) (err error) {
	Client, err = newClient(host, port, useSSL)
	return
}

//连接配置
func newClient(host string, port int, useSSL bool) (c *RpcClient, err error) {
	if len(host) == 0 {
		err = errors.New("Bad call missing argument host")
		return
	}
	var serverAddr string
	var httpClient *http.Client
	var ws *websocket.Conn
	if useSSL {
		serverAddr = "https://"
		t := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: t}
		ws, err = websocket.Dial("wss://"+host+":"+strconv.Itoa(port), "", "wss://"+host+":"+strconv.Itoa(port))
	} else {
		serverAddr = "http://"
		httpClient = &http.Client{}
		ws, err = websocket.Dial("ws://"+host+":"+strconv.Itoa(port), "", "ws://"+host+":"+strconv.Itoa(port))
	}
	if err != nil {
		log.Fatal("init sdk error:::", err)
	}
	c = &RpcClient{serverAddr: fmt.Sprintf("%s%s:%d", serverAddr, host, port), httpClient: httpClient, ws: ws, Handler: &sync.Map{}, SubscribeHandler: &sync.Map{}}
	go c.handler()
	return
}

// 超时处理
func (c *RpcClient) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.httpClient.Do(req)
		done <- result{resp, err}
	}()
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("Timeout reading data from server")
	}
}
func (c *RpcClient) handler() {
	for {
		var reply string
		ret := &RpcResp{}
		notice := &Notice{}
		if err := websocket.Message.Receive(c.ws, &reply); err == nil {
			//fmt.Println("-------------------")
			//log.Println(reply)
			if err = json.Unmarshal([]byte(reply), ret); err == nil && ret.Id != `` {
				if f, ok := c.Handler.Load(ret.Id); ok {
					//fmt.Println("-------------------")
					go f.(func(r *RpcResp) error)(ret)
					c.Handler.Delete(ret.Id)
				}
			} else if err = json.Unmarshal([]byte(reply), notice); err == nil {
				//log.Println(notice.Params[0].(string))
				if f, ok := c.SubscribeHandler.Load(notice.Params[0].(string)); ok {
					go f.(func(r *Notice) error)(notice)
				}
			} else {
				log.Println("xxxxxxxxxxx")
			}

		}
	}
}

//websocket
func (c *RpcClient) SendWithHandler(reqData *RpcRequest, f func(r *RpcResp) error) (err error) {
	reqJson := reqData.ToString()
	if err = websocket.Message.Send(c.ws, reqJson); err == nil {
		c.Handler.Store(strconv.Itoa(int(reqData.Id)), f)
	}
	return
}

func (c *RpcClient) Subscribe(subscribe string, f func(r *Notice) error) (ret *RpcResp, err error) {
	id := time.Now().UnixNano()
	reqData := CreateRpcRequest(CALL,
		[]interface{}{DATABASE_API_ID, subscribe,
			[]interface{}{id, true}})
	reqData.Id = id
	reqJson := reqData.ToString()
	if err = websocket.Message.Send(c.ws, reqJson); err == nil {
		c.SubscribeHandler.Store(strconv.Itoa(int(reqData.Id)), f)
	}
	return
}

//websocket
func (c *RpcClient) Send(reqData *RpcRequest) (ret *RpcResp, err error) {
	ret = &RpcResp{}
	reqJson := reqData.ToString()
	if err = websocket.Message.Send(c.ws, reqJson); err == nil {
		ch := make(chan *RpcResp)
		c.Handler.Store(strconv.Itoa(int(reqData.Id)), func(r *RpcResp) error {
			ch <- r
			return nil
		})
		select {
		case ret = <-ch:
			return
		case <-time.After(time.Second * 5):
			err = errors.New("rpc time out!")
			ret = nil
			return
		}
	}
	return
	//废弃的 HTTP rpc
	/*
		log.Println("rpc Send start:::", reqJson)
		connectTimer := time.NewTimer(RPCCLIENT_TIMEOUT * time.Second)
		payloadBuffer := bytes.NewReader(reqJsonByte)
		req, err := http.NewRequest("POST", c.serverAddr, payloadBuffer)
		if err != nil {
			log.Println("rpc Send error:::", err)
			return
		}
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
		req.Header.Add("Accept", "application/json")
		resp, err := c.doTimeoutRequest(connectTimer, req)
		if err != nil {
			log.Println("rpc Send error:::", err)
			return
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		log.Println("rpc Send:::", string(data))
		if err != nil {
			return
		}
		if resp.StatusCode != 200 {
			err = errors.New("HTTP error: " + resp.Status)
			return
		}
		//fmt.Println(string(data))
		json.Unmarshal(data, ret)
		return*/
}
