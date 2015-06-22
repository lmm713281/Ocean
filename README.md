Ocean
=====

## Overview
Ocean is a smart and powerful application framework and server which uses the KISS principle ("Keep it simple, stupid"). It enables you to develop (distributed) (web or Internet of Things) applications. Therefore, Ocean contains several components to support you:
* A distributed logging component with different loggers (log to console, to a database cluster, admin's web view)
* A public facing webserver
* A private facing webserver (e.g. for administration issues)
* A distributed messaging component called ICCC to communicate with all of your servers (or even with other parts at different programming languages at different servers)
* A distributed template engine for web applications
* A distributed and half-automated configuration management
* A distributed unique number generator which produces e.g. customer IDs
* A simple I18N support
* A simple database abstraction which MongoDB as database back-end

## Operation modes
* You can use Ocean just as **distributed messaging broker or e.g. as distributed logging service** by downloading and starting the executables. In this case, you have to implement your business logic somewhere else and connect that logic by the ICCC component to Ocean. This case means, Ocean is a program for you.

* The other operation mode is the **integrated mode**. In this case, you write at least some parts of your business logic with the programming language "Go" (http://golang.org) and you import the Ocean framework. It is still possible to have other parts of your business logic somewhere else with different programming languages, connected by Oceans ICCC component. This integraded mode means, Ocean is a framework for you.

## Stability
Several projects using already Ocean at production level. The experiences are very good: The projects are running fine over several months - normally, the server maintenance with system reboots happens before any issue occur. A long-term project for social media monitoring runs on a specially prepared server without maintenance window since more than 9 months without any restart and without any related issue.

## Environment Requirements
You have to setup a MongoDB (http://www.mongodb.org) database or database cluster. You can setup individual databases for the configuration management, for the customer data and for the logging service. It is also possible to use one database for all! In case of the logging database, you can also use one shared logging database for any count of Ocean projects. The databases may run on different dedicated servers, if you want.

Ocean uses a very small memory footprint which is usually between 5 and 20 MB, thus, the memory requirements are very low. Further, Ocean also tries to avoid any disk I/O after start-up and utilise for e.g. HTML templates and small static files an in-memory cache.

One assumption to the environment is the usage of static IP addresses. If you start Ocean with a different IP address as last time, you will receive the "Was not able to register this host" error. If you want to use e.g. a DHCP environment with changing addresses, you have to delete the corresponding out-dated entries from the `ICCCHosts` and `ICCCListener` collections.

## Setup
To enable the Ocean's startup, two small configuration files at the working directory are required. The first file is **project.name**: This file contains one line with the project name. If the project name is longer than 10 characters, the name will be cutted after 10 characters.

The second configuration file is **configuration.json**. It contains the configuration database's information, such as user's name, password, etc. Here is an example:

```json
{
	"ConfigDBHostname" : "127.0.0.1:27017",
	"ConfigDBDatabase" : "MyOcean",
	"ConfigDBConfigurationCollection" : "Configuration",
	"ConfigDBConfigurationCollectionUsername" : "MyOceanConfig",
	"ConfigDBConfigurationCollectionPassword" : "PWD"
}
```
The ConfigDBConfigurationCollection collection gets automatically created. After both files are present and the MongoDB (cluster) is running, Ocean is able to start. The database can and should be empty. Ocean does the configuration automatically. Due to this minimal local configuration and the distributed configuration database, the deployment of additional Ocean servers can be achieved by using e.g. data center scripts, etc.

In case you setting up an additional Ocean server, you are done. After a few minutes, the ICCC components are connected to each others. The further setup steps in case of a new environment:
* Start Ocean the first time. Ocean will connect to the configuration database. After the setup of all necessary database indexes and after the creation of the configuration collection (table), Ocean stops working. This is fine: Ocean cannot connect to the customer and logging database.
* Therefore, the second step is to use your MongoDB (GUI) tool(s) and edit the configuration collection's values. The most important values are:
   * `AdminWebServerBinding`: Where should the admin/private web server connected? You should use this only at local and trusted networks.
   * `AdminWebServerEnabled`: Is the admin/private web server enabled?
   * `PublicWebServerPort`: Which port should the public web server use? If you use another web server at the front (e.g. a hardware load balancer, nginx, etc.) you can use the default. If you want to use Ocean directly to the public, use port 80.
   * `InternalCommPassword`: Replase this value with e.g. a UUID or a good password. It will be used e.g. for the ICCC communication. Therefore, it will be used for the communication between the Ocean's servers.
   * `CustomerDBHost`: Please provide the hostname and port for the customer database. You can use the same as for the configuration database.
   * `CustomerDBDatabase`: Please provide the database name for the customer database. You can use the same as for the configuration database.
   * `CustomerDBUsername`: Please provide the customer database's username.
   * `CustomerDBPassword`: Please provide the customer database's password.
   * `LogDBHost`: Please provide the hostname and port for the logging database. You can use the same as for the configuration database.
   * `LogDBDatabase`: Please provide the database name for the logging database. You can use the same as for the configuration database.
   * `LogDBUsername`: Please provide the logging database's username.
   * `LogDBPassword`: Please provide the logging database's password.
   * `OceanUtilizeCPUs`: How many CPUs should Ocean utilise? At the moment, Ocean uses one value for all servers. Later on, this will be replaced by a per-machine configuration table/collection to enable the usage of heterogeneous servers.
   * `LogUseConsoleLogging`: Do you wan to use the console logging? Use it for your first steps and disable it for the production usage. Use the distributed database logging instead!
   * `LogUseDatabaseLogging`: Is the database logging enabled? Yes, you should enable it.
* As third step, start Ocean again. The system should now run fine. Please have a look at the ICCC startup messages.

## Thanks
"github.com/twinj/uuid"
