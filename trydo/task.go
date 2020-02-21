package trydo

// DoFn ...
type DoFn func() error

// FallbackFn ...
type FallbackFn func(err error)

// FinalizeFn ...
type FinalizeFn func()

// Task 任务
type Task struct {
	do       DoFn
	fallback FallbackFn
	finalize FinalizeFn
}

// Then 复合任务
func (pre Task) Then(task Task) Task {
	preTaskDone := false
	do := func() error {
		err := pre.do()
		if err != nil {
			return err
		}
		preTaskDone = true
		if task.do != nil {
			return task.do()
		}
		return nil
	}
	fallback := func(err error) {
		if preTaskDone && task.fallback != nil {
			task.fallback(err)
		}
		if pre.fallback != nil {
			pre.fallback(err)
		}
	}
	finalize := func() {
		if preTaskDone && task.finalize != nil {
			task.finalize()
		}
		if pre.finalize != nil {
			pre.finalize()
		}
	}
	return Task{
		do:       do,
		fallback: fallback,
		finalize: finalize,
	}
}

// ThenDo ...
func (pre Task) ThenDo(do DoFn) Task {
	return pre.Then(Task{do: do})
}

// ThenDoWithFallback ...
func (pre Task) ThenDoWithFallback(do DoFn, fallback FallbackFn) Task {
	return pre.Then(Task{
		do:       do,
		fallback: fallback,
	})
}

// OnErr ...
func (pre Task) OnErr(fn FallbackFn) Task {
	fallback := func(err error) {
		if pre.fallback != nil {
			pre.fallback(err)
		}
		fn(err)
	}
	return Task{
		do:       pre.do,
		fallback: fallback,
		finalize: pre.finalize,
	}
}

// Finally ...
func (pre Task) Finally(fn FinalizeFn) Task {
	finalize := func() {
		if pre.finalize != nil {
			pre.finalize()
		}
		fn()
	}
	return Task{
		do:       pre.do,
		fallback: pre.fallback,
		finalize: finalize,
	}
}

// Do ...
func (pre Task) Do() error {
	return Try(pre.do, pre.fallback, pre.finalize)
}

// NewTask 创建Task
func NewTask(do DoFn, fallback FallbackFn, finalize FinalizeFn) Task {
	return Task{
		do:       do,
		fallback: fallback,
		finalize: finalize,
	}
}

// NewDo ...
func NewDo(do DoFn) Task {
	return Task{
		do: do,
	}
}

// NewDoWithFallback ...
func NewDoWithFallback(do DoFn, fallback FallbackFn) Task {
	return Task{
		do:       do,
		fallback: fallback,
	}
}

// TryTask 执行任务
func TryTask(task Task) error {
	return Try(task.do, task.fallback, task.finalize)
}

// Try ...
func Try(do DoFn, fallback FallbackFn, finalize FinalizeFn) error {
	err := do()
	if err != nil && fallback != nil {
		fallback(err)
	}
	if finalize != nil {
		finalize()
	}
	return err
}
