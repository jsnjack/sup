package sup

import (
	"github.com/jsnjack/sshconfig"
)

var extractedHostSSHConfig map[string]*sshconfig.SSHHost

func ParseAndLoadSSHConfig(sshConfig string) (map[string]*sshconfig.SSHHost, error) {
	if sshConfig != "" {
		confHosts, err := sshconfig.ParseSSHConfig(ResolvePath(sshConfig))
		if err != nil {
			return nil, err
		}

		extractedHostSSHConfig = make(map[string]*sshconfig.SSHHost)

		// flatten Host -> *SSHHost, not the prettiest
		// but will do
		for _, conf := range confHosts {
			for _, host := range conf.Host {
				extractedHostSSHConfig[host] = conf
			}
		}
		return extractedHostSSHConfig, nil
	}
	return nil, nil
}
