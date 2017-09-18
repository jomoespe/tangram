# Tangram. An Universal Edge-side Composition server


## Building the application

### Environment

To build the application you need an environment with:

  - [Git](https://git-scm.com/) as version control system.
  - [make](https://www.gnu.org/software/make/) as build automation tool.
  - [Go](https://golang.org) version 1.9+. To build the product.
  - [Dep](https://github.com/golang/dep), to handle dependencies.


### How to build

You can compile or build a full independent Linux binary. The executable will be in ```build/``` folder. To compile the application only need to:

```
$ make 
```

To upgrade the dependencies:

```
$ make dependencies 
```

To create a full independent binary application, you will:

```
$ make clean build
```

Once the application is build, simply execute:

```
./build/tangram
```

## Running the application

This version does not have a external configuration and have a fixed routing from **/dachop/** to **http://localhost:81/**, if in ```http://localhost:81``` there are any page serving could be using as root component of a composition. 
