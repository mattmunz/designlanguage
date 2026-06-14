# designlanguage
Nonzero Sum Design Language definition and tools

# Overview

Software can be seen to have structural elements as well as aspects that are more fluid. It can be useful when defining the structure of an application to write down its design. This design might then be implemented in multiple languages and platforms.

The Nonzero Sum Design Language is a part of the [Nonzero Sum Stack](https://nonzerosumsolutions.com/stack/index.html). The Nonzero Sum Stack is a set of recommended technologies and engineering processes for sustainable software development. For an example of an app built using the Stack (including Design Languiage) see [helloapp](https://github.com/mattmunz/helloapp).

Design Language is a language and toolset for working with software designs. By automating the process of developing software structure from an abstract design, the cost of developing and maintaining an application can be reduced.

This workflow can nicely complement a workflow that includes AI code generation, by removing from the AI coder the task of generating the application structure. This allows the agent and the AI to focus on the high-level design of the application without wasting effort on the generation of the application structure, which can be easily automated with the design language tools. Once the structure is laid down, both human and AI coders can fill in the details, confident in the accuracy and reliability of the overall application structure.

## Getting started

### 1. Create your design

Here's an example design for a commandline application.

```nzsd
User

Console
* Version String
* CommandHandlers []CommandHandler

CommandHandler
* Description String
* Interpret (Command String) -> (Response String)
* GetConfigs (Time DateTime) -> (Configs []String)

Browser
* GetElement (ID String) -> (Element Element)
```

### 2. Save your design

Your design should be saved to a path matching `$PROJECT_DIR/documentation/design/$DESIGN_NAME/$DESIGN_NAME.nzsd.txt`. For example, see the example [Geometry design](test/data/project1/documentation/design/geometry/Geometry.nzsd.txt).

### 3. Install

```bash
go install github.com/mattmunz/designlanguage
```

### 4. Go to your project directory

```bash
cd $PROJECT_DIR
```

### 3. Test run

First print the code to the commandline.

```bash
designlanguage gen 
```

Review the code and address any issues in the design that arise.

### 4. Generate code

```bash
designlanguage gen -d=false
```

### 5. Update your code to use the generated model code

(The generated code can be found under `$PROJECT_DIR/model/gen/$DESIGN_NAME)`.)

### 6. Compile and run your application

Presumably this is something like...

```bash
go build
go install
myapp 
```

## FAQ

### What target languages are supported?

Currently, the only target language is Go but support for other languages is planned. Most likely the next one will be Python. If you have a language that you would like to be supported please [get in touch with us](contact).

### Can changes be made to the syntax?

Sure. Concievably multiple front ends can be used to allow multiple syntaxes. Simple changes to the defaul syntax may also be considered.

### Why not use UML?

UML is rather complex. Design Language is intended to be extremely simple, so as to avoid operational errors. 

### Doesn't OpenAPI do this?

OpenAPI is focused on data exchange. Design Language is more general, focused on application structure. It is possible to integrate Design Language with OpenAPI, for example by generating OpenAPI model object from Design Language designs.

### What about product X, which does something similar?

It would be great for us to learn about similar products. Please [contact us](contact) to tell us all about it.

## Contributing

Contributions are welcome! Please [contact us](contact) to learn about the best ways to contribute, report bugs, submit fixes, plan new features, etc.

## Contact

See the [Nonzero Sum contact page](https://nonzerosumsolutions.com/contact.html) for ways to get in touch with us.

## TODO

* Release
* Implement designlanguage-based flow for ClimateComms, PersonalKB, etc.
* Consider making parser less strict
  * Allow trailing space
  * Allow no space before parens

## Done

* Make documentation: X Getting started, X Contriubuting, etc.
* As part of this, make example repo of a hello world service. hellocli
* X Use it for an existing project -- appkit -- generate model interfaces
* X ANTLR parser no longer worth the trouble. 
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
