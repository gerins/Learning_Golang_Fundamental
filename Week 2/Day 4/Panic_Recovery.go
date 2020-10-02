package main

import "fmt"

func main() {
	totalGaji, err := salaryCalculation()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total Gaji Bulan ini", totalGaji)
}

type appError struct {
	customMessage string
}

func (a appError) Error() string {
	return a.customMessage
}

func newAppError(message string) *appError {
	return &appError{customMessage: message}
}

func salaryCalculation() (t float64, e error) {
	var totalGajiTemp float64
	var aErr error = nil
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, `error123`)
			fmt.Println(`Masuk if`)
			e = newAppError(fmt.Sprintf("%v", err))
			t = totalGajiTemp
		}
	}()
	for i := 0; i <= 10; i++ {
		if i == 3 {
			fmt.Println("Masuk if panic")
			panic("Fail Panic")
		}
		totalGajiTemp++
	}
	fmt.Println(`Masuk Sini`)
	return totalGajiTemp, aErr
}
