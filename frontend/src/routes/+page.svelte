<script lang="ts">
	import { GreetService } from "$lib/wailsjs/learn-wails";
	import { Events } from "@wailsio/runtime";

	let name$: HTMLInputElement;
	let time$: HTMLElement;
	let result$: HTMLElement;

	function handleGreet() {
		let name = name$.value;
		if (!name) {
			name = "anonymous";
		}
		GreetService.Greet(name)
			.then((result) => {
				result$.innerText = result;
			})
			.catch((err) => {
				console.log(err);
			});
	}

	Events.On("time", (time: any) => {
		time$.innerText = time.data;
	});
</script>

<div class="container">
	<div>
		<a href="https://wails.io">
			<img src="/wails.png" class="logo" alt="Wails logo" />
		</a>
		<a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript">
			<img src="/javascript.svg" class="logo vanilla" alt="JavaScript logo" />
		</a>
	</div>
	<h1>Wails + Javascript</h1>
	<div class="card">
		<div bind:this={result$} class="result" id="result">Please enter your name below 👇</div>
		<div class="input-box" id="input">
			<input bind:this={name$} class="input" id="name" type="text" autocomplete="off" />
			<button class="btn" onclick={handleGreet}>Greet</button>
		</div>
	</div>
	<div class="footer">
		<div><p>Click on the Wails logo to learn more</p></div>
		<div><p bind:this={time$} id="time">Listening for Time event...</p></div>
	</div>
</div>
