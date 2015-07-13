/*
This is the "[I]nter-[C]omponent [C]ommunication [C]hannel". It is a minimal messaging service to connect different servers or even different parts of huge systems across programming languages.

The basis idea is to create such messaging service on top of HTTP, because every programming language is able to process HTTP. Therefore, all messages are transformed to HTTP form values (with URL encoding).

To be able to marshal / parse the data back to objects, some additional information is added:

Example 01:
name=str:Surname
value=U29tbWVy

The HTTP form name is 'str:Surname' and the value is 'Sommer' as base64 encoded string. The 'str' is the indicator for the data type, in this case it is a string.

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

Formatting of the corresponding values (each value is at the end a base64 string).
* str := the plain UTF8 string as URL encoded. These bytes are getting base64 encoded.
* int := the little endian representation of the int. These bytes are getting base64 encoded.
* f64 := the little endian representation of the float. These bytes are getting base64 encoded.
* bool := the byte 0x1 or 0x0 for true and false. These byte will be base64 encoded.
* ui8 := These byte will be base64 encoded.
* ui8[] := These bytes are getting base64 encoded.
* int[] := the little endian representation of the integers. These bytes are getting base64 encoded.
* bool[] := the bools are getting converted to bytes (0x1 or 0x0 for true and false). These bytes are getting base64 encoded.
* str[] := each string will be URL encoded. Afterwards, join all strings by \n. These bytes are getting base64 encoded.
* f64[] := the little endian representation of the floats. These bytes are getting base64 encoded.

The format of a message is:
command=COMMAND
channel=CHANNEL
[any count of data tuples]
checksum=[the SHA512 checksum of the message]

Constrains to the environment:
The three field names 'command', 'channel' and 'checksum' are reserved, thus, you cannot use these names.
*/
package ICCC
