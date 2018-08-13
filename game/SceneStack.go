package game

type Node struct {
	value GameScene
	next  *Node
}

func newStack() *Stack {
	return &Stack{}
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(item GameScene) {
	n := &Node{value: item}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}

	s.size++
}

func (s *Stack) Pop() GameScene {
	var n *Node
	if s.head != nil {
		n = s.head
		s.head = n.next
		s.size--
	}

	if n == nil {
		return nil
	}

	return n.value

}

func (s *Stack) Peek() GameScene {
	n := s.head
	if n == nil || n.value == nil {
		return nil
	}

	return n.value
}

type SceneStack struct {
	scenes *Stack
}

func NewSceneStack() *SceneStack {
	return &SceneStack{
		scenes: newStack(),
	}
}

func (s *SceneStack) PushScene(scene GameScene, payload interface{}) {
	s.scenes.Push(scene)
	scene.OnStart(payload)
}

func (s *SceneStack) Update() {
	s.scenes.Peek().Update()
}

func (s *SceneStack) Draw() {
	s.scenes.Peek().Draw()
}
