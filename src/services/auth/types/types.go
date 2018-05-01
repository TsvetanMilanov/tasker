package types

import (
	"github.com/TsvetanMilanov/tasker/src/common/cdeclarations"
)

// CallbackHandler ...
type CallbackHandler struct {
	Config     cdeclarations.IConfig     `di:""`
	HTTPClient cdeclarations.IHTTPClient `di:""`
}
