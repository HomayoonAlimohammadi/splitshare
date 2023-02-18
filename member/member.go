package member

type Member struct {
	Name string

	// maps member name to your debts to them
	Debts map[string]float64
}

func New(name string) *Member {
	return &Member{
		Name:  name,
		Debts: make(map[string]float64),
	}
}

// Increase member's debts to the spender
func (m *Member) Borrow(from string, amount float64) {
	m.Debts[from] += amount
}

// Calculates the amount you owe the other member
func (m *Member) Owes(other *Member) float64 {
	iOwe := m.Debts[other.Name]
	sheOwes := other.Debts[m.Name]
	return iOwe - sheOwes
}
