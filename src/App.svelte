<script lang="ts">
  import { onMount } from "svelte";

  import { Router, Route } from "svelte-routing";
  import { API_BASE_URL } from "./constants";
  import Contact from "./pages/Contact.svelte";
  import Home from "./pages/Home.svelte";
  import { data } from "./store";

  onMount(() => {
    fetch(API_BASE_URL + "/data")
      .then((data) => data.json())
      .then((apiData) => {
        console.log(apiData);
        data.set(apiData);
      })
      .catch(console.error);

    return () => {
      fetch(API_BASE_URL + "/data", {
        method: "POST",
        body: JSON.stringify(data),
      }).catch(console.error);
    };
  });
</script>

<Router>
  <Route path="/"><Home /></Route>
  <Route path="contact/:id"><Contact /></Route>
</Router>
