package main

type BicycleBuilder interface {
	WithBodyType(bodyType string) BicycleBuilder
	WithTotalGear(frontGear, backGear int) BicycleBuilder
	Build() *Bicycle
}

type bicycleBuilder struct {
	kind      Kind
	bodyType  string
	totalGear int
}

func (b *bicycleBuilder) WithBodyType(bodyType string) BicycleBuilder {
	b.bodyType = bodyType
	return b
}

func (b *bicycleBuilder) WithTotalGear(frontGear, backGear int) BicycleBuilder {
	b.totalGear = frontGear * backGear
	return b
}

func (b *bicycleBuilder) Build() *Bicycle {
	return &Bicycle{
		Kind:      b.kind,
		BodyType:  b.bodyType,
		TotalGear: b.totalGear,
	}
}

func NewBicycleWithBP(kind Kind) BicycleBuilder {
	return &bicycleBuilder{
		kind: kind,
	}
}
