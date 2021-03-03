<script>
  import Channel from "./Channel.svelte";
  import { pageTitle, subTitle, listTitle } from "./store";
  export let site;

  let view = "site";
  let selectedChannel = {};

  const goToChannel = (e) => {
    view = "channel";
    // console.log(site.channels[0].id)
    selectedChannel = site.channels.find((c) => c.id == e.target.id);
    pageTitle.set(selectedChannel.name);
    listTitle.set("PLAYLIST LIST");
  };
</script>

{#if view == "site"}
  <div class="channels">
    {#if site.channels != null}
      {#each site.channels as channel}
        <div class="channel" on:click={goToChannel} id={channel.id}>
          {channel.name}
        </div>
      {/each}
    {:else}
      no channels
    {/if}
  </div>
{:else}
  <Channel channel={selectedChannel} />
{/if}

<style>
  .channels {
    border: 1px solid grey;
    height: 320px;
    overflow-y: scroll;
    scrollbar-width: none;
  }

  .channels::-webkit-scrollbar {
    display: none;
  }

  .channel {
    background-color: grey;
    border: 1px solid black;
    height: 70px;
    display: flex;
    justify-content: space-evenly;
    align-content: center;
    flex-direction: row;
  }

  .channel > img {
    vertical-align: middle;
  }

  .channel:nth-child(odd) {
    background-color: green;
  }
</style>
