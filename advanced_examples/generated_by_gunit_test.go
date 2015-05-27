//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func TestBowlingGameScoringTests(t *testing.T) {
	fixture := gunit.NewFixture(t)

	test0 := &BowlingGameScoringTests{Fixture: fixture}
	test0.RunTestCase__(test0.TestAfterAllGutterBallsTheScoreShouldBeZero, "After all gutter balls the score should be zero")

	test1 := &BowlingGameScoringTests{Fixture: fixture}
	test1.RunTestCase__(test1.TestAfterAllOnesTheScoreShouldBeTwenty, "After all ones the score should be twenty")

	test2 := &BowlingGameScoringTests{Fixture: fixture}
	test2.RunTestCase__(test2.TestSpareReceivesSingleRollBonus, "Spare receives single roll bonus")

	test3 := &BowlingGameScoringTests{Fixture: fixture}
	test3.RunTestCase__(test3.TestStrikeRecievesDoubleRollBonus, "Strike recieves double roll bonus")

	test4 := &BowlingGameScoringTests{Fixture: fixture}
	test4.RunTestCase__(test4.TestPerfectGame, "Perfect game")

	fixture.Finalize()
}

func (self *BowlingGameScoringTests) RunTestCase__(test func(), description string) {
	self.Describe(description)
	self.Setup()
	test()
}

///////////////////////////////////////////////////////////////////////////////

func TestEnvironmentControllerFixture(t *testing.T) {
	fixture := gunit.NewFixture(t)

	test0 := &EnvironmentControllerFixture{Fixture: fixture}
	test0.RunTestCase__(test0.TestEverythingTurnedOffAtStartup, "Everything turned off at startup")

	test1 := &EnvironmentControllerFixture{Fixture: fixture}
	test1.RunTestCase__(test1.TestEverythingOffWhenComfortable, "Everything off when comfortable")

	test2 := &EnvironmentControllerFixture{Fixture: fixture}
	test2.RunTestCase__(test2.TestCoolerAndBlowerWhenHot, "Cooler and blower when hot")

	test3 := &EnvironmentControllerFixture{Fixture: fixture}
	test3.RunTestCase__(test3.TestHeaterAndBlowerWhenCold, "Heater and blower when cold")

	test4 := &EnvironmentControllerFixture{Fixture: fixture}
	test4.RunTestCase__(test4.TestHighAlarmOnIfAtThreshold, "High alarm on if at threshold")

	test5 := &EnvironmentControllerFixture{Fixture: fixture}
	test5.RunTestCase__(test5.TestLowAlarmOnIfAtThreshold, "Low alarm on if at threshold")

	fixture.Finalize()
}

func (self *EnvironmentControllerFixture) RunTestCase__(test func(), description string) {
	self.Describe(description)
	self.Setup()
	test()
}

///////////////////////////////////////////////////////////////////////////////

func init() {
	gunit.Assertions = map[string]map[int]string{"environment_controller_hvac_fake_test.go": map[int]string{}, "environment_controller_test.go": map[int]string{32: "self.So(self.hardware.String(), should.Equal, \"heater BLOWER COOLER low high\")", 37: "self.So(self.hardware.String(), should.Equal, \"HEATER BLOWER cooler low high\")", 42: "self.So(self.hardware.String(), should.Equal, \"heater BLOWER COOLER low HIGH\")", 47: "self.So(self.hardware.String(), should.Equal,", 81: "self.So(self.hardware.String(), should.Equal, \"heater blower cooler low high\")"}, "bowling_game_test.go": map[int]string{35: "self.So(self.game.Score(), should.Equal, 21)", 43: "self.So(self.game.Score(), should.Equal, 24)", 48: "self.So(self.game.Score(), should.Equal, 300)", 22: "self.So(self.game.Score(), should.Equal, 0)", 27: "self.So(self.game.Score(), should.Equal, 20)"}}
	gunit.Validate("b86dec445670a43b72d678682518024c")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
