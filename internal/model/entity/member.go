// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure for table member.
type Member struct {
	Id          uint64      `json:"id"          description:""`
	Username    string      `json:"username"    description:""`
	Realname    string      `json:"realname"    description:""`
	Nickname    string      `json:"nickname"    description:""`
	Idcard      string      `json:"idcard"      description:""`
	Group       string      `json:"group"       description:""`
	Bigclass    string      `json:"bigclass"    description:""`
	Smallclass  string      `json:"smallclass"  description:""`
	Mobile      string      `json:"mobile"      description:""`
	WebAuth     string      `json:"webAuth"     description:""`
	Pwd         string      `json:"pwd"         description:""`
	Avatar      string      `json:"avatar"      description:""`
	Sex         int         `json:"sex"         description:"0woman1man"`
	VerifyType  int         `json:"verifyType"  description:""`
	VerifyPhoto string      `json:"verifyPhoto" description:""`
	VerifyTime  *gtime.Time `json:"verifyTime"  description:""`
	Status      int         `json:"status"      description:"2dongjie"`
	Openid      string      `json:"openid"      description:""`
}
