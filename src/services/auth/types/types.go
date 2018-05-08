package types

import (
	"github.com/TsvetanMilanov/tasker-common/common/cdeclarations"
)

// CallbackHandler ...
type CallbackHandler struct {
	Config     cdeclarations.IConfig     `di:""`
	HTTPClient cdeclarations.IHTTPClient `di:""`
}
