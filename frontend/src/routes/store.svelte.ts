class CalculatorStore {
	number = $state("0");
	lastResult = $state(0);

	reset() {
		this.number = "0";
		this.lastResult = 0;
	}
}

export const calculatorStore = new CalculatorStore();
