<script lang="ts">
  import { onMount } from "svelte";
  import { EventsOn } from "../wailsjs/runtime/runtime"
  import {SetPort, StartServer, StopServer} from "../wailsjs/go/server/ServerHandler.js";

  let port: number = -1;
  let rxMessages: string[] = [];
  let consoleDiv: HTMLDivElement;

  async function setPort(): Promise<void> {
    await SetPort(port)
      .then(() => {
        document.getElementById("start-btn").removeAttribute("disabled");
      })
      .catch((error) => {
        rxMessages = [...rxMessages, error];
      });
  }

  async function startServer(): Promise<void> {
    await StartServer()
      .then(() => {
        document.getElementById("stop-btn").removeAttribute("disabled");
        document.getElementById("start-btn").setAttribute("disabled", "disabled");
      })
      .catch((error) => {
        rxMessages = [...rxMessages, error];
      });
  }

  async function stopServer(): Promise<void> {
    await StopServer()
      .then(() => {
        document.getElementById("start-btn").removeAttribute("disabled");
        document.getElementById("stop-btn").setAttribute("disabled", "disabled");
      })
  }

  onMount(() => {
    EventsOn("message-rx", (message: string) => {
      rxMessages = [...rxMessages, message];
      setTimeout(() => {
        if (consoleDiv) {
          consoleDiv.scrollTop = consoleDiv.scrollHeight;
        }
      }, 0);
    });
  });
</script>

<main>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={port} class="input" id="port" type="number" placeholder="port" min=0/>
    <button class="btn" on:click={setPort}>Set</button>
  </div>
  <div class="start-stop" id="start-stop">
    <button class="start-btn" id="start-btn" disabled on:click={startServer}>Start Server</button>
    <button class="stop-btn" id="stop-btn" disabled on:click={stopServer}>Stop Server</button>
  </div>
  <div class="console" bind:this={consoleDiv}>
    {#each rxMessages as message}
      <div class="message">{message}</div>
    {/each}
  </div>
     
</main>

<style>

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 5px 5px 5px 5px;
    padding: 0 10px;
    background-color: #fff5e9;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-color: #ffffff;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: #fff5e9;
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    background-color: #ffffff;
  }

  .input-box .input:focus {
    background-color: #ffffff;
  }

  .start-btn {
    width: 125px;
    height: 30px;
    font-weight: 600;
    background-color: #587792;
    color: 1c1b1c;
    border: none;
    border-radius: 3px;
  }

  .start-btn:hover{
    background-color: #6687a3;
  }

  .stop-btn {
    width: 125px;
    height: 30px;
    font-weight: 600;
    background-color: #b44545;
    color: 1c1b1c;
    border: none;
    border-radius: 3px;
  }

  .stop-btn:hover{
    background-color: #c15c5c;
  }

  .console {
    background: #151415;
    border-radius: 10px;
    color: #fff5e9;
    font-family: "Courier New", monospace;
    padding: 10px 10px;
    height: 400px;
    margin: 30px 10px 10px 10px;
    overflow-y: auto;
    border: 2px solid #fff5e9;
    text-align: left;
  }

  .message {
    margin: 2px 0;
  }

</style>
