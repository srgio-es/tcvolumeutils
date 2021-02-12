package output

import "fmt"

type VerboseOutput struct {
	Verbose bool
}

func (v *VerboseOutput) Println(message string) {

	if (v.Verbose) {
		fmt.Println(message)
	}

}

func (v *VerboseOutput) Printf(message string, a... interface{}) {
	if (v.Verbose) {
		fmt.Printf(message, a...)
	}
}