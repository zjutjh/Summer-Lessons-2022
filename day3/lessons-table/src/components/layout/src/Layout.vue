<template>
  <section class="k-layout">
    <slot></slot>
  </section>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Layout",
  setup() {},
});
</script>

<style lang="scss" scoped>
.k-layout {
  display: flex;
  flex: auto;
  flex-direction: column;
  background: #f0f2f5;
  height: 100vh;
}
</style>
