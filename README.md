# MD5 Magic File Generator
[![Release Mode](https://img.shields.io/badge/Release%20Mode-Stable-green.svg)](https://github.com/Cryptogenic/MD5-Magic-File-Generator/)  [![Maintenance](https://img.shields.io/badge/Maintained%3F-Yes-green.svg)](https://github.com/Cryptogenic/MD5-Magic-File-Generator/)  [![Version](https://img.shields.io/badge/Version-1.0-brightgreen.svg)](https://github.com/Cryptogenic/MD5-Magic-File-Generator/) 

The MD5 Magic File Generator project, is a tool developed in Golang to generate files with a given prefix that produce a magic MD5 hash to abuse type-juggling vulnerabilities in languages such as PHP. This is done by appending data to the end of the prefix until a magic hash is found. An example of such a magic hash is `0e578667258278439511436926540546`.

The tool not only allows you to specify the prefix for the magic file, but also the maximum size of the file (at least 5kb is recommended). You can also specify the number of threads you wish to run per-core to increase the likelihood of a quicker result. In my tests, a magic collision was found on average in 10-30 minutes running single-threaded. Multi-threaded yielded times of around 5-15 minutes.

I created this tool because I went to abuse a type juggling vulnerability in a CTF challenge, however after doing some researching I couldn't find any tools to generate your own magic hashes, only small strings that produced magic hashes. This tool could be helpful in a CTF toolbox, but could also be used for pentesting other systems outside a CTF environment that may improperly compare file hashes.

## Getting Started
Below are some instructions to get the project building on your machine. A bash script is also provided to easily build the project. Commands in the setup instructions assume you run a debian-based system!

### Prerequisites
- Golang
- PHP (if you want to run the test script to 100% verify your file is magic)

## Building
### Installing prerequisites
Firstly, if you don't have Golang and/or PHP, you can install them with the following commands respectively:

```
sudo apt-get update
sudo apt-get install golang-go
sudo apt-get install php
```

### Building the project
Finally, you can run the build.sh script to build the project.

```
./build.sh
```

You should now have an executable file named "md5magic" in the root directory. You can add this into your PATH environment variable or copy it to the system binary directory to be used anywhere.

### Adding to PATH
Only works for the logged-in user:
```
export PATH=$PATH:/path/to/project_root_dir
source ~/.bashrc
```

### System-wide Usage
Keep in mind, this requires SUDO privileges:
```
sudo cp md5magic /usr/bin/md5magic
```

## Usage
The MD5 Magic File Generator includes five potential CLI options (currently), one of which is required, being the name of the prefix file. The other options have default values and are mainly for tweaking performance.

`f` - Path to the prefix file - **required**

`m` - Maximum file size for the output file (defaults to 10000 or ~10kb)

`o` - The path to the output file (defaults to "file.out")

`t` - use threaded mode (if not specified, runs single-threaded)

`c` - Number of threads per-core (only evaluated if the `-t` option is specified - defaults to 1)

### Example Output
```
specter@ubuntu:~/md5magic$ time ./md5magic -f test.txt -o test1.out
[+] Checking md5sum of file test.txt...
[*] Hash: 0c50f5832db7e1aa9a6194a9cc431b3e
[+] Checking if file is already a magic file...
[+] File is not a magic file
[+] Generating a magic file - this may take some time...
[*] Found a magic collision!
[*]     Output File: test1.out
[*]     MD5 Hash: 0e578667258278439511436926540546
[*] Done.

real    4m9.220s
user    4m8.128s
sys     0m0.032s
```

## License
Specter (Cryptogenic) - [@SpecterDev](https://twitter.com/SpecterDev)

This project is licensed under the MIT license - see the [LICENSE.md](LICENSE.md) file for details.
