<script lang="ts">
  import { onMount } from "svelte";
  import { Router, Route } from "svelte-routing";
  import { API_BASE_URL } from "./constants";
  import Contact from "./pages/Contact.svelte";
  import EditContact from "./pages/EditContact.svelte";
  import Home from "./pages/Home.svelte";
  import NewContact from "./pages/NewContact.svelte";
  import { store } from "./store";

  onMount(() => {
    fetch(API_BASE_URL + "/data")
      .then((data) => data.json())
      .then((apiData) => {
        store.update(() => apiData);

        store.subscribe((value) => {
          fetch(API_BASE_URL + "/data", {
            method: "POST",
            body: JSON.stringify(value),
          });
        });
      })
      .catch(() => {
        fetch(API_BASE_URL + "/data", {
          method: "POST",
          body: JSON.stringify({}),
        });
      });

    return () => {
      fetch(API_BASE_URL + "/data", {
        method: "POST",
        body: JSON.stringify(store),
      }).catch(console.error);
    };
  });
</script>

<Router>
  <Route path="/"><Home /></Route>
  <Route path="contact/:id" let:params><Contact id={params.id} /></Route>
  <Route path="edit/:id" let:params><EditContact id={params.id} /></Route>
  <Route path="new"><NewContact /></Route>
</Router>
