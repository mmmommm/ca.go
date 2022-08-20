package main

type Option func(*Bicycle)

func WithBodyType(bodyType string) Option {
	return func(b *Bicycle) {
		b.BodyType = bodyType
	}
}

func WithTotalGear(frontGear, backGear int) Option {
	return func(b *Bicycle) {
		b.TotalGear = frontGear * backGear
	}
}

func NewBicycleWithFOP(kind Kind, ops ...Option) *Bicycle {
	b := Bicycle{Kind: kind}
	for _, option := range ops {
		option(&b)
	}
	return &b
}
