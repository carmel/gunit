package gunit

import (
	"reflect"
	"runtime"
	"testing"
	"strings"
)

func Run(fixture interface{}, t TT) { // TODO: add Parallel() to TT and call it
	id, _, _, ok := runtime.Caller(1)
	if !ok {
		t.Fail()
		t.Log("No can do (can't resolve caller)")
		return
	}
	fullTestName := runtime.FuncForPC(id).Name()
	testCaseName := fullTestName[strings.LastIndex(fullTestName, ".")+1:]
	fixtureValue := reflect.ValueOf(fixture)
	testCase := fixtureValue.MethodByName(testCaseName)
	if !testCase.IsValid() {
		t.Fail()
		t.Log("No can do (can't resolve test method called '" + testCaseName + "')")
		return
	}

	fixtureField := fixtureValue.Elem().FieldByName("Fixture")
	if !fixtureField.IsValid() {
		t.Fail()
		t.Log("No can do (couldn't find embedded *gunit.Fixture)")
		return
	}

	if !fixtureField.CanSet() {
		t.Fail()
		t.Log("No can do (CanSet() == false)")
		return
	}

	innerFixture := NewFixture(t, testing.Verbose(), "")
	if fixtureField.Type() != reflect.TypeOf(innerFixture) {
		t.Fail()
		t.Log("No can do (Fixture field isn't a *gunit.Fixture)")
		return
	}

	fixtureField.Set(reflect.ValueOf(innerFixture))

	defer innerFixture.Finalize()

	if setup, ok := fixture.(FixtureSetup); ok {
		setup.Setup()
	}

	testCase.Call(nil) // what if the method isn't niladic?

	if teardown, ok := fixture.(FixtureTeardown); ok {
		teardown.Teardown()
	}

	// TODO: When calls to So end in failure, read the source code on disk to display an error.
}

type FixtureSetup interface {
	Setup()
}

type FixtureTeardown interface {
	Teardown()
}
