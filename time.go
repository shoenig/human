// Author hoenig
// License MIT

package human

import (
	"bytes"
	"fmt"
	"time"
)

// Age returns a human readable string for a Duration composed of
// days, hours, and minutes. Not meant to be consumed by a computer or be precise.
// For example, time.Duration.String() might return "175h39m2.344s",
// wheras human.Duration would return "7d7h39m4s".
func Age(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) - (days * 24)
	minutes := int(d.Minutes()) - ((days * 24 * 60) + (hours * 60))
	seconds := int(d.Seconds()) - ((days * 24 * 60 * 60) + (hours * 60 * 60) + (minutes * 60))

	var b bytes.Buffer
	if days > 0 {
		b.WriteString(fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		b.WriteString(fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 {
		b.WriteString(fmt.Sprintf("%dm", minutes))
	}
	b.WriteString(fmt.Sprintf("%ds", seconds))
	return b.String()
}
