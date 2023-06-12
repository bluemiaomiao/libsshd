package config

// DaemonConfiguration sshd daemon configuration.
type DaemonConfiguration struct {
	location *Value
}

// Load the default sshd daemon configuration will be loaded.
func (ctx *DaemonConfiguration) Load() error {
	return nil
}

func (ctx *DaemonConfiguration) setLocation(location *Value) {
	ctx.location = location
}

func (ctx *DaemonConfiguration) WithFile(file string) *DaemonConfiguration {
	v := &Value{}
	v.SetValue(file)
	ctx.setLocation(v)
	return ctx
}

// Store the new sshd daemon configuration will be saved.
func (ctx *DaemonConfiguration) Store() error {
	return nil
}
