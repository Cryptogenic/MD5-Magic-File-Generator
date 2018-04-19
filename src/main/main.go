package main

import(
	"flag"
	"fmt"
	"sync"
	"runtime"
)

// Variables that are set by CLI options
var(prefixFile string)
var(maxOutputSize int)
var(outputFile string)
var(threadedMode bool)
var(numCoreThread int)

// Terminate all threads if magic file found
var(foundMagicFile bool)

var checkmtx sync.Mutex

/**
  * @Type: Function
  * @Name: init
  * @Purpose: Runs before main(), reads CLI options into variables
  **/
func init() {
	flag.StringVar(&prefixFile, "f", "", "Path to the prefix file")
	flag.IntVar(&maxOutputSize, "m", 10000, "Maximum file size")
	flag.StringVar(&outputFile, "o", "file.out", "Output file")
	flag.BoolVar(&threadedMode, "t", false, "Use threaded mode")
	flag.IntVar(&numCoreThread, "c", 1, "Number of threads to use per core")
	flag.Parse()
}

/**
  * @Type: Function
  * @Name: checkRequiredFlags
  * @Purpose: Check that mandatory options are set
  **/
func checkRequiredFlags() bool {
	if prefixFile == "" {
		return false
	}

	return true
}

/**
  * @Type: Function
  * @Name: generateMagicFile
  * @Purpose: Attempts to create a magic file - threaded by main()
  **/
func generateMagicFile(prefix string) {
	for {
		/*
		 * Due to threading, we need to check if some other thread may have succeeded,
		 * if so, bail on all threads
		 */
		checkmtx.Lock()

		if foundMagicFile {
			break
		}

		checkmtx.Unlock()

		// Let's try to create a magic file...
		if md5CreateMagicFilePrefixed(prefix) {
			break
		}
	}
}

/**
  * @Type: Function
  * @Name: main
  **/
func main() {
	foundMagicFile = false

	if !checkRequiredFlags() {
		fmt.Println("[*] The '-f' file option is required! Done.")
		return
	}

	fmt.Printf("[+] Checking md5sum of file %s...\n", prefixFile)

	md5hash, err := md5sumf(prefixFile)

	if err != nil {
		fmt.Println("[*] File could not be hashed! Done.\n")
		return
	}

	fmt.Printf("[*] Hash: %s\n", md5hash)

	fmt.Println("[+] Checking if file is already a magic file...")

	if md5magic(md5hash) {
		fmt.Println("[*] File is already magic! Done.")
		return
	}

	fmt.Println("[+] File is not a magic file")

	fmt.Println("[+] Generating a magic file - this may take some time...")

	prefix, err := readFileContents(prefixFile)

	if err != nil {
		fmt.Println("[*] Error opening the prefix file! Done.")
		return
	}

	// If threaded mode specified, launch some threads, otherwise, single-thread execution
	if threadedMode {
		// Get number of cores and multiply by specified option
		numThreads := runtime.NumCPU() * numCoreThread

		for i := 0; i < numThreads; i++ {
			fmt.Printf("[*] Launching thread %d...\n", i)
			go generateMagicFile(prefix)
		}

		// Wait until complete
		for {
			checkmtx.Lock()

			if foundMagicFile {
				break
			}

			checkmtx.Unlock()
		}
	} else {
		generateMagicFile(prefix)
	}
}