run_bench() {
    tag=$1
    curve=$2
    echo
    echo "Running with -tags=$tag and -curve=$curve"
    echo
    (
        cd ffi/go/mcl && \
        env PATH=$PATH:"../../../lib" \
        LD_LIBRARY_PATH="../../../lib" \
        DYLD_LIBRARY_PATH="../../../lib" \
        CGO_CFLAGS="-I`greadlink -f ../../../include`" \
        CGO_LDFLAGS="-L../../../lib" \
        go test init.go mcl.go mcl_bench_test.go -v -bench . -test.tags $tag -curve $curve
    )
}

#run_bench bn256 fp382-1         # this one fails: "Error initializing library: ERR mclBn_init curve=1" 
#echo; echo;
run_bench bn384 fp382-1          # a bit slower than BLS12-381: 105 mus/G1 exp, 198 mus/G2 exp, 734 mus/pairing 
#echo; echo;
#run_bench bn384_256 fp382-1     # this one fails: "Error initializing library: ERR mclBn_init curve=1"

#run_bench bn256 fp382-2         # this one fails: "Error initializing library: ERR mclBn_init curve=2" 
#echo; echo;
run_bench bn384 fp382-2          # same as fp382-1: a bit slower than BLS12-381
#echo; echo;
#run_bench bn384_256 fp382-2     # this one fails: "Error initializing library: ERR mclBn_init curve=2"

run_bench bn256 bn254           # these 3 have the same times (faster than bls12-381)
#echo; echo;
#run_bench bn384 bn254           # these 3 have the same times
#echo; echo;
#run_bench bn384_256 bn254       # these 3 have the same times

#run_bench bn256 bls12-381      # this one fails: "Error initializing library: ERR mclBn_init curve=5"
run_bench bn384 bls12-381       # these two have the same times
run_bench bn384_256 bls12-381   # these two have the same times
