# checkroot

## About

checkroot is a very small library that provides functions for checking whether
the program is running with root privileges on a GNU/Linux system.

## Usage

checkroot provides two simple functions: `Isroot()` returns a bool that implies
whether it is running as root or not. It also returns an error. `RootOrExit()`
checks whether it is running as root, and if it is not, it prints an error
message and exits the program with exit code 1.
