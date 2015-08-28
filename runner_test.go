package gunit

import (
	"bytes"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestManualRunner(t *testing.T) {
	fakeT := &FakeTT{}
	output := new(bytes.Buffer)
	fixture := NewFixture(fakeT, output, true)
	example := &ExampleFixture{}
	runFixture(example, fixture)

	assert := assertions.New(t)

	assert.So(fakeT.failed, should.BeTrue)
	assert.So(fakeT.skipped, should.BeTrue)
	assert.So(test3, should.Equal, 0) // skipped

	assert.So(setup, should.BeGreaterThan, 0)
	assert.So(test1, should.BeGreaterThan, 0)
	assert.So(test2, should.BeGreaterThan, 0)
	assert.So(teardown, should.BeGreaterThan, 0)
}

////////////////////////////////////////////////////////////////////////////////
var (
	counter  int
	setup    int
	test1    int
	test2    int
	test3    int
	teardown int
)

type ExampleFixture struct {
	*Fixture
}

func (this *ExampleFixture) Setup() {
	counter++
	setup = counter
}
func (this *ExampleFixture) Teardown() {
	counter++
	teardown = counter
}
func (this *ExampleFixture) Test1() {
	counter++
	test1 = counter
}
func (this *ExampleFixture) Test2() {
	counter++
	this.So(true, should.BeFalse)
	test2 = counter
}
func (this *ExampleFixture) SkipTest3() {
	counter++
	test3 = counter
}
