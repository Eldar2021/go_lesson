/*
Commands
Command is the central point of the application. Each interaction that 
the application supports will be contained in a Command. A command can 
have children commands and optionally run an action.

  $ hugo server --port=1313
  
In the example above, ‘server’ is the command.

see more https://pkg.go.dev/github.com/spf13/cobra#Command
*/