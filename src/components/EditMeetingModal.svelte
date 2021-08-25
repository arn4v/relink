<script lang="ts">
  import { onMount } from "svelte";
  import { store } from "../store";
  import EditIcon from "./EditIcon.svelte";
  import Modal from "./Modal.svelte";

  export let id: string;
  export let contactId: string;

  let state: Meeting;

  let isOpen = false;
  const onOpen = () => {
    isOpen = true;
  };

  const onClose = () => {
    isOpen = false;
  };

  const onSubmit = () => {
    store.update((val) => {
      val.contacts[contactId].meetings[id] = state;
      return val;
    });
    onClose();
  };

  onMount(() => {
    const unsubscribe = store.subscribe((val) => {
      if (!state) {
        state = val.contacts[contactId].meetings[id];
      } else {
        unsubscribe();
      }
    });
  });
</script>

<button
  class="p-1.5 hover:bg-gray-100 border border-gray-200 transition bg-white"
  on:click={onOpen}
>
  <EditIcon class="h-5 w-5" />
</button>
<Modal {isOpen} {onClose}>
  <div
    class="bg-white w-[25vw] z-[200] rounded-md relative flex flex-col items-start justify-between p-8 space-y-8"
  >
    <h1 class="text-xl font-bold mt-0">Edit meeting</h1>
    <form on:submit={onSubmit} class="flex flex-col space-y-4 w-full">
      <label for="date">Date</label>
      <input id="date" type="date" bind:value={state.date} />
      <label for="note">Note</label>
      <textarea
        id="note"
        bind:value={state.note}
        rows={3}
        class="px-4 w-full focus:outline-none focus:ring-2 ring-black transition duration-100 h-auto shadow"
      />
    </form>
    <div class="flex items-center justify-between w-full">
      <button
        class="border border-gray-200 px-4 py-2 hover:bg-gray-100 transition font-semibold"
        on:click={onClose}
      >
        Cancel
      </button>
      <button
        class="px-4 py-2 btn-primary font-semibold text-white"
        on:click={onSubmit}>Save meeting</button
      >
    </div>
  </div>
</Modal>
