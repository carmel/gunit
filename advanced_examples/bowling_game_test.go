package examples

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func TestGutterBalls(t *testing.T) { gunit.Run(&BowlingFixture{}, t) }
func TestAllOnes(t *testing.T)     { gunit.Run(&BowlingFixture{}, t) }
func TestSpare(t *testing.T)       { gunit.Run(&BowlingFixture{}, t) }
func TestStrike(t *testing.T)      { gunit.Run(&BowlingFixture{}, t) }
func TestPerfection(t *testing.T)  { gunit.Run(&BowlingFixture{}, t) }

///////////////////////////////////////////////////////////////////////////////

func (this *BowlingFixture) TestGutterBalls() {
	this.rollMany(20, 0)
	this.So(this.game.Score(), should.Equal, 0)
}

func (this *BowlingFixture) TestAllOnes() {
	this.rollMany(20, 1)
	this.So(this.game.Score(), should.Equal, 20)
}

func (this *BowlingFixture) TestSpare() {
	this.rollSpare()
	this.game.Roll(3)
	this.game.Roll(1)
	this.rollMany(16, 0)

	this.So(this.game.Score(), should.Equal, 17)
}

func (this *BowlingFixture) TestStrike() {
	this.rollStrike()
	this.game.Roll(3)
	this.game.Roll(4)
	this.rollMany(16, 0)
	this.So(this.game.Score(), should.Equal, 24)
}

func (this *BowlingFixture) TestPerfection() {
	this.rollMany(12, 10)
	this.So(this.game.Score(), should.Equal, 300)
}

///////////////////////////////////////////////////////////////////////////////

type BowlingFixture struct {
	*gunit.Fixture
	game *Game
}

func (this *BowlingFixture) Setup()      { this.game = NewGame() }
func (self *BowlingFixture) rollStrike() { self.game.Roll(10) }
func (self *BowlingFixture) rollSpare()  { self.rollMany(2, 5) }
func (this *BowlingFixture) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		this.game.Roll(pins)
	}
}
