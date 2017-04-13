// Package cloudsysfs provides a function for detecting the current host's cloud provider, based on the contents of the /sys filesystem.
package cloudsysfs

import "github.com/erichs/cloudsysfs/providers"

var cloudProviders = [...]func(chan<- string){
	providers.AWS,
	providers.Azure,
	providers.DigitalOcean,
	providers.GCE,
	providers.OpenStack,
}

// Detect tries to detect the current cloud provider a host is using.
// It returns a lowercase string identifying the provider if found, or empty string if none were detected.
func Detect() string {
	sysfscheck := make(chan string)
	for _, cloud := range cloudProviders {
		go cloud(sysfscheck)
	}

	provider := ""
	for i := 0; i < len(cloudProviders); i++ {
		v := <-sysfscheck
		if v != "" {
			provider = v
		}
	}

	return provider
}
