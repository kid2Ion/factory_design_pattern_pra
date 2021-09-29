package main

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "MAVERICK gun",
			power: 5,
		},
	}
}
