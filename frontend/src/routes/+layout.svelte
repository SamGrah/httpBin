<script>
	import '../app.css';
  import Nav from './Nav.svelte'
</script>

<Nav />
<div class="h-full flex flex-col">
	<div class="mt-24 mx-8 flex justify-center">
		<slot />
	</div>
</div>
