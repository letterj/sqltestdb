

## sqltestdb

sqltestdb db transcount

Example

```
sqltestdb /mnt/cfsdrive00/foo.db   10 
```


Test 01

```
echo -e "\nSleep 5 seconds.\n"; sleep 5 && time ./sqltestdb /mnt/cfsdrive00/foov7 5 && echo -e "\nSLEEP 7 seconds\n"; sleep 7 && time ./sqltestdb /mnt/cfsdrive00/foov7 5
```
