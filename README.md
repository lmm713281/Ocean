Ocean
=====

## Overview
Ocean is a smart and powerful application framework and server which uses the KISS principle ("Keep it simple, stupid"). It enables you to develop (distributed) (web or Internet of Things) applications. Therefore, Ocean contains several components to support you:
* A distributed logging component with different loggers (log to console, to a database cluster, admin's web view)
* A public facing webserver
* A private facing webserver (e.g. for administration issues)
* A messaging component called ICCC to communicate with all of your servers (or even with other parts at different programming languages at different servers)
* A distributed template engine for web applications
* A distributed half-automated configuration management
* A distributed number generator which produces e.g. customer IDs
* A simple I18N support
* A simple database abstraction which MongoDB as database back-end

## Operation modes
* You can use Ocean just as **messaging broker or e.g. as logging service** (both, centralised or distributed) by downloading and starting the executables. In this case, you have to implement your business logic somewhere else and connect that logic by ICCC component to Ocean. This case means, Ocean is a program for you.

* The other operation mode is the **integrated mode**. In this case, you write at least some parts of your business logic with the programming language "Go" and you import the Ocean framework. It is still possible to have other parts of your business logic somewhere else with different programming languages, connected by Oceans ICCC component. This integraded mode means, Ocean is a framework for you.

## Stability
Several smaller and a few bigger projects using already Ocean at production level. The experiences are very good: The projects are running fine over several months - normally, the server maintenance with system reboots happens before any issue occur. A long-term project for social media monitoring runs on a specially prepared server without maintenance window since more than 9 months without any restart and without any related issue. Other projects are running across multiple servers with automatic load balancing and fail-over (load balancing and fail-over are provided by Amazon AWS, and not yet by Ocean).

## Environment Requirements
You have to setup a MongoDB (http://www.mongodb.org) database or database cluster. You can setup individual databases for the configuration management, for the customer data and for the logging service. It is also possible to use one database for all! In case of the logging database, you can also use one shared logging database for any count of Ocean projects. The databases may run on different dedicated servers, if you want.

Ocean uses a very small memory footprint between xxx and 20 MB, thus, the memory requirements are very low. Further, Ocean also tries to avoid any disk I/O after start-up and utilise for e.g. HTML templates and small static files an in-memory cache.

## Thanks
"github.com/twinj/uuid"