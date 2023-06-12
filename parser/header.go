package parser

const HeaderLine_1 = "# <!!!WARN!!!> Automatically managed by libsshd, please do not manually modify."
const HeaderLine_2Prefix = " Library Version: "

// Header is a clear configuration header.
type Header struct {
	version string
}

func (ctx *Header) SetVersion(version string) {
	ctx.version = version
}

func (ctx *Header) GetVersion() string {
	return ctx.version
}
