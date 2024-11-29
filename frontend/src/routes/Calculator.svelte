<script lang="ts">
	import { type EventHandler } from "svelte/elements";
	import { Calculator } from "$lib/wailsjs/learn-wails/services/calculator";
	import { calculatorStore } from "./store.svelte";

	const log = (x: string) => () => {
		console.log(x);
	};

	const evaluate = async () => {
		calculatorStore.lastResult = await Calculator.Evaluate();
		calculatorStore.number = "0";
	};

	const handleOperatorInput: EventHandler<Event, HTMLButtonElement> = async (ev) => {
		const input = ev.currentTarget.dataset["key"];
		if (input == null) return;

		// Instead of multiple IPC calls, it's better to create a backend API that can take more inputs at once.
		await Calculator.Insert(calculatorStore.number);
		await Calculator.Insert(input);
		await evaluate();
	};

	const handleNumberInput: EventHandler<Event, HTMLButtonElement> = (ev) => {
		const input = ev.currentTarget.dataset["key"];
		if (input == null) return;

		if (calculatorStore.number === "0") {
			calculatorStore.number = input
		} else {
			calculatorStore.number += input;
		}
	};
	const resetHandler: EventHandler<Event, HTMLButtonElement> = async (ev) => {
		await Calculator.Reset();
		calculatorStore.reset();
	};

	const buttonListenerMap = {
		fib: [log("fib"), "a"],
		"√x": [log("square root"), "b"],
		"^2": [log("sqaure"), "c"],
		"/": [handleOperatorInput, "d"],
		AC: [resetHandler, "e"],
		"7": [handleNumberInput, "f"],
		"8": [handleNumberInput, "g"],
		"9": [handleNumberInput, "h"],
		"×": [handleOperatorInput, "i"],
		C: [log("Clear"), "j"],
		"4": [handleNumberInput, "k"],
		"5": [handleNumberInput, "l"],
		"6": [handleNumberInput, "m"],
		"-": [handleOperatorInput, "n"],
		exp: [log("Expression"), "o"],
		"1": [handleNumberInput, "p"],
		"2": [handleNumberInput, "q"],
		"3": [handleNumberInput, "r"],
		"+": [handleOperatorInput, "s"],
		"𝑥": [log("Variable `𝑥`"), "t"],
		"0": [handleNumberInput, "u"],
		".": [handleNumberInput, "v"],
		"=": [
			async () => {
				await Calculator.Insert(calculatorStore.number);
				await evaluate();
				calculatorStore.number = calculatorStore.lastResult.toString();
				await Calculator.Reset();
			},
			"w"
		],
		"𝑦": [log("Variable `𝑦`"), "x"]
	} as Record<string, [(ev?: any) => void, string]>;
</script>

<div class="grid">
	{#each Object.entries(buttonListenerMap) as [k, v] (k)}
		<button onclick={v[0]} style:--area={v[1]} data-key={k}>{k}</button>
	{/each}
</div>

<style>
	.grid {
		display: grid;
		grid-template-areas:
			"a b c d e"
			"f g h i j"
			"k l m n o"
			"p q r s t"
			"u v w s x";
		gap: 0.5rem;
		height: 100%;
	}

	button {
		all: unset;
		font-size: 1.25rem;
		grid-area: var(--area);
		font-weight: bold;
		border: 1px solid #ccc;
		border-radius: 0.25rem;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: #efefef;
		cursor: pointer;

		&:hover, &:focus {
			transition: background-color 300ms cubic-bezier(0.445, 0.05, 0.55, 0.95);
			background-color: #ddd;
		}
	}
</style>
