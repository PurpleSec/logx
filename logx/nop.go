package logx

import "os"

var (
	//Nop is a Log interface that can consume all types of
	// logs but does not do anything (a NOP). This can be used similar
	// to context.TODO or if logging is not needed and nil log instances are not
	// avaliable.
	Nop = nop(false)
)

type nop bool

func (nop) Fatal(_ string, _ ...interface{}) {
	os.Exit(1)
}
func (nop) SetLevel(_ Level)                   {}
func (nop) SetPrefix(_ string)                 {}
func (nop) Info(_ string, _ ...interface{})    {}
func (nop) Error(_ string, _ ...interface{})   {}
func (nop) Trace(_ string, _ ...interface{})   {}
func (nop) Debug(_ string, _ ...interface{})   {}
func (nop) Printf(_ string, _ ...interface{})  {}
func (nop) Warning(_ string, _ ...interface{}) {}
