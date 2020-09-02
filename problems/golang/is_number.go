package leetcode

const (
	stateInit int = iota
	stateSig
	stateInt
	statePoint
	statePointWithoutLeftInt
	stateFraction
	stateExp
	stateExpSig
	stateExpInt
	stateEnd
	stateIllegal
)

const (
	charTypeNumber int = iota
	charTypeExp
	charTypePoint
	charTypeSign
	charTypeSpace
	charTypeIllegal
)

func getCharType(c byte) int {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return charTypeNumber
	case 'e', 'E':
		return charTypeExp
	case '.':
		return charTypePoint
	case '+', '-':
		return charTypeSign
	case ' ':
		return charTypeSpace
	default:
		return charTypeIllegal
	}
}

func nextState(state int, charType int) int {
	switch state {
	case stateInit:
		switch charType {
		case charTypeSpace:
			return stateInit
		case charTypeSign:
			return stateSig
		case charTypePoint:
			return statePointWithoutLeftInt
		case charTypeNumber:
			return stateInt
		}
	case stateSig:
		switch charType {
		case charTypeNumber:
			return stateInt
		case charTypePoint:
			return statePointWithoutLeftInt
		}
	case stateInt:
		switch charType {
		case charTypeNumber:
			return stateInt
		case charTypeSpace:
			return stateEnd
		case charTypeExp:
			return stateExp
		case charTypePoint:
			return statePoint
		}
	case statePoint:
		switch charType {
		case charTypeNumber:
			return stateFraction
		case charTypeSpace:
			return stateEnd
		case charTypeExp:
			return stateExp
		}
	case statePointWithoutLeftInt:
		switch charType {
		case charTypeNumber:
			return stateFraction
		}
	case stateFraction:
		switch charType {
		case charTypeNumber:
			return stateFraction
		case charTypeExp:
			return stateExp
		case charTypeSpace:
			return stateEnd
		}
	case stateExp:
		switch charType {
		case charTypeSign:
			return stateExpSig
		case charTypeNumber:
			return stateExpInt
		}
	case stateExpSig:
		switch charType {
		case charTypeNumber:
			return stateExpInt
		}
	case stateExpInt:
		switch charType {
		case charTypeNumber:
			return stateExpInt
		case charTypeSpace:
			return stateEnd
		}
	case stateEnd:
		switch charType {
		case charTypeSpace:
			return stateEnd
		}
	}
	return stateIllegal
}

func isNumber(s string) bool {
	state := stateInit
	for _, c := range []byte(s) {
		state = nextState(state, getCharType(c))
		if state == stateIllegal {
			return false
		}
	}
	return state == stateInt || state == statePoint || state == stateFraction || state == stateExpInt || state == stateEnd
}
