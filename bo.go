package main

type ApplicationBuilder interface {
	WithSupport(flg bool) ApplicationBuilder
	WithMovie(flg bool) ApplicationBuilder
	WithBackupService(flg bool) ApplicationBuilder
	Build() *Application
}

type appBuilder struct {
	course Course
	subSupport bool
	subMovie bool
	subBackup bool
}

func (b *appBuilder) WithSupport(flg bool) ApplicationBuilder {
	b.subSupport = flg
	return b
}

func (b *appBuilder) WithMovie(flg bool) ApplicationBuilder {
	b.subMovie = flg
	return b
}

func (b *appBuilder) WithBackupService(flg bool) ApplicationBuilder {
	b.subBackup = flg
	return b
}

func (b *appBuilder) Build() *Application {
	return &Application{
		Course: b.course,
		SubscribeSupportService: b.subSupport,
		SubscribeMovieService: b.subMovie,
		SubscribeBackupService: b.subBackup,
	}
}

func NewApplicationWithBP(course Course) ApplicationBuilder {
	return &appBuilder{
		course: course,
	}
}
