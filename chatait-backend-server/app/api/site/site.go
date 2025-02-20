// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

// Site 站点控制器
type Site struct {
}

// Index 首页
func (s *Site) Index(r *ghttp.Request) {
	notice.Write(r, notice.NoError, "ChatAIT后台")
}
