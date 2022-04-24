// Build for cmount for unsupported platforms to stop go complaining
// about "no buildable Go source files "

//go:build (!linux && !darwin && !freebsd && !windows) || !brew || !cmount
// +build !linux,!darwin,!freebsd,!windows !brew !cmount

package cmount
