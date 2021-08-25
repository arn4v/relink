<script lang="ts">
  import { store } from "../store";
  import DeleteIcon from "./DeleteIcon.svelte";
  import DeleteModal from "./DeleteModal.svelte";

  export let id: string;
  export let contactId: string;

  let isOpen = false;
  const onOpen = () => {
    isOpen = true;
  };

  const onClose = () => {
    isOpen = false;
  };

  const onSubmit = () => {
    store.update((val) => {
      delete val.contacts[contactId].meetings[id];
      return val;
    });
  };
</script>

<button
  class="p-1.5 hover:bg-gray-100 border border-gray-200 transition bg-white"
  on:click={onOpen}
>
  <DeleteIcon class="h-5 w-5" />
</button>
<DeleteModal {isOpen} {onClose} {onSubmit} />
