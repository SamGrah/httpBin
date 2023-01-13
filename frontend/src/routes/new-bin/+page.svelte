<script>
  import axios from 'axios'
  import CodeSnippets from "./CodeSnippets.svelte"
  import { scrollEnabled } from '../../store.js';
  
  async function createNewBin() {
    const { data } = await axios.post('http://localhost:8080/new-bin')
    return data.bin.Id
  }

  $scrollEnabled = true;
  let newBin = createNewBin();
</script>

{#await newBin}
  <div></div>
{:then newBin}
  <div class="m-11 grid grid-cols-1 gap-1 justify-items-center w-4/6">
    <p class="mb-4 text-center font-bold text-3xl">Bin <i>{newBin}</i> Has Been Created</p>
    <div class="mb-10 grid grid-cols-2 divide-x divide-gray-100">
      <div class="p-4 w-full rounded-md">
        <p class="mt-3 mb-1 text-center">HTTP requests made to this endpoint will be logged</p><input
          class="mb-3 bg-green-50 border border-gray w-full outline-none text-gray-500 rounded-md p-1 text-lg text-center"
          value="localhost/bin/{newBin}" disabled="">
      </div>
      <div class="p-4 w-full rounded-md">
        <p class="mt-3 mb-1 text-center">Visit this endpoint to review logged HTTP requests</p><input
          class="mb-3 bg-green-50 border border-gray w-full outline-none text-gray-500 rounded-md p-1 text-lg text-center"
          value="localhost/view-bin/{newBin}" disabled="">
      </div>
    </div>
    <CodeSnippets bin={newBin}/>
  </div>
{/await}