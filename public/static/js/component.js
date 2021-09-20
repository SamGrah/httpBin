class CreateBinComponent extends HTMLElement {
  connectedCallback() {
    this.innerHTML = /*html*/`
      <div class="text-green-600">
        <h2 class=""><span class="">Inspect HTTP Requests</span></h2>
        <h3 class="banner-subheader">
          This is a clone of the <a href="https://requestbin.net">RequestBin</a> tool.<br />
          RequestBin gives you a URL that will collect requests made to it and let you inspect them in a human-friendly way.<br>
          Use RequestBin to see what your HTTP client is sending or to inspect and debug webhook requests.</h3>
        <form class="form-inline">
          <p class="banner-button">
            <a class="btn btn-success btn-large" onclick="createBin()"><i class="icon-plus-sign"></i> Create a RequestBin</a>
          </p>
        </form>
      </div>
      `
  }
}

customElements.define('create-bin-element', CreateBinComponent)