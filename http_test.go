package main

import (
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestNewMultipleHostsReverseProxy(t *testing.T) {
	parseURL, err := gurl.ParseURL("lb://mogu-admin/file/getList", -1)
	if err != nil {
		return
	}
	g.Dump(parseURL)
}
