// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// UserLevelFlow is the golang structure for table c_user_level_flow.
type UserLevelFlow struct {
	Id            int64       `orm:"id,primary"      json:"id"`            // id
	UserId        int64       `orm:"user_id"         json:"userId"`        // 会员ID
	OldLevelId    int         `orm:"old_level_id"    json:"oldLevelId"`    // 原级别
	NewLevelId    int         `orm:"new_level_id"    json:"newLevelId"`    // 新级别
	OldExpireDate *gtime.Time `orm:"old_expire_date" json:"oldExpireDate"` // 原有效期
	NewExpireDate *gtime.Time `orm:"new_expire_date" json:"newExpireDate"` // 新有效期
	AdminName     string      `orm:"admin_name"      json:"adminName"`     // 操作管理员名称
	Remark        string      `orm:"remark"          json:"remark"`        // 变更描述
	CreatedAt     int         `orm:"created_at"      json:"createdAt"`     // 创建时间
}