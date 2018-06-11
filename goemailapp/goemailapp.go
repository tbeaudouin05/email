package main

import (
	"github.com/thomas-bamilo/email/goemail"
)

// GoEmail function reads the config.yaml file in the same folder as the application and sends the email according to config.yaml
// To understand how config.yaml is built, please look for yaml file format on Google and also look at goemail/emailconf/emailconf.go
func main() {

	goemail.GoEmail()

}
