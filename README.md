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

You should calculate the number of unique addresses in this file using as
little memory and time as possible. There is a "naive" algorithm for solving
this problem (read line by line, put lines into HashSet). It's better if your
implementation is more complicated and faster than this naive algorithm.

# Solution

- Run in docker
- perfmon.csv has performance stats as `seconds,allocs(MB),mallocs(MB),cpus,goroutines`

## 1. Naive and blunt

- Read IPs in blocks
- Use iterator to process IP one by one
- Store as strings at `map[string]bool`
- Get len

Performance on 1GB IPs:

- 113 seconds, 12GB RAM. Bad as expected

## 2. Naive with hash instead of string

- Read IPs in blocks in a routine
- Use buffered channel to send IPs
- Store as uint32 at `map[uin32]bool`
- Get len

Performance on 1GB IPs:

- 26 seconds, 14.5GB RAM

Analysis:

- Map works much faster with uint32 representation of IP
- I should manage memory better

Attempt #1: Manual runtime.GC didn't work. debug.FreeOSMemory didn't work. A memory leak?

Attempt #2: Optimized file read. 13.5 GB. Still not great.

Attempt #3: Optimize file read even more. Read and parse by bite. 3GB RAM, 27 seconds

Performance on 3GB IPs:

- 388 seconds, about 10 GB

## 3. Map of map of map of map

- Read IPs in blocks in a routine
- Use buffered channel to send IPs
- Store ips at `map[byte]map[byte]map[byte]map[byte]bool`
- Sum lens

Performance on 1GB IPs:

- 53 seconds, 2GB

 Performance on 3GB IPs:

- 180 seconds, 3GB, kind of better

## 4. A better solution with prefix tree

- Read IPs in blocks in a routine
- Use buffered channel to send IPs
- Store ips at a tree of `Node { data string, children [10]Node }`
- Count leafs
