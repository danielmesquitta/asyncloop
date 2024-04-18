// Package asyncloop provides a number of commonly used
// methods for concurrent looping over slices.
//
// This package is a response to the experimental range-over-function feature,
// which is unnecessary since we can take inspiration from the slices in the standard library
// and create our own iterator functions.
//
// This code is inspired by the github.com/dreamsofcode-io/loop package,
// which addresses the same problems using the GOEXPERIMENT=rangefunc setting.
package asyncloop
