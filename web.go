// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2021/6/2.

package lauvinson

import (
	"fmt"
	"milkomeda.org/pi"
	"net/http"
	"sync"
	"time"
)

var stat *pi.InterposerStat

func init() {
	var once sync.Once
	once.Do(func() {
		var createdAt = time.Date(2021, 5, 21, 0, 0, 0, 0, &time.Location{})
		stat = &pi.InterposerStat{
			Name:      "CMS",
			Author:    "vinson",
			CreatedAt: &createdAt,
		}
	})
}

type DefaultRouter struct {
}

func (engine *DefaultRouter) Stat() *pi.InterposerStat {
	return stat
}

func (engine DefaultRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintln(w, "welcome to the vinson's cms!")
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
