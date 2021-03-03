<script>
  import { onMount } from "svelte";
  import { pageTitle, subTitle, listTitle } from "./store";
  import Site from "./Site.svelte";

  export let sites;
  // let listTitle = "SITES LIST";
  let view = "sites";
  let selectedSite = {};

  const updateTitles = (subtitle) => {
    subTitle.set(subtitle);
    pageTitle.set("SITE PAGE");
    listTitle.set("CHANNELS LIST");
  };

  const goToSite = (e) => {
    view = "channels";
    selectedSite = sites.find((s) => s.id == e.target.id);
    updateTitles(e.target.getAttribute("name"));
  };

  onMount(() => {
    if (sites.length == 1) {
      view = "channels";
      selectedSite = sites[0];
      updateTitles(sites[0].name);
    }
  });
</script>

<div class="outer">
  <h4>{$listTitle}</h4>
  {#if view == "sites"}
    <div class="sites">
      {#if sites.length > 0}
        {#each sites as site}
          <div
            class="site"
            id={`${site.id}`}
            name={`${site.name}`}
            on:click={goToSite}
          >
            <span>{site.name}</span><img
              src="./img/placeholder.jpg"
              alt=""
              width="50"
              height="50"
            />
          </div>
        {/each}
      {:else}
        No Sites found
      {/if}
    </div>
  {:else}
    <Site site={selectedSite} />
  {/if}
</div>

<style>
  .outer {
    border-top: 1px solid grey;
    text-align: center;
    height: 380px;
    position: relative;
    border-left: 1px solid grey;
    border-right: 1px solid grey;
    border-image: linear-gradient(
      to bottom,
      grey 50%,
      rgba(0, 0, 0, 0) 50%
    ); /* to top - at 50% transparent */
    border-image-slice: 1;
  }

  /* .outer:after {
    background: grey;
    position: absolute;
    bottom: 0;
    left: 0;
    height: 50%;
    width: 1px;
  } */
  .sites {
    border: 1px solid grey;
    height: 320px;
    overflow-y: scroll;
    scrollbar-width: none;
  }

  .site {
    background-color: grey;
    border: 1px solid black;
    height: 70px;
    display: flex;
    justify-content: space-evenly;
    align-content: center;
    flex-direction: row;
  }

  .site > img {
    vertical-align: middle;
  }

  .site:nth-child(odd) {
    background-color: green;
  }

  .sites::-webkit-scrollbar {
    display: none;
  }

  /* .site >span {
    vertical-align: middle;
    display: inline-block;
  } */
</style>
