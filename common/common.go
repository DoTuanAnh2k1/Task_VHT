package common

// Define constants
const (
	PATH_INPUT                     = "data/input/number.txt"
	PATH_OUTPUT_MERGESORT_ONLY     = "data/output/numberOnly.txt"
	PATH_OUTPUT_MERGESORT_MULTI    = "data/output/numberMulti.txt"
	PATH_OUTPUT_MERGESORT_EXTERNAL = "data/output/numberExternal.txt"
	PATH_TEMP                      = "data/temp"
	MAXVALUE                       = 400_000_000
	MINVALUE                       = 1
	NUMBER_OF_NUMBER               = 400_000_000
	NUMBER_OF_GOROUTINE            = 50
	CHUNK_SIZE                     = 1_000_000
	NUMBER_OF_CHUCKS_FILE          = NUMBER_OF_NUMBER / CHUNK_SIZE
)
