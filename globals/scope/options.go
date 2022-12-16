package scope

import "github.com/taubyte/go-sdk/globals/internal"

func Application(i *internal.Instance) {
	i.Application = 1
}

func Function(i *internal.Instance) {
	i.Function = 1
}
