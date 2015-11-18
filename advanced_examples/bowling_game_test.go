package examples

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type BowlingGameFixture struct {
	*gunit.Fixture

	game *Game
}

func (self *BowlingGameFixture) Setup() {
	self.game = NewGame()
}

func (self *BowlingGameFixture) TestAfterAllGutterBallsTheScoreShouldBeZero() {
	self.rollMany(20, 0)
	self.So(self.game.Score(), should.Equal, 0)
}

func (self *BowlingGameFixture) TestAfterAllOnesTheScoreShouldBeTwenty() {
	self.rollMany(20, 1)
	self.So(self.game.Score(), should.Equal, 20)
}

func (self *BowlingGameFixture) TestSpareReceivesSingleRollBonus() {
	self.rollSpare()
	self.game.Roll(4)
	self.game.Roll(3)
	self.rollMany(16, 0)
	self.So(self.game.Score(), should.Equal, 21)
}

func (self *BowlingGameFixture) TestStrikeReceivesDoubleRollBonus() {
	self.rollStrike()
	self.game.Roll(4)
	self.game.Roll(3)
	self.rollMany(16, 0)
	self.So(self.game.Score(), should.Equal, 24)
}

func (self *BowlingGameFixture) TestPerfectGame() {
	self.rollMany(12, 10)
	self.So(self.game.Score(), should.Equal, 300)
}

func (self *BowlingGameFixture) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		self.game.Roll(pins)
	}
}
func (self *BowlingGameFixture) rollSpare() {
	self.game.Roll(5)
	self.game.Roll(5)
}
func (self *BowlingGameFixture) rollStrike() {
	self.game.Roll(10)
}

///////////////////////////////////////////////////////////////////////////////

func TestAfterAllGutterBallsTheScoreShouldBeZero(t *testing.T) { gunit.Run(&BowlingGameFixture{}, t) }
func TestAfterAllOnesTheScoreShouldBeTwenty(t *testing.T)      { gunit.Run(&BowlingGameFixture{}, t) }
func TestSpareReceivesSingleRollBonus(t *testing.T)            { gunit.Run(&BowlingGameFixture{}, t) }
func TestStrikeReceivesDoubleRollBonus(t *testing.T)           { gunit.Run(&BowlingGameFixture{}, t) }
func TestPerfectGame(t *testing.T)                             { gunit.Run(&BowlingGameFixture{}, t) }
