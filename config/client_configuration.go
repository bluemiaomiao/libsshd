package config

// ClientConfiguration ssh client configuration.
type ClientConfiguration struct {
	location                         *Value `name:"Location"`
	host                             *Value `name:"Host"`
	match                            *Value `name:"Match"`
	addKeysToAgent                   *Value `name:"AddKeysToAgent"`
	addressFamily                    *Value `name:"AddressFamily"`
	batchMode                        *Value `name:"BatchMode"`
	bindAddress                      *Value `name:"BindAddress"`
	bindInterface                    *Value `name:"BindInterface"`
	canonicalDomain                  *Value `name:"CanonicalDomain"`
	canonicalizeFallbackLocal        *Value `name:"CanonicalizeFallbackLocal"`
	canonicalizeHostname             *Value `name:"CanonicalizeHostname"`
	canonicalizeMaxDots              *Value `name:"CanonicalizeMaxDots"`
	canonicalizePermittedCNAMEs      *Value `name:"CanonicalizePermittedCNAMEs"`
	cASignatureAlgorithms            *Value `name:"CASignatureAlgorithms"`
	certificateFile                  *Value `name:"CertificateFile"`
	checkHostIP                      *Value `name:"CheckHostIP"`
	ciphers                          *Value `name:"Ciphers"`
	clearAllForwardings              *Value `name:"ClearAllForwardings"`
	compression                      *Value `name:"Compression"`
	connectionAttempts               *Value `name:"ConnectionAttempts"`
	connectTimeout                   *Value `name:"ConnectTimeout"`
	controlMaster                    *Value `name:"ControlMaster"`
	controlPath                      *Value `name:"ControlPath"`
	controlPersist                   *Value `name:"ControlPersist"`
	dynamicForward                   *Value `name:"DynamicForward"`
	enableEscapeCommandline          *Value `name:"EnableEscapeCommandline"`
	enableSSHKeysign                 *Value `name:"EnableSSHKeysign"`
	escapeChar                       *Value `name:"EscapeChar"`
	exitOnForwardFailure             *Value `name:"ExitOnForwardFailure"`
	fingerprintHash                  *Value `name:"FingerprintHash"`
	forkAfterAuthentication          *Value `name:"ForkAfterAuthentication"`
	forwardAgent                     *Value `name:"ForwardAgent"`
	forwardX11                       *Value `name:"ForwardX11"`
	forwardX11Timeout                *Value `name:"ForwardX11Timeout"`
	forwardX11Trusted                *Value `name:"ForwardX11Trusted"`
	gatewayPorts                     *Value `name:"GatewayPorts"`
	globalKnownHostsFile             *Value `name:"GlobalKnownHostsFile"`
	gSSAPIAuthentication             *Value `name:"GSSAPIAuthentication"`
	gSSAPIDelegateCredentials        *Value `name:"GSSAPIDelegateCredentials"`
	hashKnownHosts                   *Value `name:"HashKnownHosts"`
	hostbasedAcceptedAlgorithms      *Value `name:"HostbasedAcceptedAlgorithms"`
	hostbasedAuthentication          *Value `name:"HostbasedAuthentication"`
	hostKeyAlgorithms                *Value `name:"HostKeyAlgorithms"`
	hostKeyAlias                     *Value `name:"HostKeyAlias"`
	hostname                         *Value `name:"Hostname"`
	identitiesOnly                   *Value `name:"IdentitiesOnly"`
	identityAgent                    *Value `name:"IdentityAgent"`
	identityFile                     *Value `name:"IdentityFile"`
	ignoreUnknown                    *Value `name:"IgnoreUnknown"`
	include                          *Value `name:"Include"`
	iPQoS                            *Value `name:"IPQoS"`
	kbdInteractiveAuthentication     *Value `name:"KbdInteractiveAuthentication"`
	kbdInteractiveDevices            *Value `name:"KbdInteractiveDevices"`
	kexAlgorithms                    *Value `name:"KexAlgorithms"`
	knownHostsCommand                *Value `name:"KnownHostsCommand"`
	localCommand                     *Value `name:"LocalCommand"`
	localForward                     *Value `name:"LocalForward"`
	logLevel                         *Value `name:"LogLevel"`
	logVerbose                       *Value `name:"LogVerbose"`
	mACs                             *Value `name:"MACs"`
	noHostAuthenticationForLocalhost *Value `name:"NoHostAuthenticationForLocalhost"`
	numberOfPasswordPrompts          *Value `name:"NumberOfPasswordPrompts"`
	passwordAuthentication           *Value `name:"PasswordAuthentication"`
	permitLocalCommand               *Value `name:"PermitLocalCommand"`
	permitRemoteOpen                 *Value `name:"PermitRemoteOpen"`
	pKCS11Provider                   *Value `name:"PKCS11Provider"`
	port                             *Value `name:"Port"`
	preferredAuthentications         *Value `name:"PreferredAuthentications"`
	proxyCommand                     *Value `name:"ProxyCommand"`
	proxyJump                        *Value `name:"ProxyJump"`
	proxyUseFdpass                   *Value `name:"ProxyUseFdpass"`
	pubkeyAuthentication             *Value `name:"PubkeyAuthentication"`
	rekeyLimit                       *Value `name:"RekeyLimit"`
	remoteCommand                    *Value `name:"RemoteCommand"`
	remoteForward                    *Value `name:"RemoteForward"`
	requestTTY                       *Value `name:"RequestTTY"`
	requiredRSASize                  *Value `name:"RequiredRSASize"`
	revokedHostKeys                  *Value `name:"RevokedHostKeys"`
	securityKeyProvider              *Value `name:"SecurityKeyProvider"`
	sendEnv                          *Value `name:"SendEnv"`
	serverAliveCountMax              *Value `name:"ServerAliveCountMax"`
	serverAliveInterval              *Value `name:"ServerAliveInterval"`
	sessionType                      *Value `name:"SessionType"`
	setEnv                           *Value `name:"SetEnv"`
	stdinNull                        *Value `name:"StdinNull"`
	streamLocalBindMask              *Value `name:"StreamLocalBindMask"`
	streamLocalBindUnlink            *Value `name:"StreamLocalBindUnlink"`
	strictHostKeyChecking            *Value `name:"StrictHostKeyChecking"`
	syslogFacility                   *Value `name:"SyslogFacility"`
	tCPKeepAlive                     *Value `name:"TCPKeepAlive"`
	tunnel                           *Value `name:"Tunnel"`
	tunnelDevice                     *Value `name:"TunnelDevice"`
	updateHostKeys                   *Value `name:"UpdateHostKeys"`
	user                             *Value `name:"User"`
	userKnownHostsFile               *Value `name:"UserKnownHostsFile"`
	verifyHostKeyDNS                 *Value `name:"VerifyHostKeyDNS"`
	visualHostKey                    *Value `name:"VisualHostKey"`
	xAuthLocation                    *Value `name:"XAuthLocation"`
}

// Load the default ssh client configuration will be loaded.
func (ctx *ClientConfiguration) Load() error {
	return nil
}

func (ctx *ClientConfiguration) setLocation(location *Value) {
	ctx.location = location
}

func (ctx *ClientConfiguration) WithFile(file string) *ClientConfiguration {
	v := &Value{}
	v.SetValue(file)
	return ctx
}

// Store the new ssh client configuration will be saved.
func (ctx *ClientConfiguration) Store() error {
	return nil
}
