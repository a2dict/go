package trydo

import (
	"errors"
	"fmt"
	"testing"
	"time"
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

	err := TryTask(task1.Then(task2).Then(task3).Then(task4))
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

func TestChainTask(t *testing.T) {
	err := NewDo(func() error {
		fmt.Println("task 1")
		return nil
	}).ThenDo(func() error {
		fmt.Println("task 2")
		return nil
	}).ThenDoWithFallback(func() error {
		fmt.Println("task 3")
		return errors.New("task3 err")
	}, func(err error) {
		fmt.Println("task 3 fallback", err)
	}).ThenDo(func() error {
		fmt.Println("task 4")
		return nil
	}).OnErr(func(err error) {
		fmt.Println("onErr. some err happened.", err)
	}).Finally(func() {
		fmt.Println("finally")
	}).Do()
	fmt.Println(err)
}

func TestTryWithIntervals(t *testing.T) {
	do := func() error {
		now := time.Now()
		fmt.Println("do...", now)
		return errors.New("xx")
	}
	err := TryWithIntervals(do, 3*time.Second, 5*time.Second, 5*time.Second)
	fmt.Println(err)
}

func TestErrNil(t *testing.T) {
	var e error
	fmt.Println(e)
	fmt.Println(e == nil)
}
