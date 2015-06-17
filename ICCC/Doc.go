/*
This is the "[I]nter-[C]omponent [C]ommunication [C]hannel". It is a minimal
messaging service to connect different servers or even different parts of
huge systems across programming languages.

The basis idea is to create such messaging service on top of HTTP, because
every programming language is able to process HTTP. Therefore, all messages
are transformed to HTTP form values (with URL encoding).

To be able to marshal / parse the data back to objects, some additional
information is added:

Example 01:
name=str:Surname
value=Sommer

The HTTP form name is 'str:Surname' and the value is 'Sommer'. The 'str' is
the indicator for the data type, in this case it is a string.

Known data types are:
* str := string
* int := 64 bit integer (means for some languages long, e.g. .NET/C#)
* f64 := 64 bit float (means for some languages double, e.g. .NET/C#)
* bool := boolean
* ui8 := 8 bit unsigned integer i.e. a byte
* ui8[] := byte array
* int[] := 64 bit integer array
* bool[] := boolean array
* str[] := string array
* f64[] := 64 bit float array

Formatting of the corresponding values (each value is a string => HTTP). Plase note:
For the arrays, the name will repeated for each value.
* str := the plain UTF8 string
* int := the integer e.g. '13894612'
* f64 := the float with nine digits e.g. 9.48 gets '9.480000000'
* bool := 'true' or 'false' (lower case)
* ui8 := the byte as hexadecimal string e.g. 255 gets 'ff'
* ui8[] := the bytes as hexdecimal strings e.g. 0 255 0 gets ui8[]:name:00 ui8[]:name:ff ui8[]:name:00
* int[] := 64 bit integer array e.g. 1 2 gets int[]:name:1 int[]:name:2
* bool[] := a boolean array e.g. true true gets bool[]:name:true bool[]:name:true
* str[] := string array e.g. 'a' 'abc' gets str[]:name:a str[]:name:abc
* f64[] := 64 bit float array e.g. 1.1 1.2 gets f64[]:name:1.100000000 f64[]:name:1.2000000000

The format of a message is:
command=COMMAND
channel=CHANNEL
[any count of data tuples]
InternalCommPassword=[configured communication password e.g. an UUID etc.]

If you want to build a distributed system across the Internet, please use e.g. SSH tunnels
to keep things secret.
*/
package ICCC
