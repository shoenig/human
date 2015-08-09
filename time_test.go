// Author hoenig
// License MIT

package human

import (
	"testing"
	"time"
)

func Test_Age(t *testing.T) {
	{
		// 7 days, 7 hours, 39 minutes, 4 seconds
		in := 632344944 * time.Millisecond
		out := Age(in)
		exp := "7d7h39m4s"
		if out != exp {
			t.Error("expected", exp, "got", out)
		}
	}

	{
		// 16 minutes 27 seconds
		in := 987654 * time.Millisecond
		out := Age(in)
		exp := "16m27s"
		if out != exp {
			t.Error("expected", exp, "got", out)
		}
	}
}
