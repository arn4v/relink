<script lang="ts">
  import { store } from "../store";
  import { Link, navigate } from "svelte-routing";
  import * as uuid from "uuid";

  let state: Contact = {
    id: uuid.v4(),
    name: "",
    location: "",
    now: "",
    background: "",
    meetings: {},
  };

  const onSubmit = (e: Event) => {
    e.preventDefault();
    store.update((value) => {
      if (!value.contacts) {
        value.contacts = {};
      }

      value.contacts[state.id] = state;

      return value;
    });
    navigate("/");
  };
</script>

<template>
  <div class="h-screen w-screen bg-white">
    <div class="flex flex-col mx-auto w-1/2 py-8 h-full space-y-8">
      <div class="flex items-center justify-between space-x-8 h-10">
        <Link to="/">
          <h1 class="text-2xl uppercase font-bold font-sans">relink</h1>
        </Link>
      </div>
      <div class="flex items-center justify-start space-x-8 w-full">
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
        <h2 class="text-xl font-bold">Create new contact</h2>
      </div>
      <form class="flex flex-col space-y-8" on:submit={onSubmit}>
        <div class="flex flex-col space-y-4">
          <label for="name" class="font-medium">Name</label>
          <input
            id="name"
            bind:value={state.name}
            class="px-4 w-full focus:outline-none focus:ring-2 ring-black transition duration-100 h-10 shadow"
          />
        </div>
        <div class="flex flex-col space-y-4">
          <label for="location" class="font-medium">Location</label>
          <input
            id="location"
            bind:value={state.location}
            class="px-4 w-full focus:outline-none focus:ring-2 ring-black transition duration-100 h-10 shadow"
          />
        </div>
        <div class="flex flex-col space-y-4">
          <label for="background" class="font-medium">Background</label>
          <textarea
            id="background"
            bind:value={state.background}
            rows={3}
            class="px-4 w-full focus:outline-none focus:ring-2 ring-black transition duration-100 h-auto shadow"
          />
        </div>
        <div class="flex flex-col space-y-4">
          <label for="now" class="font-medium">Now</label>
          <textarea
            id="now"
            bind:value={state.now}
            rows={3}
            class="px-4 w-full focus:outline-none focus:ring-2 ring-black transition duration-100 h-auto shadow"
          />
        </div>
        <button
          class="px-4 bg-black hover:bg-gray-600 text-white font-medium whitespace-nowrap transition ml-auto h-10"
        >
          Submit
        </button>
      </form>
    </div>
  </div>
</template>
