package env

import (
	"os"
	"github.com/qiniu/x/token/protected"
)
//line env/env.gop:9:1
func env(key string, defval string) string {
//line env/env.gop:10:1
	v := os.Getenv(key)
//line env/env.gop:11:1
	if v == "" {
//line env/env.gop:12:1
		v = defval
	}
//line env/env.gop:14:1
	return v
}
//line env/env.gop:17
func init() {
//line env/env.gop:17:1
	protected.EnvKeyName = "YAPSERVE_KEY"
//line env/env.gop:18:1
	protected.KeySalt = env("YAPSERVE_SALT", "%8aqHVK9J@8Y%7Or")
}
