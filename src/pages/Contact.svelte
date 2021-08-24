<script lang="ts">
  import { store } from "../store";
  import { Link, navigate } from "svelte-routing";
  import Spinner from "../components/Spinner.svelte";
  import Logo from "../components/Logo.svelte";
  import { onDestroy, onMount } from "svelte";
  import { get } from "svelte/store";

  export let id: string;

  let contact: Contact;

  onMount(() => {
    contact = get(store).contacts[id];
  });

  onDestroy(() => {
    contact = undefined;
  });
</script>

<div class="h-screen w-screen bg-white">
  <div class="flex flex-col mx-auto w-1/2 py-8 h-full space-y-12">
    <div class="flex items-center justify-between space-x-12 h-10">
      <Logo />
      <Link
        to={`/edit/${id}`}
        class="px-4 h-full whitespace-nowrap transition button-anchor grid place-items-center"
      >
        Edit contact
      </Link>
    </div>
    {#if !!contact}
      <template>
        <div class="flex items-center justify-between w-full">
          <div class="flex items-center justify-center space-x-8">
            <Link to="/">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M10 19l-7-7m0 0l7-7m-7 7h18"
                />
              </svg>
            </Link>
            <h2 class="text-xl font-bold">{contact.name}</h2>
          </div>
          <Link to={`/edit/${id}`}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"
              />
            </svg>
          </Link>
        </div>
        <div class="flex flex-col w-full">
          <dl class="flex flex-col rounded shadow py-4">
            <div class="grid grid-cols-3 divide-x divide-gray-400">
              <dt class="px-8 place-self-end">Name</dt>
              <dl class="px-8">{contact?.name}</dl>
            </div>
          </dl>
        </div>
      </template>
    {:else}
      <Spinner />
    {/if}
  </div>
</div>
