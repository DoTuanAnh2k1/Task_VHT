package common

// Define constants
const (
	// Path file input
	PATH_INPUT = "data/input/number.txt"

	// Path file output
	PATH_OUTPUT_MERGESORT_ONLY     = "data/output/numberOnly.txt"
	PATH_OUTPUT_MERGESORT_MULTI    = "data/output/numberMulti.txt"
	PATH_OUTPUT_MERGESORT_EXTERNAL = "data/output/numberExternal.txt"
	PATH_OUTPUT_LIB_SORT           = "data/output/numberLib.txt"

	// Path temp create chunks file for external sort
	PATH_TEMP = "data/temp"

	// Max value in file
	MAXVALUE = 1_000_000_000_000_000

	// Min value in file
	MINVALUE = 1

	// Number of num in file
	NUMBER_OF_NUMBER = 4_000_000_000
	// NUMBER_OF_NUMBER = 10_000

	// Number of goroutine running
	NUMBER_OF_GOROUTINE = 16

	// Number of num in one file chunk
	CHUNK_SIZE = 50_000_000
	// CHUNK_SIZE = 5_000

	// Number of chunks
	NUMBER_OF_CHUCKS_FILE = NUMBER_OF_NUMBER / CHUNK_SIZE

	// Buffer byte per file
	BYTES_BUFF_FILE = 2048

	// number of element in buffer to write final output file
	COUNT_BUFFER = 50

	// Size of pool object
	POOL_SIZE = NUMBER_OF_CHUCKS_FILE * (BYTES_BUFF_FILE / 8) * 2
)
