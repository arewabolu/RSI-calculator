package rsi

//would need redo almost completely

func Chng(lstprice int, tf []int) []int {
	changeSli := make([]int, 0)
	for _, elem := range tf {
		change := lstprice - elem
		changeSli = append(changeSli, change)
	}
	return changeSli
}

//returns average price differences
// for the price changes over a certain period.
func Average(changeSli []int) ([]int, []int) {
	//avgUp for when difference is positive
	avgUp := make([]int, 0)
	//avgUp for when difference is negative
	avgDown := make([]int, 0)
	for _, val := range changeSli {
		switch {
		case val < 0:
			val = -1 * val
			avgDown = append(avgDown, val)
		default:
			avgUp = append(avgUp, val)
		}
	}
	return avgUp, avgDown
}

func RsValue(avgUp, avgDown []int) {
	rs := make([]int, 0)
	for indUp := range avgUp {
		for indDwn := range avgUp {
			if indUp == indDwn {
				relStr := avgUp[indUp] / avgDown[indDwn]
				rs = append(rs, relStr)
			}

		}
	}
	return
}

func RSI(rs []int) {
	rsi := make([]int, 0)
	for _, elem := range rs {
		rsiVal := 100 - (100 / (1 + elem))
		rsi = append(rsi, rsiVal)
	}

}
