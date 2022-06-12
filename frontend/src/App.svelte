<script>
  import { Router, Route } from "svelte-navigator"
  import Nav from "./components/nav.svelte";
  import Intro from "./components/intro.svelte";
  import NewBinInfo from "./components/new_bin_info.svelte";
  import ViewBin from "./components/view_bin.svelte"
</script>

<div class="h-full flex flex-col">
  <Router>
  <Nav />
  <div class='mt-24 mx-8 flex justify-center'>
      <Route path="/"><Intro /></Route>
      <Route path="new-bin"><NewBinInfo /></Route>
      <Route path="view-bin/:binId" let:params>
        <ViewBin binId={params.binId} />
      </Route>
  </div>
</Router>
</div>