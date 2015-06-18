package NumGen

import (
	"math/big"
	"time"
)

// Internal function to generate the first part of the ID.
func generateTimeID() (id int64) {

	//
	// Please have a look to the main documentation of this package,
	// if you are interested about how to calculate the parts.
	//

	// Using for all calculations big integers to get a precise result!

	bigResult := new(big.Int)
	bigBase := new(big.Int)
	bigNext := new(big.Int)
	bigBase.SetString("-9223372036854775808", 10)
	bigResult = bigBase

	t1 := time.Now().UTC()
	year := int64(t1.Year())
	month := int64(t1.Month())
	day := int64(t1.Day())
	hours := int64(t1.Hour())
	minutes := int64(t1.Minute())
	seconds := int64(t1.Second())
	milliseconds := int64(float64(t1.Nanosecond()) / 1000000.0)

	bigNext = big.NewInt(year)
	bigYear := new(big.Int)
	bigYear.SetString("1000000000000000", 10)
	bigYear = bigYear.Mul(bigYear, bigNext)
	bigResult = bigResult.Add(bigResult, bigYear)

	bigNext = big.NewInt(month)
	bigMonth := new(big.Int)
	bigMonth.SetString("10000000000000", 10)
	bigMonth = bigMonth.Mul(bigMonth, bigNext)
	bigResult = bigResult.Add(bigResult, bigMonth)

	bigNext = big.NewInt(day)
	bigDay := new(big.Int)
	bigDay.SetString("100000000000", 10)
	bigDay = bigDay.Mul(bigDay, bigNext)
	bigResult = bigResult.Add(bigResult, bigDay)

	bigNext = big.NewInt(hours)
	bigHours := new(big.Int)
	bigHours.SetString("1000000000", 10)
	bigHours = bigHours.Mul(bigHours, bigNext)
	bigResult = bigResult.Add(bigResult, bigHours)

	bigNext = big.NewInt(minutes)
	bigMinutes := new(big.Int)
	bigMinutes.SetString("10000000", 10)
	bigMinutes = bigMinutes.Mul(bigMinutes, bigNext)
	bigResult = bigResult.Add(bigResult, bigMinutes)

	bigNext = big.NewInt(seconds)
	bigSeconds := new(big.Int)
	bigSeconds.SetString("100000", 10)
	bigSeconds = bigSeconds.Mul(bigSeconds, bigNext)
	bigResult = bigResult.Add(bigResult, bigSeconds)

	bigNext = big.NewInt(milliseconds)
	bigMilliseconds := new(big.Int)
	bigMilliseconds.SetString("100", 10)
	bigMilliseconds = bigMilliseconds.Mul(bigMilliseconds, bigNext)
	bigResult = bigResult.Add(bigResult, bigMilliseconds)

	id = bigResult.Int64()
	return
}
