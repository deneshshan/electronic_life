package engine_entities

type Carrot struct {
	Entity
}

func (c *Carrot) Act() {
	//1. FSM: grow
	//2. FSM: reproduce

	for {
		select {
		case <-tick:
			HandleAction()
		case <-done:
			Destroy()
		}
	}
}

func (c Carrot) HandleAction() {
	switch fsm.Current() {
	case "grow":
		c.healthUpdate <- 10
	case "reproduce":
		c.NextAction <- Action{Type: Spawn, CurrentPosition: c.Position}
	case "reproduced":
		c.healthUpdate <- -50
	}
}

func (c *Carrot) updateHealth() {
	for change := range c.healthUpdate {
		health = health + change

		if health > 100 {
			c.fsm.Event("reproduce")
		}
	}
}

func (c *Carrot) Destroy() {
	close(c.tick)
	close(c.done)
}
