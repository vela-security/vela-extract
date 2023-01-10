package extract

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/export"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

func helperL(L *lua.LState, convert func(string) error) {
	n := L.GetTop()
	if n == 0 {
		return
	}

	for i := 1; i <= n; i++ {
		v := L.CheckString(i)
		if err := convert(v); err != nil {
			L.RaiseError("#%d convert query fail %v", i, err)
		}
	}
}

func WithEnv(env assert.Environment) {
	xEnv = env
	kv := lua.NewUserKV()
	kv.Set("jq", lua.NewFunction(jqL))
	kv.Set("xpath", lua.NewFunction(xpathL))
	kv.Set("regex", lua.NewFunction(regexL))
	kv.Set("ipv4", ipv4)
	kv.Set("number", num)
	kv.Set("url", uRL)
	kv.Set("phone", phone)
	xEnv.Set("extract", export.New("vela.extract.export", export.WithTable(kv)))
}
