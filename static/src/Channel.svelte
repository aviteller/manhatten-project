<script>
  import { onMount } from "svelte";
  import { currentPlaylist } from "./store";

  export let channel;
  let view = "channel";
  let selectedPlaylist = {};

  const setCurrentPlaylist = async (id) => {
    await currentPlaylist.set();
    // console.log(id);
    await currentPlaylist.set(channel.playlists.find((c) => c.id == id));
  };

  onMount(() => {
    if (channel.playlists != null) setCurrentPlaylist(channel.playlists[0].id);
  });
</script>

{#if view == "channel"}
  {#if channel.playlists != null}
    <div class="playlists">
      {#each channel.playlists as playlist}
        <div
          class="playlist"
          id={playlist.id}
          on:click={() => setCurrentPlaylist(playlist.id)}
        >
          {playlist.name}
        </div>
      {/each}
    </div>
  {:else}
    no channels
  {/if}
{:else}
  <!-- <Channel channel={selectedChannel} /> -->
{/if}

<style>
  .playlists {
    border: 1px solid grey;
    height: 200px;
    overflow-y: scroll;
    scrollbar-width: none;
  }

  .playlists::-webkit-scrollbar {
    display: none;
  }

  .playlist {
    background-color: grey;
    border: 1px solid black;
    height: 70px;
    display: flex;
    justify-content: space-evenly;
    align-content: center;
    flex-direction: row;
  }

  .playlist > img {
    vertical-align: middle;
  }

  .playlist:nth-child(odd) {
    background-color: green;
  }
</style>
