<p class="text-center font-bold text-3xl">Execute Requests With The Following Code</p>
<div class="mt-4 w-4/6"><b>cURL</b>
  <pre
    class="p-2 mt-2 border-gray-300 border-2 whitespace-normal break-all bg-gray-100"><code>curl -X POST -d "fizz=buzz" localhost/bin/7tqid7</code></pre>
</div>
<div class="mt-4 w-4/6"><b>PowerShell</b>
  <pre
    class="p-2 mt-2 border-gray-300 border-2 whitespace-pre-wrap break-all bg-gray-100"><code>powershell -NoLogo -Command "(New-Object System.Net.WebClient).DownloadFile('localhost/bin/7tqid7', 'C:\Windows\Temp\ednze13v.txt')"</code></pre>
</div>
<div class="mt-4 w-4/6"><b>Python (with Requests)</b>
  <pre class="p-2 mt-2 border-gray-300 border-2 whitespace-pre-wrap break-all bg-gray-100"><code>{`import requests, time
r = requests.post('localhost/bin/7tqid7', data={"ts":time.time()})
print r.status_code
print r.content`}</code></pre>
</div>
<div class="mt-4 w-4/6"><b>Node.js (with request)</b>
  <pre class="p-2 mt-2 border-gray-300 border-2 whitespace-pre-wrap break-all bg-gray-100"><code>
    {`{var request = require('request');
var url ='localhost/bin/7tqid7'
request(url, function (error, response, body) {
  if (!error) {
    console.log(body);
  }
});`}</code></pre>
</div>
<div class="mt-4 w-4/6"><b>Ruby</b>
  <pre class="p-2 mt-2 border-gray-300 border-2 whitespace-pre-wrap break-all bg-gray-100"><code>
    {`require 'open-uri'
result = open('localhost/bin/7tqid7')
result.lines { |f| f.each_line {|line| p line} }</code></pre>
</div>
<div class="mt-4 w-4/6"><b>C# / .NET</b>
  <pre class="p-2 mt-2 border-gray-300 border-2 whitespace-pre-wrap break-all bg-gray-100"><code>using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace RequestBinExample
{
class Program
{
  static void Main(string[] args)
  {
  MakeRequest();
  }

  private static async Task MakeRequest()
  {
  var httpClient = new HttpClient();
  var response = await httpClient.GetAsync(new Uri("localhost/bin/7tqid7"));
  var body = await response.Content.ReadAsStringAsync();
  Console.WriteLine(body);
}
}
}`}</code></pre>
</div>
<div class="mt-4 w-4/6"><b>Java</b>
  <pre class="p-2 mt-2 border-gray-300 border-2 whitespace-pre-wrap break-all bg-gray-100"><code>{
  `import org.apache.commons.httpclient.*;
import org.apache.commons.httpclient.methods.*;
import org.apache.commons.httpclient.params.HttpMethodParams;

import java.io.*;

public class RequestBinTutorial {
public static void main(String[] args) {
  HttpClient client = new HttpClient();
  GetMethod method = new GetMethod("localhost/bin/7tqid7");
  try {
    int statusCode = client.executeMethod(method);
    byte[] responseBody = method.getResponseBody();
    System.out.println(new String(responseBody));
  } catch (Exception e) {
    System.err.println("Fatal error: " + e.getMessage());
    e.printStackTrace();
  } finally {
    method.releaseConnection();
  }
}
}`}</code></pre>
</div>
<div class="mt-4 w-4/6"><b>PHP</b>
  <pre class="p-2 mt-2 border-gray-300 border-2 whitespace-pre-wrap break-all bg-gray-100"><code>&lt;?php
  $result = file_get_contents('localhost/bin/7tqid7');
  echo $result;
?/&gt;</code></pre>
</div>