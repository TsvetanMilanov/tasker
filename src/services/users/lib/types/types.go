package types

import "github.com/TsvetanMilanov/tasker-common/common/cdeclarations"

// BaseHandler ...
type BaseHandler struct {
	Config     cdeclarations.IConfig     `di:""`
	HTTPClient cdeclarations.IHTTPClient `di:""`
}
