# Gopher
> Gopher is a modern, light weight, and extensible web framework for GO

[![Build Status](https://travis-ci.org/gopherlabs/gopher.svg)](https://travis-ci.org/gopherlabs/gopher)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)
[![Coverage Status](https://img.shields.io/coveralls/gopherlabs/gopher.svg)](https://coveralls.io/r/gopherlabs/gopher)
[![GoDoc](https://godoc.org/github.com/gopherlabs/gopher?status.svg)](https://godoc.org/github.com/gopherlabs/gopher)
[![Made with heart](https://img.shields.io/badge/made%20with-%E2%99%A5-orange.svg)](https://github.com/ricardo-rossi)
[![Join the chat at https://gitter.im/gopherlabs/gopher](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/gopherlabs/gopher?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

```
   _____             
  / ____|           | |                            
 | |  __  ___  _ __ | |__   ___ _ __               
 | | |_ |/ _ \| '_ \| '_ \ / _ \ '__|              
 | |__| | (_) | |_) | | | |  __/ |                 
  \_____|\___/| .__/|_| |_|\___|_|                 
              | |                                  
              |_|                                  
```

## Table of Contents

* [Overview](#overview)
* [Heroic Features](#heroic-features)
* [Installation](#installation)
* The Basics
  * Routing
  * Request Handlers
  * Middleware
  * Application Context
  * Logging
  * Views &amp; Templates
  * Responses
* Architecture
  * IoC Container
  * Contracts
  * Facades  
  * Service Providers
* Advanced Topics
  * Creating Service Providers
  * Custom Bootstrapping 
* [API Documentation](#api-documentation)
* Performance &amp; Benchmarks
* [Roadmap](#roadmap)
* [Support](#support)
* [Contribution Guide](#contribution-guide)
* [Authors](#authors)
* [License](#license)

## Overview 

## Heroic Features

* **Simple**: Straightforward, clean *Idiomatic* Go syntax. 
* **Intuitive**: Beautiful APIs for maximum coding happiness and productivity.
* **Exposed**: No reflection or hidden magic. Just clean code.
* **Modern**: Features an IoC Container, nested Middleware, flexible Routing, and more.
* **Extensible**: Easy to add Service Providers or even replace the built-in ones.
* **Comprehensive**: Routing, Handlers, Middleware, Logging, Views, and much more.
* **Speedy**: Gopher is blazing fast. See our benchmarks.
* **Documented**: APIs are thoroughly explained in both [GoDoc](https://godoc.org/github.com/gopherlabs/gopher-framework) and on this page.    

## Installation

//TODO

# The Basics

//TODO

## Routing

//TODO

## Request Handlers

//TODO

## Middleware

//TODO

## Application Context

//TODO

## Logging

//TODO

## Views &amp; Templates

//TODO

## Responses

//TODO

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

# Advanced Topics

//TODO

## Creating Service Providers

//TODO

## Custom Bootstrapping 

//TODO

## API Documentation

//TODO
link to godoc

## Performance &amp; Benchmarks

//TODO

## Roadmap

// TODO: roadmap table

#### Road to v1.0

// TODO: roadmap table

## Support

**Forum**   
https://groups.google.com/d/forum/gopher-framework

**Mailing List**   
gopher-framework@googlegroups.com

**Twitter**   
For Gopher announcements, follow us at [@gopherweb](https://twitter.com/gopherweb) 

**Chat**   
Join the conversation at [![Join the chat at https://gitter.im/gopherlabs/gopher](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/gopherlabs/gopher?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)


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

Created with ♥ by [Ricardo Rossi](https://github.com/ricardo-rossi), founder of Gopher Labs   
 ✉ ricardo@gopherlabs.org - Twitter [@ricardojrossi](https://twitter.com/ricardojrossi)

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