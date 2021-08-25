<script>
  import Modal from "./Modal.svelte";

  export let isOpen;
  export let onClose;
  export let onSubmit;
</script>

<Modal {isOpen} {onClose}>
  <div
    class="bg-white h-[15vh] w-[25vw] z-[200] rounded-md relative flex flex-col items-center justify-between p-8"
  >
    <h1 class="text-xl font-bold">Do you really want to delete this?</h1>
    <div class="flex items-center justify-between w-full">
      <button
        class="border border-gray-200 px-4 py-2 hover:bg-gray-100 transition font-semibold"
      >
        Cancel
      </button>
      <button
        class="px-4 py-2 bg-red-500 font-semibold text-white hover:bg-red-700 transition"
        on:click={onSubmit}>Confirm</button
      >
    </div>
  </div>
</Modal>
