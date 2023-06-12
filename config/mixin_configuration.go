package config

// MixinConfiguration contains ssh client configuration and sshd daemon configuration.
type MixinConfiguration struct {
	clientConfiguration *ClientConfiguration
	daemonConfiguration *DaemonConfiguration
}

// WithClientFile specific the ssh client configuration file location.
func (ctx *MixinConfiguration) WithClientFile(file string) *MixinConfiguration {

	ctx.clientConfiguration.WithFile(file)

	return ctx
}

// WithDaemonFile specific the sshd daemon configuration file location.
func (ctx *MixinConfiguration) WithDaemonFile(file string) *MixinConfiguration {

	ctx.daemonConfiguration.WithFile(file)

	return ctx
}

// Load the default ssh and sshd configuration will be loaded.
func (ctx *MixinConfiguration) Load() error {

	err := ctx.clientConfiguration.Load()
	if err != nil {
		return err
	}

	err = ctx.daemonConfiguration.Load()
	if err != nil {
		return err
	}

	return nil
}

// Store the new ssh and sshd configuration will be saved.
func (ctx *MixinConfiguration) Store() error {
	err := ctx.clientConfiguration.Store()
	if err != nil {
		return err
	}

	err = ctx.daemonConfiguration.Store()
	if err != nil {
		return err
	}

	return nil
}

func (ctx *MixinConfiguration) GetClientConfiguration() *ClientConfiguration {
	return ctx.clientConfiguration
}

func (ctx *MixinConfiguration) GetDaemonConfiguration() *DaemonConfiguration {
	return ctx.daemonConfiguration
}
