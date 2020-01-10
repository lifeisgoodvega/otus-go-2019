package hellonow

import (
	"fmt"

	"github.com/beevik/ntp"
)

var cNtpHost string = "0.beevik-ntp.pool.ntp.org"

// Now - function prints current time using ntp module
func Now() error {
	time, err := ntp.Time(cNtpHost)
	if err == nil {
		fmt.Printf("%s\n", time.String())
	}

	return err
}
