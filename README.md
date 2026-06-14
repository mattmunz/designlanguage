# designlanguage
Nonzero Sum Design Language definition and tools

## TODO

* ANTLR parser no longer worth the trouble. 
  * commit
* NZS design tools -- A set of tools for working with the NZS design language
  * Use it for an existing project -- appkit -- generate model interfaces    
    * Make try the command-line on the appkit lib
  * Make documentation: How To, Getting started, Contriubuting, etc.
* Make test for antlr parsing a document with a method with no arrow.
* Develop use cases in ClimateComms documentation and use it to drive PersonalKB and other projects
* Address all TODOs
* Release
* Consider making parser less strict
  * Allow trailing space
  * Allow no space before parens

## Done

* X Add inline documentation in the source file ending up in the output source code
* X comment of type in first line of body
* X comments after -- mark on line
* X Document comment including author attribution
* X Implement all the appkit syntax. TDD
* X Make a single-pass parser, line by line, with context
* X TDD test: entity with super interface, object with superinterface, renders to superType as first field
* Make a parser interface
* Move appkit to its own repository, and later rework sundries to use it.
* X Make an ANTLR grammar file
* X Generate a Go parser
* X Make a simple test and get it to work
* X Make more complex tests and get them to work
* X Make a gen command to parse design to model, and then write out interfaces
