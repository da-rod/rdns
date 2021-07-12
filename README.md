# rdns

Simple and pretty fast CLI tool to query PTR records of IP addresses and/or ranges.  
Handy when performing network investigations ;)

# Installation

```bash
$ go install github.com/da-rod/rdns@latest
```

# Usage

Takes an IP address or range as argument:

```bash
$ rdns 140.82.118.0/28
140.82.118.0 []
140.82.118.1 [lb-140-82-118-1-ams.github.com.]
140.82.118.2 [lb-140-82-118-2-ams.github.com.]
140.82.118.3 [lb-140-82-118-3-ams.github.com.]
140.82.118.4 [lb-140-82-118-4-ams.github.com.]
140.82.118.5 [lb-140-82-118-5-ams.github.com.]
140.82.118.6 [lb-140-82-118-6-ams.github.com.]
140.82.118.7 [lb-140-82-118-7-ams.github.com.]
140.82.118.8 [lb-140-82-118-8-ams.github.com.]
140.82.118.9 [lb-140-82-118-9-ams.github.com.]
140.82.118.10 [lb-140-82-118-10-ams.github.com.]
140.82.118.11 [lb-140-82-118-11-ams.github.com.]
140.82.118.12 [lb-140-82-118-12-ams.github.com.]
140.82.118.13 [lb-140-82-118-13-ams.github.com.]
140.82.118.14 [lb-140-82-118-14-ams.github.com.]
140.82.118.15 []
```

Reads data from stdin:

```bash
$ cat /tmp/ips.txt | rdns -
216.58.204.0 [lhr35s07-in-f0.1e100.net.]
216.58.204.1 [lhr35s07-in-f1.1e100.net.]
216.58.204.2 [lhr35s07-in-f2.1e100.net.]
216.58.204.3 [lhr48s21-in-f3.1e100.net.]
[...]
```

Notes:

* Filters input which is not (valid, duh) **strict** IP or IP range
* When working from stdin, it filters out empty and commented lines (redundant, I know...)
* Accepts "short" IP ranges like `x.y.z` and treats it as `x.y.z.0/24`

# (Known) Alternative

* http://www.spamshield.org/fast-rdns.pl
