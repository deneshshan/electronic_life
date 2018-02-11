package engine_entities

type Rabbit struct {
	Entity
	foodType EntityType
}

func (r *Rabbit) Act() {
	//0. FSM: check to reproduce -> reproducing -> scan for food
	//1. FSM: scan for food. callback: food spotted move to food
	//2. FSM: move towards food
	//3. FSM: near food. callback feed
	//4. FSM: feeding -> continue feeding

	for _ := range r.tick {
	}
}
