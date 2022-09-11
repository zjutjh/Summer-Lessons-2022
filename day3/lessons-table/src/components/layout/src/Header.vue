<template>
  <header class="k-header">
    <slot></slot>
  </header>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  setup() {},
});
</script>

<style lang="scss">
.k-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
  padding: 0 50px;
  color: #000000d9;
  line-height: 64px;
  background: #001529;
  h1 {
    margin: 0;
    color: #fff;
  }
}
</style>
