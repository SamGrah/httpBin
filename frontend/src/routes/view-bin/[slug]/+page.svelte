<script>
  // @ts-nocheck
	export let data;
  const { binContents, hostname, binId } = data
  for (const request of binContents) {
    console.log(request.contents)
    console.log(request.recieved)
  }
  
  const calculateTimeStr = (request) => {
    const msDiff = Date.now() - (request.recieved.seconds)*1000
    const secsDiff = Math.floor(msDiff / 1000)
    const minsDiff = Math.floor(secsDiff / 60)
    const hoursDiff= Math.floor(minsDiff / 60)
    const daysDiff = Math.floor(hoursDiff / 24)
    let timeStr = `${secsDiff} secs`
    if (minsDiff >= 1) timeStr = `${minsDiff} mins`
    if (hoursDiff >= 1) timeStr = `${hoursDiff} hours`
    if (daysDiff >= 1) timeStr = `${daysDiff} days`
    return timeStr
  }
</script>

<!-- var ws = new WebSocket(`wss://!{hostname}/wss`);
    ws.onopen = function () {
        ws.send('!{bin}')
    }

    ws.onmessage = function (serverMsg) {
        console.log(serverMsg.data);
        if (serverMsg.data == 'refresh') {
          location.reload()
        }
    } -->
{#if !binContents?.length}
	<div class="max-w-md py-4 px-8 bg-white shadow-lg rounded-lg my-20">
		<div>
			<h2 class="text-gray-800 text-3xl font-semibold">Bin {binId} is Empty</h2>
			<p class="mt-4 text-gray-600">
				No HTTP requests have been recieved by bin {binId}.<br />
				A request of any type (<i>GET</i>, <i>DELETE</i>, etc) can be added to this bin by making a
				request to the following address.
			</p>
			<br />
			<div class="flex justify-center mt-4.mb-3">
				<a class="text-xl font-medium text-blue-900" href=".">{hostname}/bin/{binId}</a>
			</div>
		</div>
	</div>
{:else}
  <ul>
    {#each binContents as request}
        <li class="m-6 grid grid-cols-3 border-2 border-gray-300">
            <!-- <div class="p-2 bg-gray-100" style="white-space:pre;">https://{request.hostIp}</div> -->
            <!-- <div class="p-2 bg-gray-100">{request.requestType} {request.recieved.originalUrl}</div> -->
            <div class="p-2 text-right bg-gray-100" style="white-space:pre;">{calculateTimeStr(request)} ago | {request.hostIp}</div>
        </li>
      {/each}
	<!-- ul each request, idx in requests li.m-6.grid.grid-cols-3.border-2.border-gray-300
	div.p-2.bg-gray-100(style="white-space:pre;") | https://#{request.headers.host}
	| #[b #{request.requestType}] #{request.originalUrl}
	div.p-2.bg-gray-100 #{request.headers['content-type']}
	div.p-2.text-right.bg-gray-100(style='white-space:pre;') | #{timeStr} ago | #{request.senderIP}
	div.p-2(style='white-space:pre;') span.font-bold.text-gray-500 FORM/POST PARAMETERS | ul each value,
	key in request.body li.whitespace-normal.break-all #[b #{key}]: #{value}
	else li None div.p-2.col-span-2(style='white-space:pre;') span.font-bold.text-gray-500 HEADERS ul each
	value, key in request.headers li.whitespace-normal.break-all #[b #{key}]: #{value}
	else li None div.p-2.col-span-3(style='white-space:pre;') span.font-bold.text-gray-500 RAW BODY if
	Object.entries(request.body).length !== 0 .whitespace-normal.break-all | |#{JSON.stringify(
		request.body
	)}
	else | | None else" -->
  </ul>
{/if}
