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

# Solutions

- Run in docker
- perfmon.csv has performance stats as `seconds,allocs(MB),mallocs(MB),cpus,goroutines`

## 7. Table v2 `[256][256][256][64]byte`

- Read IPs in blocks in a routine
- Represent IP as `[4]byte`, where each byte is a segment
- Use buffered channel to send IPs
- Store IPs at a table of `[256][256][256][32]byte`
- For the last byte, mod 32 to get target index, use binary operations to read/write it
- Count insertions

Performance on 1GB IPs:

- 13 seconds
- ALWAYS 1.5GB Ram

Pros/cons: same as #6

## 6. Table `[256][256][256][256]byte`

- Read IPs in blocks in a routine
- Represent IP as `[4]byte`, where each byte is a segment
- Use buffered channel to send IPs
- Store IPs at a table of `[256][256][256][256]byte`
- Count insertions

Performance on 1GB IPs:

- 16 seconds
- ALWAYS 4GB Ram

Pros:

- Quick
- Uses the same amount of memory(4GB) despite the size of the file

Cons:

- Uses the same amount of memory(4GB) despite the size of the file

Improvements:

- We store just one value in 8 bits. We can compress that.

## 5. Radix tree

- Read IPs in blocks in a routine
- Represent IP as `[10]byte`, where each item is a digit of uint32 representation
- Use buffered channel to send IPs
- Store ips at a tree of `Node { digits []byte, children []*Node }`
- Count insertions

Performance on 1GB IPs:

- 937 seconds, 7 GB Ram. Another bad idea. Well, no point to optimize RAM usage then

## 4. Prefix tree using segment as a node

- Read IPs in blocks in a routine
- Represent IP as `4[byte]`, where each item is a segment
- Use buffered channel to send IPs
- Store ips at a tree of `Node { children map[byte]*Node }`
- Count insertions

Performance on 1GB IPs:

- another bad idea

## 3. Map of map of map of map

- Read IPs in blocks in a routine
- Represent IP as `4[byte]`, where each item is a segment
- Use buffered channel to send IPs
- Store ips at `map[byte]map[byte]map[byte]map[byte]bool`
- Sum lens

Performance:

- that was a bad idea

## 2. Naive with hash instead of string

- Read IPs in blocks in a routine
- Use buffered channel to send IPs
- Store as uint32 at `map[uin32]bool`
- Get len

Performance on 1GB IPs:

- 26 seconds, 1.5GB RAM

Performance on 3GB IPs:

- 211 seconds, about 6 GB

## 1. Naive and blunt

- Read IPs in blocks
- Use iterator to process IP one by one
- Store as strings at `map[string]bool`
- Get len

Performance on 1GB IPs:

- 113 seconds, 6.7GB RAM. Bad as expected

## Bonus. Theoretical solution on how to solve the task if we have less than 512MB RAM

- Read IPs in 100MB blocks
- Given that each character in a string is 2 bytes, we can represent any IP with with just 2 characters "ab"
- Use `buffer [50*1024*1024]uint16` (50 MB) to store IPs, store it's actual len at `blen`
- Insert IPs into buffer. Use binary search to find a place for insertion, `memcpy(pos+2, pos, blen)` block
- Whenever buffer is full - flush to disk with name like `aazz-uuid`, where `aa` is the first IP in the buffer and `zz` is the last one

IP lookup process:

- convert IP to 2 uint16 (example: `er` )
- check if IP is in buffer already (binary search)
- look for files on disk that has the range of IPs that could include the IP (example: `er` fits into range of `aazz` and `bfur`, but not `vfgh`)
- binary search the file (don't have to load it in memory)
- not found on disk - add to buffer
- buffer full - flush to disk, add `uuid` to avoid collisions
- `count-uniq-IPs = total-files-size / 4 + blen / 2`

Possibility for improvements:

- Use goroutines to search for IP in multiple files in parallel
- If we expect memory to be highly fragmented - we can divide in-memory buffer to a smaller chunks
