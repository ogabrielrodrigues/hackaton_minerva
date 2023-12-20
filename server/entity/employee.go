package entity

type employee struct {
	registry      string
	name          string
	email         string
	sector        string
	unit          string
	administrator bool
}

type Employee interface {
	SetRegistry(string)
	GetRegistry() string
	SetName(string)
	GetName() string
	SetEmail(string)
	GetEmail() string
	SetSector(string)
	GetSector() string
	SetUnit(string)
	GetUnit() string
	SetAdministrator(bool)
	GetAdministrator() bool
}

func NewEmployee(name, email, sector, unit string, administrator bool) Employee {
	registry := ""
	return &employee{
		registry,
		name,
		email,
		sector,
		unit,
		administrator,
	}
}

func (u *employee) SetRegistry(registry string) {
	u.registry = registry
}

func (u *employee) GetRegistry() string {
	return u.registry
}

func (u *employee) SetName(name string) {
	u.name = name
}

func (u *employee) GetName() string {
	return u.name
}

func (u *employee) SetEmail(email string) {
	u.email = email
}

func (u *employee) GetEmail() string {
	return u.email
}

func (u *employee) SetSector(sector string) {
	u.sector = sector
}

func (u *employee) GetSector() string {
	return u.sector
}

func (u *employee) SetUnit(unit string) {
	u.unit = unit
}

func (u *employee) GetUnit() string {
	return u.unit
}

func (u *employee) SetAdministrator(administrator bool) {
	u.administrator = administrator
}

func (u *employee) GetAdministrator() bool {
	return u.administrator
}
