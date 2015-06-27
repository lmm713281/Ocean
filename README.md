Ocean
=====

## Overview
Ocean is a smart and powerful application framework and server which uses the KISS principle ("Keep it simple, stupid"). It enables you to develop (distributed) e.g. web, Internet of Things, distributed data / text mining applications. Therefore, Ocean contains several components to support you:
* A distributed logging component with different loggers (log to console, to a database cluster, admin's web view)
* A public facing webserver
* A private facing webserver (e.g. for administration issues)
* A distributed messaging component called ICCC to communicate with all of your servers (or even with other parts at different programming languages at different servers)
* A distributed template engine for web applications
* A distributed and half-automated configuration management with HTML5 admin's interface
* A distributed unique number generator which produces e.g. customer IDs
* A distributed logging service with HTML5 admin's logging viewer
* A simple I18N support
* A simple database abstraction which MongoDB as database back-end

## Operation modes
* You can use Ocean as **distributed messaging broker or e.g. as distributed logging service** by downloading and starting the executables. In this case, you have to implement your business logic somewhere else and connect that logic by the ICCC component to Ocean. This case means, Ocean is a program for you. This case is also for you e.g. if you want to get the strengths of different languages and avoid their weaknesses.

* The other operation mode is the **integrated mode**. In this case, you write at least some parts of your business logic with the programming language "Go" (http://golang.org) and you import the Ocean framework. It is still possible to have other parts of your business logic somewhere else with different programming languages, connected by Oceans ICCC component. This integraded mode means, Ocean is a framework for you.

## Stability
Several projects using already Ocean at production level. The experiences are very good: The projects are running fine over several months - normally, the server maintenance with system reboots happens before any issue occur. A long-term project for social media monitoring runs on a specially prepared server without maintenance window since more than 9 months without any restart and without any related issue.

## Environment Requirements
You have to setup a MongoDB (http://www.mongodb.org) database or database cluster. You can setup individual databases for the configuration management, for the customer data and for the logging service. It is also possible to use one database for all! In case of the logging database, you can also use one shared logging database for any count of Ocean projects. The databases may run on different dedicated servers, if you want.

Ocean uses a very small memory footprint which is usually between 5 and 20 MB, thus, the memory requirements are very low. Further, Ocean also tries to avoid any disk I/O after start-up and utilise for e.g. HTML templates and small static files an in-memory cache.

One assumption to the environment is the usage of static IP addresses. If you start Ocean with a different IP address as last time, you will receive the "*Was not able to register this host*" error. If you want to use e.g. a DHCP environment with changing addresses, you have to delete the corresponding out-dated entries from the `ICCCHosts` and `ICCCListener` collections.

## Setup
### MongoDB steps
After you have installed MongoDB and it runs as service, you have to add a database and provide a user for Ocean. Here is how to do this without any GUI tool. Assumption: The database is installed and no user was added and the authentication is not activated til now.
```
mongo # opens the command-line interface of the MongoDB database
use admin # Change to admin database
db.createUser({user: "root", pwd: "PASSWORD", roles: [ "root" ]}) # Adds an root user for all databases
use Ocean # Change to the non-existing Ocean database
db.createUser({ user: "Ocean", pwd: "PASSWORD", roles: [{ role: "dbOwner", db: "Ocean" }]}) # Create an admin user for the Ocean database
exit
```
You must activate the authentication for the MongoDB installtion now in order to allow Ocean to start.

### Common steps
To enable the Ocean's startup, two small configuration files at the working directory are required. The first file is **project.name**: This file contains one line with the project name. If the project name is longer than 10 characters, the name will be cutted after 10 characters.

The second configuration file is **configuration.json**. It contains the configuration database's information, such as user's name, password, etc. Here is an example:

```JSON
{
	"ConfigDBHostname" : "127.0.0.1:27017",
	"ConfigDBDatabase" : "MyOcean",
	"ConfigDBConfigurationCollection" : "Configuration",
	"ConfigDBConfigurationCollectionUsername" : "MyOceanConfig",
	"ConfigDBConfigurationCollectionPassword" : "PWD"
}
```
The `ConfigDBConfigurationCollection` collection gets automatically created. After both files are present and the MongoDB database or database cluster is running, Ocean is able to start. The database can and should be empty. Ocean does the configuration automatically. Due to this minimal local configuration and the distributed configuration database, the deployment of additional Ocean servers can be achieved by using e.g. data center scripts, etc.

In case you setting up an additional Ocean server, you are done. After a few minutes, the ICCC components are connected to each others. The further setup steps in case of a new environment:
* Start Ocean the first time. Ocean will connect to the configuration database. After the setup of all necessary database indexes and after the creation of the configuration collection (i.e. table), Ocean stops working. This is fine: Ocean cannot connect to the customer and logging database.
* Therefore, the second step is to edit the configuration collection's values.
* To update these values, you can use the built-in MongoDB command-line shell, instead of using a GUI tool. Please see http://docs.mongodb.org/manual/reference/method/db.collection.update/ about how to update single values. To identify the documents which you want to update it is fine to use the names.
* Example, about how to update the `CustomerDBHost` field:
```
mongo # opens the command-line interface of the MongoDB database
use Ocean # Change to the Ocean database for the authentication
db.auth("root", "PASSWORD") # Your user name and password
db.Configuration.update({Name: "CustomerDBHost"},{$set: {Value: "127.0.0.1:27017"}})
exit
```
* The most important values and the corresponding MongoDB update-codes are:
   * `CustomerDBHost`: Please provide the hostname and port for the customer database. You can use the same as for the configuration database. `db.Configuration.update({Name: "CustomerDBHost"},{$set: {Value: "127.0.0.1:27017"}})`
   * `CustomerDBDatabase`: Please provide the database name for the customer database. You can use the same as for the configuration database. `db.Configuration.update({Name: "CustomerDBDatabase"},{$set: {Value: "Ocean"}})`
   * `CustomerDBUsername`: Please provide the customer database's username. `db.Configuration.update({Name: "CustomerDBUsername"},{$set: {Value: "Ocean"}})`
   * `CustomerDBPassword`: Please provide the customer database's password. `db.Configuration.update({Name: "CustomerDBPassword"},{$set: {Value: "PASSWORD"}})`
   * `LogDBHost`: Please provide the hostname and port for the logging database. You can use the same as for the configuration database. `db.Configuration.update({Name: "LogDBHost"},{$set: {Value: "127.0.0.1:27017"}})`
   * `LogDBDatabase`: Please provide the database name for the logging database. You can use the same as for the configuration database. `db.Configuration.update({Name: "LogDBDatabase"},{$set: {Value: "Ocean"}})`
   * `LogDBUsername`: Please provide the logging database's username. `db.Configuration.update({Name: "LogDBUsername"},{$set: {Value: "Ocean"}})`
   * `LogDBPassword`: Please provide the logging database's password. `db.Configuration.update({Name: "LogDBPassword"},{$set: {Value: "PASSWORD"}})`
* Start Ocean the second time. Ocean should start-up now and provides you the admin's web interface at http://127.0.0.1:60000/. With this interface, you can now conveniently adjust additional configuration values. May you consider these values:
   * `AdminWebServerBinding`: Where should the admin/private web server connected? You should use this only at local and trusted networks.
   * `AdminWebServerEnabled`: Is the admin/private web server enabled?
   * `PublicWebServerPort`: Which port should the public web server use? If you use another web server at the front (e.g. a hardware load balancer, nginx, etc.) you can use the default. If you want to use Ocean directly to the public, use port 80.
   * `InternalCommPassword`: Replase this value with e.g. a UUID or a good password. It will be used e.g. for the ICCC communication. Therefore, it will be used for the communication between the Ocean's servers.
   * `OceanUtilizeCPUs`: How many CPUs should Ocean utilise? At the moment, Ocean uses one value for all servers. Later on, this will be replaced by a per-machine configuration table/collection to enable the usage of heterogeneous servers.
   * `LogUseConsoleLogging`: Do you wan to use the console logging? Use it for your first steps and disable it for the production usage. Use the distributed database logging instead!
   * `LogUseDatabaseLogging`: Is the database logging enabled? Yes, you should enable it.
* In case of using Ocean to provide some websites: Use the admin's web interface (http://127.0.0.1:60000/) to upload the necessary files, e.g. `templates.zip` (your HTML templates, see http://golang.org/pkg/html/template/), `staticFiles.zip` (e.g. to provide static images, CSS, JavaScript, XML, JSON data) and a `web.zip` (to separately manage web frameworks like e.g. jQuery, Bootstrap, etc.) See also https://github.com/SommerEngineering/Example003 for an example.
* Finnaly, start Ocean again. The system should now run fine. Please have a look at the ICCC startup logging messages.

### Use Ocean as distributed messaging broker
For this case, you can now start your other ICCC components. This are e.g. some Java, Python or C# programs which are using the Ocean ICCC drivers. These drivers currently under development and they will be available soon.

### Use Ocean as framework e.g. for websites
For this case, an example project with documentation is available here: https://github.com/SommerEngineering/Example003

## License
Ocean's source code is available under a BSD 2-clause license. The used third-party components have different licenses:

Component | License
--------- | -------
[UUID](https://github.com/twinj/uuid) | [Read](https://github.com/twinj/uuid/blob/master/LICENSE)
[Source Code Pro](https://github.com/adobe-fonts/source-code-pro) | [Read](https://github.com/adobe-fonts/source-code-pro/blob/master/LICENSE.txt)
[jQuery](https://github.com/jquery/jquery) | [Read](https://github.com/jquery/jquery/blob/master/LICENSE.txt)
[Modernizr](https://github.com/Modernizr/Modernizr) | [Read](https://github.com/Modernizr/Modernizr/blob/master/LICENSE)
[Webflow's front-end library](https://webflow.com/) | [Read](https://github.com/SommerEngineering/Ocean/blob/master/Admin/Assets/JSWebflow.go)
[mgo](https://github.com/go-mgo/mgo/) | [Read](https://github.com/go-mgo/mgo/blob/v2-unstable/LICENSE)
