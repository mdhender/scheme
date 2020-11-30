// Package xii contains helpers for 12 Factor.
package xii

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BoolOpts specifies if a value is required or has a default value.
type BoolOpts struct {
	Required     bool
	DefaultValue bool
	Help         string // short help message if required and not found
}

// AsBool retrieves a boolean value from the environment.
// Returns an error if the value is missing or invalid.
func AsBool(key string, opts BoolOpts) (bool, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		if opts.Required {
			if opts.Help == "" {
				return opts.DefaultValue, fmt.Errorf("%s: must be exported", key)
			}
			return opts.DefaultValue, fmt.Errorf("%s: %s", key, opts.Help)
		}
		return opts.DefaultValue, nil
	}

	trimmedVal := strings.TrimSpace(val)
	if opts.Required && trimmedVal == "" {
		return opts.DefaultValue, fmt.Errorf("%s must be set", key)
	} else if val != trimmedVal {
		return opts.DefaultValue, fmt.Errorf("%s: must not contain leading or trailing spaces", key)
	}

	switch trimmedVal {
	case "false":
		return false, nil
	case "no":
		return false, nil
	case "true":
		return true, nil
	case "yes":
		return true, nil
	}
	return opts.DefaultValue, fmt.Errorf("%s: must be a valid boolean", key)
}

// IntOpts is
type IntOpts struct {
	Required     bool
	DefaultValue int
	Help         string // short help message if required and not found
}

// AsInt is retrieves an integer value from the environment.
// Returns an error if the value is missing or invalid.
func AsInt(key string, opts IntOpts) (int, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		if opts.Required {
			if opts.Help == "" {
				return 0, fmt.Errorf("%s: must be exported", key)
			}
			return 0, fmt.Errorf("%s: %s", key, opts.Help)
		}
		return opts.DefaultValue, nil
	}

	trimmedVal := strings.TrimSpace(val)
	if opts.Required && trimmedVal == "" {
		return 0, fmt.Errorf("%s must be set", key)
	} else if val != trimmedVal {
		return 0, fmt.Errorf("%s: must not contain leading or trailing spaces", key)
	}

	integer, err := strconv.Atoi(trimmedVal)
	if err != nil || trimmedVal != fmt.Sprintf("%d", integer) {
		return 0, fmt.Errorf("%s: must be a valid integer", key)
	}

	return integer, nil
}

// StringOpts is
type StringOpts struct {
	Required     bool
	DefaultValue string
	Help         string // short help message if required and not found
}

// AsString retrieves a string from the environment.
// Returns an error if the string is not set.
func AsString(key string, opts StringOpts) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		if opts.Required {
			if opts.Help == "" {
				return "", fmt.Errorf("%s: must be exported", key)
			}
			return "", fmt.Errorf("%s: %s", key, opts.Help)
		}
		return opts.DefaultValue, nil
	}

	trimmedVal := strings.TrimSpace(val)
	if opts.Required && trimmedVal == "" {
		return "", fmt.Errorf("%s: must be set", key)
	} else if val != trimmedVal {
		return "", fmt.Errorf("%s: must not contain leading or trailing spaces", key)
	}

	return val, nil
}
