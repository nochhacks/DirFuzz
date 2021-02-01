# DirFuzz
Basic URL fuzzer written in GOLANG

# Prerequisites
* A GOLANG installation
* Sudo / Admin privileges

# Installation
Firstly, you need the script. You can do this one of two ways:

1) Clone the repository:

<pre>git clone https://github.com/nochhacks/DirFuzz.git</pre>

2) Diriectly download the ZIP file from the "Code" section above the landing page.


<b>(Linux Only!)</b>

Open the DirFuzz directory and make sure that the script is executable:
<pre>chmod +x dirfuzz.go</pre>

# HELP PAGE:

<pre>USAGE: ./dirfuzz.go [options]
  
  Options:
  
    -u
        Specify base URL, include http(s):// (default "http://X.X.X.X/")
    -v    
        Toggle verbose setting, ouputs all responses including non 200(OK)
    -w
        Specify path to directory wordlist (default "wordlist.txt")

   Example: ./dirfuzz.go -u http://site.com -w /path/to/wordlist.txt -v
</pre>
