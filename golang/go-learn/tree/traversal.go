package tree

func (receiver *Node) PreOrderTraversal() {
	if receiver == nil {
		return
	}
	receiver.Print()
	receiver.Left.PreOrderTraversal()
	receiver.Right.PreOrderTraversal()
}

func (receiver *Node) OrderTraversal() {
	if receiver == nil {
		return
	}
	receiver.Left.OrderTraversal()
	receiver.Print()
	receiver.Right.OrderTraversal()
}

func (receiver *Node) PostOrderTraversal() {
	if receiver == nil {
		return
	}
	receiver.Left.PostOrderTraversal()
	receiver.Right.PostOrderTraversal()
	receiver.Print()
}

func (receiver *Node) TraversalFunctionInput(f func(*Node)) {
	if receiver == nil {
		return
	}
	receiver.Left.TraversalFunctionInput(f)
	receiver.Right.TraversalFunctionInput(f)
	f(receiver)
}
