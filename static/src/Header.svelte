<script>
  import { pageTitle } from "./store";
  // export let title;
  export let loggedIn;
  import { createEventDispatcher, getContext } from "svelte";

  const dispatch = createEventDispatcher();

  const logout = () => dispatch("logOut");
</script>

<header>
  <h1>
    {$pageTitle}
    {#if loggedIn}
      <a href style="height:5px;width:5px" on:click={logout}>X</a>
    {/if}
  </h1>
</header>

<style>
  header {
    text-align: center;
    text-transform: uppercase;
    max-height: 35px;
  }
</style>
