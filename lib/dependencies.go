package lib

import "fmt"

type dependency struct {
	name     string
	children []dependency
}

// Есть n модулей, для каждого известны зависимости,
// которые нужны для его установки.
// Есть m запросов — на каждой запрос нужно установить
// какой-то модуль, предварительно поставив все его
// зависимости, игнорируя уже установленные модули.
func InstallDependencies(modules []dependency) {
	stack := modules
	visited := map[string]bool{}

	for len(stack) > 0 {
		v := stack[0]
		stack = append(make([]dependency, 0), stack[1:]...)
		_, isExists := visited[v.name]

		if !isExists {
			fmt.Println("Installing", v.name, "...")
			visited[v.name] = true
			stack = append(v.children, stack...)
		}
	}
}

func Check() {
	a := dependency{
		"a",
		[]dependency{
			{
				"first-level-a_1",
				[]dependency{
					{
						"first-level-a_inner-1",
						make([]dependency, 0),
					},
					{
						"first-level-a_inner-2",
						[]dependency{
							{
								"first-level-a_inner-2_1",
								make([]dependency, 0),
							},
						},
					},
				},
			},
			{
				"first-level-a_2",
				[]dependency{
					{
						"first-level-a_inner-1",
						make([]dependency, 0),
					},
					{
						"first-level-a_inner-3",
						[]dependency{
							{
								"first-level-a_inner-3_1",
								make([]dependency, 0),
							},
						},
					},
				},
			},
		},
	}
	b := dependency{
		"b",
		[]dependency{
			{
				"first-level-b_1",
				[]dependency{
					{
						"first-level-b_inner-1",
						make([]dependency, 0),
					},
					{
						"first-level-b_inner-2",
						[]dependency{
							{
								"first-level-b_inner-2_1",
								make([]dependency, 0),
							},
						},
					},
				},
			},
			{
				"first-level-b_2",
				[]dependency{
					{
						"first-level-b_inner-1",
						make([]dependency, 0),
					},
					{
						"first-level-b_inner-3",
						[]dependency{
							{
								"first-level-b_inner-3_1",
								make([]dependency, 0),
							},
						},
					},
				},
			},
		},
	}

	dependencies := []dependency{a, b}
	InstallDependencies(dependencies)
}
