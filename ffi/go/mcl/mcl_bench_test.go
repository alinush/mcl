package mcl

import "testing"
import "fmt"
import "os"

//import "log"
import "crypto/rand"
import "flag"

func BenchmarkG1mul(b *testing.B) {
	var a Fr
	var R, P G1

	//fmt.Printf("P=%s\n", P.GetString(16))
	//fmt.Printf("a=%s\n", a.GetString(16))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		var bytes = make([]byte, 32)
		_, err := rand.Read(bytes)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		P.HashAndMapTo(bytes)
		a.SetHashOf(bytes)

		b.StartTimer()
		G1Mul(&R, &P, &a)
	}
}

func BenchmarkG2mul(b *testing.B) {
	var a Fr
	var R, P G2

	//fmt.Printf("P=%s\n", P.GetString(16))
	//fmt.Printf("a=%s\n", a.GetString(16))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		var bytes = make([]byte, 32)
		_, err := rand.Read(bytes)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		P.HashAndMapTo(bytes)
		a.SetHashOf(bytes)

		b.StartTimer()
		G2Mul(&R, &P, &a)
	}
}

func BenchmarkPairing(b *testing.B) {
	var e GT
	var Q G1
	var P G2

	//fmt.Printf("P=%s\n", P.GetString(16))
	//fmt.Printf("a=%s\n", a.GetString(16))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		var bytes = make([]byte, 32)
		_, err := rand.Read(bytes)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		P.HashAndMapTo(bytes)
		Q.HashAndMapTo(bytes)

		b.StartTimer()
		Pairing(&e, &Q, &P)
	}
}

//func testVecG1(t *testing.T) {
//	N := 50
//	xVec := make([]G1, N)
//	yVec := make([]Fr, N)
//	xVec[0].HashAndMapTo([]byte("aa"))
//	var R1, R2 G1
//	for i := 0; i < N; i++ {
//		if i > 0 {
//			G1Dbl(&xVec[i], &xVec[i-1])
//		}
//		yVec[i].SetByCSPRNG()
//		G1Mul(&R1, &xVec[i], &yVec[i])
//		G1Add(&R2, &R2, &R1)
//	}
//	G1MulVec(&R1, xVec, yVec)
//	if !R1.IsEqual(&R2) {
//		t.Errorf("wrong G1MulVec")
//	}
//}
//
//func testVecG2(t *testing.T) {
//	N := 50
//	xVec := make([]G2, N)
//	yVec := make([]Fr, N)
//	xVec[0].HashAndMapTo([]byte("aa"))
//	var R1, R2 G2
//	for i := 0; i < N; i++ {
//		if i > 0 {
//			G2Dbl(&xVec[i], &xVec[i-1])
//		}
//		yVec[i].SetByCSPRNG()
//		G2Mul(&R1, &xVec[i], &yVec[i])
//		G2Add(&R2, &R2, &R1)
//	}
//	G2MulVec(&R1, xVec, yVec)
//	if !R1.IsEqual(&R2) {
//		t.Errorf("wrong G2MulVec")
//	}
//}
//
//func testVecPairing(t *testing.T) {
//	N := 50
//	xVec := make([]G1, N)
//	yVec := make([]G2, N)
//	var e1, e2 GT
//	e1.SetInt64(1)
//	for i := 0; i < N; i++ {
//		xVec[0].HashAndMapTo([]byte("aa"))
//		yVec[0].HashAndMapTo([]byte("aa"))
//		Pairing(&e2, &xVec[i], &yVec[i])
//		GTMul(&e1, &e1, &e2)
//	}
//	MillerLoopVec(&e2, xVec, yVec)
//	FinalExp(&e2, &e2)
//	if !e1.IsEqual(&e2) {
//		t.Errorf("wrong MillerLoopVec")
//	}
//}

var curve = flag.String("curve", "", "Type of curve to test: bn254, fp382-1, fp382-2, bls12-381")

func TestMain(m *testing.M) {
	flag.Parse()

	var c int
	switch *curve {
	case "":
		// so that when make.sh runs this we just exit
		fmt.Printf("No argument given, exiting benchmark...\n")
		os.Exit(0)
	case "bn254":
		c = CurveFp254BNb
	case "fp382-1":
		c = CurveFp382_1
	case "fp382-2":
		c = CurveFp382_2
	case "bls12-381":
		c = BLS12_381
	default:
		fmt.Printf("Unknown curve type: %s\n", *curve)
		os.Exit(1)
	}

	fmt.Printf("Testing/benchmarking curve: %s (%d)\n", *curve, c)
	fmt.Printf("GetMaxOpUnitSize() = %d\n", GetMaxOpUnitSize())
	fmt.Printf("GetFrUnitSize() = %d\n", GetFrUnitSize())

	err := Init(c)
	if err != nil {
		fmt.Printf("Error initializing library: %v\n", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}
