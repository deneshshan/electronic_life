package engine_factories

type RabbitFactory struct {
	EntityFactory
}

func (rf *RabbitFactory) Create() *Entity {
	//fsm := fsm.NewFSM(
	//"scan_for_food",
	//fsm.Events{
	//{Name: "scan_for_food", src:[]{"move_random", "reproduce"}, Dst: []string{"move_random", "move_to_food"}},
	//{Name: "move_random", src[]{"scan_for_food"}, Dst: "scan_for_food"},
	//{Name: "move_to_food", src[]{"scan_for_food"}, Dst: "feed"},
	//{Name: "feed", src[]{"move_to_food"}, Dst: "reproduce"},
	//{Name: "reproduce", src[]{"feed"}, Dst: "scan_for_food"},
	//},
	//)

	//tick := make(chan struct{}, 1)

	return &Rabbit{Tick: tick}
}
