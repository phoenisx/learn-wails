<script lang="ts">
	import Calculator from "./Calculator.svelte";
</script>

<div>
	<textarea readonly  value="Hello World"></textarea>
	<Calculator />
</div>

<style>
	:global(*) {
		box-sizing: border-box;
	}

	div {
		display: flex;
		flex-direction: column;
		gap: 1rem;
		height: 100%;
	}

	textarea {
		height: 50%;
		resize: none;
	}
</style>
