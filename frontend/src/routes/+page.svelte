<script lang="ts">
	import Calculator from "./Calculator.svelte";
	import { calculatorStore } from "./store.svelte";
</script>

<div class="wrapper">
	<div class="screen">
		<p>{calculatorStore.lastResult}</p>
		<p>{calculatorStore.number}</p>
	</div>
	<Calculator />
</div>

<style>
	:global(*) {
		box-sizing: border-box;
	}

	.wrapper {
		display: flex;
		flex-direction: column;
		gap: 1rem;
		height: 100%;
	}

	.screen {
		height: 50%;
		border: 1px solid #ccc;
		border-radius: 0.4rem;
		padding: 1rem;
		display: flex;
		align-items: right;
		justify-content: flex-end;
		flex-direction: column;

		& p {
			margin: 0;
			text-align: right;
			padding: 0.25rem;

			&:first-child {
				font-size: 1.2rem;
				font-weight: 500;
				color: #666;
			}

			&:not(:first-child) {
				font-size: 2rem;
				font-weight: 800;
			}
		}
	}
</style>
