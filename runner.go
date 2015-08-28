package gunit

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/smartystreets/gunit/gunit/generate"
)

func RunFixture(t TT, outer interface{}) {
	inner := NewFixture(t, os.Stdout, testing.Verbose())
	runFixture(outer, inner)
}

type TestCaseInfo struct {
	methodIndex int
	skipped     bool
	long        bool
}

func runFixture(outer interface{}, inner *Fixture) {
	defer inner.Finalize()

	fixtureValue := reflect.ValueOf(outer)
	fixtureType := fixtureValue.Type()

	var (
		setupYes    bool
		setup       int
		teardownYes bool
		teardown    int
		tests       []TestCaseInfo
	)

	for m := 0; m < fixtureType.NumMethod(); m++ {
		methodName := fixtureType.Method(m).Name
		if strings.HasPrefix(methodName, "Setup") {
			setup = m
			setupYes = true
		} else if strings.HasPrefix(methodName, "Teardown") {
			teardown = m
			teardownYes = true
		} else if strings.HasPrefix(methodName, "Test") {
			tests = append(tests, TestCaseInfo{
				methodIndex: m,
			})
		} else if strings.HasPrefix(methodName, "SkipTest") {
			tests = append(tests, TestCaseInfo{
				methodIndex: m,
				skipped:     true,
			})
		}
		// TODO: ShortTest
	}

	for _, t := range tests {
		fixture := reflect.New(fixtureType.Elem())
		fixture.Elem().FieldByName("Fixture").Set(reflect.ValueOf(inner))
		description := generate.ToSentence(fixtureType.Method(t.methodIndex).Name)
		if t.skipped {
			inner.Skip("Skipping test case: '" + description + "'")
		} else {
			// TODO: extract to func and defer teardown
			inner.Describe(description)

			if setupYes {
				fixture.Method(setup).Call(nil)
			}

			fixture.Method(t.methodIndex).Call(nil)

			if teardownYes {
				fixture.Method(teardown).Call(nil)
			}
		}
	}
}

type Finalizer interface {
	Finalize()
}
