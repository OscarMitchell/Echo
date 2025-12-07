<script lang="ts">
  import type { Message } from "../types";
  export let message: Message;
</script>

<div class="message">
  <div class="message-header">
    <span class="time-and-sender"
      >Message from: {message.sender} @ {message.rxTime}</span
    >
    <span class="size">Length: {message.size}</span>
  </div>
  <div class="message-content">
    {message.content}
  </div>
</div>

<style>
  .message {
    border-top: 1px solid #fff5e9;
    margin: 0.5rem 0;
  }

  .message-header {
    display: flex;
    justify-content: space-between;
    font-size: 0.8rem;
    color: #ffc885;
    margin: 0.5rem;
  }

  .message-content {
    color: #fff5e9;
    margin: 0.5rem;
  }
</style>
