<script lang="ts">
   const log = (x: string) => () => {
      console.log(x);
   }
   const buttonListenerMap = {
      "fib": [log("fib"), 'a'],
      "√x": [log("square root") , 'b'],
      "^2": [log("sqaure"), 'c'],
      "/": [log("divide"), 'd'],
      "AC": [log("Reset"), 'e'],
      "7": [log("Number 7"), 'f'],
      "8": [log("Number 8"), 'g'],
      "9": [log("Number 9"), 'h'],
      "×": [log("Multiply"), 'i'],
      "C": [log("Clear"), 'j'],
      "4": [log("Number 4"), 'k'],
      "5": [log("Number 5"), 'l'],
      "6": [log("Number 6"), 'm'],
      "-": [log("Subtract"), 'n'],
      "exp": [log("Expression"), 'o'],
      "1": [log("Number 1"), 'p'],
      "2": [log("Number 2"), 'q'],
      "3": [log("Number 3"), 'r'],
      "+": [log("Addition"), 's'],
      "𝑥": [log("Variable `𝑥`"), 't'],
      "0": [log("Number 0"), 'u'],
      ".": [log("Decimal"), 'v'],
      "=": [log("Result"), 'w'],
      "𝑦": [log("Variable `𝑦`"), 'x'],
   } as Record<string, [() => void, string]>
</script>

<div class="grid">
   {#each Object.entries(buttonListenerMap) as [k, v] (k)}
      <button onclick={v[0]} style:--area={v[1]}>{k}</button>
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
         "u v w s x"
      ;
      gap: 0.5rem;
      height: 100%;
   }

   button {
      /* width: 2rem;
      height: 1.2rem; */
      font-size: 1.25rem;
      grid-area: var(--area);
      font-weight: bold;
   }
</style>
