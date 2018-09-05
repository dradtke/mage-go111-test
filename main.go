// +build mage

package main

import (
	"golang.org/x/text/unicode/norm"
)

func Test() {
	print("unicode version: " + norm.Version)
}
