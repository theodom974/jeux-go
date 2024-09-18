package fight

import (
	"main/src/entity"
)

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
)

func Fight(player *entity.Player, monster *entity.Monster) {

	player.Attack(monster)
	monster.Attack(player)

	if player.Health <= 0 {
		player.IsAlive = false
		return
	} else if monster.Health <= 0 {
		monster.IsAlive = false
		player.Inventory = append(player.Inventory, monster.Loot...)
		player.Money += monster.Worth
		return
	}

	

}
