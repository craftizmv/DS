package design

type UIComposer struct {
}

type IUserState interface {
	GetDMA()
	GetTimeZone()
}

type UserState struct {
}

func (u *UserState) GetDMA() {

}

func (u *UserState) GetTimeZone() {

}

func (u *UserState) GetSomeData() {

}

func NewUIComposer() *UIComposer {
	return &UIComposer{}
}

func (c *UIComposer) Compose(state IUserState) {

}

func init() {
	composer := NewUIComposer()
	us := &UserState{}
	composer.Compose(us)
}
