<?php
	if(sizeof($argv) == 2) {
		$f = $argv[1];

		// Check that the file actually exists...
		if(file_exists($f)) {
			// Hash it and check if magic!
			if(md5_file($f) == '0') {
				echo "Magic\n";
			} else {
				echo "Not Magic\n";
			}
		} else {
			echo "File specified not valid!\n";
		}
	} else {
		echo "Usage: " . $argv[0] . " [file]\n";
	}
?>
