<script lang="ts">
  import * as uuid from "uuid";
  import { store } from "../store";
  import Modal from "./Modal.svelte";
  import PlusIcon from "./PlusIcon.svelte";

  export let id: string;

  let isOpen = false;

  let initialState = {
    note: "",
    date: "",
  };

  let state = { id: uuid.v4(), ...initialState };

  const onOpen = () => {
    isOpen = true;
  };

  const onClose = () => {
    isOpen = false;
    state = { id: uuid.v4(), ...initialState };
  };

  const onSubmit = () => {
    store.update((val) => {
      val.contacts[id].meetings[state.id] = state;
      return val;
    });
    onClose();
  };
</script>

<button
  class="p-1.5 hover:bg-gray-100 border border-gray-200 transition bg-white"
  on:click={onOpen}
>
  <PlusIcon />
</button>

<Modal {isOpen} {onClose}>
  <div
    class="bg-white w-[25vw] z-[200] rounded-md relative flex flex-col items-start justify-between p-8 space-y-8"
  >
    <h1 class="text-xl font-bold mt-0">Add meeting</h1>
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
