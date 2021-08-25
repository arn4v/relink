<script lang="ts">
  import Logo from "../components/Logo.svelte";
  import { Link, navigate } from "svelte-routing";
  import { store } from "../store";

  let searchValue = "";

  let contacts: State["contacts"];
  store.subscribe((val) => (contacts = val?.contacts ?? {}));

  $: filtered = Object.values(contacts).filter((value) =>
    value.name.includes(searchValue)
  );
</script>

<template>
  <div class="h-screen w-screen bg-white">
    <div class="flex flex-col mx-auto w-1/2 py-8 h-full space-y-12">
      <div class="flex items-center justify-between space-x-12 h-10">
        <Logo />
        <input
          placeholder="search for contacts..."
          class="shadow px-4 h-full w-full input"
          bind:value={searchValue}
        />
        <button
          class="px-4 h-full whitespace-nowrap transition btn-primary"
          on:click={() => navigate("/new")}
        >
          New contact
        </button>
      </div>
      {#if filtered.length > 0}
        <div class="grid grid-cols-3 gap-8">
          {#each filtered as contact}
            <Link to={`/contact/${contact.id}`}>
              <div
                class="px-4 py-2 bg-gray-50 shadow font-bold hover:bg-gray-200 transition"
              >
                {contact.name}
              </div>
            </Link>
          {/each}
        </div>
      {:else}
        <div
          class="font-medium text-2xl w-full grid place-items-center bg-gray-100 py-8 px-4"
        >
          No contacts found
        </div>
      {/if}
    </div>
  </div>
</template>
