<script>
  import {mainComponentToDisplay} from './store'
  import Nav from "./components/nav.svelte"
  import Intro from "./components/intro.svelte";
  import NewBinInfo from "./components/new_bin_info.svelte";
</script>

<div class="h-full flex flex-col">
  <Nav />
  <div class='mt-24 mx-8 flex justify-center'>
    {#if $mainComponentToDisplay === "Intro"}
      <Intro />
    {:else if $mainComponentToDisplay === "NewBinInfo"}
      <NewBinInfo />
    {:else}
      404: Page not Found
    {/if}
  </div>
</div>