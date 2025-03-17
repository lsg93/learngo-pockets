/*
This library is a basic implementation of logging in Go.

Instantiate a logger with pocketlog.New() and pass it a threshold, t.
This determines the level at which messages should be discarded.

The library comes with three thresholds for logging:
  - Debug : Messages for debugging, ideally turned off in production environments
  - Info : Logging of general events in the application
  - Error : Logging when errors are thrown.
*/
package pocketlog
