# DirFuzz
Basic URL fuzzer written in GO.
The program iterates over a provided wordlist and checks response codes from a webserver.
GO is a server side scripting language, which is why I chose it over Python3 for a URL fuzzer.

# Prerequisites
* A GOLANG installation
* The "go-color" module. (See Installation or this <a href="https://github.com/TwinProduction/go-color">link</a>)
* Sudo / Admin privileges

# Installation
Firstly, you need the script. You can do this one of two ways:

1) Clone the repository:

<pre>git clone https://github.com/nochhacks/DirFuzz.git</pre>

2) Directly download the ZIP file from the "Code" section above the landing page.


<b>(Linux Only!)</b>

There is a GitHub module that GO needs to run the script. You can install it like this:
<pre>go get github.com/TwinProduction/go-color</pre>

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
