package NumGen

/*
NumGen is a distributed number generator which generates unique
numbers at a distributed sysem without any centralised or master server.

Range of 64 bit integer: -9223372036854775808 til 9223372036854775807

Construction of the ID:
-----------------------

First comes the time part:

-9223372036854775808 (base number is the smallest possible int64)
+YYYY * 1000000000000000
+  MM * 10000000000000
+  DD * 100000000000
+  HH * 1000000000
+  mm * 10000000
+  ss * 100000
+ fff * 100


Second the machine part:

+ 1000000000000000000 (offset for the machine part to get a fixed length)
+ PID (capped on 822336) * 10000000000000
+ Num CPUs (capped on 99) * 100000000000
+ PageSize (capped on 999999) * 100000
+ Random 0-99 * 1000
+ Sequence 0-999 * 1


Positions:
----------

First part:

   9999 maximum year
       12 month
         31 day
           24 hours
             59 minutes
               59 seconds
                 999 milliseconds
   99991231245959999 Maximum
-9223372036854775808 Base value


Second part:

  822336 PID
        99 CPUs
          999999 Page Size
                99 Random
                  999 Sequence
  8223369999999999999 Maximum content value
 +1000000000000000000 Offset=Minimum
  9223369999999999999 Maximum of part two


The number at the year 9999, if all IDs are used, is:
99989194391184190

Therefore, this is the reserve for a further implementation:
9223372036854775807 - 99989194391184190 = 9123382842463591617

Test & Development:
http://play.golang.org/p/-wbvmFV99D

*/
