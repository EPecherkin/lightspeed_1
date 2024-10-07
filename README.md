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

- Read ips.txt in blocks
- Present IP as `[4]byte` instead of string (better for memory and performance)
- How much memory we need to store hash in memory : `256*256*256*256*4 Bytes = ~17 GB`. Let's limit memory to 2GB, store the rest on the disk
- Cache in memory
  - `[256][2097152][3]` - first byte as a key, store the rest 3 in sorted array (for binary search). `2GB / 4 = 2097152 Bytes`. TODO: Check memory usage, need to take in account pointers too.
- Need a quick way to search on disk
  - Partition to folders 0-255
  - Dump sorted array data to a file, name it to indicate what range of IPs it is storing
    - b1/s2s3s4e2e3e4-uuid
      - b1 - first byte
      - s2s3s4 - 3 last bytes of IP at start of the range
      - e2e3e4 - 3 last bytes of the end of the range
      - uuid - uniq string to prevent collision

# Algorithm

- Take an IP from ips.txt
- Convert to `[4]byte`
- Check if it is known in the memory
- Not found in memory - check on disk
- Not found on disk - add to cache partition in memory, increment counter
- Partition too big - flush to disk
- Found on memory or on disk - skip, take next IP

# Possible improvements

- Several threads to search on disk
- Separate thread to insert data to a partition
- Track cache size, allow more than 2097152 records in a partition at expense of other partitions, flush the heaviest partition
