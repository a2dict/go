package tryutil

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

// NewTask 创建Task
func NewTask(do DoFn, fallback FallbackFn, finalize FinalizeFn) Task {
	return Task{
		do:       do,
		fallback: fallback,
		finalize: finalize,
	}
}

// TryTask 执行任务
func TryTask(task Task) error {
	return Try(task.do, task.fallback, task.finalize)
}

// TryTasks 链式执行任务
func TryTasks(tasks ...Task) error {
	tasksLen := len(tasks)
	if tasksLen == 0 {
		return nil
	}
	if tasksLen == 1 {
		return TryTask(tasks[0])
	}

	t := tasks[0]
	for i := 1; i < tasksLen; i++ {
		t = t.Then(tasks[i])
	}
	return TryTask(t)
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
