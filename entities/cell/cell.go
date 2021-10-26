package cell

type cell struct {
	state string
}

func NewCell(state string) cell {
	var cell cell = cell{state: state}
	return cell
}
