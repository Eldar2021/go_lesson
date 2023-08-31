/*
About

Cobra is a CLI framework for Go. It contains a library for creating 
powerful modern CLI applications and a tool to rapidly generate Cobra 
based applications and command files.

It was created by Go team member, spf13 for hugo and has been adopted 
by the most popular Go projects.

Cobra provides:
  - Easy subcommand-based CLIs: app server, app fetch, etc.
  - Fully POSIX-compliant flags (including short & long versions)
  - Nested subcommands
  - Global, local and cascading flags
  - Easy generation of applications & commands with cobra init 
    appname & cobra add cmdname
  - Intelligent suggestions (app srver… did you mean app server?)
  - Automatic help generation for commands and flags
  - Automatic help flag recognition of -h, --help, etc.
  - Automatically generated bash autocomplete for your application
  - Automatically generated man pages for your application
  - Command aliases so you can change things without breaking them
  - The flexibility to define your own help, usage, etc.
  - Optional tight integration with viper for 12-factor apps
*/
/*
Install
Using Cobra is easy. First, use `go get` to install the latest version 
of the library. This command will install the cobra generator 
executable along with the library and its dependencies:
 
 $ go get github.com/spf13/cobra

Next, include Cobra in your application:

 import "github.com/spf13/cobra"
*/
/*
Cobra is built on a structure of commands, arguments & flags.

Commands represent actions, Args are things and Flags are modifiers 
for those actions.

The best applications will read like sentences when used. Users will 
know how to use the application because they will natively understand how to use it.

The pattern to follow is APPNAME VERB NOUN --ADJECTIVE. or APPNAME COMMAND ARG --FLAG

A few good real world examples may better illustrate this point.

In the following example, ‘server’ is a command, and ‘port’ is a flag:

 $ hugo server --port=8080
*/