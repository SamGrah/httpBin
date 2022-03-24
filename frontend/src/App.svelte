<script>
  import Nav from "./components/nav.svelte"
  import IntroPage from "./components/create_bin.svelte";
  import BinInfo from "./components/bin_info.svelte"
  import { subPageToDisplay, scrollEnabled } from './store.js';
</script>

<div class="h-full flex flex-col">
  <Nav />
  <div class='mt-24 mx-8 flex justify-center'>
    {#if $subPageToDisplay === 'create-bin'}
      <IntroPage />
    {:else if $subPageToDisplay === 'bin-info'}
      <BinInfo />
    {/if}
  </div>
</div>