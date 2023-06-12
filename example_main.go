package libsshd

func main() {
	conf := &Configuration{}
	err := conf.WithClientFile("").WithDaemonFile("").Load()
	if err != nil {
		return
	}

	client := &ClientConfiguration{}
	err = client.WithFile("").Load()
	if err != nil {
		return
	}

	daemon := &DaemonConfiguration{}
	err = daemon.WithFile("").Load()
	if err != nil {
		return
	}
}
