package common

// Define constants
const (
	// Path file input
	PATH_INPUT = "data/input/number.txt"

	// Path file output
	PATH_OUTPUT_MERGESORT_ONLY     = "data/output/numberOnly.txt"
	PATH_OUTPUT_MERGESORT_MULTI    = "data/output/numberMulti.txt"
	PATH_OUTPUT_MERGESORT_EXTERNAL = "data/output/numberExternal.txt"

	// Path temp create chunks file for external sort
	PATH_TEMP = "data/temp"

	// Max value in file
	MAXVALUE = 4_000_000_000

	// Min value in file
	MINVALUE = 1

	// Number of num in file
	NUMBER_OF_NUMBER = 4_000_000_000
	// NUMBER_OF_NUMBER = 1_000

	// Number of goroutine running
	NUMBER_OF_GOROUTINE = 50

	// Number of num in one file chunk
	CHUNK_SIZE = 1_000_000
	// CHUNK_SIZE = 600

	// Number of chunks
	NUMBER_OF_CHUCKS_FILE = NUMBER_OF_NUMBER / CHUNK_SIZE

	// Int max in Golang
	MAX_INT = ^uint(0) >> 1

	// Buffer byte per file
	BYTES_BUFF_FILE = 2048

	// number of element in buffer to write final output file
	COUNT_BUFFER = 50
)
