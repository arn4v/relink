<script lang="ts">
  import { fly } from "svelte/transition";
  import ModalOverlay from "../components/ModalOverlay.svelte";
  import { onDestroy, onMount } from "svelte";
  import { Link, navigate } from "svelte-routing";
  import Logo from "../components/Logo.svelte";
  import Spinner from "../components/Spinner.svelte";
  import { store } from "../store";

  export let id: string;
  let deleteModal = {
    isOpen: false,
    onDelete: () => {
      store.update((val) => {
        delete val.contacts[id];
        return val;
      });
      navigate("/");
    },
  };

  let contact: Contact;

  onMount(() => {
    const unsubscribe = store.subscribe((value) => {
      if (!contact) {
        contact = value.contacts[id];
      } else {
        unsubscribe();
      }
    });
  });

  onDestroy(() => {
    contact = undefined;
  });
</script>

<div class="h-screen w-screen bg-white relative">
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
          <div class="flex items-center justify-center space-x-6">
            <button
              on:click={() => (deleteModal.isOpen = true)}
              class="p-1.5 hover:bg-gray-100 border border-gray-200 transition"
            >
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
                  d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                />
              </svg>
            </button>
            <Link
              to={`/edit/${id}`}
              class="p-1.5 hover:bg-gray-100 border border-gray-200 transition"
            >
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
        </div>
        <div class="flex flex-col w-full space-y-6">
          <div>
            <dt class="contact-data-title">Name</dt>
            <dd class="contact-data-body">{contact.name}</dd>
          </div>
          <div>
            <dt class="contact-data-title">Location</dt>
            <dd class="contact-data-body">
              {contact.location !== "" ? contact.location : "Not specified"}
            </dd>
          </div>
          <div>
            <dt class="contact-data-title">Background</dt>
            <dd class="contact-data-body">
              {contact.background !== "" ? contact.background : "Not specified"}
            </dd>
          </div>
          <div>
            <dt class="contact-data-title">Now</dt>
            <dd class="contact-data-body">
              {contact.now !== "" ? contact.now : "Not specified"}
            </dd>
          </div>
        </div>
      </template>
    {:else}
      <Spinner />
    {/if}
  </div>
</div>
<ModalOverlay
  isOpen={deleteModal.isOpen}
  className=""
  onClose={() => {
    deleteModal.isOpen = false;
  }}
>
  <div
    class="bg-white h-[15vh] w-[25vw] z-[200] rounded-md relative flex flex-col items-center justify-between p-8"
  >
    <h1 class="text-xl font-bold">
      Do you really want to delete this contact?
    </h1>
    <div class="flex items-center justify-between w-full">
      <button
        class="border border-gray-200 px-4 py-2 hover:bg-gray-100 transition font-semibold"
      >
        Cancel
      </button>
      <button
        class="px-4 py-2 bg-red-500 font-semibold text-white hover:bg-red-700 transition"
        on:click={deleteModal.onDelete}>Confirm</button
      >
    </div>
  </div>
</ModalOverlay>
