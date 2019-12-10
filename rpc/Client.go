package rpc

import (
	. "CocosSDK/type"
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
	TIMEOUT     = 10 //超时时间
	HEARTBEAT   = 30 //心跳间隔
	Reconnect   = 10 //重连间隔
	RESEND_TIME = 3  //重发次数
)

// 连接参数
type RpcClient struct {
	httpClient       *http.Client
	ws               *websocket.Conn
	Handler          *sync.Map
	SubscribeHandler *sync.Map
}

var Client *RpcClient

//初始化rpc客户端
func InitClient(host string, useSSL bool, port ...int) (err error) {
	Client, err = newClient(host, useSSL, port...)
	return
}

//连接配置
func newClient(host string, useSSL bool, port ...int) (c *RpcClient, err error) {
	if len(host) == 0 {
		err = errors.New("Bad call missing argument host")
		return
	}
	var httpClient *http.Client
	var ws *websocket.Conn
	var port_str string

	if len(port) > 0 {
		port_str = fmt.Sprintf(":%d", port[0])
	}
	if useSSL {
		t := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: t}
		ws, err = websocket.Dial("wss://"+host+port_str, "", "wss://"+host+port_str)
	} else {
		httpClient = &http.Client{}
		ws, err = websocket.Dial("ws://"+host+port_str, "", "ws://"+host+port_str)
	}
	if err != nil {
		log.Fatal("init sdk error:::", err)
	}
	c = &RpcClient{httpClient: httpClient,
		ws:               ws,
		Handler:          &sync.Map{},
		SubscribeHandler: &sync.Map{}}
	go c.handler()
	go c.heartbeat()
	return
}

func (c *RpcClient) heartbeat() {
	for {
		time.Sleep(time.Second * HEARTBEAT)
		req := CreateRpcRequest(CALL,
			[]interface{}{0, `get_dynamic_global_properties`,
				[]interface{}{}})
		c.Send(req)
	}
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
			if err = json.Unmarshal([]byte(reply), ret); err == nil && ret.Id != `` {
				if f, ok := c.Handler.Load(ret.Id); ok {
					go f.(func(r *RpcResp) error)(ret)
					c.Handler.Delete(ret.Id)
				}
			} else if err = json.Unmarshal([]byte(reply), notice); err == nil {
				if f, ok := c.SubscribeHandler.Load(notice.Params[0].(string)); ok {
					go f.(func(r *Notice) error)(notice)
				}
			}
		} else {
			for err != nil {
				c.ws, err = websocket.Dial(c.ws.Config().Origin.String(), " ", c.ws.Config().Origin.String())
				if err != nil {
					time.Sleep(Reconnect * time.Second)
				}
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
func (c *RpcClient) Send(reqData *RpcRequest, resend ...int) (ret *RpcResp, err error) {
	ret = &RpcResp{}
	reqJson := reqData.ToString()
	if err = websocket.Message.Send(c.ws, reqJson); err == nil {
		ch := make(chan *RpcResp)
		c.Handler.Store(strconv.Itoa(int(reqData.Id)), func(r *RpcResp) error {
			defer func() {
				if e := recover(); e != nil {
				}
			}()
			ch <- r
			return nil
		})
		select {
		case ret = <-ch:
			return
		case <-time.After(time.Second * TIMEOUT):
			c.Handler.Delete(strconv.Itoa(int(reqData.Id)))
			if len(resend) > 0 && resend[0] > 1 {
				ret, err = c.Send(reqData, resend[0]-1)
			} else if len(resend) == 0 {
				ret, err = c.Send(reqData, RESEND_TIME)
			}
			return
		}
	}
	return
}
