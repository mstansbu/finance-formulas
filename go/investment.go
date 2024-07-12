package finance

type Period uint8

const (
	WEEKLY      Period = 52
	BIWEEKLY    Period = 26
	MONTHLY     Period = 12
	SEMIMONTHLY Period = 6
	QUARTERLY   Period = 4
	ANNUALLY    Period = 1
)

// TODO, replace return with sttring
func CompoundInterest(initialAmount uint64, contribution uint64,
	period Period, rateOfReturn uint8, numberOfPeriods uint16) uint64 {
	currentAmount := initialAmount
	compoundRoR := rateOfReturn / uint8(period)
	for i := 1; i <= int(numberOfPeriods); i++ {
		currentAmount = Interest(currentAmount, compoundRoR)
	}
}

// TODO, replace return with sttring
func Interest(initialAmount uint64, rateOfReturn float64) uint64 {
	return initialAmount * uint64((100 + rateOfReturn))
}
