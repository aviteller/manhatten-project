<script lang="ts">
  // export let name: string;
  

  let sites: Array<object> 


  let whichForm: boolean = true;
  let password: string = "";
  let username: string = "";
  let email: string = "";
  let responseText: string = "";

  const greet = () => {
    fetch("http://localhost:8080/greet", { credentials: "include" })
      .then((res) => res.text())
      .then((res) => (responseText = res));
  };
  const getSites = () => {
    fetch("http://localhost:8080/sites", { credentials: "include" })
      .then((res) => res.json())
      .then((res) => {

      });
  };

  const logout = () =>
    fetch("http://localhost:8080/logout", { credentials: "include" }).then(
      (res) => (responseText = "logged out")
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
      .then((res) => res.text())
      .then((res) => (responseText = res));
  };

  const toggleForm = () => {
    whichForm = !whichForm;
  };

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
</script>

<main>
  <button on:click={greet}>click to greet</button>
  <button on:click={logout}>click to logout</button>
  <button on:click={getSites}>get sites</button>
  <button on:click={toggleForm}>toggle</button>

  <p>{responseText}</p>

  {#if sites.length > 0}
    {#each sites as site}
  
    {/each}
  {/if}

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
    <form on:submit={login}>
      <h3>login</h3>
      <input
        type="text"
        name="username"
        placeholder="enter username"
        bind:value={username}
      />
      <input
        type="text"
        name="password"
        placeholder="enter password"
        bind:value={password}
      />
      <input type="submit" value="login" />
    </form>
  {/if}
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  /* h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  } */

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
