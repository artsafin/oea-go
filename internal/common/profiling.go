package common

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	MemProfileSuffix = "memprof"
)

func WriteMemProfile(id string) {
	memProfilePrefix := os.Getenv("MEM_PROFILE_PREFIX")

	if len(memProfilePrefix) == 0 {
		log.Println("skipped profiling!")
		return
	}

	if err := os.MkdirAll(memProfilePrefix, 0777); err != nil {
		log.Fatal("could not create memory profile dir: ", err)
	}

	file := fmt.Sprintf("%v/%v_%v.prof", memProfilePrefix, MemProfileSuffix, id)

	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("could not create memory profile %v: %v\n", file, err)
	}
	defer f.Close() // error handling omitted for example
	runtime.GC()    // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
