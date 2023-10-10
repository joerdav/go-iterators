package main

import "fmt"

type Iter[I any] func(func(yield I) bool)

func main() {
	list := []int{1, 2, 3, 4, 5}

	fmt.Println("Normal iteration")
	for _, item := range list {
		fmt.Println(item)
	}

	fmt.Println("Func iteration")
	for item := range ToIter(list) {
		fmt.Println(item)
	}

	fmt.Println("Filter iteration")
	even := func(i int) bool {
		return i%2 == 0
	}
	for item := range Filter(ToIter(list), even) {
		fmt.Println(item)
	}

	fmt.Println("Map iteration")
	toString := func(i int) string {
		return fmt.Sprintf("num - %d", i)
	}
	for item := range Map(ToIter(list), toString) {
		fmt.Println(item)
	}

	fmt.Println("Find in iteration")
	if res, ok := Find(ToIter(list), even); ok {
		fmt.Println(res)
	}
}

func ToIter[S ~[]I, I any](s S) Iter[I] {
	return func(yield func(I) bool) {
		for _, item := range s {
			yield(item)
		}
	}
}

func Filter[I any](iter Iter[I], check func(I) bool) Iter[I] {
	return func(yield func(I) bool) {
		iter(func(item I) bool {
			if check(item) {
				return yield(item)
			}
			return true
		})
	}
}

func Map[I any, T any](iter Iter[I], m func(I) T) Iter[T] {
	return func(yield func(T) bool) {
		iter(func(item I) bool {
			return yield(m(item))
		})
	}
}

func Find[I any](iter Iter[I], check func(I) bool) (result I, ok bool) {
	for item := range iter {
		if check(item) {
			return item, true
		}
	}
	return result, false
}
