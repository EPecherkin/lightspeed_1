# lightspeed_1

Technical assignment for Lightspeed

# Task

<https://github.com/Ecwid/new-job/blob/master/IP-Addr-Counter.md>

You have a simple text file with IPv4 addresses. One line is one address, line by line:

```
145.67.23.4
8.34.5.23
89.54.3.124
89.54.3.124
3.45.71.5
...
```

The file is unlimited in size and can occupy tens and hundreds of gigabytes.

You should calculate the number of unique addresses in this file using as little memory and time as possible. There is a "naive" algorithm for solving this problem (read line by line, put lines into HashSet). It's better if your implementation is more complicated and faster than this naive algorithm.

# Solution

## Start with naive

- Read IPs in blocks
- Convert to uint32 (4 bytes hash)
- Store
- Get len

Performance on 1GB IPs:

-
