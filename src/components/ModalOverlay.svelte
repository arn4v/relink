<script lang="ts">
  import { fade, fly } from "svelte/transition";

  export let isOpen: boolean;
  export let onClose: () => void;
  export let className: string;
</script>

{#if isOpen}
  <div class="fixed z-50 inset-0 h-screen w-screen">
    <div
      class="h-full w-full bg-opacity-75 bg-black z-50 fixed inset-0"
      transition:fade={{ duration: 250 }}
      on:click={onClose}
    />
    <div
      class="absolute z-[200] top-[10%] left-1/2 -translate-x-1/2"
      class:className
      transition:fly={{ duration: 300, y: -50 }}
    >
      <slot />
    </div>
  </div>
{/if}
