package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var HTTP_TIMEOUT = 15 * time.Second

const maxIdleConns = 2048
const maxIdleConnsHerHost = 2048
const idleConnTimeout = 60 * time.Second

var transport = &http.Transport{
	MaxIdleConns:        maxIdleConns,
	MaxIdleConnsPerHost: maxIdleConnsHerHost,
	IdleConnTimeout:     idleConnTimeout,
}

func SetProxy(pUrl string) {
	u, err := url.Parse(pUrl)
	if err != nil {
		log.Panicf("fail to parse url '%s' - -%v", pUrl)
	}
	transport.Proxy = http.ProxyURL(u)
}

var client = &http.Client{Transport: transport}

func HttpGet(rUrl string) ([]byte, error) {
	ctx, calcel := context.WithTimeout(context.Background(), HTTP_TIMEOUT)
	defer calcel()
	req, err := http.NewRequest(http.MethodGet, rUrl, nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, nil
	}
	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("status code %d", resp.StatusCode)
	}
	return body, nil
}
