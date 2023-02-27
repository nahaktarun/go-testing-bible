package calculator_test

import "testing"
import "github.com/nahaktarun/go-testing-bible/calculator"

type TestCase struct {
	value int
	expected bool
	actual bool
}

func TestCalulateIsArmstrong(t *testing.T){
	
	testcase := TestCase {
		value : 371,
		expected : true,
	}

	testcase.actual = calculator.CalculateIsArmstrong(testcase.value)

	if testcase.actual != testcase.expected{
		t.Fail()
	}
}

func TestNegativeCalculatorIsArmstrong(t *testing.T){

	testcase := TestCase {
		value : 350,
		expected : false,
	}

	testcase.actual = calculator.CalculateIsArmstrong(testcase.value)

	if testcase.actual != testcase.expected{
		t.Fail()
	}
}