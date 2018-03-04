package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/* fn is a function type with the signature of the mathematical op functions. */
type fn func(float64, float64) float64

func main() {
	mux := http.NewServeMux()

	// Define route handlers
	mux.HandleFunc("/", routeToOp)
	mux.HandleFunc("/man", manPage)

	// Fire it up
	http.ListenAndServe(":8000", mux)
}

/* routeToOp takes an http request and calls the appropriate mathematical op function with the operands. */
func routeToOp(w http.ResponseWriter, r *http.Request) {
	op, operands, err := getOperands(r)
	t := 0.0
	if err == nil {
		switch op {
		case "add":
			t = applyOperator(opAdd, operands)
		case "sub":
			t = applyOperator(opSubtract, operands)
		case "mul":
			t = applyOperator(opMultiply, operands)
		case "div":
			t = applyOperator(opDivide, operands)
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		fmt.Fprintf(w, "%f", t)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

}

/* getOperands takes an http request and returns the math operator and a list of operands. */
func getOperands(r *http.Request) (op string, operands []float64, err error) {
	expr := strings.Split(r.URL.Path, "/")
	op = expr[1]
	expr = expr[2:len(expr)]
	for _, e := range expr {
		v, err := strconv.ParseFloat(strings.TrimSpace(e), 64)
		if err == nil {
			operands = append(operands, v)
		} else {
			return "", nil, errors.New("Invalid operand")
		}
	}
	return op, operands, nil
}

/* applyOperator takes an operator function, applies it to a list of operands. and returns the result. */
func applyOperator(opFn fn, operands []float64) float64 {
	total := 0.0
	for i, e := range operands {
		if i == 0 {
			total = e
		} else {
			total = opFn(total, e)
		}
	}
	return total
}

/* opAdd is the + operator. */
func opAdd(x float64, y float64) float64 {
	return x + y
}

/* opSubtract is the - operator. */
func opSubtract(x float64, y float64) float64 {
	return x - y
}

/* opDivide is the / operator. */
func opDivide(x float64, y float64) float64 {
	return x / y
}

/* opMultiply is the * operator. */
func opMultiply(x float64, y float64) float64 {
	return x * y
}

/* manPage prints usage info to an http ResponseWriter. */
func manPage(w http.ResponseWriter, r *http.Request) {

	mp := `
This API is called like so: Example.
	`

	fmt.Fprintf(w, "%v", mp)
}
