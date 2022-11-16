## Getting Started

These instructions will give you a copy of the project up and running on
your local machine for development
### Prerequisites

- Go 1.17+
- Docker
- (Optional) GNU Make
- (Optional) [Ginkgo v2 binary](https://onsi.github.io/ginkgo/#installing-ginkgo)

> Note:
>   Optional means that it is not required, but it will ease your life

### Installing

A step by step series of examples that tell you how to get a development
environment running

Start with installing all the dependencies
    
``` sh
go install .
```

And then you can run the project using

``` sh
make run
```

or if you don't have GNU Make

``` sh
go run main.go
```

## About My Pattern

i don't have a specific description about the pattern, 
but i'm used this pattern because if in development and 
want to change logic one endpoint, 
other endpoints are not compromised, because in folder 
feature have their own handler.   