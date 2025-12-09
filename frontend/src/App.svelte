<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { EventsOff, EventsOn } from "../wailsjs/runtime/runtime";
  import {
    SetPort,
    StartServer,
    StopServer,
  } from "../wailsjs/go/server/ServerHandler.js";
  import Message from "./components/Message.svelte";
  import type { Message as MessageType } from "./types";

  let port: number = null;
  let messages: MessageType[] = [];
  let consoleDiv: HTMLDivElement;

  async function setPort(): Promise<void> {
    await SetPort(port)
      .then(() => {
        document.getElementById("start-btn").removeAttribute("disabled");
      })
      .catch((error) => {
        messages = [...messages, error];
      });
  }

  async function startServer(): Promise<void> {
    await StartServer()
      .then(() => {
        document.getElementById("stop-btn").removeAttribute("disabled");
        document
          .getElementById("start-btn")
          .setAttribute("disabled", "disabled");
      })
      .catch((error) => {
        messages = [...messages, error];
      });
  }

  async function stopServer(): Promise<void> {
    await StopServer().then(() => {
      document.getElementById("start-btn").removeAttribute("disabled");
      document.getElementById("stop-btn").setAttribute("disabled", "disabled");
    });
  }

  onMount(() => {
    EventsOn("console-message", (message: MessageType) => {
      messages = [...messages, message];
      setTimeout(() => {
        if (consoleDiv) {
          consoleDiv.scrollTop = consoleDiv.scrollHeight;
        }
      }, 0);
    });
  });

  onDestroy(() => {
    EventsOff("console-message");
  });
</script>

<main>
  <div class="input-box" id="input">
    <input
      autocomplete="off"
      bind:value={port}
      class="input"
      id="port"
      type="number"
      placeholder="choose port number..."
      min="0"
    />
    <button class="btn" on:click={setPort}>Set</button>
  </div>
  <div class="start-stop" id="start-stop">
    <button class="start-btn" id="start-btn" disabled on:click={startServer}
      >Start Server</button
    >
    <button class="stop-btn" id="stop-btn" disabled on:click={stopServer}
      >Stop Server</button
    >
  </div>
  <div class="console" bind:this={consoleDiv}>
    {#each messages as msg}
      <Message message={msg} />
    {/each}
  </div>
</main>

<style>
  main {
    height: 100%;
    display: flex;
    flex-direction: column;
  }

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

  .start-btn:hover {
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

  .stop-btn:hover {
    background-color: #c15c5c;
  }

  .console {
    display: flex;
    flex-direction: column;
    flex: 1;
    overflow: auto;
    background: #151415;
    border-radius: 10px;
    color: #fff5e9;
    font-family: "Courier New", monospace;
    padding: 10px 10px;
    margin: 30px 10px 10px 10px;
    overflow-y: auto;
    border: 2px solid #fff5e9;
    text-align: left;
  }
</style>
