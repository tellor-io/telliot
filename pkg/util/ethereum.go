// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var ethAddressRE *regexp.Regexp = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

// ValidateAddress checks if an ethereum URL is valid?
func ValidateAddress(ethereumURL string) (string, error) {
	// Get ethereum address from ethereum URL.
	// https://eips.ethereum.org/EIPS/eip-681
	parts := strings.Split(ethereumURL, ":")
	if len(parts) != 2 {
		return "", errors.New("invalid ethereum URL")
	}
	if match := ethAddressRE.MatchString(parts[1]); !match {
		return "", errors.New("invalid ethereum address")
	}
	return parts[1], nil
}
