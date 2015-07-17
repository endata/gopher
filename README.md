# Gopher
> Gopher is a modern, light weight, and extensible web framework for Go

[![Build Status](https://travis-ci.org/gopherlabs/gopher.svg)](https://travis-ci.org/gopherlabs/gopher)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)
[![Coverage Status](https://img.shields.io/coveralls/gopherlabs/gopher.svg)](https://coveralls.io/r/gopherlabs/gopher)
[![GoDoc](https://godoc.org/github.com/gopherlabs/gopher?status.svg)](https://godoc.org/github.com/gopherlabs/gopher)
[![Made with heart](https://img.shields.io/badge/made%20with-%E2%99%A5-orange.svg)](https://github.com/ricardo-rossi)
[![Join the chat at https://gitter.im/gopherlabs/gopher](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/gopherlabs/gopher?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

```go
Route.Get("/", func(w http.ResponseWriter, r *http.Request) {
    Render.Text(w, "Hello, Gopher!")
})
ListenAndServe()
```
```awk
> go run server.go

INFO[0000] |----------------------------------------|   
INFO[0000] |    _____                                   
INFO[0000] |   / ____|           | |                    
INFO[0000] |  | |  __  ___  _ __ | |__   ___ _ __       
INFO[0000] |  | | |_ |/ _ \| '_ \| '_ \ / _ \ '__|      
INFO[0000] |  | |__| | (_) | |_) | | | |  __/ |         
INFO[0000] |   \_____|\___/| .__/|_| |_|\___|_|         
INFO[0000] |               | |                          
INFO[0000] |               |_|                          
INFO[0000] |----------------------------------------|   
INFO[0000] | GOPHER READY FOR ACTION ON PORT 3000             
INFO[0000] |----------------------------------------|   
```
```awk
curl http://localhost:3000/
Hello, Gopher!
```

## Table of Contents

* [Overview](#overview)
* [Heroic Features](#heroic-features)
* [Getting Started](#getting-started)
* [The Basics](#the-basics)
  * [Routing](#routing)
  * [Middleware](#middleware)
  * [Context](#context)
  * [Logging](#logging)
  * [Views &amp; Templates](#views--templates)
  * [Responses](#responses)
  * [Configuration](#configuration)
* [Architecture](#architecture)
  * [IoC Container](#ioc-container)
  * [Contracts](#contracts)
  * [Facades](#facades)  
  * [Service Providers](#service-providers)  
[]([Advanced Topics](#advanced-topics)[Creating Service Providers](#creating-service-providers)[Custom Bootstrapping](#custom-bootstrapping)[API Documentation](#api-documentation)[Performance &amp; Benchmarks](#performance--benchmarks))
* [Roadmap](#roadmap)
* [Frequently Asked Questions](#frequently-asked-questions)
* [Support](#support)
* [Contribution Guide](#contribution-guide)
* [Authors](#authors)
* [License](#license)

## Overview 

//TODO

## Heroic Features

* **Simple**: Straightforward, clean *Idiomatic* Go syntax. 
* **Intuitive**: Beautiful APIs for maximum coding happiness and productivity.
* **Exposed**: No reflection, dependency injection, or hidden magic. Just clean Go interfaces.
* **Modern**: Features an IoC Container, nested Middleware, flexible Routing, and more.
* **Extensible**: Easy to add Service Providers or even replace the built-in ones.
* **Comprehensive**: Routing, Handlers, Middleware, Logging, Views, and much more.
* **Speedy**: Gopher is blazing fast. See our benchmarks.
* **Documented**: Thoroughly detailed APIs

## Getting Started

Let's create our first "Hello, Gopher!" example.

#### 1. Install Gopher

1. [Install Go](https://golang.org/dl/) and set your [GOPATH](http://golang.org/doc/code.html#GOPATH) (if you haven't already).
2. Then, from your GOPATH, type this to install Gopher and its dependencies:

```
go get github.com/gopherlabs/gopher
```

#### 2. Create your server.go file

```go
package main

import (
	"net/http"
	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, "Hello, Gopher!")
	})
	ListenAndServe()
}
```
See this example at: [routes/01_hello.go](https://github.com/gopherlabs/gopher-examples/blob/master/routes/01_hello.go)

#### 3. Run your server

```shell
go run server.go
```

You will now have a Gopher web server running on localhost:3000 (default port) and you should 
see the following output:

```awk
INFO[0000] |----------------------------------------|   
INFO[0000] | LOADING SERVICE PROVIDERS ...              
INFO[0000] |----------------------------------------|   
INFO[0000] | * LOGGER ✓                                 
INFO[0000] | * MAPPER ✓                                 
INFO[0000] | * ROUTER ✓                                 
INFO[0000] | * RENDERER ✓                                                     
INFO[0000] |----------------------------------------|   
INFO[0000] |    _____                                   
INFO[0000] |   / ____|           | |                    
INFO[0000] |  | |  __  ___  _ __ | |__   ___ _ __       
INFO[0000] |  | | |_ |/ _ \| '_ \| '_ \ / _ \ '__|      
INFO[0000] |  | |__| | (_) | |_) | | | |  __/ |         
INFO[0000] |   \_____|\___/| .__/|_| |_|\___|_|         
INFO[0000] |               | |                          
INFO[0000] |               |_|                          
INFO[0000] |----------------------------------------|   
INFO[0000] | GOPHER READY FOR ACTION ON PORT 3000             
INFO[0000] |----------------------------------------|   
 
```

#### 4. Now, try it!

```awk
curl http://localhost:3000/
Hello, Gopher!
```
Awesome, it worked!   

Next, let's take a look at some of the basic concepts of Gopher:

# The Basics

## Routing

//TODO

```go
func(rw http.ResponseWriter, req *http.Request)
```
#### Routes

Get(), Post(), etc..

#### Request Handlers

//


#### Registering A Route For Multiple Verbs

Sometimes you may need to register a route that responds to multiple HTTP verbs. 

You may do so using the [*Match()*](https://godoc.org/github.com/gopherlabs/gopher-framework#RouteFacade.Match) 
method of the [Route Facade](https://godoc.org/github.com/gopherlabs/gopher-framework#RouteFacade) as shown 
on this example:

```go
router.Match("/users", MyHandler, []string{"GET", "POST", "DELETE"}, MyAppMiddleWareRouteHanlder)
```

Or, you may even register a route that responds to all HTTP verbs using the 
[*All()*](https://godoc.org/github.com/gopherlabs/gopher-framework#RouteFacade.All) method:

```go
router.All("/products", MyHandler)
```

#### Route Parameters

//TODO

#### Not Found Router

//TODO


## Middleware

//TODO

## Context

//TODO

## Logging

//TODO

## Views &amp; Templates

//TODO

## Responses

//TODO

## Configuration

// TODO

# Architecture

//TODO

## IoC Container

//TODO

## Contracts

//TODO

## Facades  

//TODO

## Service Providers

//TODO

## Roadmap

#### v0.9 (Beta Release 07/31) - ON TRACK

- [x] Routing APIs
- [x] Request Handlers  
- [x] Nested Middleware 
- [x] Application Context
- [x] Logging
- [x] Views &amp; Templates
- [x] Responses
- [x] IoC Container
- [x] Contracts  
- [x] Facades 
- [x] Initial Documentation  

#### v1.0 (Target Release 08/31) - ON TRACK

- [ ] Extensibility APIs
- [ ] Performance Benchmarks 
- [ ] Enhanced Documentation  
- [ ] Enhanced Test Cases  

#### v2.0 (Target Release Q4 2015)

- [ ] DB/ORM 
- [ ] Mail
- [ ] Sessions 
- [ ] Caching
- [ ] Hashing
- [ ] Authentication
- [ ] Queues? 

#### v3.0 (Target Release H1 2016)

- [ ] Micro-Gopher Docker Containers
- [ ] Health &amp; Service Instrumentation
- [ ] LTS Support Options

## Frequently Asked Questions

* **This looks great and I would like to start using Gopher right now, but how stable is it?** 

   Gopher is currently in alpha (pre v0.9 release) but by all means use it now! We don't have any planned breaking API changes 
   (although we can't guarantee it won't happen) from now until the the v1.0 release as we are mainly focusing on performance, 
   testing, bug fixes, and extensibility until then.  

* **Is anyone using Gopher yet?** 

   At Gopher Labs, we are using Gopher for several customer projects. If you, or anyone you know is using Gopher please let us
   know and we can include them here. 

* **Will you provide Long-term support (LTS)?** 

   Yes, LTS options are planned for Gopher v3 (Target Release H1 2016). Please see our [Roadmap](#roadmap) for details.

## Support

**Forum:** https://groups.google.com/d/forum/gopher-framework

**Mailing List:** gopher-framework@googlegroups.com

**Twitter:** For Gopher announcements, follow us at [@gopherweb](https://twitter.com/gopherweb) 

**Chat:** Join the conversation at [![Join the chat at https://gitter.im/gopherlabs/gopher](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/gopherlabs/gopher?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Contribution Guide

Thank you for considering contributing to Gopher!

To encourage active collaboration, we strongly encourage [pull requests](https://github.com/gopherlabs/gopher/pulls), 
not just bug reports. "Bug reports" may also be sent in the form of a 
pull request containing a failing test.

However, if you file a bug report, your [issue](https://github.com/gopherlabs/gopher/issues) should contain a title 
and a clear description of the problem. You should also include as much relevant information as possible 
and a code sample that demonstrates the issue. The goal of a bug report is to make it easy to replicate 
the bug and develop a fix.

#### Opening a Pull Request

1. Fork the appropriate Gopher repository: [gopher](https://github.com/gopherlabs/gopher), 
[gopher-framework](https://github.com/gopherlabs/gopher-framework), 
[gopher-services](https://github.com/gopherlabs/gopher-services)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new [Pull Request](https://github.com/gopherlabs/gopher/pulls)

## Authors 

Created with ♥ by [Ricardo Rossi](https://github.com/ricardo-rossi) 
[@ricardojrossi](https://twitter.com/ricardojrossi), founder of Gopher Labs ✉ ricardo@gopherlabs.org

Maintained with care by all its [Contributors](https://github.com/gopherlabs/gopher/graphs/contributors) 

Built-in Service Providers include the following vendored projects:

  * gorilla/mux [gorilla/mux](https://github.com/gorilla/mux)
  * sirupsen/logrus [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
  * unrolled/render [https://github.com/unrolled/render](https://github.com/unrolled/render)

## License 

The Gopher framework is open-sourced software licensed under the [MIT license](mit-license.md)

Copyright (c) 2015 Ricardo Rossi (Gopher Labs) - ricardo@gopherlabs.org 


## Thank you for using Gopher!

We hope you enjoy using Gopher. Please Star it, Watch it, &amp; Share this repo.

Follow us on Twitter [@gopherweb](https://twitter.com/gopherweb) 

```
   _____             
  / ____|           | |                                	                        ..   ..   . ..
 | |  __  ___  _ __ | |__   ___ _ __                   	                 .  .?8MMMMMNOZ$ZONMMMN~..
 | | |_ |/ _ \| '_ \| '_ \ / _ \ '__|                  	            .   8M8~.                   .OMI. . ... .
 | |__| | (_) | |_) | | | |  __/ |                     	    . .   ...MD    . ..          .    ..    7M ?MMMMMO
  \_____|\___/| .__/|_| |_|\___|_|                     	   MM:. :DNM ...MI.. 8Z.         +I.   .. N   :M      ?M
              | |                                      	 :M .    M..  ?.       .M       $         . =  .M ~,   ?.
              |_|                                      	 M.  .7IM   .            $.   .7 .               MMM.   8.
  _    _      _                                        	.M   8MM    +.....        M   ~ =MMN         D.   MM   .I.
 | |  | |    | |                                       	.M   .Z:   ~ MMMMN.           Z+MMMMM.       D     M  .M.
 | |  | | ___| |__                                     	. M  .M.   . MMM.M        ~     MMMII.       D     Z,M7
 | |/\| |/ _ \ '_ \                                    	  .OMO$.   .O,MMN        .M    D.  .              . M..
 \  /\  /  __/ |_) |                                   	     M      .8.          M .++:.O .        $.       M .
  \/  \/ \___|_.__/                                    	     M       .D .     .=, MMMMMM. 77. .. M          ?7
 ______                                           _    	   .,8          .8MNN7  ,MMMMMMO8?. ... .           .M
 |  ___|                                         | |   	   .$=                 M .. ...  ..                  M
 | |_ _ __ __ _ _ __ ___   _____      _____  _ __| | __	    N                  N    .  . . =                 M
 |  _| '__/ _` | '_ ` _ \ / _ \ \ /\ / / _ \| '__| |/ /	   .M.                  $OO..N ?MD=                  M
 | | | | | (_| | | | | | |  __/\ V  V / (_) | |  |   < 	    M                   . .  7  N                    M
 \_| |_|  \__,_|_| |_| |_|\___| \_/\_/ \___/|_|  |_|\_\	   .M.                  . . .N  N                    M.

```