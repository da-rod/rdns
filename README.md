# rdns

Simple CLI tool to query PTR records of IP addresses and/or CIDRs.  
Handy when performing network investigations ;)

# Installation

```bash
$ go install github.com/da-rod/rdns/v2@latest
```

# Usage

```bash
$ rdns -h
```

Notes:

* Filters input (comments, empty lines, invalid entries, ...)
* "Short" IP ranges like `x.y.z` are OK and treated as `x.y.z.0/24`
* Handles defanged entries (eg: a[.]b[.]c[.]d) automatically

# Examples

From command-line argument:

```bash
$ rdns 140.82.112.32/29
140.82.112.32 [in-10.smtp.github.com.]
140.82.112.33 [lb-140-82-112-33-iad.github.com.]
140.82.112.34 [lb-140-82-112-34-iad.github.com.]
140.82.112.35 [lb-140-82-112-35-iad.github.com.]
140.82.112.36 [lb-140-82-112-36-iad.github.com.]
140.82.112.37 [lb-140-82-112-37-iad.github.com.]
140.82.112.38 [lb-140-82-112-38-iad.github.com.]
140.82.112.39 []
```

From stdin:

```bash
$ cat /tmp/ips.txt | rdns
216.58.204.0 [lhr48s21-in-f0.1e100.net.]
216.58.204.1 [lhr35s07-in-f1.1e100.net.]
216.58.204.2 [lcwawa-ad-in-f2.1e100.net.]
216.58.204.3 [lcwawa-ad-in-f3.1e100.net.]
[...]
```

# (Known) Alternative

* http://www.spamshield.org/fast-rdns.pl
