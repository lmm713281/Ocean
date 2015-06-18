package NumGen

import (
	"math/big"
	"math/rand"
	"os"
	"runtime"
)

// Internal function to generate the second part of the ID.
func generateMachineID() (id int64) {

	//
	// Please have a look to the main documentation of this package,
	// if you are interested about how to calculate the parts.
	//

	// Using for all calculations big integers to get a precise result!

	bigResult := new(big.Int)
	bigBase := new(big.Int)
	bigNext := new(big.Int)
	bigBase.SetString("1000000000000000000", 10)
	bigResult = bigBase

	pid := int64(os.Getpid())
	cpus := int64(runtime.NumCPU())
	pageSize := int64(os.Getpagesize())
	rnd := int64(rand.Intn(100))

	if pid > 822336 {
		pid = int64(822336)
	}

	if cpus > 99 {
		cpus = int64(99)
	}

	if pageSize > 999999 {
		pageSize = int64(999999)
	}

	bigNext = big.NewInt(pid)
	bigPID := new(big.Int)
	bigPID.SetString("10000000000000", 10)
	bigPID = bigPID.Mul(bigPID, bigNext)
	bigResult = bigResult.Add(bigResult, bigPID)

	bigNext = big.NewInt(cpus)
	bigCPU := new(big.Int)
	bigCPU.SetString("100000000000", 10)
	bigCPU = bigCPU.Mul(bigCPU, bigNext)
	bigResult = bigResult.Add(bigResult, bigCPU)

	bigNext = big.NewInt(pageSize)
	bigPage := new(big.Int)
	bigPage.SetString("100000", 10)
	bigPage = bigPage.Mul(bigPage, bigNext)
	bigResult = bigResult.Add(bigResult, bigPage)

	bigNext = big.NewInt(rnd)
	bigRND := new(big.Int)
	bigRND.SetString("1000", 10)
	bigRND = bigRND.Mul(bigRND, bigNext)
	bigResult = bigResult.Add(bigResult, bigRND)

	id = bigResult.Int64()
	return
}
