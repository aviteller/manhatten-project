<script>
  import { onMount } from "svelte";
  import { currentPlaylist } from "./store";

  let currentTrack;
  let songValue = 0;

  const formatToMin = (time) => {
    let seconds = time % 60;
    let min = Math.floor(time / 60);

    return `${min.toString().length == 1 ? `0${min}` : min}:${
      seconds.toString().length == 1 ? `0${seconds}` : seconds
    }`;
  };

  const getRandomInt = (min, max) => {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
  };

  const simulateTrack = async (fromBegin = false) => {
    currentTrack = await $currentPlaylist.tracks[
      Math.floor(Math.random() * $currentPlaylist.tracks.length)
    ];
    currentTrack.songLength = getRandomInt(120, 260);
    currentTrack.songLengthFormatted = formatToMin(currentTrack.songLength);
    if (!fromBegin) {
      currentTrack.randStart = getRandomInt(0, currentTrack.songLength);
    } else {
      currentTrack.randStart = 0;
    }
    // console.log(currentTrack)

    songValue = currentTrack.randStart;
    let interval = await setInterval(() => {
      songValue++;
      if (songValue == currentTrack.songLength) {
        simulateTrack(true);
        clearInterval(interval);
        return;
      }
    }, 1000);
  };

  onMount(async () => {
    if ($currentPlaylist) {
      await simulateTrack();
    }
  });
</script>

{#if $currentPlaylist}
  <div id="playing">
    <span>{$currentPlaylist.name}</span>
    {#if currentTrack != null}
      <!-- <span>{currentTrack.name}</span> -->
      <label for="track"
        >{currentTrack.name}
        {formatToMin(songValue)}/{currentTrack.songLengthFormatted}</label
      >

      <progress id="track" value={songValue} max={currentTrack.songLength} />
    {/if}
  </div>
{/if}

<style>
  #playing {
    height: 60px;
    width: 100%;
    border-top: 1px solid grey;
    border-bottom: 1px solid grey;
    display: flex;
    flex-direction: column;
    text-align: center;
    /* justify-content: space-evenly; */
  }

  progress {
    width: 100%;
  }
</style>
