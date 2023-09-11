package shared_lib

import (
	"os/exec"
)

func Perl_inc_paths() []string {
	cmd := exec.Command("perl", "-e", `perl -e 'print join "\n", sort @INC'`)
	out, err := cmd.Output()
	paths := []string{}
	if err != nil {
		return paths
	}
	return Split(`\n`, string(out), -1)
}
