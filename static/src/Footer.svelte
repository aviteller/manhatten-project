<script>
   import { createEventDispatcher } from "svelte";

const dispatch = createEventDispatcher();

const goHome = () => dispatch("gohome")
const getHelp = () => dispatch("gethelp")
</script>


<footer>
  <button id="home" on:click={goHome}>HOME</button>
  <button id="help" on:click={getHelp}>HELP</button>
</footer>

<style>
  footer {
    /* background-color: orange; */
    position: absolute;
    left: 0;
    bottom: -29px;
    /* height: 100px; */
    width: 100%;
    overflow: hidden;
    display: grid;
    grid-template-columns: auto auto;
  }

  button {
    height: 80px;
  }

  #home {
    background-color: honeydew;
  }

  #help {
    background-color: goldenrod;
  }
</style>
