
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

test-rbst:
	- go test recursive_binary_search_tree_test.go -v -count=1

test-bdfs:
	- go test bfs_dfs_test.go -v -count=1

test-sort:
	- go test sort_test.go -v -count=1

ctest:
	- g++ $(FILE).cc -std=c++20 -g -o $(basename $(FILE)).exe
	- $(FILE).exe
	- rm -f $(FILE).exe