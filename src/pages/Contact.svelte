<script lang="ts">
  import { format } from "date-fns";
  import Modal from "../components/Modal.svelte";
  import { onDestroy, onMount } from "svelte";
  import { Link, navigate } from "svelte-routing";
  import Logo from "../components/Logo.svelte";
  import Spinner from "../components/Spinner.svelte";
  import { store } from "../store";
  import AddMeetingModal from "../components/AddMeetingModal.svelte";
  import type { Unsubscriber } from "svelte/store";

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

  type MeetingWithClientData = Meeting & {
    formattedDate: string;
    parsedDate: Date;
  };
  type MeetingArray = Array<MeetingWithClientData>;
  let contact: Contact;
  let meetings: MeetingArray;

  let unsubscribe: Unsubscriber;

  onMount(() => {
    unsubscribe = store.subscribe((value) => {
      contact = value.contacts[id];
      meetings = Object.values(contact.meetings).map(
        (meeting: MeetingWithClientData) => {
          meeting.parsedDate = new Date(meeting.date);
          meeting.formattedDate = format(meeting.parsedDate, "");
          return meeting;
        }
      );
    });
  });

  onDestroy(() => {
    contact = undefined;
    unsubscribe();
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
        <div class="h-divider w-full" />
        <div class="w-full flex items-center justify-between">
          <h2 class="text-xl font-bold">Meetings</h2>
          <AddMeetingModal id={contact.id} />
        </div>
        {#if meetings?.length > 0}
          <div class="grid grid-cols-2">
            {#each meetings as meeting}
              <div class="px-6 py-6 flex flex-col items-start justify-center">
                <span>{meeting.formattedDate}</span>
              </div>
            {/each}
          </div>
        {:else}
          <div
            class="font-medium text-2xl w-full grid place-items-center bg-gray-100 py-8 px-4"
          >
            No meetings found
          </div>
        {/if}
      </template>
    {:else}
      <Spinner />
    {/if}
  </div>
</div>
<Modal
  isOpen={deleteModal.isOpen}
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
</Modal>
