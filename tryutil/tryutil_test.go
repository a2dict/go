package tryutil

import (
	"errors"
	"fmt"
	"testing"
)

func TestTry(t *testing.T) {
	do := func() error {
		fmt.Println("--\ndo")
		return nil
	}
	doWithErr := func() error {
		fmt.Println("--\ndo with err")
		return errors.New("some err")
	}
	fallback := func(err error) {
		fmt.Println("failback err:", err)
	}
	finalize := func() {
		fmt.Println("finalize")
	}
	err := Try(do, fallback, finalize)
	if err != nil {
		fmt.Println("ret err:", err)
	}

	err = Try(doWithErr, fallback, finalize)
	if err != nil {
		fmt.Println("ret err:", err)
	}
}

func TestTryTasks(t *testing.T) {
	task1 := NewTask(func() error {
		fmt.Println("task1 do")
		return nil
	}, func(err error) {
		fmt.Println("task1 failback err:", err)
	}, func() {
		fmt.Println("task1 finalize")
	})

	task2 := NewTask(func() error {
		fmt.Println("task2 do")
		return nil
	}, func(err error) {
		fmt.Println("task2 failback err:", err)
	}, func() {
		fmt.Println("task2 finalize")
	})

	task3 := NewTask(func() error {
		fmt.Println("task3 do")
		return errors.New("task3 err")
	}, func(err error) {
		fmt.Println("task3 failback err:", err)
	}, func() {
		fmt.Println("task3 finalize")
	})

	task4 := NewTask(func() error {
		fmt.Println("task4 do")
		return nil
	}, func(err error) {
		fmt.Println("task4 failback err:", err)
	}, func() {
		fmt.Println("task4 finalize SHOULD NOT run")
	})

	err := TryTasks(task1, task2, task3, task4)
	fmt.Println(err)
}

func TestTask_Then(t *testing.T) {
	task1 := NewTask(func() error {
		fmt.Println("task1 do")
		return nil
	}, func(err error) {
		fmt.Println("task1 failback err:", err)
	}, func() {
		fmt.Println("task1 finalize")
	})

	task2 := NewTask(func() error {
		fmt.Println("task2 do")
		return nil
	}, func(err error) {
		fmt.Println("task2 failback err:", err)
	}, func() {
		fmt.Println("task2 finalize")
	})
	task3 := NewTask(func() error {
		fmt.Println("task3 do")
		//return errors.New("task3 err")
		return nil
	}, func(err error) {
		fmt.Println("task3 failback err:", err)
	}, func() {
		fmt.Println("task3 finalize")
	})

	tt := task1.Then(task2)
	tt = tt.Then(task3)
	TryTask(tt)
}
