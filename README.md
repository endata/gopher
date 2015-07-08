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
* **Extensible**: Easy to add a new Service Providers or to replace a built-in one.
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

twitter: 
chat: [![Join the chat at https://gitter.im/gopherlabs/gopher](https://img.shields.io/badge/GITTER-join%20chat-green.svg)](https://gitter.im/gopherlabs/gopher?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Contribution Guide

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Authors 

Created with ♥ by [Ricardo Rossi](https://github.com/ricardo-rossi), founder of Gopher Labs ✉ ricardo@gopherlabs.org

Powered by all [Contributors](https://github.com/gopherlabs/gopher/graphs/contributors) 

## License 

The Gopher framework is open-sourced software licensed under the [MIT license](mit-license.md)

Copyright (c) 2015 Ricardo Rossi (Gopher Labs) - ricardo@gopherlabs.org 

## Thank you for using Gopher!

Star it, watch it, share it.

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

[[top]](#gopher) - [[contents]](#[[top]](#gopher))