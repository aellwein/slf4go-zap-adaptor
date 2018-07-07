package slf4go_zap_adaptor

import "github.com/aellwein/slf4go"

func init() {
	slf4go.SetLoggerFactory(newZapLoggerFactory())
}
