package repositories

package controllers

import "fmt"

func BMIPercentile(age int, height float64, weight float64, gender string) (float64, error) {

	if age < 0 || height <= 0 || weight <= 0 {
		return 0, fmt.Errorf("Invalid input for BMIPercentile")
	}
	if gender != "Male" || gender != "Female" {
		return 0, fmt.Errorf("Invalid input for gender should be male or female only")
	}
	heightMeters := height * 0.01
	bmi := weight / (heightMeters * heightMeters)

	if gender == "Male" {
		return GetBMIMale(age, bmi)
	}

	if gender == "Female" {
		return GetBMIFemale(age, bmi)
	}

	return 1, nil
}

func GetBMIMale(age int, bmi float64) (float64, error) {
	if age < 20 {
		cdcChartValues := map[int]map[float64]float64{
			2:  {15.0: 5.0, 16.0: 10.0, 17.0: 15.0 /* ... */},
			3:  {15.0: 10.0, 16.0: 20.0, 17.0: 30.0 /* ... */},
			4:  {1: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			5:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			6:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			7:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			8:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			9:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			10: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			11: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			12: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			13: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			14: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			15: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			16: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			17: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			18: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			19: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
		}
		percentiles := cdcChartValues[age]
		estimatedBMI := GetClosestChartBMI(bmi, percentiles)
		CdcPercentile := percentiles[estimatedBMI]

		return CdcPercentile, nil
	}

	return bmi, nil
}

func GetBMIFemale(age int, bmi float64) (float64, error) {
	if age < 20 {
		cdcChartValues := map[int]map[float64]float64{
			2:  {15.0: 5.0, 16.0: 10.0, 17.0: 15.0 /* ... */},
			3:  {15.0: 10.0, 16.0: 20.0, 17.0: 30.0 /* ... */},
			4:  {1: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			5:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			6:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			7:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			8:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			9:  {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			10: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			11: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			12: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			13: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			14: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			15: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			16: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			17: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			18: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
			19: {15.0: 20.0, 16.0: 40.0, 17.0: 60.0 /* ... */},
		}

		percentiles := cdcChartValues[age]
		estimatedBMI := GetClosestChartBMI(bmi, percentiles)
		CdcPercentile := percentiles[estimatedBMI]

		return CdcPercentile, nil
	}

	return bmi, nil
}

func GetClosestChartBMI(targetBMI float64, percentiles map[float64]float64) float64 {
	closestBMI := -1.0
	closestDiff := -1.0

	for bmi := range percentiles {
		diff := targetBMI - bmi
		if closestDiff < 0 || diff < closestDiff {
			closestDiff = diff
			closestBMI = bmi
		}
	}

	return closestBMI
}
