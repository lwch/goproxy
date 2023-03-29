package utils

import "github.com/lwch/logging"

func Recover(label string) {
	if err := recover(); err != nil {
		logging.Error("%s: %v", label, err)
	}
}
