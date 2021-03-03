<script>
  import { pageTitle, subTitle, currentPlaylist } from "./store";
  import Header from "./Header.svelte";
  import Help from "./Help.svelte";
  import Footer from "./Footer.svelte";
  import Sites from "./Sites.svelte";
  import CurrentPlaying from "./CurrentPlaying.svelte";
  import { onMount } from "svelte";
  // export let name: string;
  let view = "sites"
  let sites = [];
  let loading = true;
  let whichForm = true;
  let password = "";
  let username = "";
  let email = "";
  let responseText = "";
  // let pageTitle = "Manhatten Project";
  let subtitle = "login";
  let loggedIn = false;

  const greet = () => {
    fetch("http://localhost:8080/greet", { credentials: "include" })
      .then((res) => res.text())
      .then((res) => (responseText = res));
  };
  const getSites = () => {
    loading = true;
    fetch("http://localhost:8080/sites", { credentials: "include" })
      .then((res) => res.json())
      .then((res) => {
        if (res.data == null) return;
        sites = res.data;
        loading = false;
        // console.log(sites.length);
      });
  };

  const logout = () =>
    fetch("http://localhost:8080/logout", { credentials: "include" }).then(
      (res) => {
        loggedIn = false;
        responseText = "logged out";
      }
    );

  const login = (e) => {
    e.preventDefault();
    let user = {
      username,
      password,
    };

    fetch("http://localhost:8080/login", {
      method: "POST",
      body: JSON.stringify(user),
      credentials: "include",
    })
      .then((res) => res.json())
      .then((res) => {
        if (res.status) {
          pageTitle.set("HOME PAGE");
          // subtitle = "sites";
          getSites();
          loggedIn = true;
          subTitle.set(res.data.username);
        }
      });
  };

  const toggleForm = () => (whichForm = !whichForm);
  const signup = (e) => {
    e.preventDefault();
    let user = {
      username,
      email,
      password,
    };

    fetch("http://localhost:8080/signup", {
      method: "POST",
      body: JSON.stringify(user),
    })
      .then((res) => res.text())
      .then((res) => (responseText = res));
  };

  const updateTitles = (e) => {
    // subtitle = e.detail.subtitle;
    // pageTitle = e.detail.title;
  };

  const checkLoggedOn = () => {
    fetch("http://localhost:8080/logedin", {
      method: "GET",
      credentials: "include",
    })
      .then((res) => res.json())
      .then((res) => {
        if (res.status) {
          pageTitle.set("HOME PAGE");
          view = "sites"
          currentPlaylist.set();
          // subtitle = "sites";
          getSites();
          loggedIn = true;
          subTitle.set(res.data.username);
        }
      });
  };

  const goHome = () => checkLoggedOn();
  const getHelp = () => {
    view = "help"
  };

  onMount(async () => {
    checkLoggedOn();
  });
</script>

<div class="smartphone">
  <div class="content">
    <Header on:logOut={logout} {loggedIn} />

    <main>
      <h3 class="sub-title">{$subTitle}</h3>
      {#if !loggedIn}
        {#if !whichForm}
          <form on:submit={signup}>
            <h3>signup</h3>
            <input
              type="text"
              name="username"
              placeholder="enter username"
              bind:value={username}
            />
            <input
              type="text"
              name="email"
              placeholder="enter email"
              bind:value={email}
            />
            <input
              type="text"
              name="password"
              placeholder="enter password"
              bind:value={password}
            />
            <input type="submit" value="signup" />
          </form>
        {:else}
          <form on:submit={login} class="login-form">
            <div>
              <input
                type="text"
                name="username"
                placeholder="USERNAME"
                bind:value={username}
              />
            </div>
            <div>
              <input
                type="password"
                name="password"
                placeholder="PASSWORD"
                bind:value={password}
              />
            </div>
            <div>
              <input type="submit" value="login" />
            </div>
          </form>
          <a href="#g">FORGOT PASSWORD?</a>
        {/if}
      {:else}
    

        {#if !loading}
          {#if view == "sites"}
          <Sites {sites} on:updatetitles={updateTitles} />
          {:else if view == "help"}
          <Help/>
          {/if}
        {/if}
      {/if}
    </main>
    {#if $currentPlaylist}
    <CurrentPlaying />
    {/if}
    {#if loggedIn}
      <Footer on:gohome={goHome} on:gethelp={getHelp} />
    {/if}
  </div>
</div>

<style>
  .response {
    word-wrap: break-word;
  }

  .sub-title {
    text-transform: uppercase;
    text-align: left;
  }

  .smartphone {
    position: relative;
    width: 360px;
    height: 640px;
    margin: auto;
    border: 16px black solid;
    border-top-width: 60px;
    border-bottom-width: 60px;
    border-radius: 36px;
  }

  /* The horizontal line on the top of the device */
  .smartphone:before {
    content: "";
    display: block;
    width: 60px;
    height: 5px;
    position: absolute;
    top: -30px;
    left: 50%;
    transform: translate(-50%, -50%);
    background: #333;
    border-radius: 10px;
  }

  /* The circle on the bottom of the device */
  .smartphone:after {
    /* content: "";
    display: block;
    width: 35px;
    height: 35px;
    position: absolute;
    left: 50%;
    bottom: -65px;
    transform: translate(-50%, -50%);
    background: #333;
    border-radius: 50%; */
  }

  /* The screen (or content) of the device */
  .smartphone .content {
    width: 360px;
    height: 640px;
    background: white;
  }

  input {
    text-align: center;
    /* text-transform: uppercase; */
    border-radius: 0%;
    width: 100%;
  }

  input::placeholder {
    text-align: center;
  }

  main {
    padding: 0.5em;
    text-align: center;
    /* overflow-y: hidden; */
  }
</style>
