package shandong_bill

import (
	"crontab/task/shandong_bill/service"
	"reflect"

	"codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/tool"
)

type task struct {
}

func Type() reflect.Type {
	return reflect.ValueOf(task{}).Type()
}

var (
	limit int
	test  int
)

func init() {
	tool.RegFlagVar(&limit, "limit", 100000, "数量限制，默认10")
	tool.RegFlagVar(&test, "test", 0, "数量限制，默认0")
}

func (t *task) Run() {
	limit = tool.GetFlagVar("limit").(int)
	test = tool.GetFlagVar("test").(int)
	if test == 0 {
		service.NewBillSending().Send(limit)
	} else {
		service.NewBillSending().SendTest()
	}
	return
}
