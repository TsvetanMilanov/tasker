package types

import "github.com/TsvetanMilanov/tasker/src/common/cdeclarations"

// InfoHandler ...
type InfoHandler struct {
	HTTPClient cdeclarations.IHTTPClient `di:""`
	Config     cdeclarations.IConfig     `di:""`
}
