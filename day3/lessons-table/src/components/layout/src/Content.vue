<template>
  <main class="k-content">
    <slot></slot>
  </main>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  setup() {},
});
</script>
<style lang="scss" scoped>
.k-content {
  flex: auto;
}
</style>
