package main

import (
	"fmt"
	"log"
	"strconv"

	"example.arainay.com/lib"
)

func StringToInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return intValue
}

func main() {
	if true {
		list := []string{"24:00:00-23:59:59"}
		fmt.Println(lib.ValidateTimePeriods(list))
	} else {
		lib.Example()

		var costs = []int{10, 20, 30, 40, 10, 20, 10, 30, 20, 10}
		fmt.Println(lib.GetFullCost(costs))

		fmt.Println(lib.Register("_1SdAefaasddasdsdfasdas"))
		fmt.Println(lib.Register("A_1SdAefaasddasdsdfasdass"))

		a := [2]string{"John", "8-950-111-22-33"}
		b := [2]string{"Jack", "8-904-000-99-88"}
		c := [2]string{"Jim", "8-902-555-66-77"}
		d := [2]string{"John", "8-950-111-33-55"}
		e := [2]string{"Jack", "8-904-000-88-66"}
		f := [2]string{"Jim", "8-902-555-77-99"}
		g := [2]string{"John", "8-950-123-45-67"}
		h := [2]string{"Lenny", "8-904-000-99-88"}
		i := [2]string{"Jim", "8-902-123-45-67"}
		j := [2]string{"John", "8-950-987-65-43"}
		k := [2]string{"Jack", "8-904-000-99-88"}
		l := [2]string{"Jim", "8-902-555-66-77"}
		m := [2]string{"John", "8-950-192-83-74"}
		n := [2]string{"Jack", "8-904-000-99-88"}
		o := [2]string{"Jim", "8-902-555-66-77"}
		p := [2]string{"John", "8-950-333-22-11"}
		q := [2]string{"Jack", "8-904-000-99-88"}
		r := [2]string{"Jim", "8-902-555-66-77"}
		s := [2]string{"John", "8-950-111-22-33"}
		t := [2]string{"Jack", "8-904-000-99-88"}
		u := [2]string{"Jim", "8-902-555-66-77"}
		v := [2]string{"John", "8-950-111-22-33"}
		w := [2]string{"Jack", "8-904-000-99-88"}
		x := [2]string{"Jim", "8-902-555-66-77"}
		y := [2]string{"John", "8-950-111-22-33"}
		z := [2]string{"Jack", "8-904-000-99-88"}
		aa := [2]string{"Jim", "8-902-555-66-77"}
		ab := [2]string{"John", "8-950-111-22-33"}
		ac := [2]string{"Jack", "8-904-000-99-88"}
		ad := [2]string{"Jim", "8-902-555-66-77"}

		contacts := [][2]string{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, aa, ab, ac, ad}
		fmt.Println(lib.GetLastCallsForEveryCallMaker(contacts))

		lib.Check()

		lib.Sum()
	}
}
