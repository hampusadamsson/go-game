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
