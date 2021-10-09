package continent

// Continent defines the structure of table continents
type Continent struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Store defines the interface for Continent storage
type Store interface {
	ReadAll() ([]*Continent, error)
}

// Table provides implementation of Continent store
type Table struct {
	store Store
}

func NewTable(s Store) *Table {
	return &Table{store: s}
}

func (t *Table) ReadAll()
