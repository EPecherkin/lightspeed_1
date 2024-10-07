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

- Run in docker
- perfmon.csv has performance stats as `seconds,allocs(MB),mallocs(MB),cpus,goroutines`

## 1. Naive and blunt

- Read IPs in blocks
- Use iterator to process IP one by one
- Store as strings at `map[string]bool`
- Get len

Performance on 1GB IPs:

- Took 60 seconds to stuck on 75%. It stuck because it ate 10GB of my 16GB RAM. Expected
