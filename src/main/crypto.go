package main 

import(
	"io"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

/**
  * @Type: Function
  * @Name: md5sum
  * @Purpose: Returns the md5 hash of a string
  **/
func md5sum(s string) string {
	h := md5.New()

	io.WriteString(h, s)

	return hex.EncodeToString(h.Sum(nil)[:16])
}

/**
  * @Type: Function
  * @Name: md5sumf
  * @Purpose: Returns the md5 hash of a file
  **/
func md5sumf(path string) (string, error) {
	contents, err := readFileContents(path)

	if err != nil {
		return "", err
	}

	return md5sum(contents), err
}

/**
  * @Type: Function
  * @Name: md5magic
  * @Purpose: Returns true if the hash is magic, false otherwise
  **/
func md5magic(hash string) bool {
	if hash[0:2] != "0e" {
		return false
	}

	return isNumeric(hash[2:])
}

/**
  * @Type: Function
  * @Name: md5sum
  * @Purpose: Internal function that attempts to create a magic file
  **/
func md5CreateMagicFilePrefixed(prefix string) bool {
	/*
	 * Loop until one of the following conditions occur:
	 * 1) The md5 hash is evaluated to be magic
	 * 2) The maximum file size limit has been reached
	 * In case 1, the function returns successfully,
	 * in case 2, the function returns false and is recalled.
	 */
	sizeOfGarbage := maxOutputSize - len(prefix)
	endResult := prefix

	// Ensure we have sufficient room
	if sizeOfGarbage <= 0 {
		fmt.Println("[*] Not enough room between prefix and max file size to work with! Done.")

		/*
		 * Technically, it's not "found" and we should return false here, but because the script cannot
		 * continue, return true and set "found" to true to terminate!
		 */
		checkmtx.Lock()
		foundMagicFile = true
		checkmtx.Unlock()

		return true
	}

	for i := 0; i < sizeOfGarbage; i++ {
		// Check if we have a magic hash
		md5hash := md5sum(endResult)

		if md5magic(md5hash) {
			writeFileContents(outputFile, endResult)
			fmt.Println("[*] Found a magic collision!")
			fmt.Printf("[*] \tOutput File: %s\n", outputFile)
			fmt.Printf("[*] \tMD5 Hash: %s\n", md5hash)
			fmt.Println("[*] Done.")

			checkmtx.Lock()
			foundMagicFile = true
			checkmtx.Unlock()

			return true
		}

		// No dice, add more garbage
		endResult += genRandomString(1) // Only add 1 char at a time, since 1 char can change a hash dramatically
	}

	// If we reached here, we didn't succeed and used all our room, retry
	return false
}