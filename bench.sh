(cd ffi/go/mcl && env PATH=$PATH:"../../../lib" LD_LIBRARY_PATH="../../../lib" DYLD_LIBRARY_PATH="../../../lib" CGO_CFLAGS="-I`greadlink -f ../../../include`" CGO_LDFLAGS="-L../../../lib" go test -v -tags bn384_256 -bench .)
