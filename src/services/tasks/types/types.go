package types

import "github.com/TsvetanMilanov/tasker-common/common/cdeclarations"

// InfoHandler ...
type InfoHandler struct {
	HTTPClient cdeclarations.IHTTPClient `di:""`
	Config     cdeclarations.IConfig     `di:""`
}
