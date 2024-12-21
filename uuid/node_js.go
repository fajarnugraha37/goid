//go:build js
// +build js

package uuid

// getHardwareInterface returns nil values for the JS version of the code.
// This removes the "net" dependency, because it is not used in the browser.
// Using the "net" library inflates the size of the transpiled JS code by 673k bytes.
func getHardwareInterface(name string) (string, []byte) { return "", nil }