/*
Flags
A flag is a way to modify the behavior of a command. 
Cobra supports fully POSIX-compliant flags as well as the 
Go flag package. A Cobra command can define flags that persist 
through to children commands and flags that are only 
available to that command.

  $ hugo server --port=1313

In the example above, ‘port’ is the flag.

Flag functionality is provided by the pflag library, a fork of the flag 
standard library which maintains the same interface while adding POSIX compliance.
*/