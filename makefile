
test-linked-list:
	- go test linked_list_test.go -v -count=1

test-doubly-linked-list:
	- go test double_linked_list_test.go -v -count=1

test-stack:
	- go test stack_test.go -v -count=1

test-queue:
	- go test queue_test.go -v -count=1

test-tree:
	- go test binary_search_tree_test.go -v -count=1

test-hashtable:
	- go test hashtable_test.go -v -count=1

test-graph:
	- go test graph_test.go -v -count=1

test-heap:
	- go test heap_test.go -v -count=1