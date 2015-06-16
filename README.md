Ocean
=====

## Overview
Ocean is a smart and powerful application framework and server which uses the KISS principle ("Keep it simple, stupid"). It enables you to develop (distributed) (web or Internet of Things) applications. Therefore, Ocean contains several components to support you:
* A distributed logging component with different loggers (log to console, to a database cluster, admin's web view)
* A public facing webserver
* A private facing webserver (e.g. for administration issues)
* A messaging component called ICCC to communicate with all of your servers (or even with other parts at different programming languages at different servers)
* A template engine for web applications
* A half-automated configuration management
* A simple I18N support
* A simple database abstraction which MongoDB as database back-end

## Operation modes
You can use Ocean just as *messaging broker or e.g. as logging service* (both, centralised or distributed) by downloading and starting the executables. In this case, you have to implement your business logic somewhere else and connect that logic by ICCC component to Ocean. This case means, Ocean is a program for you.

The other operation mode is the *integrated mode*. In this case, you write at least some parts of your business logic with the programming language "Go" and you import the Ocean framework. It is still possible to have other parts of your business logic somewhere else with different programming languages, connected by Oceans ICCC component. This integraded mode means, Ocean is a framework for you.
