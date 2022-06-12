package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOwner(t *testing.T) {
	p1 := Player{}
	u1 := Unit{Owner: &p1}
	u2 := Unit{Owner: &p1}
	u3 := Unit{Owner: &Player{}}
	assert.True(t, u1.sameOwner(&u2))
	assert.True(t, u2.sameOwner(&u1))
	assert.False(t, u1.sameOwner(&u3))
}

func TestCantFight(t *testing.T) {
	p1 := Player{}
	u1 := Unit{Owner: &p1}
	u2 := Unit{Owner: &p1}
	r, err := u1.fight(&u2)
	assert.NotNil(t, err)
	assert.False(t, r)
}

func TestFightDeath(t *testing.T) {
	u1 := Unit{Owner: &Player{}, Damage: 10}
	u2 := Unit{Owner: &Player{}, HP: 10}
	r, err := u1.fight(&u2)
	assert.True(t, r)
	assert.Nil(t, err)
	assert.Equal(t, u2.HP, 0)
}

func TestFightSurvived(t *testing.T) {
	u1 := Unit{Owner: &Player{}, Damage: 9}
	u2 := Unit{Owner: &Player{}, HP: 10}
	r, err := u1.fight(&u2)
	assert.False(t, r)
	assert.Nil(t, err)
	assert.Equal(t, u2.HP, 1)
}

func TestUnit_adjacentAttack1(t *testing.T) {
	u := Unit{x: 0, y: 2, attackRange: 1}
	a := u.getAllAttackCoords()
	assert.Contains(t, a, coord{0, 3})
}

func TestUnit_adjacentAttack(t *testing.T) {
	u := Unit{x: 5, y: 5, attackRange: 1}
	a := u.getAllAttackCoords()
	assert.Contains(t, a, coord{4, 5})
	assert.Contains(t, a, coord{5, 6})
}

func TestUnit_getAttackCoordsOne(t *testing.T) {
	u := Unit{x: 5, y: 5, attackRange: 1}
	a := u.getAllAttackCoords()
	assert.Equal(t, 4, len(a))
}

func TestUnit_getAttackCoordsTwo(t *testing.T) {
	u := Unit{x: 5, y: 5, attackRange: 2}
	a := u.getAllAttackCoords()
	assert.Equal(t, 12, len(a))
	assert.Contains(t, a, coord{6, 6})
	assert.NotContains(t, a, coord{7, 6})
}

func TestUnit_getAttackCoordsThree(t *testing.T) {
	u := Unit{x: 5, y: 5, attackRange: 3}
	a := u.getAllAttackCoords()
	assert.Equal(t, 24, len(a))
	assert.Contains(t, a, coord{6, 6})
	assert.Contains(t, a, coord{7, 6})
}
